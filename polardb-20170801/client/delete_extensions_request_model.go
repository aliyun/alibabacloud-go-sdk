// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteExtensionsRequest
	GetClientToken() *string
	SetDBClusterId(v string) *DeleteExtensionsRequest
	GetDBClusterId() *string
	SetDBNames(v string) *DeleteExtensionsRequest
	GetDBNames() *string
	SetExtensions(v string) *DeleteExtensionsRequest
	GetExtensions() *string
	SetOwnerAccount(v string) *DeleteExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteExtensionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteExtensionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteExtensionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteExtensionsRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DeleteExtensionsRequest
	GetVpcId() *string
}

type DeleteExtensionsRequest struct {
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// auto_test_db
	DBNames *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pase,hstore
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
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-************
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
	// vpc-25cdvfeq58pl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtensionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteExtensionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteExtensionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteExtensionsRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *DeleteExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *DeleteExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteExtensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExtensionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteExtensionsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteExtensionsRequest) SetClientToken(v string) *DeleteExtensionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExtensionsRequest) SetDBClusterId(v string) *DeleteExtensionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteExtensionsRequest) SetDBNames(v string) *DeleteExtensionsRequest {
	s.DBNames = &v
	return s
}

func (s *DeleteExtensionsRequest) SetExtensions(v string) *DeleteExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *DeleteExtensionsRequest) SetOwnerAccount(v string) *DeleteExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteExtensionsRequest) SetOwnerId(v int64) *DeleteExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteExtensionsRequest) SetRegionId(v string) *DeleteExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExtensionsRequest) SetResourceGroupId(v string) *DeleteExtensionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteExtensionsRequest) SetResourceOwnerAccount(v string) *DeleteExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteExtensionsRequest) SetResourceOwnerId(v int64) *DeleteExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteExtensionsRequest) SetVpcId(v string) *DeleteExtensionsRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
