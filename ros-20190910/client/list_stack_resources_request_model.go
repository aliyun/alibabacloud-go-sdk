// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListStackResourcesRequest
	GetRegionId() *string
	SetStackId(v string) *ListStackResourcesRequest
	GetStackId() *string
}

type ListStackResourcesRequest struct {
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the region to which the stack belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ListStackResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListStackResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackResourcesRequest) GetStackId() *string {
	return s.StackId
}

func (s *ListStackResourcesRequest) SetRegionId(v string) *ListStackResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackResourcesRequest) SetStackId(v string) *ListStackResourcesRequest {
	s.StackId = &v
	return s
}

func (s *ListStackResourcesRequest) Validate() error {
	return dara.Validate(s)
}
