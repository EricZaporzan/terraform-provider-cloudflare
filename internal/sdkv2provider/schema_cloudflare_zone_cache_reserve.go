package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneCacheReserveSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: consts.ZoneIDSchemaDescription,
			ValidateFunc: func(value any, key string) (_ []string, errs []error) {
				// Ensure that a valid Zone ID was passed.
				if err := validateZoneID(value.(string)); err != nil {
					errs = append(errs, err)
				}
				return
			},
		},
		"enabled": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Whether to enable or disable Cache Reserve support for a given zone.",
		},
	}
}
