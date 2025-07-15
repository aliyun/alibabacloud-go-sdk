// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociatePhysicalConnectionFromVirtualBorderRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetResourceOwnerId() *int64
	SetVbrId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	GetVbrId() *string
}

type UnassociatePhysicalConnectionFromVirtualBorderRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1qrb3044eqixog****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
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
	// The ID of the VBR that you want to disassociate from the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp16ksp61j7e0tkn*****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GoString() string {
	return s.String()
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetClientToken(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetOwnerAccount(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetOwnerId(v int64) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetPhysicalConnectionId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetRegionId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetResourceOwnerAccount(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetResourceOwnerId(v int64) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) SetVbrId(v string) *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest {
	s.VbrId = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) Validate() error {
	return dara.Validate(s)
}
