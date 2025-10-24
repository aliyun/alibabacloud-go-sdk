// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateCustomEntityRequest
	GetAlgorithm() *string
	SetCustomEntityInfo(v string) *CreateCustomEntityRequest
	GetCustomEntityInfo() *string
	SetCustomEntityName(v string) *CreateCustomEntityRequest
	GetCustomEntityName() *string
	SetCustomGroupId(v string) *CreateCustomEntityRequest
	GetCustomGroupId() *string
	SetOwnerAccount(v string) *CreateCustomEntityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCustomEntityRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateCustomEntityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCustomEntityRequest
	GetResourceOwnerId() *int64
}

type CreateCustomEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// { "finegrainName":"examplName" }
	CustomEntityInfo *string `json:"CustomEntityInfo,omitempty" xml:"CustomEntityInfo,omitempty"`
	// This parameter is required.
	CustomEntityName *string `json:"CustomEntityName,omitempty" xml:"CustomEntityName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomGroupId        *string `json:"CustomGroupId,omitempty" xml:"CustomGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateCustomEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomEntityRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateCustomEntityRequest) GetCustomEntityInfo() *string {
	return s.CustomEntityInfo
}

func (s *CreateCustomEntityRequest) GetCustomEntityName() *string {
	return s.CustomEntityName
}

func (s *CreateCustomEntityRequest) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *CreateCustomEntityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCustomEntityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCustomEntityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCustomEntityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCustomEntityRequest) SetAlgorithm(v string) *CreateCustomEntityRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateCustomEntityRequest) SetCustomEntityInfo(v string) *CreateCustomEntityRequest {
	s.CustomEntityInfo = &v
	return s
}

func (s *CreateCustomEntityRequest) SetCustomEntityName(v string) *CreateCustomEntityRequest {
	s.CustomEntityName = &v
	return s
}

func (s *CreateCustomEntityRequest) SetCustomGroupId(v string) *CreateCustomEntityRequest {
	s.CustomGroupId = &v
	return s
}

func (s *CreateCustomEntityRequest) SetOwnerAccount(v string) *CreateCustomEntityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCustomEntityRequest) SetOwnerId(v int64) *CreateCustomEntityRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCustomEntityRequest) SetResourceOwnerAccount(v string) *CreateCustomEntityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCustomEntityRequest) SetResourceOwnerId(v int64) *CreateCustomEntityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCustomEntityRequest) Validate() error {
	return dara.Validate(s)
}
