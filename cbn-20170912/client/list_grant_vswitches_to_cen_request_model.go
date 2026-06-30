// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantVSwitchesToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListGrantVSwitchesToCenRequest
	GetCenId() *string
	SetEnabledIpv6(v bool) *ListGrantVSwitchesToCenRequest
	GetEnabledIpv6() *bool
	SetOwnerAccount(v string) *ListGrantVSwitchesToCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListGrantVSwitchesToCenRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListGrantVSwitchesToCenRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGrantVSwitchesToCenRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListGrantVSwitchesToCenRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListGrantVSwitchesToCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListGrantVSwitchesToCenRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *ListGrantVSwitchesToCenRequest
	GetVpcId() *string
	SetZoneId(v string) *ListGrantVSwitchesToCenRequest
	GetZoneId() *string
}

type ListGrantVSwitchesToCenRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-44m0p68spvlrqq****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Specifies whether IPv6 is enabled. Valid values: true and false. If you do not specify this parameter, the query is not filtered by this property.
	//
	// example:
	//
	// true
	EnabledIpv6  *bool   `json:"EnabledIpv6,omitempty" xml:"EnabledIpv6,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the VPC is deployed.
	//
	// Call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp12ge2tq5gzdc915****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// - If you specify a zone ID, only the vSwitches in that zone are queried.
	//
	// - If you do not specify a zone ID, the vSwitches in all zones of the VPC are queried by default.
	//
	// Call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListGrantVSwitchesToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchesToCenRequest) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListGrantVSwitchesToCenRequest) GetEnabledIpv6() *bool {
	return s.EnabledIpv6
}

func (s *ListGrantVSwitchesToCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListGrantVSwitchesToCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListGrantVSwitchesToCenRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGrantVSwitchesToCenRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGrantVSwitchesToCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGrantVSwitchesToCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListGrantVSwitchesToCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListGrantVSwitchesToCenRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGrantVSwitchesToCenRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListGrantVSwitchesToCenRequest) SetCenId(v string) *ListGrantVSwitchesToCenRequest {
	s.CenId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetEnabledIpv6(v bool) *ListGrantVSwitchesToCenRequest {
	s.EnabledIpv6 = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetOwnerAccount(v string) *ListGrantVSwitchesToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetOwnerId(v int64) *ListGrantVSwitchesToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetPageNumber(v int32) *ListGrantVSwitchesToCenRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetPageSize(v int32) *ListGrantVSwitchesToCenRequest {
	s.PageSize = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetRegionId(v string) *ListGrantVSwitchesToCenRequest {
	s.RegionId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetResourceOwnerAccount(v string) *ListGrantVSwitchesToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetResourceOwnerId(v int64) *ListGrantVSwitchesToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetVpcId(v string) *ListGrantVSwitchesToCenRequest {
	s.VpcId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) SetZoneId(v string) *ListGrantVSwitchesToCenRequest {
	s.ZoneId = &v
	return s
}

func (s *ListGrantVSwitchesToCenRequest) Validate() error {
	return dara.Validate(s)
}
