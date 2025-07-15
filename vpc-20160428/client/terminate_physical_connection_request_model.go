// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminatePhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *TerminatePhysicalConnectionRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *TerminatePhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TerminatePhysicalConnectionRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *TerminatePhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *TerminatePhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *TerminatePhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TerminatePhysicalConnectionRequest
	GetResourceOwnerId() *int64
}

type TerminatePhysicalConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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
	// pc-119mfjzm****
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
}

func (s TerminatePhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminatePhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *TerminatePhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TerminatePhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TerminatePhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TerminatePhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *TerminatePhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TerminatePhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TerminatePhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TerminatePhysicalConnectionRequest) SetClientToken(v string) *TerminatePhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetOwnerAccount(v string) *TerminatePhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetOwnerId(v int64) *TerminatePhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetPhysicalConnectionId(v string) *TerminatePhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetRegionId(v string) *TerminatePhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *TerminatePhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetResourceOwnerId(v int64) *TerminatePhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
