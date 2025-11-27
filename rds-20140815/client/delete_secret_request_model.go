// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSecretRequest
	GetClientToken() *string
	SetDbInstanceId(v string) *DeleteSecretRequest
	GetDbInstanceId() *string
	SetEngine(v string) *DeleteSecretRequest
	GetEngine() *string
	SetOwnerId(v int64) *DeleteSecretRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSecretRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteSecretRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteSecretRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSecretRequest
	GetResourceOwnerId() *int64
	SetSecretArn(v string) *DeleteSecretRequest
	GetSecretArn() *string
	SetSecretName(v string) *DeleteSecretRequest
	GetSecretName() *string
}

type DeleteSecretRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// >  If you specify this parameter, you must also specify the **SecretName*	- parameter. parameter.
	//
	// example:
	//
	// rm-sfjdlsjxxxxx
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
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
	// The region ID. You can call the DescribeSecrets operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the credential for the created Data API account. You can call the CreateSecret operation to obtain the value of this parameter.
	//
	// >  You must specify one of the SecretArn and **SecretName*	- parameters.
	//
	// example:
	//
	// acs:rds:cn-hangzhou:1335786***:dbInstance/rm-bp1m7l3j63****
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the credential.
	//
	// > 	- You must specify one of **SecretArn*	- and SecretName.
	//
	// > 	- If you specify this parameter, you must also specify **DbInstanceId**.
	//
	// example:
	//
	// Foo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DeleteSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSecretRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DeleteSecretRequest) GetEngine() *string {
	return s.Engine
}

func (s *DeleteSecretRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSecretRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSecretRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteSecretRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSecretRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSecretRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *DeleteSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *DeleteSecretRequest) SetClientToken(v string) *DeleteSecretRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSecretRequest) SetDbInstanceId(v string) *DeleteSecretRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DeleteSecretRequest) SetEngine(v string) *DeleteSecretRequest {
	s.Engine = &v
	return s
}

func (s *DeleteSecretRequest) SetOwnerId(v int64) *DeleteSecretRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSecretRequest) SetRegionId(v string) *DeleteSecretRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecretRequest) SetResourceGroupId(v string) *DeleteSecretRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteSecretRequest) SetResourceOwnerAccount(v string) *DeleteSecretRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSecretRequest) SetResourceOwnerId(v int64) *DeleteSecretRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSecretRequest) SetSecretArn(v string) *DeleteSecretRequest {
	s.SecretArn = &v
	return s
}

func (s *DeleteSecretRequest) SetSecretName(v string) *DeleteSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretRequest) Validate() error {
	return dara.Validate(s)
}
