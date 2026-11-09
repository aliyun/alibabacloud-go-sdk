// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageDeductionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *DescribePackageDeductionsRequest
	GetAgentType() *string
	SetAgentTypes(v []*string) *DescribePackageDeductionsRequest
	GetAgentTypes() []*string
	SetBizType(v string) *DescribePackageDeductionsRequest
	GetBizType() *string
	SetConsumeWyId(v string) *DescribePackageDeductionsRequest
	GetConsumeWyId() *string
	SetConsumeWyIds(v []*string) *DescribePackageDeductionsRequest
	GetConsumeWyIds() []*string
	SetDeductionRequestId(v string) *DescribePackageDeductionsRequest
	GetDeductionRequestId() *string
	SetDeductionTypes(v []*string) *DescribePackageDeductionsRequest
	GetDeductionTypes() []*string
	SetEndTime(v int64) *DescribePackageDeductionsRequest
	GetEndTime() *int64
	SetGroupResourceTypes(v []*string) *DescribePackageDeductionsRequest
	GetGroupResourceTypes() []*string
	SetInstanceId(v string) *DescribePackageDeductionsRequest
	GetInstanceId() *string
	SetInstanceIdType(v string) *DescribePackageDeductionsRequest
	GetInstanceIdType() *string
	SetInstanceIds(v []*string) *DescribePackageDeductionsRequest
	GetInstanceIds() []*string
	SetModel(v string) *DescribePackageDeductionsRequest
	GetModel() *string
	SetPackageIds(v []*string) *DescribePackageDeductionsRequest
	GetPackageIds() []*string
	SetPageNum(v int32) *DescribePackageDeductionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribePackageDeductionsRequest
	GetPageSize() *int32
	SetResourceType(v string) *DescribePackageDeductionsRequest
	GetResourceType() *string
	SetResourceTypes(v []*string) *DescribePackageDeductionsRequest
	GetResourceTypes() []*string
	SetRunId(v string) *DescribePackageDeductionsRequest
	GetRunId() *string
	SetRunIdList(v []*string) *DescribePackageDeductionsRequest
	GetRunIdList() []*string
	SetStartTime(v int64) *DescribePackageDeductionsRequest
	GetStartTime() *int64
	SetWyId(v string) *DescribePackageDeductionsRequest
	GetWyId() *string
}

