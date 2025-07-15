// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ConfirmPhysicalConnectionRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ConfirmPhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ConfirmPhysicalConnectionRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *ConfirmPhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *ConfirmPhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ConfirmPhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConfirmPhysicalConnectionRequest
	GetResourceOwnerId() *int64
}

type ConfirmPhysicalConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e0****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-119mf****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
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
}

func (s ConfirmPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *ConfirmPhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConfirmPhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ConfirmPhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConfirmPhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *ConfirmPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfirmPhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConfirmPhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConfirmPhysicalConnectionRequest) SetClientToken(v string) *ConfirmPhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfirmPhysicalConnectionRequest) SetOwnerAccount(v string) *ConfirmPhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ConfirmPhysicalConnectionRequest) SetOwnerId(v int64) *ConfirmPhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfirmPhysicalConnectionRequest) SetPhysicalConnectionId(v string) *ConfirmPhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *ConfirmPhysicalConnectionRequest) SetRegionId(v string) *ConfirmPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *ConfirmPhysicalConnectionRequest) SetResourceOwnerAccount(v string) *ConfirmPhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConfirmPhysicalConnectionRequest) SetResourceOwnerId(v int64) *ConfirmPhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConfirmPhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
