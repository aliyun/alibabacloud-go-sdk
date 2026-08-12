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
	// Details about access denied errors.
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Returned data.
	Data []*ListOperationHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Error code.
	//
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of records.
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
	// Value after the operation.
	//
	// example:
	//
	// FE enable = true
	AfterValue *string `json:"AfterValue,omitempty" xml:"AfterValue,omitempty"`
	// Value before the operation.
	//
	// example:
	//
	// FE enable = false
	BeforeValue *string `json:"BeforeValue,omitempty" xml:"BeforeValue,omitempty"`
	// Start time of the operation.
	//
	// example:
	//
	// 1742179008000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// End time of the operation.
	//
	// example:
	//
	// 1742179008000
	GmtEnd *int64 `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// c-cd7a3a6f2186d5c9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Operation details.
	//
	// example:
	//
	// 库存不足，操作失败，已退款
	OperationDetail *string `json:"OperationDetail,omitempty" xml:"OperationDetail,omitempty"`
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
	// upgrade_version
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// Operation progress.
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
