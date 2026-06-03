// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListSchemeRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListSchemeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSchemeRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListSchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListSchemeRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *ListSchemeRequest
	GetSearchKey() *string
}

type ListSchemeRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SearchKey            *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s ListSchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeRequest) GoString() string {
	return s.String()
}

func (s *ListSchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSchemeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSchemeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListSchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListSchemeRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListSchemeRequest) SetOwnerId(v int64) *ListSchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSchemeRequest) SetPageNumber(v int32) *ListSchemeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSchemeRequest) SetPageSize(v int32) *ListSchemeRequest {
	s.PageSize = &v
	return s
}

func (s *ListSchemeRequest) SetResourceOwnerAccount(v string) *ListSchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSchemeRequest) SetResourceOwnerId(v int64) *ListSchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSchemeRequest) SetSearchKey(v string) *ListSchemeRequest {
	s.SearchKey = &v
	return s
}

func (s *ListSchemeRequest) Validate() error {
	return dara.Validate(s)
}
