package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func expandRunbookRetentionPeriod(d *schema.ResourceData, key string) *octopusdeploy.RunbookRetentionPeriod {
	v, ok := d.GetOk(key)
	if ok {
		runbookRetentionPeriod := v.([]interface{})
		if len(runbookRetentionPeriod) == 1 {
			tfRetentionItem := runbookRetentionPeriod[0].(map[string]interface{})
			retention := octopusdeploy.RunbookRetentionPeriod{
				QuantityToKeep:    int32(tfRetentionItem["quantity_to_keep"].(int)),
				ShouldKeepForever: tfRetentionItem["should_keep_forever"].(bool),
			}
			return &retention
		}
	}

	return nil
}

func flattenRunbookRetentionPeriod(r octopusdeploy.RunbookRetentionPeriod) []interface{} {
	runbookRetentionPeriod := make(map[string]interface{})
	runbookRetentionPeriod["quantity_to_keep"] = int(r.QuantityToKeep)
	runbookRetentionPeriod["should_keep_forever"] = r.ShouldKeepForever
	return []interface{}{runbookRetentionPeriod}
}

func getRunbookRetentionPeriodDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"quantity_to_keep": {
			Computed:    true,
			Description: "The number of days/releases to keep. The default value is `30`. If `0` then all are kept.",
			Type:        schema.TypeInt,
		},
		"should_keep_forever": {
			Computed:    true,
			Description: "Indicates if items should never be deleted. The default value is `false`.",
			Type:        schema.TypeBool,
		},
	}
}

func getRunbookRetentionPeriod() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"quantity_to_keep": {
			Default:          30,
			Description:      "The number of days/releases to keep. The default value is `30`. If `0` then all are kept.",
			Optional:         true,
			Type:             schema.TypeInt,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
		},
		"should_keep_forever": {
			Default:     false,
			Description: "Indicates if items should never be deleted. The default value is `false`.",
			Optional:    true,
			Type:        schema.TypeBool,
		},
	}
}
