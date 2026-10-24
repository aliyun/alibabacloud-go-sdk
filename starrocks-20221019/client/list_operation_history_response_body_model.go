// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListOperationHistoryResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*ListOperationHistoryResponseBodyData) *ListOperationHistoryResponseBody
	GetData() []*ListOperationHistoryResponseBodyData
	SetErrCode(v string) *ListOperationHistoryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListOperationHistoryResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListOperationHistoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListOperationHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOperationHistoryResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListOperationHistoryResponseBody
	GetTotal() *int32
}

type ListOperationHistoryResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The returned data.
	Data []*ListOperationHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 832
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOperationHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationHistoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListOperationHistoryResponseBody) GetData() []*ListOperationHistoryResponseBodyData {
	return s.Data
}

func (s *ListOperationHistoryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListOperationHistoryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListOperationHistoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOperationHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOperationHistoryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListOperationHistoryResponseBody) SetAccessDeniedDetail(v string) *ListOperationHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetData(v []*ListOperationHistoryResponseBodyData) *ListOperationHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListOperationHistoryResponseBody) SetErrCode(v string) *ListOperationHistoryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetErrMessage(v string) *ListOperationHistoryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetHttpStatusCode(v int32) *ListOperationHistoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetRequestId(v string) *ListOperationHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetSuccess(v bool) *ListOperationHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListOperationHistoryResponseBody) SetTotal(v int32) *ListOperationHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *ListOperationHistoryResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationHistoryResponseBodyData struct {
	// The value after the operation.
	//
	// example:
	//
	// FE enable = true
	AfterValue *string `json:"AfterValue,omitempty" xml:"AfterValue,omitempty"`
	// The value before the operation.
	//
	// example:
	//
	// FE enable = false
	BeforeValue *string `json:"BeforeValue,omitempty" xml:"BeforeValue,omitempty"`
	// The start time of the operation.
	//
	// example:
	//
	// 1742179008000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The end time of the operation.
	//
	// example:
	//
	// 1742179008000
	GmtEnd *int64 `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// c-cd7a3a6f2186d5c9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation details.
	//
	// example:
	//
	// Insufficient inventory. The operation failed and a refund has been issued
	OperationDetail *string `json:"OperationDetail,omitempty" xml:"OperationDetail,omitempty"`
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
	// upgrade_version
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The operation progress.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s ListOperationHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOperationHistoryResponseBodyData) GetAfterValue() *string {
	return s.AfterValue
}

func (s *ListOperationHistoryResponseBodyData) GetBeforeValue() *string {
	return s.BeforeValue
}

func (s *ListOperationHistoryResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListOperationHistoryResponseBodyData) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *ListOperationHistoryResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationHistoryResponseBodyData) GetOperationDetail() *string {
	return s.OperationDetail
}

func (s *ListOperationHistoryResponseBodyData) GetOperationId() *string {
	return s.OperationId
}

func (s *ListOperationHistoryResponseBodyData) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *ListOperationHistoryResponseBodyData) GetOperationType() *string {
	return s.OperationType
}

func (s *ListOperationHistoryResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *ListOperationHistoryResponseBodyData) SetAfterValue(v string) *ListOperationHistoryResponseBodyData {
	s.AfterValue = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetBeforeValue(v string) *ListOperationHistoryResponseBodyData {
	s.BeforeValue = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetGmtCreate(v int64) *ListOperationHistoryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetGmtEnd(v int64) *ListOperationHistoryResponseBodyData {
	s.GmtEnd = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetInstanceId(v string) *ListOperationHistoryResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationDetail(v string) *ListOperationHistoryResponseBodyData {
	s.OperationDetail = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationId(v string) *ListOperationHistoryResponseBodyData {
	s.OperationId = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationStatus(v string) *ListOperationHistoryResponseBodyData {
	s.OperationStatus = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetOperationType(v string) *ListOperationHistoryResponseBodyData {
	s.OperationType = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) SetProgress(v int32) *ListOperationHistoryResponseBodyData {
	s.Progress = &v
	return s
}

func (s *ListOperationHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
