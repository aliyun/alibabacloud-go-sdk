// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitId(v string) *ListUsersRequest
	GetOrganizationalUnitId() *string
	SetPageNumber(v int32) *ListUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUsersRequest
	GetPageSize() *int32
}

type ListUsersRequest struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersRequest) SetOrganizationalUnitId(v string) *ListUsersRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
