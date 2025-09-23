// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVSwitchListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVSwitchListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVSwitchListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVSwitchListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVSwitchListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVSwitchListRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeVSwitchListRequest
	GetSecurityToken() *string
	SetVSwitchIds(v []*string) *DescribeVSwitchListRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *DescribeVSwitchListRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeVSwitchListRequest
	GetZoneId() *string
}

type DescribeVSwitchListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VSwitchIds           []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-25cdvfeq58pl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVSwitchListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVSwitchListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVSwitchListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVSwitchListRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeVSwitchListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchListRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchListRequest) SetOwnerAccount(v string) *DescribeVSwitchListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetOwnerId(v int64) *DescribeVSwitchListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetPageNumber(v int32) *DescribeVSwitchListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetPageSize(v int32) *DescribeVSwitchListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetRegionId(v string) *DescribeVSwitchListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetResourceOwnerAccount(v string) *DescribeVSwitchListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetResourceOwnerId(v int64) *DescribeVSwitchListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetSecurityToken(v string) *DescribeVSwitchListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetVSwitchIds(v []*string) *DescribeVSwitchListRequest {
	s.VSwitchIds = v
	return s
}

func (s *DescribeVSwitchListRequest) SetVpcId(v string) *DescribeVSwitchListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetZoneId(v string) *DescribeVSwitchListRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchListRequest) Validate() error {
	return dara.Validate(s)
}
