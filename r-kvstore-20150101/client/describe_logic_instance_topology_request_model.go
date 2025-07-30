// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogicInstanceTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeLogicInstanceTopologyRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeLogicInstanceTopologyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLogicInstanceTopologyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLogicInstanceTopologyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLogicInstanceTopologyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeLogicInstanceTopologyRequest
	GetSecurityToken() *string
}

type DescribeLogicInstanceTopologyRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLogicInstanceTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogicInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLogicInstanceTopologyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLogicInstanceTopologyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLogicInstanceTopologyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLogicInstanceTopologyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLogicInstanceTopologyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLogicInstanceTopologyRequest) SetInstanceId(v string) *DescribeLogicInstanceTopologyRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetOwnerAccount(v string) *DescribeLogicInstanceTopologyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetOwnerId(v int64) *DescribeLogicInstanceTopologyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetResourceOwnerAccount(v string) *DescribeLogicInstanceTopologyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetResourceOwnerId(v int64) *DescribeLogicInstanceTopologyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) SetSecurityToken(v string) *DescribeLogicInstanceTopologyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLogicInstanceTopologyRequest) Validate() error {
	return dara.Validate(s)
}
