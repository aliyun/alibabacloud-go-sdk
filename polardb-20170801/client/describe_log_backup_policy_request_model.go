// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLogBackupPolicyRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeLogBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLogBackupPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLogBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLogBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeLogBackupPolicyRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query all the information about the available clusters in the target region, including the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLogBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLogBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLogBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLogBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLogBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLogBackupPolicyRequest) SetDBClusterId(v string) *DescribeLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetOwnerAccount(v string) *DescribeLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetOwnerId(v int64) *DescribeLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
