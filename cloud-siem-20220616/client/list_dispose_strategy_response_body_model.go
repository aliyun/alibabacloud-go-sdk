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
	// The request status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request return value.
	//
	// example:
	//
	// 123456
	Data *ListDisposeStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request return message.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDisposeStrategyResponseBodyData struct {
	Groups []*ListDisposeStrategyResponseBodyDataGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
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

func (s *ListDisposeStrategyResponseBodyData) GetGroups() []*ListDisposeStrategyResponseBodyDataGroups {
	return s.Groups
}

func (s *ListDisposeStrategyResponseBodyData) GetPageInfo() *ListDisposeStrategyResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *ListDisposeStrategyResponseBodyData) GetResponseData() []*ListDisposeStrategyResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *ListDisposeStrategyResponseBodyData) SetGroups(v []*ListDisposeStrategyResponseBodyDataGroups) *ListDisposeStrategyResponseBodyData {
	s.Groups = v
	return s
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
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseData != nil {
		for _, item := range s.ResponseData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDisposeStrategyResponseBodyDataGroups struct {
	FailedCount         *int64                                              `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	FirstOccurrenceTime *int64                                              `json:"FirstOccurrenceTime,omitempty" xml:"FirstOccurrenceTime,omitempty"`
	GroupBy             *string                                             `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	GroupKey            *string                                             `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	GroupMeta           *ListDisposeStrategyResponseBodyDataGroupsGroupMeta `json:"GroupMeta,omitempty" xml:"GroupMeta,omitempty" type:"Struct"`
	GroupName           *string                                             `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupTitle          *string                                             `json:"GroupTitle,omitempty" xml:"GroupTitle,omitempty"`
	LastOccurrenceTime  *int64                                              `json:"LastOccurrenceTime,omitempty" xml:"LastOccurrenceTime,omitempty"`
	LatestModifiedTime  *int64                                              `json:"LatestModifiedTime,omitempty" xml:"LatestModifiedTime,omitempty"`
	RunningCount        *int64                                              `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	SuccessCount        *int64                                              `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	TotalCount          *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataGroups) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataGroups) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetFirstOccurrenceTime() *int64 {
	return s.FirstOccurrenceTime
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetGroupKey() *string {
	return s.GroupKey
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetGroupMeta() *ListDisposeStrategyResponseBodyDataGroupsGroupMeta {
	return s.GroupMeta
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetGroupTitle() *string {
	return s.GroupTitle
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetLastOccurrenceTime() *int64 {
	return s.LastOccurrenceTime
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetLatestModifiedTime() *int64 {
	return s.LatestModifiedTime
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetRunningCount() *int64 {
	return s.RunningCount
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *ListDisposeStrategyResponseBodyDataGroups) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetFailedCount(v int64) *ListDisposeStrategyResponseBodyDataGroups {
	s.FailedCount = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetFirstOccurrenceTime(v int64) *ListDisposeStrategyResponseBodyDataGroups {
	s.FirstOccurrenceTime = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetGroupBy(v string) *ListDisposeStrategyResponseBodyDataGroups {
	s.GroupBy = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetGroupKey(v string) *ListDisposeStrategyResponseBodyDataGroups {
	s.GroupKey = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetGroupMeta(v *ListDisposeStrategyResponseBodyDataGroupsGroupMeta) *ListDisposeStrategyResponseBodyDataGroups {
	s.GroupMeta = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetGroupName(v string) *ListDisposeStrategyResponseBodyDataGroups {
	s.GroupName = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetGroupTitle(v string) *ListDisposeStrategyResponseBodyDataGroups {
	s.GroupTitle = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetLastOccurrenceTime(v int64) *ListDisposeStrategyResponseBodyDataGroups {
	s.LastOccurrenceTime = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetLatestModifiedTime(v int64) *ListDisposeStrategyResponseBodyDataGroups {
	s.LatestModifiedTime = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetRunningCount(v int64) *ListDisposeStrategyResponseBodyDataGroups {
	s.RunningCount = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetSuccessCount(v int64) *ListDisposeStrategyResponseBodyDataGroups {
	s.SuccessCount = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) SetTotalCount(v int64) *ListDisposeStrategyResponseBodyDataGroups {
	s.TotalCount = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroups) Validate() error {
	if s.GroupMeta != nil {
		if err := s.GroupMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDisposeStrategyResponseBodyDataGroupsGroupMeta struct {
	GroupInfo interface{} `json:"GroupInfo,omitempty" xml:"GroupInfo,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataGroupsGroupMeta) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataGroupsGroupMeta) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataGroupsGroupMeta) GetGroupInfo() interface{} {
	return s.GroupInfo
}

func (s *ListDisposeStrategyResponseBodyDataGroupsGroupMeta) SetGroupInfo(v interface{}) *ListDisposeStrategyResponseBodyDataGroupsGroupMeta {
	s.GroupInfo = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataGroupsGroupMeta) Validate() error {
	return dara.Validate(s)
}

type ListDisposeStrategyResponseBodyDataPageInfo struct {
	// The current page number of the list.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of records returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
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
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The alert UUID.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The SIEM primary account ID associated with the policy.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The policy status. Valid values:
	//
	// example:
	//
	// 0
	EffectiveStatus *int32 `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	// The entity details in JSON array format.
	//
	// example:
	//
	// [{"ip":"1.1.1.1"}]
	Entity []interface{} `json:"Entity,omitempty" xml:"Entity,omitempty" type:"Repeated"`
	// The entity ID.
	//
	// example:
	//
	// 123456789
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity type. Valid values:
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The failure summary of the task.
	//
	// example:
	//
	// DisposalEntity failed which description is Aegis Quarantine File , return_info failed which description is Check Aegis Process Result , [ERROR DETAIL] *******.php:file not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The finish time of the task.
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
	// The modification time.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The incident name.
	//
	// example:
	//
	// Multiple type of alerts, including Miner Network, Command line download and run malicious files, Backdoor Process, etc
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The globally unique UUID of the incident.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The unique identifier name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The playbook type. Valid values:
	//
	// - system: manual handling
	//
	// - custom: event-triggered playbook
	//
	// - custom_alert: alert-triggered playbook
	//
	// - soar-manual: manually run playbook
	//
	// - soar-mdr: MDR-run playbook
	//
	// example:
	//
	// system
	PlaybookType *string `json:"PlaybookType,omitempty" xml:"PlaybookType,omitempty"`
	// The playbook UUID.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The disposition scope.
	//
	// example:
	//
	// [{ aliUid: 1766185894104675 }]
	Scope []interface{} `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The SOAR response policy ID.
	//
	// example:
	//
	// 577bbf90-a770-44a7-8154-586aa2d3****
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The playbook invocation status. Valid values:
	//
	// example:
	//
	// 10
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Alibaba Cloud account ID that configured the policy.
	//
	// example:
	//
	// 176555323***
	SubAliuid *int64 `json:"SubAliuid,omitempty" xml:"SubAliuid,omitempty"`
	// The playbook trigger parameters in JSON format.
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
	// The playbook URL.
	//
	// example:
	//
	// {"playbookUuid":"system_aliyun_aegis_stop_container_book","requestUuid":"e8924356-448b-4301-aee9-*******"}
	TaskUrl *string `json:"TaskUrl,omitempty" xml:"TaskUrl,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetAlertName() *string {
	return s.AlertName
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

func (s *ListDisposeStrategyResponseBodyDataResponseData) GetErrorCode() *string {
	return s.ErrorCode
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

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetAlertName(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.AlertName = &v
	return s
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

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetErrorCode(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.ErrorCode = &v
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
