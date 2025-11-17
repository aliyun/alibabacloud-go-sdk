// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionLibsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *ListRecognitionLibsRequest
	GetAlgorithm() *string
	SetLibId(v string) *ListRecognitionLibsRequest
	GetLibId() *string
	SetOwnerAccount(v string) *ListRecognitionLibsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListRecognitionLibsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListRecognitionLibsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecognitionLibsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListRecognitionLibsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListRecognitionLibsRequest
	GetResourceOwnerId() *int64
}

type ListRecognitionLibsRequest struct {
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
	// 	- label
	//
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
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
	// The number of entries per page. Valid values: 1 to 50.
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

func (s ListRecognitionLibsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionLibsRequest) GoString() string {
	return s.String()
}

func (s *ListRecognitionLibsRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListRecognitionLibsRequest) GetLibId() *string {
	return s.LibId
}

func (s *ListRecognitionLibsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListRecognitionLibsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListRecognitionLibsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecognitionLibsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecognitionLibsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListRecognitionLibsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListRecognitionLibsRequest) SetAlgorithm(v string) *ListRecognitionLibsRequest {
	s.Algorithm = &v
	return s
}

func (s *ListRecognitionLibsRequest) SetLibId(v string) *ListRecognitionLibsRequest {
	s.LibId = &v
	return s
}

func (s *ListRecognitionLibsRequest) SetOwnerAccount(v string) *ListRecognitionLibsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListRecognitionLibsRequest) SetOwnerId(v int64) *ListRecognitionLibsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRecognitionLibsRequest) SetPageNumber(v int32) *ListRecognitionLibsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecognitionLibsRequest) SetPageSize(v int32) *ListRecognitionLibsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecognitionLibsRequest) SetResourceOwnerAccount(v string) *ListRecognitionLibsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRecognitionLibsRequest) SetResourceOwnerId(v int64) *ListRecognitionLibsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRecognitionLibsRequest) Validate() error {
	return dara.Validate(s)
}
