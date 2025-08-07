// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalDatabaseNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateGlobalDatabaseNetworkRequest
	GetDBClusterId() *string
	SetEnableGlobalDomainName(v bool) *CreateGlobalDatabaseNetworkRequest
	GetEnableGlobalDomainName() *bool
	SetGDNDescription(v string) *CreateGlobalDatabaseNetworkRequest
	GetGDNDescription() *string
	SetGDNVersion(v string) *CreateGlobalDatabaseNetworkRequest
	GetGDNVersion() *string
	SetOwnerAccount(v string) *CreateGlobalDatabaseNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateGlobalDatabaseNetworkRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateGlobalDatabaseNetworkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateGlobalDatabaseNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateGlobalDatabaseNetworkRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateGlobalDatabaseNetworkRequest
	GetSecurityToken() *string
}

type CreateGlobalDatabaseNetworkRequest struct {
	// The ID of the primary cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1q76364ird*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to create a global domain name.
	//
	// example:
	//
	// false
	EnableGlobalDomainName *bool `json:"EnableGlobalDomainName,omitempty" xml:"EnableGlobalDomainName,omitempty"`
	// The description of the GDN. The description must meet the following requirements:
	//
	// 	- It cannot start with [http:// or https://.](http://https://。)
	//
	// 	- It must start with a letter.
	//
	// 	- It can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- It must be 2 to 126 characters in length.
	//
	// example:
	//
	// GDN-fortest
	GDNDescription *string `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
	GDNVersion     *string `json:"GDNVersion,omitempty" xml:"GDNVersion,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateGlobalDatabaseNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalDatabaseNetworkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateGlobalDatabaseNetworkRequest) GetEnableGlobalDomainName() *bool {
	return s.EnableGlobalDomainName
}

func (s *CreateGlobalDatabaseNetworkRequest) GetGDNDescription() *string {
	return s.GDNDescription
}

func (s *CreateGlobalDatabaseNetworkRequest) GetGDNVersion() *string {
	return s.GDNVersion
}

func (s *CreateGlobalDatabaseNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateGlobalDatabaseNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGlobalDatabaseNetworkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGlobalDatabaseNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateGlobalDatabaseNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateGlobalDatabaseNetworkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateGlobalDatabaseNetworkRequest) SetDBClusterId(v string) *CreateGlobalDatabaseNetworkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetEnableGlobalDomainName(v bool) *CreateGlobalDatabaseNetworkRequest {
	s.EnableGlobalDomainName = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetGDNDescription(v string) *CreateGlobalDatabaseNetworkRequest {
	s.GDNDescription = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetGDNVersion(v string) *CreateGlobalDatabaseNetworkRequest {
	s.GDNVersion = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *CreateGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *CreateGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetResourceGroupId(v string) *CreateGlobalDatabaseNetworkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *CreateGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *CreateGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *CreateGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) Validate() error {
	return dara.Validate(s)
}
