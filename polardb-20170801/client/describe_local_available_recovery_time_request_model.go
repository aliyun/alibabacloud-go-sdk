// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalAvailableRecoveryTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetSecurityToken() *string
}

type DescribeLocalAvailableRecoveryTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze3ngi149b313***
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLocalAvailableRecoveryTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalAvailableRecoveryTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetDBClusterId(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetOwnerAccount(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetResourceOwnerAccount(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetResourceOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetSecurityToken(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) Validate() error {
	return dara.Validate(s)
}
