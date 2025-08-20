// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListStacksRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListStacksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStacksRequest
	GetPageSize() *int64
	SetParentStackId(v string) *ListStacksRequest
	GetParentStackId() *string
	SetRegionId(v string) *ListStacksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListStacksRequest
	GetResourceGroupId() *string
	SetShowNestedStack(v bool) *ListStacksRequest
	GetShowNestedStack() *bool
	SetStackId(v string) *ListStacksRequest
	GetStackId() *string
	SetStackIds(v []*string) *ListStacksRequest
	GetStackIds() []*string
	SetStackName(v []*string) *ListStacksRequest
	GetStackName() []*string
	SetStartTime(v string) *ListStacksRequest
	GetStartTime() *string
	SetStatus(v []*string) *ListStacksRequest
	GetStatus() []*string
	SetTag(v []*ListStacksRequestTag) *ListStacksRequest
	GetTag() []*ListStacksRequestTag
}

type ListStacksRequest struct {
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T15:16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	ParentStackId *string `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.\\
	//
	// For more information about resource groups, see the "Resource Group" section of the [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html) topic.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to return nested stacks. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// > If you specify ParentStackId, you must set ShowNestedStack to true.
	//
	// example:
	//
	// true
	ShowNestedStack *bool `json:"ShowNestedStack,omitempty" xml:"ShowNestedStack,omitempty"`
	// The stack ID. You can specify this parameter to query only the stack ID. If you want to query the detailed information about the stack, call the GetStack operation.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The IDs of the stacks.
	StackIds []*string `json:"StackIds,omitempty" xml:"StackIds,omitempty" type:"Repeated"`
	// The names of the stacks.
	//
	// example:
	//
	// MyStack
	StackName []*string `json:"StackName,omitempty" xml:"StackName,omitempty" type:"Repeated"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-01T15:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the stack.
	//
	// example:
	//
	// CREATE_COMPLETE
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags of the stack.
	Tag []*ListStacksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListStacksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStacksRequest) GoString() string {
	return s.String()
}

func (s *ListStacksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListStacksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStacksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStacksRequest) GetParentStackId() *string {
	return s.ParentStackId
}

func (s *ListStacksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStacksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListStacksRequest) GetShowNestedStack() *bool {
	return s.ShowNestedStack
}

func (s *ListStacksRequest) GetStackId() *string {
	return s.StackId
}

func (s *ListStacksRequest) GetStackIds() []*string {
	return s.StackIds
}

func (s *ListStacksRequest) GetStackName() []*string {
	return s.StackName
}

func (s *ListStacksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListStacksRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListStacksRequest) GetTag() []*ListStacksRequestTag {
	return s.Tag
}

func (s *ListStacksRequest) SetEndTime(v string) *ListStacksRequest {
	s.EndTime = &v
	return s
}

func (s *ListStacksRequest) SetPageNumber(v int64) *ListStacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStacksRequest) SetPageSize(v int64) *ListStacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListStacksRequest) SetParentStackId(v string) *ListStacksRequest {
	s.ParentStackId = &v
	return s
}

func (s *ListStacksRequest) SetRegionId(v string) *ListStacksRequest {
	s.RegionId = &v
	return s
}

func (s *ListStacksRequest) SetResourceGroupId(v string) *ListStacksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStacksRequest) SetShowNestedStack(v bool) *ListStacksRequest {
	s.ShowNestedStack = &v
	return s
}

func (s *ListStacksRequest) SetStackId(v string) *ListStacksRequest {
	s.StackId = &v
	return s
}

func (s *ListStacksRequest) SetStackIds(v []*string) *ListStacksRequest {
	s.StackIds = v
	return s
}

func (s *ListStacksRequest) SetStackName(v []*string) *ListStacksRequest {
	s.StackName = v
	return s
}

func (s *ListStacksRequest) SetStartTime(v string) *ListStacksRequest {
	s.StartTime = &v
	return s
}

func (s *ListStacksRequest) SetStatus(v []*string) *ListStacksRequest {
	s.Status = v
	return s
}

func (s *ListStacksRequest) SetTag(v []*ListStacksRequestTag) *ListStacksRequest {
	s.Tag = v
	return s
}

func (s *ListStacksRequest) Validate() error {
	return dara.Validate(s)
}

type ListStacksRequestTag struct {
	// The key of tag N.\\
	//
	// Valid values of N: 1 to 20.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.\\
	//
	// Valid values of N: 1 to 20.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStacksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListStacksRequestTag) GoString() string {
	return s.String()
}

func (s *ListStacksRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListStacksRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListStacksRequestTag) SetKey(v string) *ListStacksRequestTag {
	s.Key = &v
	return s
}

func (s *ListStacksRequestTag) SetValue(v string) *ListStacksRequestTag {
	s.Value = &v
	return s
}

func (s *ListStacksRequestTag) Validate() error {
	return dara.Validate(s)
}
