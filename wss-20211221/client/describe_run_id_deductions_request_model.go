// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRunIdDeductionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *DescribeRunIdDeductionsRequest
	GetAgentType() *string
	SetAgentTypes(v []*string) *DescribeRunIdDeductionsRequest
	GetAgentTypes() []*string
	SetAliUid(v int64) *DescribeRunIdDeductionsRequest
	GetAliUid() *int64
	SetBizType(v string) *DescribeRunIdDeductionsRequest
	GetBizType() *string
	SetConsumeWyId(v string) *DescribeRunIdDeductionsRequest
	GetConsumeWyId() *string
	SetConsumeWyIds(v []*string) *DescribeRunIdDeductionsRequest
	GetConsumeWyIds() []*string
	SetDeductionRequestId(v string) *DescribeRunIdDeductionsRequest
	GetDeductionRequestId() *string
	SetDeductionTypes(v []*string) *DescribeRunIdDeductionsRequest
	GetDeductionTypes() []*string
	SetEndTime(v int64) *DescribeRunIdDeductionsRequest
	GetEndTime() *int64
	SetGroupByFields(v []*string) *DescribeRunIdDeductionsRequest
	GetGroupByFields() []*string
	SetGroupResourceTypes(v []*string) *DescribeRunIdDeductionsRequest
	GetGroupResourceTypes() []*string
	SetGroupSeparator(v bool) *DescribeRunIdDeductionsRequest
	GetGroupSeparator() *bool
	SetInstanceId(v string) *DescribeRunIdDeductionsRequest
	GetInstanceId() *string
	SetInstanceIdType(v string) *DescribeRunIdDeductionsRequest
	GetInstanceIdType() *string
	SetInstanceIds(v []*string) *DescribeRunIdDeductionsRequest
	GetInstanceIds() []*string
	SetMaxResults(v int32) *DescribeRunIdDeductionsRequest
	GetMaxResults() *int32
	SetModel(v string) *DescribeRunIdDeductionsRequest
	GetModel() *string
	SetNextToken(v string) *DescribeRunIdDeductionsRequest
	GetNextToken() *string
	SetPackageIds(v []*string) *DescribeRunIdDeductionsRequest
	GetPackageIds() []*string
	SetPageNum(v int32) *DescribeRunIdDeductionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeRunIdDeductionsRequest
	GetPageSize() *int32
	SetResourceType(v string) *DescribeRunIdDeductionsRequest
	GetResourceType() *string
	SetResourceTypes(v []*string) *DescribeRunIdDeductionsRequest
	GetResourceTypes() []*string
	SetRunId(v string) *DescribeRunIdDeductionsRequest
	GetRunId() *string
	SetRunIdList(v []*string) *DescribeRunIdDeductionsRequest
	GetRunIdList() []*string
	SetStartTime(v int64) *DescribeRunIdDeductionsRequest
	GetStartTime() *int64
	SetWyId(v string) *DescribeRunIdDeductionsRequest
	GetWyId() *string
}

type DescribeRunIdDeductionsRequest struct {
	// The agent type. Valid values: CREDIT_PACKAGE, JVS_CLAW, OPEN_CLAW, and JVS_COPILOT.
	//
	// example:
	//
	// JVSCopilot、JVSClaw、OpenClaw
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The list of agent types.
	AgentTypes []*string `json:"AgentTypes,omitempty" xml:"AgentTypes,omitempty" type:"Repeated"`
	// The Alibaba Cloud UID.
	//
	// example:
	//
	// 1457450820614624
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The business type.
	//
	// example:
	//
	// ENTERPRISE、BUSINESS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The consumption WyId.
	//
	// example:
	//
	// 2892766fb8521756
	ConsumeWyId *string `json:"ConsumeWyId,omitempty" xml:"ConsumeWyId,omitempty"`
	// The Wuying user IDs. You can obtain the IDs by calling the GetLoginToken logon operation. The returned parameter is WyId.
	ConsumeWyIds []*string `json:"ConsumeWyIds,omitempty" xml:"ConsumeWyIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// uuid
	DeductionRequestId *string `json:"DeductionRequestId,omitempty" xml:"DeductionRequestId,omitempty"`
	// The deduction types. Do not specify this parameter for non-knowledge base scenarios.
	DeductionTypes []*string `json:"DeductionTypes,omitempty" xml:"DeductionTypes,omitempty" type:"Repeated"`
	// The end time of the query. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1785205179000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of additional grouping fields when aggregating by runId, such as model to distinguish between different models.
	GroupByFields []*string `json:"GroupByFields,omitempty" xml:"GroupByFields,omitempty" type:"Repeated"`
	// The list of grouped resource types.
	GroupResourceTypes []*string `json:"GroupResourceTypes,omitempty" xml:"GroupResourceTypes,omitempty" type:"Repeated"`
	// Specifies whether to group the results by deduction type.
	//
	// example:
	//
	// false
	GroupSeparator *bool `json:"GroupSeparator,omitempty" xml:"GroupSeparator,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// mdp-0c62ayep0nk4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance ID type. Do not specify this parameter for non-knowledge base scenarios.
	//
	// example:
	//
	// KnowledgeSpaceId、AgentId
	InstanceIdType *string `json:"InstanceIdType,omitempty" xml:"InstanceIdType,omitempty"`
	// The list of cloud computer IDs. If this field is specified, the PackageIds field is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The model name.
	//
	// example:
	//
	// qwen-3.8
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The pagination token. Leave this parameter empty for the first request. For subsequent requests, use the nextToken value returned in the previous response.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of core-hour package IDs in JSON format.
	PackageIds []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The Settings for the number of entries per page during paging in a paged query.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Enterprise Edition: CreditPackage, Business Edition: BusinessCreditPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of resource types in JSON array format.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The run ID.
	//
	// example:
	//
	// t-xxx
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The list of run IDs.
	RunIdList []*string `json:"RunIdList,omitempty" xml:"RunIdList,omitempty" type:"Repeated"`
	// The start time of the query. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1785205179000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The globally unique ID of the Wuying user.
	//
	// example:
	//
	// wy-123456****
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s DescribeRunIdDeductionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunIdDeductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRunIdDeductionsRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeRunIdDeductionsRequest) GetAgentTypes() []*string {
	return s.AgentTypes
}

func (s *DescribeRunIdDeductionsRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeRunIdDeductionsRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeRunIdDeductionsRequest) GetConsumeWyId() *string {
	return s.ConsumeWyId
}

func (s *DescribeRunIdDeductionsRequest) GetConsumeWyIds() []*string {
	return s.ConsumeWyIds
}

func (s *DescribeRunIdDeductionsRequest) GetDeductionRequestId() *string {
	return s.DeductionRequestId
}

func (s *DescribeRunIdDeductionsRequest) GetDeductionTypes() []*string {
	return s.DeductionTypes
}

func (s *DescribeRunIdDeductionsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeRunIdDeductionsRequest) GetGroupByFields() []*string {
	return s.GroupByFields
}

func (s *DescribeRunIdDeductionsRequest) GetGroupResourceTypes() []*string {
	return s.GroupResourceTypes
}

func (s *DescribeRunIdDeductionsRequest) GetGroupSeparator() *bool {
	return s.GroupSeparator
}

func (s *DescribeRunIdDeductionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRunIdDeductionsRequest) GetInstanceIdType() *string {
	return s.InstanceIdType
}

func (s *DescribeRunIdDeductionsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeRunIdDeductionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRunIdDeductionsRequest) GetModel() *string {
	return s.Model
}

func (s *DescribeRunIdDeductionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRunIdDeductionsRequest) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *DescribeRunIdDeductionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeRunIdDeductionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRunIdDeductionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRunIdDeductionsRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribeRunIdDeductionsRequest) GetRunId() *string {
	return s.RunId
}

func (s *DescribeRunIdDeductionsRequest) GetRunIdList() []*string {
	return s.RunIdList
}

func (s *DescribeRunIdDeductionsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeRunIdDeductionsRequest) GetWyId() *string {
	return s.WyId
}

func (s *DescribeRunIdDeductionsRequest) SetAgentType(v string) *DescribeRunIdDeductionsRequest {
	s.AgentType = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetAgentTypes(v []*string) *DescribeRunIdDeductionsRequest {
	s.AgentTypes = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetAliUid(v int64) *DescribeRunIdDeductionsRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetBizType(v string) *DescribeRunIdDeductionsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetConsumeWyId(v string) *DescribeRunIdDeductionsRequest {
	s.ConsumeWyId = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetConsumeWyIds(v []*string) *DescribeRunIdDeductionsRequest {
	s.ConsumeWyIds = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetDeductionRequestId(v string) *DescribeRunIdDeductionsRequest {
	s.DeductionRequestId = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetDeductionTypes(v []*string) *DescribeRunIdDeductionsRequest {
	s.DeductionTypes = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetEndTime(v int64) *DescribeRunIdDeductionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetGroupByFields(v []*string) *DescribeRunIdDeductionsRequest {
	s.GroupByFields = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetGroupResourceTypes(v []*string) *DescribeRunIdDeductionsRequest {
	s.GroupResourceTypes = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetGroupSeparator(v bool) *DescribeRunIdDeductionsRequest {
	s.GroupSeparator = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetInstanceId(v string) *DescribeRunIdDeductionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetInstanceIdType(v string) *DescribeRunIdDeductionsRequest {
	s.InstanceIdType = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetInstanceIds(v []*string) *DescribeRunIdDeductionsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetMaxResults(v int32) *DescribeRunIdDeductionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetModel(v string) *DescribeRunIdDeductionsRequest {
	s.Model = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetNextToken(v string) *DescribeRunIdDeductionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetPackageIds(v []*string) *DescribeRunIdDeductionsRequest {
	s.PackageIds = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetPageNum(v int32) *DescribeRunIdDeductionsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetPageSize(v int32) *DescribeRunIdDeductionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetResourceType(v string) *DescribeRunIdDeductionsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetResourceTypes(v []*string) *DescribeRunIdDeductionsRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetRunId(v string) *DescribeRunIdDeductionsRequest {
	s.RunId = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetRunIdList(v []*string) *DescribeRunIdDeductionsRequest {
	s.RunIdList = v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetStartTime(v int64) *DescribeRunIdDeductionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) SetWyId(v string) *DescribeRunIdDeductionsRequest {
	s.WyId = &v
	return s
}

func (s *DescribeRunIdDeductionsRequest) Validate() error {
	return dara.Validate(s)
}
