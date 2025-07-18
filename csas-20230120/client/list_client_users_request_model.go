// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListClientUsersRequest
	GetCurrentPage() *int64
	SetDepartmentId(v string) *ListClientUsersRequest
	GetDepartmentId() *string
	SetEmail(v string) *ListClientUsersRequest
	GetEmail() *string
	SetIdpConfigId(v string) *ListClientUsersRequest
	GetIdpConfigId() *string
	SetMobileNumber(v string) *ListClientUsersRequest
	GetMobileNumber() *string
	SetPageSize(v int64) *ListClientUsersRequest
	GetPageSize() *int64
	SetStatus(v string) *ListClientUsersRequest
	GetStatus() *string
	SetUsername(v string) *ListClientUsersRequest
	GetUsername() *string
}

type ListClientUsersRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10785
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1071
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 18980976559
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Enabled
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListClientUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientUsersRequest) GoString() string {
	return s.String()
}

func (s *ListClientUsersRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListClientUsersRequest) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *ListClientUsersRequest) GetEmail() *string {
	return s.Email
}

func (s *ListClientUsersRequest) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *ListClientUsersRequest) GetMobileNumber() *string {
	return s.MobileNumber
}

func (s *ListClientUsersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListClientUsersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListClientUsersRequest) GetUsername() *string {
	return s.Username
}

func (s *ListClientUsersRequest) SetCurrentPage(v int64) *ListClientUsersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClientUsersRequest) SetDepartmentId(v string) *ListClientUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListClientUsersRequest) SetEmail(v string) *ListClientUsersRequest {
	s.Email = &v
	return s
}

func (s *ListClientUsersRequest) SetIdpConfigId(v string) *ListClientUsersRequest {
	s.IdpConfigId = &v
	return s
}

func (s *ListClientUsersRequest) SetMobileNumber(v string) *ListClientUsersRequest {
	s.MobileNumber = &v
	return s
}

func (s *ListClientUsersRequest) SetPageSize(v int64) *ListClientUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClientUsersRequest) SetStatus(v string) *ListClientUsersRequest {
	s.Status = &v
	return s
}

func (s *ListClientUsersRequest) SetUsername(v string) *ListClientUsersRequest {
	s.Username = &v
	return s
}

func (s *ListClientUsersRequest) Validate() error {
	return dara.Validate(s)
}
