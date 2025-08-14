// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *ListRecognitionEntitiesRequest
	GetAlgorithm() *string
	SetLibId(v string) *ListRecognitionEntitiesRequest
	GetLibId() *string
	SetOwnerAccount(v string) *ListRecognitionEntitiesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListRecognitionEntitiesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListRecognitionEntitiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecognitionEntitiesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListRecognitionEntitiesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListRecognitionEntitiesRequest
	GetResourceOwnerId() *int64
}

type ListRecognitionEntitiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1965304870001
	LibId        *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListRecognitionEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListRecognitionEntitiesRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListRecognitionEntitiesRequest) GetLibId() *string {
	return s.LibId
}

func (s *ListRecognitionEntitiesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListRecognitionEntitiesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRecognitionEntitiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecognitionEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecognitionEntitiesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListRecognitionEntitiesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListRecognitionEntitiesRequest) SetAlgorithm(v string) *ListRecognitionEntitiesRequest {
	s.Algorithm = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) SetLibId(v string) *ListRecognitionEntitiesRequest {
	s.LibId = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) SetOwnerAccount(v string) *ListRecognitionEntitiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) SetOwnerId(v int64) *ListRecognitionEntitiesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) SetPageNumber(v int32) *ListRecognitionEntitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) SetPageSize(v int32) *ListRecognitionEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) SetResourceOwnerAccount(v string) *ListRecognitionEntitiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) SetResourceOwnerId(v int64) *ListRecognitionEntitiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRecognitionEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
