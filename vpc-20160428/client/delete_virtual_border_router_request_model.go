// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVirtualBorderRouterRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetVbrId(v string) *DeleteVirtualBorderRouterRequest
	GetVbrId() *string
}

type DeleteVirtualBorderRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VBR. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp12mw1f8k3jgygk9****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s DeleteVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVirtualBorderRouterRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *DeleteVirtualBorderRouterRequest) SetClientToken(v string) *DeleteVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetOwnerAccount(v string) *DeleteVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetOwnerId(v int64) *DeleteVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetRegionId(v string) *DeleteVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *DeleteVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *DeleteVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) SetVbrId(v string) *DeleteVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *DeleteVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
