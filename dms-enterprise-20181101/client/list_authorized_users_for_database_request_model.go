// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *ListAuthorizedUsersForDatabaseRequest
	GetDbId() *string
	SetLogic(v bool) *ListAuthorizedUsersForDatabaseRequest
	GetLogic() *bool
	SetPageNumber(v string) *ListAuthorizedUsersForDatabaseRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAuthorizedUsersForDatabaseRequest
	GetPageSize() *string
	SetSearchKey(v string) *ListAuthorizedUsersForDatabaseRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAuthorizedUsersForDatabaseRequest
	GetTid() *int64
}

type ListAuthorizedUsersForDatabaseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 135***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// poc_test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAuthorizedUsersForDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetDbId() *string {
	return s.DbId
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAuthorizedUsersForDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetDbId(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.DbId = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetLogic(v bool) *ListAuthorizedUsersForDatabaseRequest {
	s.Logic = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetPageNumber(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetPageSize(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetSearchKey(v string) *ListAuthorizedUsersForDatabaseRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) SetTid(v int64) *ListAuthorizedUsersForDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *ListAuthorizedUsersForDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
