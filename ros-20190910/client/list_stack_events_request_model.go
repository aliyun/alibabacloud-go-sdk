// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogicalResourceId(v []*string) *ListStackEventsRequest
	GetLogicalResourceId() []*string
	SetPageNumber(v int64) *ListStackEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStackEventsRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListStackEventsRequest
	GetRegionId() *string
	SetResourceType(v []*string) *ListStackEventsRequest
	GetResourceType() []*string
	SetStackId(v string) *ListStackEventsRequest
	GetStackId() *string
	SetStatus(v []*string) *ListStackEventsRequest
	GetStatus() []*string
}

type ListStackEventsRequest struct {
	// The logical IDs of the resources.
	//
	// example:
	//
	// WebServer
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
	// The number of the page to return.\\
	//
	// Pages start from page 1.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\\
	//
	// Maximum value: 50.\\
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource types.
	//
	// example:
	//
	// ALIYUN::ECS::Instance
	ResourceType []*string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
	// The stack ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// CREATE_IN_PROGRESS
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListStackEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackEventsRequest) GoString() string {
	return s.String()
}

func (s *ListStackEventsRequest) GetLogicalResourceId() []*string {
	return s.LogicalResourceId
}

func (s *ListStackEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStackEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStackEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackEventsRequest) GetResourceType() []*string {
	return s.ResourceType
}

func (s *ListStackEventsRequest) GetStackId() *string {
	return s.StackId
}

func (s *ListStackEventsRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListStackEventsRequest) SetLogicalResourceId(v []*string) *ListStackEventsRequest {
	s.LogicalResourceId = v
	return s
}

func (s *ListStackEventsRequest) SetPageNumber(v int64) *ListStackEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackEventsRequest) SetPageSize(v int64) *ListStackEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackEventsRequest) SetRegionId(v string) *ListStackEventsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackEventsRequest) SetResourceType(v []*string) *ListStackEventsRequest {
	s.ResourceType = v
	return s
}

func (s *ListStackEventsRequest) SetStackId(v string) *ListStackEventsRequest {
	s.StackId = &v
	return s
}

func (s *ListStackEventsRequest) SetStatus(v []*string) *ListStackEventsRequest {
	s.Status = v
	return s
}

func (s *ListStackEventsRequest) Validate() error {
	return dara.Validate(s)
}
