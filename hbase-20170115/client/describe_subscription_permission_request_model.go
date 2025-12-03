// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeSubscriptionPermissionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSubscriptionPermissionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSubscriptionPermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSubscriptionPermissionRequest
	GetResourceOwnerId() *int64
}

type DescribeSubscriptionPermissionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSubscriptionPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSubscriptionPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSubscriptionPermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSubscriptionPermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSubscriptionPermissionRequest) SetOwnerId(v int64) *DescribeSubscriptionPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionPermissionRequest) SetRegionId(v string) *DescribeSubscriptionPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionPermissionRequest) SetResourceOwnerAccount(v string) *DescribeSubscriptionPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSubscriptionPermissionRequest) SetResourceOwnerId(v int64) *DescribeSubscriptionPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSubscriptionPermissionRequest) Validate() error {
	return dara.Validate(s)
}
