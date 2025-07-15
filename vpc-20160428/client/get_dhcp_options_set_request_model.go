// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDhcpOptionsSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDhcpOptionsSetId(v string) *GetDhcpOptionsSetRequest
	GetDhcpOptionsSetId() *string
	SetOwnerAccount(v string) *GetDhcpOptionsSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetDhcpOptionsSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetDhcpOptionsSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetDhcpOptionsSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetDhcpOptionsSetRequest
	GetResourceOwnerId() *int64
}

type GetDhcpOptionsSetRequest struct {
	// The ID of the DHCP options set.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the DHCP options set that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetDhcpOptionsSetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDhcpOptionsSetRequest) GoString() string {
	return s.String()
}

func (s *GetDhcpOptionsSetRequest) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *GetDhcpOptionsSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetDhcpOptionsSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetDhcpOptionsSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDhcpOptionsSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetDhcpOptionsSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetDhcpOptionsSetRequest) SetDhcpOptionsSetId(v string) *GetDhcpOptionsSetRequest {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *GetDhcpOptionsSetRequest) SetOwnerAccount(v string) *GetDhcpOptionsSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetDhcpOptionsSetRequest) SetOwnerId(v int64) *GetDhcpOptionsSetRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDhcpOptionsSetRequest) SetRegionId(v string) *GetDhcpOptionsSetRequest {
	s.RegionId = &v
	return s
}

func (s *GetDhcpOptionsSetRequest) SetResourceOwnerAccount(v string) *GetDhcpOptionsSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDhcpOptionsSetRequest) SetResourceOwnerId(v int64) *GetDhcpOptionsSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDhcpOptionsSetRequest) Validate() error {
	return dara.Validate(s)
}
