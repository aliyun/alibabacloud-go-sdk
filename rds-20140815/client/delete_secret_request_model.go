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
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// This parameter is required.
	Engine  *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretArn            *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	SecretName           *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
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
