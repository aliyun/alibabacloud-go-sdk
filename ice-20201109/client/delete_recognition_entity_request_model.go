// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DeleteRecognitionEntityRequest
	GetAlgorithm() *string
	SetEntityId(v string) *DeleteRecognitionEntityRequest
	GetEntityId() *string
	SetLibId(v string) *DeleteRecognitionEntityRequest
	GetLibId() *string
	SetOwnerAccount(v string) *DeleteRecognitionEntityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRecognitionEntityRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteRecognitionEntityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRecognitionEntityRequest
	GetResourceOwnerId() *int64
}

type DeleteRecognitionEntityRequest struct {
	// The type of recognition algorithm associated with the entity. Valid values:
	//
	// 	- landmark
	//
	// 	- object
	//
	// 	- logo
	//
	// 	- face
	//
	// 	- label
	//
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the entity to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// **************544cb84754************
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The ID of the recognition library.
	//
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

func (s DeleteRecognitionEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionEntityRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DeleteRecognitionEntityRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *DeleteRecognitionEntityRequest) GetLibId() *string {
	return s.LibId
}

func (s *DeleteRecognitionEntityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRecognitionEntityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRecognitionEntityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRecognitionEntityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRecognitionEntityRequest) SetAlgorithm(v string) *DeleteRecognitionEntityRequest {
	s.Algorithm = &v
	return s
}

func (s *DeleteRecognitionEntityRequest) SetEntityId(v string) *DeleteRecognitionEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteRecognitionEntityRequest) SetLibId(v string) *DeleteRecognitionEntityRequest {
	s.LibId = &v
	return s
}

func (s *DeleteRecognitionEntityRequest) SetOwnerAccount(v string) *DeleteRecognitionEntityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRecognitionEntityRequest) SetOwnerId(v int64) *DeleteRecognitionEntityRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRecognitionEntityRequest) SetResourceOwnerAccount(v string) *DeleteRecognitionEntityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRecognitionEntityRequest) SetResourceOwnerId(v int64) *DeleteRecognitionEntityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRecognitionEntityRequest) Validate() error {
	return dara.Validate(s)
}
