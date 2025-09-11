// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeBackupRegionsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeBackupRegionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupRegionsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeBackupRegionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupRegionsRequest
	GetResourceOwnerId() *int64
}

type DescribeBackupRegionsRequest struct {
	// example:
	//
	// pc-uf64u64fln9039***
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupRegionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupRegionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupRegionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupRegionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupRegionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupRegionsRequest) SetDBClusterId(v string) *DescribeBackupRegionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupRegionsRequest) SetOwnerAccount(v string) *DescribeBackupRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupRegionsRequest) SetOwnerId(v int64) *DescribeBackupRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupRegionsRequest) SetResourceOwnerAccount(v string) *DescribeBackupRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupRegionsRequest) SetResourceOwnerId(v int64) *DescribeBackupRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupRegionsRequest) Validate() error {
	return dara.Validate(s)
}
