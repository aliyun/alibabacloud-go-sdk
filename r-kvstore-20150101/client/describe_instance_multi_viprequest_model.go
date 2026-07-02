// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMultiVIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceMultiVIPRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceMultiVIPRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceMultiVIPRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeInstanceMultiVIPRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceMultiVIPRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceMultiVIPRequest struct {
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
}

func (s DescribeInstanceMultiVIPRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMultiVIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMultiVIPRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceMultiVIPRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceMultiVIPRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceMultiVIPRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceMultiVIPRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceMultiVIPRequest) SetInstanceId(v string) *DescribeInstanceMultiVIPRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMultiVIPRequest) SetOwnerAccount(v string) *DescribeInstanceMultiVIPRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceMultiVIPRequest) SetOwnerId(v int64) *DescribeInstanceMultiVIPRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceMultiVIPRequest) SetResourceOwnerAccount(v string) *DescribeInstanceMultiVIPRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceMultiVIPRequest) SetResourceOwnerId(v int64) *DescribeInstanceMultiVIPRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceMultiVIPRequest) Validate() error {
	return dara.Validate(s)
}
