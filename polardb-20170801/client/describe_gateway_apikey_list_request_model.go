// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayApikeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeGatewayApikeyListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGatewayApikeyListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGatewayApikeyListRequest
	GetRegionId() *string
}

type DescribeGatewayApikeyListRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGatewayApikeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayApikeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayApikeyListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGatewayApikeyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGatewayApikeyListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGatewayApikeyListRequest) SetPageNumber(v int32) *DescribeGatewayApikeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayApikeyListRequest) SetPageSize(v int32) *DescribeGatewayApikeyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayApikeyListRequest) SetRegionId(v string) *DescribeGatewayApikeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGatewayApikeyListRequest) Validate() error {
	return dara.Validate(s)
}
