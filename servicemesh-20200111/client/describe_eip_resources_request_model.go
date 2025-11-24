// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *DescribeEipResourcesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeEipResourcesRequest
	GetPageSize() *int32
	SetServiceMeshId(v string) *DescribeEipResourcesRequest
	GetServiceMeshId() *string
}

type DescribeEipResourcesRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the ASM instance.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeEipResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeEipResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEipResourcesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeEipResourcesRequest) SetPageNum(v int32) *DescribeEipResourcesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeEipResourcesRequest) SetPageSize(v int32) *DescribeEipResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEipResourcesRequest) SetServiceMeshId(v string) *DescribeEipResourcesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeEipResourcesRequest) Validate() error {
	return dara.Validate(s)
}
