// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceForTerraformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2InstanceForTerraformRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormV2InstanceForTerraformRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormV2InstanceForTerraformRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetLindormV2InstanceForTerraformRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormV2InstanceForTerraformRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormV2InstanceForTerraformRequest
	GetSecurityToken() *string
}

type GetLindormV2InstanceForTerraformRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2InstanceForTerraformRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceForTerraformRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceForTerraformRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceForTerraformRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormV2InstanceForTerraformRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormV2InstanceForTerraformRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormV2InstanceForTerraformRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormV2InstanceForTerraformRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormV2InstanceForTerraformRequest) SetInstanceId(v string) *GetLindormV2InstanceForTerraformRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformRequest) SetOwnerAccount(v string) *GetLindormV2InstanceForTerraformRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformRequest) SetOwnerId(v int64) *GetLindormV2InstanceForTerraformRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformRequest) SetResourceOwnerAccount(v string) *GetLindormV2InstanceForTerraformRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformRequest) SetResourceOwnerId(v int64) *GetLindormV2InstanceForTerraformRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformRequest) SetSecurityToken(v string) *GetLindormV2InstanceForTerraformRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormV2InstanceForTerraformRequest) Validate() error {
	return dara.Validate(s)
}
