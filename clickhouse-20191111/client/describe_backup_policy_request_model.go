// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeBackupPolicyRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeBackupPolicyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupPolicyRequest) SetDBClusterId(v string) *DescribeBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
