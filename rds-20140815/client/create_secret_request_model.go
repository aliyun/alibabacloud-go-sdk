// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSecretRequest
	GetClientToken() *string
	SetDbInstanceId(v string) *CreateSecretRequest
	GetDbInstanceId() *string
	SetDbNames(v string) *CreateSecretRequest
	GetDbNames() *string
	SetDescription(v string) *CreateSecretRequest
	GetDescription() *string
	SetEngine(v string) *CreateSecretRequest
	GetEngine() *string
	SetOwnerId(v int64) *CreateSecretRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateSecretRequest
	GetPassword() *string
	SetRegionId(v string) *CreateSecretRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSecretRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateSecretRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSecretRequest
	GetResourceOwnerId() *int64
	SetSecretName(v string) *CreateSecretRequest
	GetSecretName() *string
	SetUsername(v string) *CreateSecretRequest
	GetUsername() *string
}

type CreateSecretRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance. You can call the DescribeDBInstances operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-sdfljk123***
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// users
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// The description of the credential.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The engine of the database.
	//
	// > Only MySQL is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine  *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password that is used to access the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the instance. You can call the DescribeDBInstanceAttribute operation to query the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. You can call the DescribeDBInstanceAttribute operation to query the ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmxypivk***
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the credential.
	//
	// example:
	//
	// Foo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The username that is used to access the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_jack
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSecretRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateSecretRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *CreateSecretRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecretRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateSecretRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSecretRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateSecretRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSecretRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSecretRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSecretRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateSecretRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateSecretRequest) SetClientToken(v string) *CreateSecretRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecretRequest) SetDbInstanceId(v string) *CreateSecretRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateSecretRequest) SetDbNames(v string) *CreateSecretRequest {
	s.DbNames = &v
	return s
}

func (s *CreateSecretRequest) SetDescription(v string) *CreateSecretRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretRequest) SetEngine(v string) *CreateSecretRequest {
	s.Engine = &v
	return s
}

func (s *CreateSecretRequest) SetOwnerId(v int64) *CreateSecretRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSecretRequest) SetPassword(v string) *CreateSecretRequest {
	s.Password = &v
	return s
}

func (s *CreateSecretRequest) SetRegionId(v string) *CreateSecretRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecretRequest) SetResourceGroupId(v string) *CreateSecretRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecretRequest) SetResourceOwnerAccount(v string) *CreateSecretRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSecretRequest) SetResourceOwnerId(v int64) *CreateSecretRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSecretRequest) SetSecretName(v string) *CreateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretRequest) SetUsername(v string) *CreateSecretRequest {
	s.Username = &v
	return s
}

func (s *CreateSecretRequest) Validate() error {
	return dara.Validate(s)
}
