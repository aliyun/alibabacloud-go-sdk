// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *TerminateVirtualBorderRouterRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *TerminateVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TerminateVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TerminateVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *TerminateVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TerminateVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetVbrId(v string) *TerminateVirtualBorderRouterRequest
	GetVbrId() *string
}

type TerminateVirtualBorderRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
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
	// vbr-bp15zckdt37pq72****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s TerminateVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *TerminateVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TerminateVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TerminateVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TerminateVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TerminateVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TerminateVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TerminateVirtualBorderRouterRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *TerminateVirtualBorderRouterRequest) SetClientToken(v string) *TerminateVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetOwnerAccount(v string) *TerminateVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetOwnerId(v int64) *TerminateVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetRegionId(v string) *TerminateVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *TerminateVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *TerminateVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) SetVbrId(v string) *TerminateVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *TerminateVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
