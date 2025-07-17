// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersForInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAuthorizedUsersForInstanceRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListAuthorizedUsersForInstanceRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAuthorizedUsersForInstanceRequest
	GetPageSize() *string
	SetSearchKey(v string) *ListAuthorizedUsersForInstanceRequest
	GetSearchKey() *string
	SetTid(v int64) *ListAuthorizedUsersForInstanceRequest
	GetTid() *int64
}

type ListAuthorizedUsersForInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 174****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListAuthorizedUsersForInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersForInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersForInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizedUsersForInstanceRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAuthorizedUsersForInstanceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAuthorizedUsersForInstanceRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAuthorizedUsersForInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListAuthorizedUsersForInstanceRequest) SetInstanceId(v string) *ListAuthorizedUsersForInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetPageNumber(v string) *ListAuthorizedUsersForInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetPageSize(v string) *ListAuthorizedUsersForInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetSearchKey(v string) *ListAuthorizedUsersForInstanceRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) SetTid(v int64) *ListAuthorizedUsersForInstanceRequest {
	s.Tid = &v
	return s
}

func (s *ListAuthorizedUsersForInstanceRequest) Validate() error {
	return dara.Validate(s)
}
