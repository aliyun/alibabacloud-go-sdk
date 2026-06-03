// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListEnterpriseInfoRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListEnterpriseInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEnterpriseInfoRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListEnterpriseInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListEnterpriseInfoRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *ListEnterpriseInfoRequest
	GetSearchKey() *string
}

type ListEnterpriseInfoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SearchKey            *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s ListEnterpriseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseInfoRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListEnterpriseInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEnterpriseInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnterpriseInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListEnterpriseInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListEnterpriseInfoRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListEnterpriseInfoRequest) SetOwnerId(v int64) *ListEnterpriseInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *ListEnterpriseInfoRequest) SetPageNumber(v int32) *ListEnterpriseInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnterpriseInfoRequest) SetPageSize(v int32) *ListEnterpriseInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnterpriseInfoRequest) SetResourceOwnerAccount(v string) *ListEnterpriseInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListEnterpriseInfoRequest) SetResourceOwnerId(v int64) *ListEnterpriseInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListEnterpriseInfoRequest) SetSearchKey(v string) *ListEnterpriseInfoRequest {
	s.SearchKey = &v
	return s
}

func (s *ListEnterpriseInfoRequest) Validate() error {
	return dara.Validate(s)
}
