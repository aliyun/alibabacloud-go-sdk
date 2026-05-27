// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeWuyingServerResponseBodyData) *DescribeWuyingServerResponseBody
	GetData() *DescribeWuyingServerResponseBodyData
	SetRequestId(v string) *DescribeWuyingServerResponseBody
	GetRequestId() *string
}

type DescribeWuyingServerResponseBody struct {
	Data *DescribeWuyingServerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerResponseBody) GetData() *DescribeWuyingServerResponseBodyData {
	return s.Data
}

func (s *DescribeWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWuyingServerResponseBody) SetData(v *DescribeWuyingServerResponseBodyData) *DescribeWuyingServerResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWuyingServerResponseBody) SetRequestId(v string) *DescribeWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWuyingServerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWuyingServerResponseBodyData struct {
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 10
	EniPrivateIpAddressQuantity *int32 `json:"EniPrivateIpAddressQuantity,omitempty" xml:"EniPrivateIpAddressQuantity,omitempty"`
	// example:
	//
	// 2027-01-01T00:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// img-bp1234567890abcde
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// Ubuntu 22.04
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// example:
	//
	// 10.0.0.1
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-abc123
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// 默认工作区
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// example:
	//
	// Simple
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// example:
	//
	// Linux
	OsType        *string                                              `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PrivateIpSets []*DescribeWuyingServerResponseBodyDataPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Repeated"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cloud_essd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// example:
	//
	// aig-bp1234567890abcde
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
	// example:
	//
	// my-dev-server
	WuyingServerName *string `json:"WuyingServerName,omitempty" xml:"WuyingServerName,omitempty"`
}

func (s DescribeWuyingServerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerResponseBodyData) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeWuyingServerResponseBodyData) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeWuyingServerResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeWuyingServerResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeWuyingServerResponseBodyData) GetEniPrivateIpAddressQuantity() *int32 {
	return s.EniPrivateIpAddressQuantity
}

func (s *DescribeWuyingServerResponseBodyData) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeWuyingServerResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeWuyingServerResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeWuyingServerResponseBodyData) GetNetworkInterfaceIp() *string {
	return s.NetworkInterfaceIp
}

func (s *DescribeWuyingServerResponseBodyData) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeWuyingServerResponseBodyData) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeWuyingServerResponseBodyData) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeWuyingServerResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *DescribeWuyingServerResponseBodyData) GetPrivateIpSets() []*DescribeWuyingServerResponseBodyDataPrivateIpSets {
	return s.PrivateIpSets
}

func (s *DescribeWuyingServerResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeWuyingServerResponseBodyData) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeWuyingServerResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeWuyingServerResponseBodyData) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *DescribeWuyingServerResponseBodyData) GetWuyingServerName() *string {
	return s.WuyingServerName
}

func (s *DescribeWuyingServerResponseBodyData) SetBandwidth(v int32) *DescribeWuyingServerResponseBodyData {
	s.Bandwidth = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetBizRegionId(v string) *DescribeWuyingServerResponseBodyData {
	s.BizRegionId = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetChargeType(v string) *DescribeWuyingServerResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetCreateTime(v string) *DescribeWuyingServerResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetEniPrivateIpAddressQuantity(v int32) *DescribeWuyingServerResponseBodyData {
	s.EniPrivateIpAddressQuantity = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetExpiredTime(v string) *DescribeWuyingServerResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetImageId(v string) *DescribeWuyingServerResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetImageName(v string) *DescribeWuyingServerResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetNetworkInterfaceIp(v string) *DescribeWuyingServerResponseBodyData {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetOfficeSiteId(v string) *DescribeWuyingServerResponseBodyData {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetOfficeSiteName(v string) *DescribeWuyingServerResponseBodyData {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetOfficeSiteType(v string) *DescribeWuyingServerResponseBodyData {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetOsType(v string) *DescribeWuyingServerResponseBodyData {
	s.OsType = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetPrivateIpSets(v []*DescribeWuyingServerResponseBodyDataPrivateIpSets) *DescribeWuyingServerResponseBodyData {
	s.PrivateIpSets = v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetStatus(v string) *DescribeWuyingServerResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetSystemDiskCategory(v string) *DescribeWuyingServerResponseBodyData {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetSystemDiskSize(v int32) *DescribeWuyingServerResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetWuyingServerId(v string) *DescribeWuyingServerResponseBodyData {
	s.WuyingServerId = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) SetWuyingServerName(v string) *DescribeWuyingServerResponseBodyData {
	s.WuyingServerName = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyData) Validate() error {
	if s.PrivateIpSets != nil {
		for _, item := range s.PrivateIpSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWuyingServerResponseBodyDataPrivateIpSets struct {
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// example:
	//
	// 10.0.0.1
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeWuyingServerResponseBodyDataPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeWuyingServerResponseBodyDataPrivateIpSets) GoString() string {
	return s.String()
}

func (s *DescribeWuyingServerResponseBodyDataPrivateIpSets) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeWuyingServerResponseBodyDataPrivateIpSets) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeWuyingServerResponseBodyDataPrivateIpSets) SetPrimary(v bool) *DescribeWuyingServerResponseBodyDataPrivateIpSets {
	s.Primary = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyDataPrivateIpSets) SetPrivateIpAddress(v string) *DescribeWuyingServerResponseBodyDataPrivateIpSets {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeWuyingServerResponseBodyDataPrivateIpSets) Validate() error {
	return dara.Validate(s)
}
