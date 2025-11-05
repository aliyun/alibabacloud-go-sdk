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
	// Pagination, current page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Pagination, record number on each page, maximum 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Account UID of Distribution Customer. This parameter and the UserType parameter must have one filled. If this parameter is empty, then check all Distribution Customer accounts of the selected UserType.
	//
	// example:
	//
	// 1215848086704806
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// Distribution Customer\\"s Account Type:
	//
	// - 1 End User
	//
	// - 2 Enterprise
	//
	// - 3 T2 Partner
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
