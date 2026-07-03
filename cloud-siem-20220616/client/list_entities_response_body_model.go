// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListEntitiesResponseBody
	GetCode() *int32
	SetData(v *ListEntitiesResponseBodyData) *ListEntitiesResponseBody
	GetData() *ListEntitiesResponseBodyData
	SetMessage(v string) *ListEntitiesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEntitiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEntitiesResponseBody
	GetSuccess() *bool
}

type ListEntitiesResponseBody struct {
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// 123456
	Data *ListEntitiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
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
	// - true: successful.
	//
	// - false: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListEntitiesResponseBody) GetData() *ListEntitiesResponseBodyData {
	return s.Data
}

func (s *ListEntitiesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEntitiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEntitiesResponseBody) SetCode(v int32) *ListEntitiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEntitiesResponseBody) SetData(v *ListEntitiesResponseBodyData) *ListEntitiesResponseBody {
	s.Data = v
	return s
}

func (s *ListEntitiesResponseBody) SetMessage(v string) *ListEntitiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEntitiesResponseBody) SetRequestId(v string) *ListEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesResponseBody) SetSuccess(v bool) *ListEntitiesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEntitiesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEntitiesResponseBodyData struct {
	// The pagination information.
	PageInfo *ListEntitiesResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListEntitiesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListEntitiesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyData) GetPageInfo() *ListEntitiesResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *ListEntitiesResponseBodyData) GetResponseData() []*ListEntitiesResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *ListEntitiesResponseBodyData) SetPageInfo(v *ListEntitiesResponseBodyDataPageInfo) *ListEntitiesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListEntitiesResponseBodyData) SetResponseData(v []*ListEntitiesResponseBodyDataResponseData) *ListEntitiesResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *ListEntitiesResponseBodyData) Validate() error {
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

type ListEntitiesResponseBodyDataPageInfo struct {
	// The page number of the current page.
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
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEntitiesResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListEntitiesResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEntitiesResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEntitiesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListEntitiesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListEntitiesResponseBodyDataPageInfo) SetPageSize(v int32) *ListEntitiesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListEntitiesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListEntitiesResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListEntitiesResponseBodyDataResponseData struct {
	AgentAnalysisReason *string `json:"AgentAnalysisReason,omitempty" xml:"AgentAnalysisReason,omitempty"`
	// The confidence level of the entity as determined by the agent.
	//
	// example:
	//
	// 85
	AgentConfidence *string `json:"AgentConfidence,omitempty" xml:"AgentConfidence,omitempty"`
	// The disposal method recommended by the agent.
	//
	// example:
	//
	// delete_file
	AgentDisposalMethod *string `json:"AgentDisposalMethod,omitempty" xml:"AgentDisposalMethod,omitempty"`
	// The UUID of the playbook recommended by the agent for disposal.
	//
	// example:
	//
	// 12XAD-SFQ-WAF-2ca2
	AgentDisposalPlaybookUuid *string `json:"AgentDisposalPlaybookUuid,omitempty" xml:"AgentDisposalPlaybookUuid,omitempty"`
	AgentDisposalReason       *string `json:"AgentDisposalReason,omitempty" xml:"AgentDisposalReason,omitempty"`
	// The disposal suggestion recommended by the agent.
	//
	// example:
	//
	// {}
	AgentDisposalSuggestion *string `json:"AgentDisposalSuggestion,omitempty" xml:"AgentDisposalSuggestion,omitempty"`
	// The list of disposal suggestions recommended by the agent.
	AgentDisposes []*ListEntitiesResponseBodyDataResponseDataAgentDisposes `json:"AgentDisposes,omitempty" xml:"AgentDisposes,omitempty" type:"Repeated"`
	// The number of alerts associated with the entity.
	//
	// example:
	//
	// 1
	AlertNum *int32 `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	// The alert UUID.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 123456789****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The code of the cloud service provider from which the entity originates. Valid values:
	//
	// - aliyun: Alibaba Cloud
	//
	// - qcloud: Tencent Cloud
	//
	// - hcloud: Huawei Cloud.
	//
	// example:
	//
	// aliyun
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The logical ID of the entity.
	//
	// example:
	//
	// 12345****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The display information of the entity in JSON format.
	//
	// example:
	//
	// {"file_path": "c:/www/leixi.jsp","file_hash": "aa0ca926ad948cd820e0a3d9a18c****","host_uuid": "efed2cf7-0b77-45d9-a97b-d2cf246b****","malware_type": "${aliyun.siem.sas.alert_tag.webshell}","host_name": "launch-advisor-2023****"}
	EntityInfo *string `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	// The entity name.
	//
	// example:
	//
	// 123.123.123.123
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The entity type. Valid values:
	//
	// - ip: IP address
	//
	// - domain: domain name
	//
	// - url: URL
	//
	// - process: process
	//
	// - file: file
	//
	// - host: host
	//
	// - cloud_account: cloud account
	//
	// - container: container
	//
	// - bucket: Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity UUID.
	//
	// example:
	//
	// 8087b3e4aa6862852c100c8738cf****
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
	// The number of events associated with the entity.
	//
	// example:
	//
	// 1
	EventNum *int32 `json:"EventNum,omitempty" xml:"EventNum,omitempty"`
	// The time when the entity was collected.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the entity was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The entity ID.
	//
	// example:
	//
	// 123456789***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The incident UUID. You can obtain this value from the incident list operation.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Indicates whether the entity is an asset. Valid values:
	//
	// + 0: No.
	//
	// + 1: Yes.
	//
	// example:
	//
	// 1
	IsAsset *string `json:"IsAsset,omitempty" xml:"IsAsset,omitempty"`
	// Indicates whether the entity is malicious. Valid values:
	//
	// + 0: No.
	//
	// + 1: Yes.
	//
	// example:
	//
	// 0
	IsMalware *string `json:"IsMalware,omitempty" xml:"IsMalware,omitempty"`
	// The malware type of the entity.
	//
	// example:
	//
	// aliyun.siem.sas.alert_tag.webshell
	MalwareType *string `json:"MalwareType,omitempty" xml:"MalwareType,omitempty"`
	// The linked account ID associated with the entity.
	//
	// example:
	//
	// 113091674488****
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The entity tags. The value is a JSON array string:
	//
	// `"[{"tagKey1":"tagValue1"},{"tagKey2":"tagValue2"}]"`.
	//
	// example:
	//
	// [{"tagKey1":"tagValue1"},{"tagKey2":"tagValue2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListEntitiesResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAgentAnalysisReason() *string {
	return s.AgentAnalysisReason
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAgentConfidence() *string {
	return s.AgentConfidence
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAgentDisposalMethod() *string {
	return s.AgentDisposalMethod
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAgentDisposalPlaybookUuid() *string {
	return s.AgentDisposalPlaybookUuid
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAgentDisposalReason() *string {
	return s.AgentDisposalReason
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAgentDisposalSuggestion() *string {
	return s.AgentDisposalSuggestion
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAgentDisposes() []*ListEntitiesResponseBodyDataResponseDataAgentDisposes {
	return s.AgentDisposes
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAlertNum() *int32 {
	return s.AlertNum
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *ListEntitiesResponseBodyDataResponseData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *ListEntitiesResponseBodyDataResponseData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListEntitiesResponseBodyDataResponseData) GetEntityId() *string {
	return s.EntityId
}

func (s *ListEntitiesResponseBodyDataResponseData) GetEntityInfo() *string {
	return s.EntityInfo
}

func (s *ListEntitiesResponseBodyDataResponseData) GetEntityName() *string {
	return s.EntityName
}

func (s *ListEntitiesResponseBodyDataResponseData) GetEntityType() *string {
	return s.EntityType
}

func (s *ListEntitiesResponseBodyDataResponseData) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *ListEntitiesResponseBodyDataResponseData) GetEventNum() *int32 {
	return s.EventNum
}

func (s *ListEntitiesResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListEntitiesResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListEntitiesResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *ListEntitiesResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *ListEntitiesResponseBodyDataResponseData) GetIsAsset() *string {
	return s.IsAsset
}

func (s *ListEntitiesResponseBodyDataResponseData) GetIsMalware() *string {
	return s.IsMalware
}

func (s *ListEntitiesResponseBodyDataResponseData) GetMalwareType() *string {
	return s.MalwareType
}

func (s *ListEntitiesResponseBodyDataResponseData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListEntitiesResponseBodyDataResponseData) GetTags() *string {
	return s.Tags
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAgentAnalysisReason(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AgentAnalysisReason = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAgentConfidence(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AgentConfidence = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAgentDisposalMethod(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AgentDisposalMethod = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAgentDisposalPlaybookUuid(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AgentDisposalPlaybookUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAgentDisposalReason(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AgentDisposalReason = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAgentDisposalSuggestion(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AgentDisposalSuggestion = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAgentDisposes(v []*ListEntitiesResponseBodyDataResponseDataAgentDisposes) *ListEntitiesResponseBodyDataResponseData {
	s.AgentDisposes = v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAlertNum(v int32) *ListEntitiesResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAlertUuid(v string) *ListEntitiesResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetAliuid(v int64) *ListEntitiesResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetCloudCode(v string) *ListEntitiesResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityId(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityInfo(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityInfo = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityName(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityName = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityType(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityType = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEntityUuid(v string) *ListEntitiesResponseBodyDataResponseData {
	s.EntityUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetEventNum(v int32) *ListEntitiesResponseBodyDataResponseData {
	s.EventNum = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetGmtCreate(v string) *ListEntitiesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetGmtModified(v string) *ListEntitiesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetId(v int64) *ListEntitiesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetIncidentUuid(v string) *ListEntitiesResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetIsAsset(v string) *ListEntitiesResponseBodyDataResponseData {
	s.IsAsset = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetIsMalware(v string) *ListEntitiesResponseBodyDataResponseData {
	s.IsMalware = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetMalwareType(v string) *ListEntitiesResponseBodyDataResponseData {
	s.MalwareType = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetSubUserId(v int64) *ListEntitiesResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) SetTags(v string) *ListEntitiesResponseBodyDataResponseData {
	s.Tags = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseData) Validate() error {
	if s.AgentDisposes != nil {
		for _, item := range s.AgentDisposes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEntitiesResponseBodyDataResponseDataAgentDisposes struct {
	// The disposal suggestion recommended by the agent.
	//
	// example:
	//
	// {}
	AgentDisposalMethod *string `json:"AgentDisposalMethod,omitempty" xml:"AgentDisposalMethod,omitempty"`
	// The UUID of the playbook recommended by the agent for disposal.
	//
	// example:
	//
	// 12XAD-SFQ-WAF-2ca2
	AgentDisposalPlaybookUuid *string `json:"AgentDisposalPlaybookUuid,omitempty" xml:"AgentDisposalPlaybookUuid,omitempty"`
	AgentDisposalReason       *string `json:"AgentDisposalReason,omitempty" xml:"AgentDisposalReason,omitempty"`
}

func (s ListEntitiesResponseBodyDataResponseDataAgentDisposes) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesResponseBodyDataResponseDataAgentDisposes) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyDataResponseDataAgentDisposes) GetAgentDisposalMethod() *string {
	return s.AgentDisposalMethod
}

func (s *ListEntitiesResponseBodyDataResponseDataAgentDisposes) GetAgentDisposalPlaybookUuid() *string {
	return s.AgentDisposalPlaybookUuid
}

func (s *ListEntitiesResponseBodyDataResponseDataAgentDisposes) GetAgentDisposalReason() *string {
	return s.AgentDisposalReason
}

func (s *ListEntitiesResponseBodyDataResponseDataAgentDisposes) SetAgentDisposalMethod(v string) *ListEntitiesResponseBodyDataResponseDataAgentDisposes {
	s.AgentDisposalMethod = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseDataAgentDisposes) SetAgentDisposalPlaybookUuid(v string) *ListEntitiesResponseBodyDataResponseDataAgentDisposes {
	s.AgentDisposalPlaybookUuid = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseDataAgentDisposes) SetAgentDisposalReason(v string) *ListEntitiesResponseBodyDataResponseDataAgentDisposes {
	s.AgentDisposalReason = &v
	return s
}

func (s *ListEntitiesResponseBodyDataResponseDataAgentDisposes) Validate() error {
	return dara.Validate(s)
}
