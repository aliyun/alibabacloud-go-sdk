// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedInstancesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *ListAuthorizedInstancesForUserRequest
	GetDbType() *string
	SetEnvType(v string) *ListAuthorizedInstancesForUserRequest
	GetEnvType() *string
	SetPageNumber(v string) *ListAuthorizedInstancesForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAuthorizedInstancesForUserRequest
	GetPageSize() *string
	SetSearchKey(v string) *ListAuthorizedInstancesForUserRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAuthorizedInstancesForUserRequest
	GetTid() *int64
	SetUserId(v string) *ListAuthorizedInstancesForUserRequest
	GetUserId() *string
}

type ListAuthorizedInstancesForUserRequest struct {
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

func (s ListAuthorizedInstancesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedInstancesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedInstancesForUserRequest) GetDbType() *string {
	return s.DbType
}

func (s *ListAuthorizedInstancesForUserRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListAuthorizedInstancesForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAuthorizedInstancesForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAuthorizedInstancesForUserRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAuthorizedInstancesForUserRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorizedInstancesForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListAuthorizedInstancesForUserRequest) SetDbType(v string) *ListAuthorizedInstancesForUserRequest {
	s.DbType = &v
	return s
}

func (s *ListAuthorizedInstancesForUserRequest) SetEnvType(v string) *ListAuthorizedInstancesForUserRequest {
	s.EnvType = &v
	return s
}

func (s *ListAuthorizedInstancesForUserRequest) SetPageNumber(v string) *ListAuthorizedInstancesForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedInstancesForUserRequest) SetPageSize(v string) *ListAuthorizedInstancesForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedInstancesForUserRequest) SetSearchKey(v string) *ListAuthorizedInstancesForUserRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAuthorizedInstancesForUserRequest) SetTid(v int64) *ListAuthorizedInstancesForUserRequest {
	s.Tid = &v
	return s
}

func (s *ListAuthorizedInstancesForUserRequest) SetUserId(v string) *ListAuthorizedInstancesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListAuthorizedInstancesForUserRequest) Validate() error {
	return dara.Validate(s)
}
