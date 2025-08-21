// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecondaryPublicIpAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSecondaryPublicIpAddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSecondaryPublicIpAddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSecondaryPublicIpAddressesResponseBody
	GetRequestId() *string
	SetSecondaryPublicIpAddresses(v []*DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) *DescribeSecondaryPublicIpAddressesResponseBody
	GetSecondaryPublicIpAddresses() []*DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses
	SetTotalCount(v int32) *DescribeSecondaryPublicIpAddressesResponseBody
	GetTotalCount() *int32
}

type DescribeSecondaryPublicIpAddressesResponseBody struct {
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of returned secondary IP addresses.
	SecondaryPublicIpAddresses []*DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses `json:"SecondaryPublicIpAddresses,omitempty" xml:"SecondaryPublicIpAddresses,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecondaryPublicIpAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecondaryPublicIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) GetSecondaryPublicIpAddresses() []*DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	return s.SecondaryPublicIpAddresses
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) SetPageNumber(v int32) *DescribeSecondaryPublicIpAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) SetPageSize(v int32) *DescribeSecondaryPublicIpAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) SetRequestId(v string) *DescribeSecondaryPublicIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) SetSecondaryPublicIpAddresses(v []*DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) *DescribeSecondaryPublicIpAddressesResponseBody {
	s.SecondaryPublicIpAddresses = v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) SetTotalCount(v int32) *DescribeSecondaryPublicIpAddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses struct {
	// The subnet mask of the CIDR block.
	//
	// example:
	//
	// 24
	CidrMask *int32 `json:"CidrMask,omitempty" xml:"CidrMask,omitempty"`
	// The time when the secondary public IP address was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-07-25T09:43:49Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing-15
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The gateway.
	//
	// example:
	//
	// 12.XXX.XXX.1
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The version of the IP address. Valid values:
	//
	// 	- **ipv4**
	//
	// 	- **ipv6**
	//
	// example:
	//
	// ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
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
	// telecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The secondary public IP address.
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

func (s DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetCidrMask() *int32 {
	return s.CidrMask
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetGateway() *string {
	return s.Gateway
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetSecondaryPublicIpAddress() *string {
	return s.SecondaryPublicIpAddress
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) GetSecondaryPublicIpId() *string {
	return s.SecondaryPublicIpId
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetCidrMask(v int32) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.CidrMask = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetCreationTime(v string) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetEnsRegionId(v string) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetGateway(v string) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.Gateway = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetIpVersion(v string) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.IpVersion = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetIsp(v string) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.Isp = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetSecondaryPublicIpAddress(v string) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.SecondaryPublicIpAddress = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) SetSecondaryPublicIpId(v string) *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses {
	s.SecondaryPublicIpId = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponseBodySecondaryPublicIpAddresses) Validate() error {
	return dara.Validate(s)
}
