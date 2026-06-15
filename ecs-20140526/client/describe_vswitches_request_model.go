// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsDefault(v bool) *DescribeVSwitchesRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *DescribeVSwitchesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVSwitchesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVSwitchesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVSwitchesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVSwitchesRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *DescribeVSwitchesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeVSwitchesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeVSwitchesRequest
	GetZoneId() *string
}

type DescribeVSwitchesRequest struct {
	// Specifies whether the vSwitch is the default vSwitch in the zone. Valid values:
	//
	// - `true`
	//
	// - `false`
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number to return.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the vSwitches are located. To query the latest list of regions, call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVSwitchesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVSwitchesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVSwitchesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesRequest) SetIsDefault(v bool) *DescribeVSwitchesRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerId(v int64) *DescribeVSwitchesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageNumber(v int32) *DescribeVSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageSize(v int32) *DescribeVSwitchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetRegionId(v string) *DescribeVSwitchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerId(v int64) *DescribeVSwitchesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchId(v string) *DescribeVSwitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVpcId(v string) *DescribeVSwitchesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetZoneId(v string) *DescribeVSwitchesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesRequest) Validate() error {
	return dara.Validate(s)
}
