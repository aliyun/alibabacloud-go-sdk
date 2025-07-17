// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedDatabasesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *ListAuthorizedDatabasesForUserRequest
	GetDbType() *string
	SetEnvType(v string) *ListAuthorizedDatabasesForUserRequest
	GetEnvType() *string
	SetLogic(v bool) *ListAuthorizedDatabasesForUserRequest
	GetLogic() *bool
	SetPageNumber(v string) *ListAuthorizedDatabasesForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAuthorizedDatabasesForUserRequest
	GetPageSize() *string
	SetSearchKey(v string) *ListAuthorizedDatabasesForUserRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAuthorizedDatabasesForUserRequest
	GetTid() *int64
	SetUserId(v string) *ListAuthorizedDatabasesForUserRequest
	GetUserId() *string
}

type ListAuthorizedDatabasesForUserRequest struct {
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
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
	// policy_test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAuthorizedDatabasesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDatabasesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDatabasesForUserRequest) GetDbType() *string {
	return s.DbType
}

func (s *ListAuthorizedDatabasesForUserRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListAuthorizedDatabasesForUserRequest) GetLogic() *bool {
	return s.Logic
}

func (s *ListAuthorizedDatabasesForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAuthorizedDatabasesForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAuthorizedDatabasesForUserRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAuthorizedDatabasesForUserRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorizedDatabasesForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListAuthorizedDatabasesForUserRequest) SetDbType(v string) *ListAuthorizedDatabasesForUserRequest {
	s.DbType = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) SetEnvType(v string) *ListAuthorizedDatabasesForUserRequest {
	s.EnvType = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) SetLogic(v bool) *ListAuthorizedDatabasesForUserRequest {
	s.Logic = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) SetPageNumber(v string) *ListAuthorizedDatabasesForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) SetPageSize(v string) *ListAuthorizedDatabasesForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) SetSearchKey(v string) *ListAuthorizedDatabasesForUserRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) SetTid(v int64) *ListAuthorizedDatabasesForUserRequest {
	s.Tid = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) SetUserId(v string) *ListAuthorizedDatabasesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserRequest) Validate() error {
	return dara.Validate(s)
}
