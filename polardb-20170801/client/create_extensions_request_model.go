// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateExtensionsRequest
	GetAccountName() *string
	SetClientToken(v string) *CreateExtensionsRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateExtensionsRequest
	GetDBClusterId() *string
	SetDBNames(v string) *CreateExtensionsRequest
	GetDBNames() *string
	SetExtensions(v string) *CreateExtensionsRequest
	GetExtensions() *string
	SetOwnerAccount(v string) *CreateExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateExtensionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateExtensionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateExtensionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateExtensionsRequest
	GetResourceOwnerId() *int64
	SetSourceDBName(v string) *CreateExtensionsRequest
	GetSourceDBName() *string
	SetVersion(v string) *CreateExtensionsRequest
	GetVersion() *string
	SetVpcId(v string) *CreateExtensionsRequest
	GetVpcId() *string
}

type CreateExtensionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testuser
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb1,testdb2
	DBNames *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	// example:
	//
	// pg_stat_statements,pg_trgm
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// example:
	//
	// test@example.com
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 1234567890123456
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// test@example.com
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 1234567890123456
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// sourcedb
	SourceDBName *string `json:"SourceDBName,omitempty" xml:"SourceDBName,omitempty"`
	// example:
	//
	// 7.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// vpc-****************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExtensionsRequest) GoString() string {
	return s.String()
}

func (s *CreateExtensionsRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateExtensionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExtensionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateExtensionsRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *CreateExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *CreateExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateExtensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExtensionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateExtensionsRequest) GetSourceDBName() *string {
	return s.SourceDBName
}

func (s *CreateExtensionsRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateExtensionsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateExtensionsRequest) SetAccountName(v string) *CreateExtensionsRequest {
	s.AccountName = &v
	return s
}

func (s *CreateExtensionsRequest) SetClientToken(v string) *CreateExtensionsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExtensionsRequest) SetDBClusterId(v string) *CreateExtensionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateExtensionsRequest) SetDBNames(v string) *CreateExtensionsRequest {
	s.DBNames = &v
	return s
}

func (s *CreateExtensionsRequest) SetExtensions(v string) *CreateExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *CreateExtensionsRequest) SetOwnerAccount(v string) *CreateExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateExtensionsRequest) SetOwnerId(v int64) *CreateExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateExtensionsRequest) SetRegionId(v string) *CreateExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExtensionsRequest) SetResourceGroupId(v string) *CreateExtensionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExtensionsRequest) SetResourceOwnerAccount(v string) *CreateExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateExtensionsRequest) SetResourceOwnerId(v int64) *CreateExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateExtensionsRequest) SetSourceDBName(v string) *CreateExtensionsRequest {
	s.SourceDBName = &v
	return s
}

func (s *CreateExtensionsRequest) SetVersion(v string) *CreateExtensionsRequest {
	s.Version = &v
	return s
}

func (s *CreateExtensionsRequest) SetVpcId(v string) *CreateExtensionsRequest {
	s.VpcId = &v
	return s
}

func (s *CreateExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
