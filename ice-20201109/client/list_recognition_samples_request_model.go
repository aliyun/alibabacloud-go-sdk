// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionSamplesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *ListRecognitionSamplesRequest
	GetAlgorithm() *string
	SetEntityId(v string) *ListRecognitionSamplesRequest
	GetEntityId() *string
	SetLibId(v string) *ListRecognitionSamplesRequest
	GetLibId() *string
	SetOwnerAccount(v string) *ListRecognitionSamplesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListRecognitionSamplesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListRecognitionSamplesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecognitionSamplesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListRecognitionSamplesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListRecognitionSamplesRequest
	GetResourceOwnerId() *int64
}

type ListRecognitionSamplesRequest struct {
	// The type of recognition algorithm. Valid values:
	//
	// 	- landmark
	//
	// 	- object
	//
	// 	- logo
	//
	// 	- face
	//
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// e6b985c05174412dbc77c92496b7373b
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The ID of the recognition library.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxxxxxx
	LibId        *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListRecognitionSamplesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionSamplesRequest) GoString() string {
	return s.String()
}

func (s *ListRecognitionSamplesRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListRecognitionSamplesRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *ListRecognitionSamplesRequest) GetLibId() *string {
	return s.LibId
}

func (s *ListRecognitionSamplesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListRecognitionSamplesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRecognitionSamplesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecognitionSamplesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecognitionSamplesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListRecognitionSamplesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListRecognitionSamplesRequest) SetAlgorithm(v string) *ListRecognitionSamplesRequest {
	s.Algorithm = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetEntityId(v string) *ListRecognitionSamplesRequest {
	s.EntityId = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetLibId(v string) *ListRecognitionSamplesRequest {
	s.LibId = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetOwnerAccount(v string) *ListRecognitionSamplesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetOwnerId(v int64) *ListRecognitionSamplesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetPageNumber(v int32) *ListRecognitionSamplesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetPageSize(v int32) *ListRecognitionSamplesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetResourceOwnerAccount(v string) *ListRecognitionSamplesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRecognitionSamplesRequest) SetResourceOwnerId(v int64) *ListRecognitionSamplesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRecognitionSamplesRequest) Validate() error {
	return dara.Validate(s)
}
