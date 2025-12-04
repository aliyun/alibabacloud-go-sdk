// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLeniPrivateIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticNetworkInterfaceId(v string) *ListLeniPrivateIpAddressesRequest
	GetElasticNetworkInterfaceId() *string
	SetIpName(v string) *ListLeniPrivateIpAddressesRequest
	GetIpName() *string
	SetPageNumber(v int32) *ListLeniPrivateIpAddressesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLeniPrivateIpAddressesRequest
	GetPageSize() *int32
	SetPrivateIpAddress(v string) *ListLeniPrivateIpAddressesRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *ListLeniPrivateIpAddressesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListLeniPrivateIpAddressesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListLeniPrivateIpAddressesRequest
	GetStatus() *string
}

type ListLeniPrivateIpAddressesRequest struct {
	// Lingjun Elastic Network Interface ID.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP unique identifier.
	//
	// example:
	//
	// sip-8ylg****
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Lingjun Elastic Network Interface secondary private IP.
	//
	// example:
	//
	// 10.0.****
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmynvzeknccpy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the image build command risk.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLeniPrivateIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLeniPrivateIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *ListLeniPrivateIpAddressesRequest) GetIpName() *string {
	return s.IpName
}

func (s *ListLeniPrivateIpAddressesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLeniPrivateIpAddressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLeniPrivateIpAddressesRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ListLeniPrivateIpAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLeniPrivateIpAddressesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLeniPrivateIpAddressesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListLeniPrivateIpAddressesRequest) SetElasticNetworkInterfaceId(v string) *ListLeniPrivateIpAddressesRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetIpName(v string) *ListLeniPrivateIpAddressesRequest {
	s.IpName = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetPageNumber(v int32) *ListLeniPrivateIpAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetPageSize(v int32) *ListLeniPrivateIpAddressesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetPrivateIpAddress(v string) *ListLeniPrivateIpAddressesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetRegionId(v string) *ListLeniPrivateIpAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetResourceGroupId(v string) *ListLeniPrivateIpAddressesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) SetStatus(v string) *ListLeniPrivateIpAddressesRequest {
	s.Status = &v
	return s
}

func (s *ListLeniPrivateIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}
