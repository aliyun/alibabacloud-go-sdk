// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSecurityGroupConfigurationRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeSecurityGroupConfigurationRequest
	GetSecurityToken() *string
}

type DescribeSecurityGroupConfigurationRequest struct {
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

func (s DescribeSecurityGroupConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityGroupConfigurationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSecurityGroupConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSecurityGroupConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSecurityGroupConfigurationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityGroupConfigurationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSecurityGroupConfigurationRequest) SetInstanceId(v string) *DescribeSecurityGroupConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetResourceOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetResourceOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) SetSecurityToken(v string) *DescribeSecurityGroupConfigurationRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
