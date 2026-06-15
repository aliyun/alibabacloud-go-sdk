// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RecoverVirtualBorderRouterRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *RecoverVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RecoverVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RecoverVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RecoverVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RecoverVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *RecoverVirtualBorderRouterRequest
	GetUserCidr() *string
	SetVbrId(v string) *RecoverVirtualBorderRouterRequest
	GetVbrId() *string
}

type RecoverVirtualBorderRouterRequest struct {
	// A client-generated token that must be unique across requests to ensure idempotency. The token can contain only ASCII characters and must be no longer than 64 characters.
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the VBR is located. You can call the `DescribeRegions` operation to get the latest region list.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The secondary IPv4 CIDR block of your on-premises data center. This parameter is used for disaster recovery in a dual-homed configuration.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	// The ID of the VBR to recover.
	//
	// This parameter is required.
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s RecoverVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *RecoverVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RecoverVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RecoverVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RecoverVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RecoverVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RecoverVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RecoverVirtualBorderRouterRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *RecoverVirtualBorderRouterRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *RecoverVirtualBorderRouterRequest) SetClientToken(v string) *RecoverVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) SetOwnerAccount(v string) *RecoverVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) SetOwnerId(v int64) *RecoverVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) SetRegionId(v string) *RecoverVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *RecoverVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *RecoverVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) SetUserCidr(v string) *RecoverVirtualBorderRouterRequest {
	s.UserCidr = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) SetVbrId(v string) *RecoverVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
