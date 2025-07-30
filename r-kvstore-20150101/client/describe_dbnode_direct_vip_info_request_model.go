// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodeDirectVipInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDBNodeDirectVipInfoRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBNodeDirectVipInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBNodeDirectVipInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBNodeDirectVipInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBNodeDirectVipInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeDBNodeDirectVipInfoRequest struct {
	// The instance ID.
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
}

func (s DescribeDBNodeDirectVipInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodeDirectVipInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBNodeDirectVipInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDBNodeDirectVipInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBNodeDirectVipInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBNodeDirectVipInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBNodeDirectVipInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBNodeDirectVipInfoRequest) SetInstanceId(v string) *DescribeDBNodeDirectVipInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoRequest) SetOwnerAccount(v string) *DescribeDBNodeDirectVipInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoRequest) SetOwnerId(v int64) *DescribeDBNodeDirectVipInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoRequest) SetResourceOwnerAccount(v string) *DescribeDBNodeDirectVipInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoRequest) SetResourceOwnerId(v int64) *DescribeDBNodeDirectVipInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBNodeDirectVipInfoRequest) Validate() error {
	return dara.Validate(s)
}
