// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *CheckConnectionStringRequest
	GetConnectionStringPrefix() *string
	SetDBClusterId(v string) *CheckConnectionStringRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CheckConnectionStringRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckConnectionStringRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckConnectionStringRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckConnectionStringRequest
	GetResourceOwnerId() *int64
}

type CheckConnectionStringRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************.pg.polardb.pre.rds.aliyuncs.com:5432
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *CheckConnectionStringRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CheckConnectionStringRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckConnectionStringRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckConnectionStringRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckConnectionStringRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckConnectionStringRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckConnectionStringRequest) SetConnectionStringPrefix(v string) *CheckConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CheckConnectionStringRequest) SetDBClusterId(v string) *CheckConnectionStringRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckConnectionStringRequest) SetOwnerAccount(v string) *CheckConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckConnectionStringRequest) SetOwnerId(v int64) *CheckConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckConnectionStringRequest) SetResourceOwnerAccount(v string) *CheckConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckConnectionStringRequest) SetResourceOwnerId(v int64) *CheckConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