type DescribePackageDeductionsRequest struct {
	// The agent type.
	//
	// example:
	//
	// OpenClaw
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The list of agent types.
	AgentTypes []*string `json:"AgentTypes,omitempty" xml:"AgentTypes,omitempty" type:"Repeated"`
	// The business type.
	//
	// example:
	//
	// ENTERPRISE、BUSINESS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The WUYING user ID. You can obtain it by calling the GetLoginToken operation. The response parameter is WyId.
	//
	// example:
	//
	// fb4dc579c48ffbea
	ConsumeWyId *string `json:"ConsumeWyId,omitempty" xml:"ConsumeWyId,omitempty"`
	// The WUYING user ID. You can obtain it by calling the GetLoginToken operation. The response parameter is WyId.
	ConsumeWyIds []*string `json:"ConsumeWyIds,omitempty" xml:"ConsumeWyIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// UUID
	DeductionRequestId *string `json:"DeductionRequestId,omitempty" xml:"DeductionRequestId,omitempty"`
	// The deduction types.
	DeductionTypes []*string `json:"DeductionTypes,omitempty" xml:"DeductionTypes,omitempty" type:"Repeated"`
	// The end time of the query. The time is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1709117383000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of resource group types.
	GroupResourceTypes []*string `json:"GroupResourceTypes,omitempty" xml:"GroupResourceTypes,omitempty" type:"Repeated"`
	// The instance ID, such as ecd-xxx.
	//
	// example:
	//
	// ecd-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance ID type.
	//
	// example:
	//
	// ecd-xxx
	InstanceIdType *string `json:"InstanceIdType,omitempty" xml:"InstanceIdType,omitempty"`
	// The list of Cloud Desktop IDs. If this field is specified, the PackageIds field is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The model name.
	//
	// example:
	//
	// happyhorse-1.1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The list of core-hour package IDs in JSON format.
	PackageIds []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The paging Settings for the number of entries to return on each page in a paged query. Maximum value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The array of resource types.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The run ID.
	//
	// example:
	//
	// t-xxx
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The list of run IDs.
	RunIdList []*string `json:"RunIdList,omitempty" xml:"RunIdList,omitempty" type:"Repeated"`
	// The start time of the query. The time is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1709117383000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The WUYING user ID. You can obtain it by calling the GetLoginToken operation. The response parameter is WyId.
	//
	// example:
	//
	// fb4dc579c48ffbea
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s DescribePackageDeductionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageDeductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribePackageDeductionsRequest) GetAgentTypes() []*string {
	return s.AgentTypes
}

func (s *DescribePackageDeductionsRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribePackageDeductionsRequest) GetConsumeWyId() *string {
	return s.ConsumeWyId
}

func (s *DescribePackageDeductionsRequest) GetConsumeWyIds() []*string {
	return s.ConsumeWyIds
}

func (s *DescribePackageDeductionsRequest) GetDeductionRequestId() *string {
	return s.DeductionRequestId
}

func (s *DescribePackageDeductionsRequest) GetDeductionTypes() []*string {
	return s.DeductionTypes
}

func (s *DescribePackageDeductionsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePackageDeductionsRequest) GetGroupResourceTypes() []*string {
	return s.GroupResourceTypes
}

func (s *DescribePackageDeductionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePackageDeductionsRequest) GetInstanceIdType() *string {
	return s.InstanceIdType
}

func (s *DescribePackageDeductionsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePackageDeductionsRequest) GetModel() *string {
	return s.Model
}

func (s *DescribePackageDeductionsRequest) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *DescribePackageDeductionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribePackageDeductionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePackageDeductionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePackageDeductionsRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribePackageDeductionsRequest) GetRunId() *string {
	return s.RunId
}

func (s *DescribePackageDeductionsRequest) GetRunIdList() []*string {
	return s.RunIdList
}

func (s *DescribePackageDeductionsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePackageDeductionsRequest) GetWyId() *string {
	return s.WyId
}

func (s *DescribePackageDeductionsRequest) SetAgentType(v string) *DescribePackageDeductionsRequest {
	s.AgentType = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetAgentTypes(v []*string) *DescribePackageDeductionsRequest {
	s.AgentTypes = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetBizType(v string) *DescribePackageDeductionsRequest {
	s.BizType = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetConsumeWyId(v string) *DescribePackageDeductionsRequest {
	s.ConsumeWyId = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetConsumeWyIds(v []*string) *DescribePackageDeductionsRequest {
	s.ConsumeWyIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetDeductionRequestId(v string) *DescribePackageDeductionsRequest {
	s.DeductionRequestId = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetDeductionTypes(v []*string) *DescribePackageDeductionsRequest {
	s.DeductionTypes = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetEndTime(v int64) *DescribePackageDeductionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetGroupResourceTypes(v []*string) *DescribePackageDeductionsRequest {
	s.GroupResourceTypes = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetInstanceId(v string) *DescribePackageDeductionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetInstanceIdType(v string) *DescribePackageDeductionsRequest {
	s.InstanceIdType = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetInstanceIds(v []*string) *DescribePackageDeductionsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetModel(v string) *DescribePackageDeductionsRequest {
	s.Model = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPackageIds(v []*string) *DescribePackageDeductionsRequest {
	s.PackageIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageNum(v int32) *DescribePackageDeductionsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageSize(v int32) *DescribePackageDeductionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetResourceType(v string) *DescribePackageDeductionsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetResourceTypes(v []*string) *DescribePackageDeductionsRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetRunId(v string) *DescribePackageDeductionsRequest {
	s.RunId = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetRunIdList(v []*string) *DescribePackageDeductionsRequest {
	s.RunIdList = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetStartTime(v int64) *DescribePackageDeductionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetWyId(v string) *DescribePackageDeductionsRequest {
	s.WyId = &v
	return s
}

func (s *DescribePackageDeductionsRequest) Validate() error {
	return dara.Validate(s)
}
