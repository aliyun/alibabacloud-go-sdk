// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateExtensionsRequest
	GetClientToken() *string
	SetDBClusterId(v string) *UpdateExtensionsRequest
	GetDBClusterId() *string
	SetDBNames(v string) *UpdateExtensionsRequest
	GetDBNames() *string
	SetExtensions(v string) *UpdateExtensionsRequest
	GetExtensions() *string
	SetOwnerAccount(v string) *UpdateExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateExtensionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateExtensionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateExtensionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *UpdateExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateExtensionsRequest
	GetResourceOwnerId() *int64
	SetVersion(v string) *UpdateExtensionsRequest
	GetVersion() *string
	SetVpcId(v string) *UpdateExtensionsRequest
	GetVpcId() *string
}

type UpdateExtensionsRequest struct {
	// example:
	//
	// 6000170000591aed949d0f5********************
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
	// postgres_fdw
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
	// rg-**********
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
	// 7.7
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// vpc-bp1qpo0kug3a20qqe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtensionsRequest) GoString() string {
	return s.String()
}

func (s *UpdateExtensionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateExtensionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateExtensionsRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *UpdateExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *UpdateExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateExtensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateExtensionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateExtensionsRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateExtensionsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateExtensionsRequest) SetClientToken(v string) *UpdateExtensionsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExtensionsRequest) SetDBClusterId(v string) *UpdateExtensionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateExtensionsRequest) SetDBNames(v string) *UpdateExtensionsRequest {
	s.DBNames = &v
	return s
}

func (s *UpdateExtensionsRequest) SetExtensions(v string) *UpdateExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *UpdateExtensionsRequest) SetOwnerAccount(v string) *UpdateExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateExtensionsRequest) SetOwnerId(v int64) *UpdateExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateExtensionsRequest) SetRegionId(v string) *UpdateExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateExtensionsRequest) SetResourceGroupId(v string) *UpdateExtensionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateExtensionsRequest) SetResourceOwnerAccount(v string) *UpdateExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateExtensionsRequest) SetResourceOwnerId(v int64) *UpdateExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateExtensionsRequest) SetVersion(v string) *UpdateExtensionsRequest {
	s.Version = &v
	return s
}

func (s *UpdateExtensionsRequest) SetVpcId(v string) *UpdateExtensionsRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
