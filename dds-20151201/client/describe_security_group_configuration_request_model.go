// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSecurityGroupConfigurationRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSecurityGroupConfigurationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSecurityGroupConfigurationRequest
	GetResourceOwnerId() *int64
}

type DescribeSecurityGroupConfigurationRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSecurityGroupConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
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

func (s *DescribeSecurityGroupConfigurationRequest) SetDBInstanceId(v string) *DescribeSecurityGroupConfigurationRequest {
	s.DBInstanceId = &v
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

func (s *DescribeSecurityGroupConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
