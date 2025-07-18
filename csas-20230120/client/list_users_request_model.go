// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListUsersRequest
	GetCurrentPage() *int64
	SetDepartment(v string) *ListUsersRequest
	GetDepartment() *string
	SetFuzzyUsername(v string) *ListUsersRequest
	GetFuzzyUsername() *string
	SetPageSize(v int64) *ListUsersRequest
	GetPageSize() *int64
	SetPreciseUsername(v string) *ListUsersRequest
	GetPreciseUsername() *string
	SetSaseUserIds(v []*string) *ListUsersRequest
	GetSaseUserIds() []*string
	SetStatus(v string) *ListUsersRequest
	GetStatus() *string
}

type ListUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage   *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	FuzzyUsername *string `json:"FuzzyUsername,omitempty" xml:"FuzzyUsername,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize        *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PreciseUsername *string   `json:"PreciseUsername,omitempty" xml:"PreciseUsername,omitempty"`
	SaseUserIds     []*string `json:"SaseUserIds,omitempty" xml:"SaseUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListUsersRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListUsersRequest) GetFuzzyUsername() *string {
	return s.FuzzyUsername
}

func (s *ListUsersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUsersRequest) GetPreciseUsername() *string {
	return s.PreciseUsername
}

func (s *ListUsersRequest) GetSaseUserIds() []*string {
	return s.SaseUserIds
}

func (s *ListUsersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUsersRequest) SetCurrentPage(v int64) *ListUsersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUsersRequest) SetDepartment(v string) *ListUsersRequest {
	s.Department = &v
	return s
}

func (s *ListUsersRequest) SetFuzzyUsername(v string) *ListUsersRequest {
	s.FuzzyUsername = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetPreciseUsername(v string) *ListUsersRequest {
	s.PreciseUsername = &v
	return s
}

func (s *ListUsersRequest) SetSaseUserIds(v []*string) *ListUsersRequest {
	s.SaseUserIds = v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
