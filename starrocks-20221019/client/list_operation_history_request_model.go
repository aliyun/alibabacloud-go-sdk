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
	// End time of the operation.
	//
	// example:
	//
	// 1742179008000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Operation ID.
	//
	// example:
	//
	// op-f49743caa809****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// Operation status:
	//
	// - COMPLETED: completed
	//
	// - TERMINATED: terminated
	//
	// - HUMAN_PROCESSING: pending manual processing
	//
	// example:
	//
	// COMPLETED
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// Operation type. Valid values:
	//
	// - trial_to_official: upgrade from Trial Edition to Standard Edition
	//
	// - upgrade_version: upgrade version
	//
	// - update_configuration: update configuration
	//
	// - update_public_network_status: update public network status
	//
	// - create_cluster: create cluster
	//
	// - delete_cluster: delete cluster
	//
	// - disable_cluster: stop cluster
	//
	// - enable_cluster: resume cluster
	//
	// - restart_cluster: restart cluster
	//
	// - migrate_cluster: migrate cluster
	//
	// - renew_cluster: renew cluster
	//
	// - modify_charge_type: change billing method
	//
	// - UPGRADE: upgrade cluster
	//
	// - DOWNGRADE: downgrade cluster
	//
	// - create_node_group: create node group
	//
	// - delete_node_group: delete node group
	//
	// - disable_node_group: stop node group
	//
	// - enable_node_group: resume node group
	//
	// - sre_operation: O\\&M cluster
	//
	// - resource_change: resource change
	//
	// - disable_postpaid_resource: disable pay-as-you-go resources
	//
	// - enable_postpaid_resource: enable pay-as-you-go resources
	//
	// - restart_node_group: restart compute group
	//
	// - enable_ha_cluster: enable high availability (HA) for cluster
	//
	// - restart_node: restart node
	//
	// - backup: data backup
	//
	// - delete_backup: delete data backup
	//
	// - cancel_backup_task: cancel data backup
	//
	// - modify_timezone: modify system time zone
	//
	// - restore: data restoration
	//
	// - switch_az: switch primary and secondary zones
	//
	// - rollback_upgrade_version: roll back version upgrade
	//
	// - scale_out_fe: scale out FE
	//
	// - scale_in_fe: scale in FE
	//
	// - upgrade_fe_cu: upgrade FE CU specification
	//
	// - downgrade_fe_cu: downgrade FE CU specification
	//
	// - increase_fe_disk_size: increase FE disk size
	//
	// - decrease_fe_disk_size: decrease FE disk size
	//
	// - increase_fe_disk_number: increase FE disk count
	//
	// - decrease_fe_disk_number: decrease FE disk count
	//
	// - upgrade_fe_disk_performance_level: upgrade FE disk performance level
	//
	// - downgrade_fe_disk_performance_level: downgrade FE disk performance level
	//
	// - create_agent: create Agent
	//
	// - upgrade_agent_cu: upgrade Agent CU specification
	//
	// - scale_out_be: scale out BE
	//
	// - scale_in_be: scale in BE
	//
	// - upgrade_be_cu: upgrade BE CU specification
	//
	// - downgrade_be_cu: downgrade BE CU specification
	//
	// - increase_be_disk_size: increase BE disk size
	//
	// - decrease_be_disk_size: decrease BE disk size
	//
	// - increase_be_disk_number: increase BE disk count
	//
	// - decrease_be_disk_number: decrease BE disk count
	//
	// - upgrade_be_disk_performance_level: upgrade BE disk performance level
	//
	// - downgrade_be_disk_performance_level: downgrade BE disk performance level
	//
	// - upgrade_be_spec_type: upgrade BE specification type
	//
	// - downgrade_be_spec_type: downgrade BE specification type
	//
	// - scale_out_cn: scale out CN
	//
	// - scale_in_cn: scale in CN
	//
	// - upgrade_cn_cu: upgrade CN CU specification
	//
	// - downgrade_cn_cu: downgrade CN CU specification
	//
	// - increase_cn_disk_size: increase CN disk size
	//
	// - decrease_cn_disk_size: decrease CN disk size
	//
	// - increase_cn_disk_number: increase CN disk count
	//
	// - decrease_cn_disk_number: decrease CN disk count
	//
	// - upgrade_cn_disk_performance: upgrade CN disk performance level
	//
	// - downgrade_cn_disk_performance: downgrade CN disk performance level
	//
	// - upgrade_cn_spec_type: upgrade CN specification type
	//
	// - downgrade_cn_spec_type: downgrade CN specification type
	//
	// - elastic_scale_out_cn: elastically scale out CN
	//
	// - elastic_scale_in_cn: elastically scale in CN
	//
	// example:
	//
	// update_configuration
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// Page number of the current page. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page for paged queries. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start time of the operation.
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
