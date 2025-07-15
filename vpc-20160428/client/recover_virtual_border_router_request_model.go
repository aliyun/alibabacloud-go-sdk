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
	SetVbrId(v string) *RecoverVirtualBorderRouterRequest
	GetVbrId() *string
}

type RecoverVirtualBorderRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
	// vbr-bp1lhl0taikrte****
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

func (s *RecoverVirtualBorderRouterRequest) SetVbrId(v string) *RecoverVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *RecoverVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
