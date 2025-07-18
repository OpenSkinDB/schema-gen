package parsers

import (
	"context"
	"strconv"
	"time"

	"go-csitems-parser/models"
	"go-csitems-parser/modules"

	"github.com/rs/zerolog"
)

func ParseSkinCollections(ctx context.Context, ig *models.ItemsGame, t *modules.Translator) []models.BaseWeapon {
	logger := zerolog.Ctx(ctx)

	start := time.Now()
	// logger.Info().Msg("Parsing music kits...")

	client_loot_lists, err := ig.Get("client_loot_lists")

	if err != nil {
		logger.Error().Err(err).Msg("Failed to get client_loot_lists")
		return nil
	}

	var knives []models.BaseWeapon
	for _, c := range client_loot_lists.GetChilds() {
		prefab, _ := c.GetString("prefab")

		if prefab != "melee_unusual" {
			// Skip non-knife items
			continue
		}

		definition_index, _ := strconv.Atoi(c.Key)
		classname, _ := c.GetString("name")
		image_inventory, _ := c.GetString("image_inventory")
		item_name, _ := c.GetString("item_name")
		// item_description, _ := c.GetString("item_description")

		current := models.BaseWeapon{
			DefinitionIndex: definition_index,
			ClassName:       classname,
			Name:            item_name,
			// ItemName:        item_name,
			// Name:            name,
			// ItemDescription: item_description,
			// Prefab:          prefab,
			ImageInventory: image_inventory,
		}

		knives = append(knives, current)
	}

	// Save knives to the database
	duration := time.Since(start)
	logger.Info().Msgf("Parsed '%d' knives in %s", len(knives), duration.String())

	return knives
}
