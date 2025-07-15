// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelPhysicalConnectionRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CancelPhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelPhysicalConnectionRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *CancelPhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *CancelPhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelPhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelPhysicalConnectionRequest
	GetResourceOwnerId() *int64
}

type CancelPhysicalConnectionRequest struct {
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
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-119mfjzm7****
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

func (s CancelPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *CancelPhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelPhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelPhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelPhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CancelPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelPhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelPhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelPhysicalConnectionRequest) SetClientToken(v string) *CancelPhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetOwnerAccount(v string) *CancelPhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetOwnerId(v int64) *CancelPhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetPhysicalConnectionId(v string) *CancelPhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetRegionId(v string) *CancelPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetResourceOwnerAccount(v string) *CancelPhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetResourceOwnerId(v int64) *CancelPhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
