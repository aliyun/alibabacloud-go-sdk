// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretReferencesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSecretReferencesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecretReferencesRequest
	GetPageSize() *int32
}

type ListSecretReferencesRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSecretReferencesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretReferencesRequest) GoString() string {
	return s.String()
}

func (s *ListSecretReferencesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretReferencesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretReferencesRequest) SetPageNumber(v int32) *ListSecretReferencesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretReferencesRequest) SetPageSize(v int32) *ListSecretReferencesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecretReferencesRequest) Validate() error {
	return dara.Validate(s)
}
