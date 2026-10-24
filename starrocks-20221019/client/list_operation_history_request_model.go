// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListOperationHistoryRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListOperationHistoryRequest
	GetInstanceId() *string
	SetOperationId(v string) *ListOperationHistoryRequest
	GetOperationId() *string
	SetOperationStatus(v string) *ListOperationHistoryRequest
	GetOperationStatus() *string
	SetOperationType(v string) *ListOperationHistoryRequest
	GetOperationType() *string
	SetPageNumber(v int32) *ListOperationHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOperationHistoryRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListOperationHistoryRequest
	GetStartTime() *int64
}

type ListOperationHistoryRequest struct {
	// The end time of the operation. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1742179008000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// op-f49743caa809****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The operation status. Valid values:
	//
	// - COMPLETED: Completed.
	//
	// - TERMINATED: Terminated.
	//
	// - HUMAN_PROCESSING: Pending manual processing.
	//
	// example:
	//
	// COMPLETED
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// The operation type. Valid values:
	//
	// - trial_to_official: Upgrades the trial edition to the standard edition.
	//
	// - upgrade_version: Upgrades the version.
	//
	// - update_configuration: Updates the configuration.
	//
	// - update_public_network_status: Updates the public network status.
	//
	// - create_cluster: Creates a cluster.
	//
	// - delete_cluster: Deletes a cluster.
	//
	// - disable_cluster: Stops a cluster.
	//
	// - enable_cluster: Recovers a cluster.
	//
	// - restart_cluster: Restarts a cluster.
	//
	// - migrate_cluster: Migrates a cluster.
	//
	// - renew_cluster: Renews a cluster.
	//
	// - modify_charge_type: Changes the billing method.
	//
	// - UPGRADE: Upgrades a cluster.
	//
	// - DOWNGRADE: Downgrades a cluster.
	//
	// - create_node_group: Creates a node group.
	//
	// - delete_node_group: Deletes a node group.
	//
	// - disable_node_group: Stops a node group.
	//
	// - enable_node_group: Recovers a node group.
	//
	// - sre_operation: Performs O&M on a cluster.
	//
	// - resource_change: Changes resources.
	//
	// - disable_postpaid_resource: Disables pay-as-you-go resources.
	//
	// - enable_postpaid_resource: Recovers pay-as-you-go resources.
	//
	// - restart_node_group: Restarts a compute group.
	//
	// - enable_ha_cluster: Enables high availability (HA) to keep the cluster active.
	//
	// - restart_node: Restarts a node.
	//
	// - backup: Backs up data.
	//
	// - delete_backup: Deletes a data backup.
	//
	// - cancel_backup_task: Cancels a data backup task.
	//
	// - modify_timezone: Modifies the system time zone.
	//
	// - restore: Recovers data.
	//
	// - switch_az: Switches the primary and secondary zones.
	//
	// - rollback_upgrade_version: Rolls back a version upgrade.
	//
	// - scale_out_fe: Scales out FE nodes.
	//
	// - scale_in_fe: Scales in FE nodes.
	//
	// - upgrade_fe_cu: Upgrades the FE CU specifications.
	//
	// - downgrade_fe_cu: Decrease the quota of FE CU specifications.
	//
	// - increase_fe_disk_size: Increases the FE disk size.
	//
	// - decrease_fe_disk_size: Decreases the FE disk size.
	//
	// - increase_fe_disk_number: Increases the number of FE disks.
	//
	// - decrease_fe_disk_number: Decreases the number of FE disks.
	//
	// - upgrade_fe_disk_performance_level: Upgrades the FE disk performance level (PL).
	//
	// - downgrade_fe_disk_performance_level: Decrease the quota of the FE disk performance level (PL).
	//
	// - create_agent: Creates an agent.
	//
	// - upgrade_agent_cu: Upgrades the agent CU specifications.
	//
	// - scale_out_be: Scales out BE nodes.
	//
	// - scale_in_be: Scales in BE nodes.
	//
	// - upgrade_be_cu: Upgrades the BE CU specifications.
	//
	// - downgrade_be_cu: Decrease the quota of BE CU specifications.
	//
	// - increase_be_disk_size: Increases the BE disk size.
	//
	// - decrease_be_disk_size: Decreases the BE disk size.
	//
	// - increase_be_disk_number: Increases the number of BE disks.
	//
	// - decrease_be_disk_number: Decreases the number of BE disks.
	//
	// - upgrade_be_disk_performance_level: Upgrades the BE disk performance level (PL).
	//
	// - downgrade_be_disk_performance_level: Decrease the quota of the BE disk performance level (PL).
	//
	// - upgrade_be_spec_type: Upgrades the BE specification type.
	//
	// - downgrade_be_spec_type: Decrease the quota of the BE specification type.
	//
	// - scale_out_cn: Scales out CN nodes.
	//
	// - scale_in_cn: Scales in CN nodes.
	//
	// - upgrade_cn_cu: Upgrades the CN CU specifications.
	//
	// - downgrade_cn_cu: Decrease the quota of CN CU specifications.
	//
	// - increase_cn_disk_size: Increases the CN disk size.
	//
	// - decrease_cn_disk_size: Decreases the CN disk size.
	//
	// - increase_cn_disk_number: Increases the number of CN disks.
	//
	// - decrease_cn_disk_number: Decreases the number of CN disks.
	//
	// - upgrade_cn_disk_performance: Upgrades the CN disk performance level (PL).
	//
	// - downgrade_cn_disk_performance: Decrease the quota of the CN disk performance level (PL).
	//
	// - upgrade_cn_spec_type: Upgrades the CN specification type.
	//
	// - downgrade_cn_spec_type: Decrease the quota of the CN specification type.
	//
	// - elastic_scale_out_cn: Elastically scales out CN nodes.
	//
	// - elastic_scale_in_cn: Elastically scales in CN nodes.
	//
	// example:
	//
	// update_configuration
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The page number of the current page. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging in a paged query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the operation. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1742179008000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListOperationHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListOperationHistoryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationHistoryRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *ListOperationHistoryRequest) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *ListOperationHistoryRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListOperationHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOperationHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationHistoryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationHistoryRequest) SetEndTime(v int64) *ListOperationHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListOperationHistoryRequest) SetInstanceId(v string) *ListOperationHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationHistoryRequest) SetOperationId(v string) *ListOperationHistoryRequest {
	s.OperationId = &v
	return s
}

func (s *ListOperationHistoryRequest) SetOperationStatus(v string) *ListOperationHistoryRequest {
	s.OperationStatus = &v
	return s
}

func (s *ListOperationHistoryRequest) SetOperationType(v string) *ListOperationHistoryRequest {
	s.OperationType = &v
	return s
}

func (s *ListOperationHistoryRequest) SetPageNumber(v int32) *ListOperationHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationHistoryRequest) SetPageSize(v int32) *ListOperationHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationHistoryRequest) SetStartTime(v int64) *ListOperationHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *ListOperationHistoryRequest) Validate() error {
	return dara.Validate(s)
}
