// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLniPrivateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListLniPrivateIpAddressRequest
	GetDescription() *string
	SetEnablePage(v bool) *ListLniPrivateIpAddressRequest
	GetEnablePage() *bool
	SetIp(v string) *ListLniPrivateIpAddressRequest
	GetIp() *string
	SetIpName(v string) *ListLniPrivateIpAddressRequest
	GetIpName() *string
	SetNetworkInterfaceId(v string) *ListLniPrivateIpAddressRequest
	GetNetworkInterfaceId() *string
	SetPageNumber(v int32) *ListLniPrivateIpAddressRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLniPrivateIpAddressRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListLniPrivateIpAddressRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListLniPrivateIpAddressRequest
	GetResourceGroupId() *string
}

type ListLniPrivateIpAddressRequest struct {
	// The description of the variable.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether pagination is required
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// network interface controller IP address
	//
	// example:
	//
	// 10.0.98.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// IP unique identifier
	//
	// example:
	//
	// sip-tynhdh2s
	IpName *string `json:"IpName,omitempty" xml:"IpName,omitempty"`
	// Lingjun network interface controller ID
	//
	// example:
	//
	// lni-2ze4uww7n6hsfzrwq77y
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Obtain the index number of the current mouse click for an animation
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListLniPrivateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLniPrivateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressRequest) GetDescription() *string {
	return s.Description
}

func (s *ListLniPrivateIpAddressRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListLniPrivateIpAddressRequest) GetIp() *string {
	return s.Ip
}

func (s *ListLniPrivateIpAddressRequest) GetIpName() *string {
	return s.IpName
}

func (s *ListLniPrivateIpAddressRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListLniPrivateIpAddressRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLniPrivateIpAddressRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLniPrivateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLniPrivateIpAddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLniPrivateIpAddressRequest) SetDescription(v string) *ListLniPrivateIpAddressRequest {
	s.Description = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetEnablePage(v bool) *ListLniPrivateIpAddressRequest {
	s.EnablePage = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetIp(v string) *ListLniPrivateIpAddressRequest {
	s.Ip = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetIpName(v string) *ListLniPrivateIpAddressRequest {
	s.IpName = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetNetworkInterfaceId(v string) *ListLniPrivateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetPageNumber(v int32) *ListLniPrivateIpAddressRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetPageSize(v int32) *ListLniPrivateIpAddressRequest {
	s.PageSize = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetRegionId(v string) *ListLniPrivateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) SetResourceGroupId(v string) *ListLniPrivateIpAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLniPrivateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
