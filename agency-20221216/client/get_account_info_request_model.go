// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetAccountInfoRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetAccountInfoRequest
	GetPageSize() *int32
	SetUid(v int64) *GetAccountInfoRequest
	GetUid() *int64
	SetUserType(v string) *GetAccountInfoRequest
	GetUserType() *string
}

type GetAccountInfoRequest struct {
	// The current page number for pagination.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page for pagination. Maximum value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UID of the distribution customer account. Either this parameter or the UserType parameter must be specified. If this parameter is empty, all distribution customer accounts of the specified UserType are queried.
	//
	// example:
	//
	// 1215848086704806
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The type of the distribution customer account. Valid values:
	//
	// - 1: End user.
	//
	// - 2: Enterprise customer.
	//
	// - 3: T2 partner.
	//
	// example:
	//
	// 1
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAccountInfoRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAccountInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAccountInfoRequest) GetUid() *int64 {
	return s.Uid
}

func (s *GetAccountInfoRequest) GetUserType() *string {
	return s.UserType
}

func (s *GetAccountInfoRequest) SetCurrentPage(v int32) *GetAccountInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAccountInfoRequest) SetPageSize(v int32) *GetAccountInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountInfoRequest) SetUid(v int64) *GetAccountInfoRequest {
	s.Uid = &v
	return s
}

func (s *GetAccountInfoRequest) SetUserType(v string) *GetAccountInfoRequest {
	s.UserType = &v
	return s
}

func (s *GetAccountInfoRequest) Validate() error {
	return dara.Validate(s)
}
