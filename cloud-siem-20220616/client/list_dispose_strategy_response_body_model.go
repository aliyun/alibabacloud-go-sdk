// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisposeStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDisposeStrategyResponseBody
	GetCode() *int32
	SetData(v *ListDisposeStrategyResponseBodyData) *ListDisposeStrategyResponseBody
	GetData() *ListDisposeStrategyResponseBodyData
	SetMessage(v string) *ListDisposeStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDisposeStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDisposeStrategyResponseBody
	GetSuccess() *bool
}

type ListDisposeStrategyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *ListDisposeStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDisposeStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDisposeStrategyResponseBody) GetData() *ListDisposeStrategyResponseBodyData {
	return s.Data
}

func (s *ListDisposeStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDisposeStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDisposeStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDisposeStrategyResponseBody) SetCode(v int32) *ListDisposeStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetData(v *ListDisposeStrategyResponseBodyData) *ListDisposeStrategyResponseBody {
	s.Data = v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetMessage(v string) *ListDisposeStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetRequestId(v string) *ListDisposeStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetSuccess(v bool) *ListDisposeStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDisposeStrategyResponseBodyData struct {
	// The pagination information.
	PageInfo *ListDisposeStrategyResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListDisposeStrategyResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListDisposeStrategyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyData) GetPageInfo() *ListDisposeStrategyResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *ListDisposeStrategyResponseBodyData) GetResponseData() []*ListDisposeStrategyResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *ListDisposeStrategyResponseBodyData) SetPageInfo(v *ListDisposeStrategyResponseBodyDataPageInfo) *ListDisposeStrategyResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListDisposeStrategyResponseBodyData) SetResponseData(v []*ListDisposeStrategyResponseBodyDataResponseData) *ListDisposeStrategyResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *ListDisposeStrategyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDisposeStrategyResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetPageSize(v int32) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetTotalCount(v int64) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListDisposeStrategyResponseBodyDataResponseData struct {
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the policy in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- 0: invalid
	//
	// 	- 1: valid
	//
	// example:
	//
	// 0
	EffectiveStatus *int32 `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	// The details of the entity. The value is a JSON array.
	//
	// example:
	//
	// [{"ip":"1.1.1.1"}]
	Entity []interface{} `json:"Entity,omitempty" xml:"Entity,omitempty" type:"Repeated"`
	// The ID of the entity.
	//
	// example:
	//
	// 123456789
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. Valid values:
	//
	// 	- ip
	//
	// 	- process
	//
	// 	- file
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The summary information about the failed task.
	//
	// example:
	//
	// DisposalEntity failed which description is Aegis Quarantine File , return_info failed which description is Check Aegis Process Result , [ERROR DETAIL] *******.php:file not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2021-08-10 21:34:07
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The name of the playbook, which is the unique identifier of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- system: user-triggered playbook
	//
	// 	- custom: event-triggered playbook
	//
	// 	- custom_alert: alert-triggered playbook
	//
	// 	- soar-manual: user-run playbook
	//
	// 	- soar-mdr: MDR-run playbook
	//
	// example:
	//
	// system
	PlaybookType *string `json:"PlaybookType,omitempty" xml:"PlaybookType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The scope of the policy.
	//
	// example:
	//
	// [{ aliUid: 1766185894104675 }]
	Scope []interface{} `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The ID of the SOAR handling policy.
	//
	// example:
	//
	// 577bbf90-a770-44a7-8154-586aa2d318fa
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The running status of the playbook. Valid values:
	//
	// 	- 200: successful
	//
	// 	- 10: deleted
	//
	// 	- 5: failed
	//
	// 	- 0: initial
	//
	// example:
	//
	// 10
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba account that is used to configure the policy.
	//
	// example:
	//
	// 176555323***
	SubAliuid *int64 `json:"SubAliuid,omitempty" xml:"SubAliuid,omitempty"`
	// The parameters that are used to trigger the playbook. The value is in the JSON format.
	//
	// example:
	//
	// {
	//
	//       "file": {
	//
	//             "op_code": "2",
	//
	//             "file_path": "/root/alert0913/a886.jsp",
	//
	//             "entity_type": "file",
	//
	//             "entity_name": "a886.jsp",
	//
	//             "file_name": "a886.jsp",
	//
	//             "file_owner": "USER:,GROUP:",
	//
	//             "hash_value": "5def10c9a4287d0920d86b42420b20b0",
	//
	//             "op_level": "2",
	//
	//             "entity_id": "/root/alert0913/a886.jsp",
	//
	//             "host_uuid": {
	//
	//                   "entity_type": "host",
	//
	//                   "entity_name": "N/A",
	//
	//                   "is_comprised": "1",
	//
	//                   "os_type": "linux",
	//
	//                   "entity_id": "5f58ef67-8803-4314-8d67-c87dc92b****",
	//
	//                   "host_uuid": "5f58ef67-8803-4314-8d67-c87dc92b****",
	//
	//                   "host_name": "N/A"
	//
	//             },
	//
	//             "malware_type": "${aliyun.siem.sas.alert_tag.webshell}"
	//
	//       },
	//
	//       "_sys_siem": {
	//
	//             "cloudCode": "aliyun",
	//
	//             "alertId": "89416745494****"
	//
	//       },
	//
	//       "scope": [
	//
	//             {
	//
	//                   "aliUid": 1766185894104****
	//
	//             }
	//
	//       ]
	//
	// }
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	TaskUrl   *string `json:"TaskUrl,omitempty" xml:"TaskUrl,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetEffectiveStatus() *int32 {
	return s.EffectiveStatus
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetEntity() []interface{} {
	return s.Entity
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetEntityType() *string {
	return s.EntityType
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetIncidentName() *string {
	return s.IncidentName
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetPlaybookType() *string {
	return s.PlaybookType
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetScope() []interface{} {
	return s.Scope
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetSophonTaskId() *string {
	return s.SophonTaskId
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetStatus() *int32 {
	return s.Status
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetSubAliuid() *int64 {
	return s.SubAliuid
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetTaskParam() *string {
	return s.TaskParam
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetTaskUrl() *string {
	return s.TaskUrl
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetAlertUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetAliuid(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEffectiveStatus(v int32) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EffectiveStatus = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntity(v []interface{}) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Entity = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntityId(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntityType(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EntityType = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetErrorMessage(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetFinishTime(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.FinishTime = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetGmtCreate(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetGmtModified(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetId(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetIncidentName(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.IncidentName = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetIncidentUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookName(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookName = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookType(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookType = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetScope(v []interface{}) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Scope = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetSophonTaskId(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.SophonTaskId = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetStatus(v int32) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetSubAliuid(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.SubAliuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetTaskParam(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.TaskParam = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetTaskUrl(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.TaskUrl = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}
