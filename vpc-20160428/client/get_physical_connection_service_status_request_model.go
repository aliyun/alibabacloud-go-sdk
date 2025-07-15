// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalConnectionServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetPhysicalConnectionServiceStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetPhysicalConnectionServiceStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetPhysicalConnectionServiceStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetPhysicalConnectionServiceStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPhysicalConnectionServiceStatusRequest
	GetResourceOwnerId() *int64
}

type GetPhysicalConnectionServiceStatusRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region for which you want to query the status of outbound data transfer billing.
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

func (s GetPhysicalConnectionServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalConnectionServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalConnectionServiceStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetPhysicalConnectionServiceStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPhysicalConnectionServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPhysicalConnectionServiceStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPhysicalConnectionServiceStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPhysicalConnectionServiceStatusRequest) SetOwnerAccount(v string) *GetPhysicalConnectionServiceStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusRequest) SetOwnerId(v int64) *GetPhysicalConnectionServiceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusRequest) SetRegionId(v string) *GetPhysicalConnectionServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusRequest) SetResourceOwnerAccount(v string) *GetPhysicalConnectionServiceStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusRequest) SetResourceOwnerId(v int64) *GetPhysicalConnectionServiceStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhysicalConnectionServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
