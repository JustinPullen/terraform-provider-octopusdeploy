package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func expandRunbook(d *schema.ResourceData) *octopusdeploy.Runbook {
	name := d.Get("name").(string)
	projectID := d.Get("project_id").(string)

	runbook := octopusdeploy.NewRunbook(name, projectID)
	runbook.ID = d.Id()

	if v, ok := d.GetOk("connectivity_policy"); ok {
		runbook.ConnectivityPolicy = expandConnectivityPolicy(v.([]interface{}))
	}

	if v, ok := d.GetOk("default_guided_failure_mode"); ok {
		runbook.DefaultGuidedFailureMode = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		runbook.Description = v.(string)
	}

	if v, ok := d.GetOk("environment_scope"); ok {
		runbook.EnvironmentScope = v.(string)
	}

	if v, ok := d.GetOk("environment"); ok {
		runbook.Environments = getSliceFromTerraformTypeList(v)
	}

	if v, ok := d.GetOk("multi_tenancy_mode"); ok {
		runbook.MultiTenancyMode = v.(string)
	}

	if v, ok := d.GetOk("published_runbook_snapshot_id"); ok {
		runbook.PublishedRunbookSnapshotID = v.(string)
	}

	runbookRetentionPeriod := expandRunbookRetentionPeriod(d, "runbook_retention_period")
	if runbookRetentionPeriod != nil {
		runbook.RunRetentionPolicy = *&runbookRetentionPeriod
	}

	if v, ok := d.GetOk("runbook_process_id"); ok {
		runbook.RunbookProcessID = v.(string)
	}

	if v, ok := d.GetOk("runbook_space_id"); ok {
		runbook.SpaceID = v.(string)
	}

	return runbook
}
