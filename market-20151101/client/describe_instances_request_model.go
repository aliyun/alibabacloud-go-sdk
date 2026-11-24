// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodes(v string) *DescribeInstancesRequest
	GetCodes() *string
	SetExceptCodes(v string) *DescribeInstancesRequest
	GetExceptCodes() *string
	SetPageNumber(v int32) *DescribeInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesRequest
	GetPageSize() *int32
	SetProductType(v string) *DescribeInstancesRequest
	GetProductType() *string
}

type DescribeInstancesRequest struct {
	// The commodity codes. Separate multiple codes with commas (,).
	//
	// example:
	//
	// cmgj000112,cmgj000113
	Codes *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
	// The commodity codes to exclude. This parameter cannot be specified together with the Codes parameter. Separate multiple codes with commas (,).
	//
	// example:
	//
	// cmgj000114,cmgj000115
	ExceptCodes *string `json:"ExceptCodes,omitempty" xml:"ExceptCodes,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The commodity type. Valid values:
	//
	// - APP: application
	//
	// - SERVICE: service
	//
	// - MIRROR: image
	//
	// - DOWNLOAD: download
	//
	// - API_SERVICE: API
	//
	// - ROS: resource orchestration
	//
	// - DOCKER: container
	//
	// - IOT: IoT.
	//
	// example:
	//
	// 1
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetCodes() *string {
	return s.Codes
}

func (s *DescribeInstancesRequest) GetExceptCodes() *string {
	return s.ExceptCodes
}

func (s *DescribeInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInstancesRequest) SetCodes(v string) *DescribeInstancesRequest {
	s.Codes = &v
	return s
}

func (s *DescribeInstancesRequest) SetExceptCodes(v string) *DescribeInstancesRequest {
	s.ExceptCodes = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetProductType(v string) *DescribeInstancesRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	return dara.Validate(s)
}
