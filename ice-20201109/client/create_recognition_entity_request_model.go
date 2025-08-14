// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateRecognitionEntityRequest
	GetAlgorithm() *string
	SetEntityInfo(v string) *CreateRecognitionEntityRequest
	GetEntityInfo() *string
	SetEntityName(v string) *CreateRecognitionEntityRequest
	GetEntityName() *string
	SetLibId(v string) *CreateRecognitionEntityRequest
	GetLibId() *string
	SetOwnerAccount(v string) *CreateRecognitionEntityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRecognitionEntityRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateRecognitionEntityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRecognitionEntityRequest
	GetResourceOwnerId() *int64
}

type CreateRecognitionEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm  *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	EntityInfo *string `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	// This parameter is required.
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *************24b47865c6**************
	LibId                *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateRecognitionEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateRecognitionEntityRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateRecognitionEntityRequest) GetEntityInfo() *string {
	return s.EntityInfo
}

func (s *CreateRecognitionEntityRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *CreateRecognitionEntityRequest) GetLibId() *string {
	return s.LibId
}

func (s *CreateRecognitionEntityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRecognitionEntityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRecognitionEntityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRecognitionEntityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRecognitionEntityRequest) SetAlgorithm(v string) *CreateRecognitionEntityRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateRecognitionEntityRequest) SetEntityInfo(v string) *CreateRecognitionEntityRequest {
	s.EntityInfo = &v
	return s
}

func (s *CreateRecognitionEntityRequest) SetEntityName(v string) *CreateRecognitionEntityRequest {
	s.EntityName = &v
	return s
}

func (s *CreateRecognitionEntityRequest) SetLibId(v string) *CreateRecognitionEntityRequest {
	s.LibId = &v
	return s
}

func (s *CreateRecognitionEntityRequest) SetOwnerAccount(v string) *CreateRecognitionEntityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRecognitionEntityRequest) SetOwnerId(v int64) *CreateRecognitionEntityRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRecognitionEntityRequest) SetResourceOwnerAccount(v string) *CreateRecognitionEntityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRecognitionEntityRequest) SetResourceOwnerId(v int64) *CreateRecognitionEntityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRecognitionEntityRequest) Validate() error {
	return dara.Validate(s)
}
