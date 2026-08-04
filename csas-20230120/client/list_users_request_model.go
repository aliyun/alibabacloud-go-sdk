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
	// The page number. Valid values: 1 to 10,000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The user\\"s department. The value must be 1 to 128 characters long and can contain Chinese characters, letters, digits, periods (.), commas (,), semicolons (;), hyphens (-), underscores (_), slashes (/), at signs (@), and spaces.
	//
	// example:
	//
	// 测试部
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The username for a fuzzy match. The value must be 1 to 128 characters long and can contain Chinese characters, letters, digits, periods (.), underscores (_), hyphens (-), asterisks (\\*), at signs (@), and spaces.
	//
	// example:
	//
	// 王先生
	FuzzyUsername *string `json:"FuzzyUsername,omitempty" xml:"FuzzyUsername,omitempty"`
	// The number of entries per page. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The username for an exact match. The value must be 1 to 128 characters long and can contain Chinese characters, letters, digits, periods (.), underscores (_), hyphens (-), asterisks (\\*), at signs (@), and spaces.
	//
	// example:
	//
	// 王先生@alibaba.com
	PreciseUsername *string `json:"PreciseUsername,omitempty" xml:"PreciseUsername,omitempty"`
	// An array of user IDs.
	SaseUserIds []*string `json:"SaseUserIds,omitempty" xml:"SaseUserIds,omitempty" type:"Repeated"`
	// The status of the user. Valid values:
	//
	// - **Enabled**: The user is enabled.
	//
	// - **Disabled**: The user is disabled.
	//
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
