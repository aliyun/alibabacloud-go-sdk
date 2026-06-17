// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DescribeGatewayListRequest
	GetGwClusterId() *string
	SetGwDescription(v string) *DescribeGatewayListRequest
	GetGwDescription() *string
	SetPageNumber(v string) *DescribeGatewayListRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeGatewayListRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeGatewayListRequest
	GetRegionId() *string
}

type DescribeGatewayListRequest struct {
	// The ID of the gateway instance.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The description of the gateway instance.
	//
	// example:
	//
	// test
	GwDescription *string `json:"GwDescription,omitempty" xml:"GwDescription,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGatewayListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayListRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeGatewayListRequest) GetGwDescription() *string {
	return s.GwDescription
}

func (s *DescribeGatewayListRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeGatewayListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGatewayListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGatewayListRequest) SetGwClusterId(v string) *DescribeGatewayListRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeGatewayListRequest) SetGwDescription(v string) *DescribeGatewayListRequest {
	s.GwDescription = &v
	return s
}

func (s *DescribeGatewayListRequest) SetPageNumber(v string) *DescribeGatewayListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayListRequest) SetPageSize(v string) *DescribeGatewayListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayListRequest) SetRegionId(v string) *DescribeGatewayListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGatewayListRequest) Validate() error {
	return dara.Validate(s)
}
