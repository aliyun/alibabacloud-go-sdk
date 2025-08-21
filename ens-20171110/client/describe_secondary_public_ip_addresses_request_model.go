// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecondaryPublicIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeSecondaryPublicIpAddressesRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeSecondaryPublicIpAddressesRequest
	GetEnsRegionIds() []*string
	SetIsp(v string) *DescribeSecondaryPublicIpAddressesRequest
	GetIsp() *string
	SetPageNumber(v int32) *DescribeSecondaryPublicIpAddressesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSecondaryPublicIpAddressesRequest
	GetPageSize() *int32
	SetSecondaryPublicIpAddress(v string) *DescribeSecondaryPublicIpAddressesRequest
	GetSecondaryPublicIpAddress() *string
	SetSecondaryPublicIpId(v string) *DescribeSecondaryPublicIpAddressesRequest
	GetSecondaryPublicIpId() *string
}

type DescribeSecondaryPublicIpAddressesRequest struct {
	// The ID of the edge node.
	//
	// example:
	//
	// cn-hangzhou-44
	EnsRegionId  *string   `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The Internet service provider. Valid values:
	//
	// 	- cmcc: China Mobile.
	//
	// 	- unicom: China Unicom.
	//
	// 	- telecom: China Telecom.
	//
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The secondary IP address.
	//
	// example:
	//
	// 12.XXX.XXX.4
	SecondaryPublicIpAddress *string `json:"SecondaryPublicIpAddress,omitempty" xml:"SecondaryPublicIpAddress,omitempty"`
	// The ID of the secondary public IP address.
	//
	// example:
	//
	// spi-5wys0pio93c9f9ukzj2f2fwyr
	SecondaryPublicIpId *string `json:"SecondaryPublicIpId,omitempty" xml:"SecondaryPublicIpId,omitempty"`
}

func (s DescribeSecondaryPublicIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecondaryPublicIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecondaryPublicIpAddressesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeSecondaryPublicIpAddressesRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeSecondaryPublicIpAddressesRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSecondaryPublicIpAddressesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSecondaryPublicIpAddressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecondaryPublicIpAddressesRequest) GetSecondaryPublicIpAddress() *string {
	return s.SecondaryPublicIpAddress
}

func (s *DescribeSecondaryPublicIpAddressesRequest) GetSecondaryPublicIpId() *string {
	return s.SecondaryPublicIpId
}

func (s *DescribeSecondaryPublicIpAddressesRequest) SetEnsRegionId(v string) *DescribeSecondaryPublicIpAddressesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesRequest) SetEnsRegionIds(v []*string) *DescribeSecondaryPublicIpAddressesRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesRequest) SetIsp(v string) *DescribeSecondaryPublicIpAddressesRequest {
	s.Isp = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesRequest) SetPageNumber(v int32) *DescribeSecondaryPublicIpAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesRequest) SetPageSize(v int32) *DescribeSecondaryPublicIpAddressesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesRequest) SetSecondaryPublicIpAddress(v string) *DescribeSecondaryPublicIpAddressesRequest {
	s.SecondaryPublicIpAddress = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesRequest) SetSecondaryPublicIpId(v string) *DescribeSecondaryPublicIpAddressesRequest {
	s.SecondaryPublicIpId = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}
