// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseUsersResponseBodyData) *DescribeLangfuseUsersResponseBody
	GetData() *DescribeLangfuseUsersResponseBodyData
	SetRequestId(v string) *DescribeLangfuseUsersResponseBody
	GetRequestId() *string
}

type DescribeLangfuseUsersResponseBody struct {
	// The returned result.
	Data *DescribeLangfuseUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C342F3DD-1FF7-55E9-A1A1-098DE07CD1A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUsersResponseBody) GetData() *DescribeLangfuseUsersResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseUsersResponseBody) SetData(v *DescribeLangfuseUsersResponseBodyData) *DescribeLangfuseUsersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseUsersResponseBody) SetRequestId(v string) *DescribeLangfuseUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseUsersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseUsersResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of users.
	Users []*DescribeLangfuseUsersResponseBodyDataUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeLangfuseUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUsersResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLangfuseUsersResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLangfuseUsersResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLangfuseUsersResponseBodyData) GetUsers() []*DescribeLangfuseUsersResponseBodyDataUsers {
	return s.Users
}

func (s *DescribeLangfuseUsersResponseBodyData) SetPageNumber(v int64) *DescribeLangfuseUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseUsersResponseBodyData) SetPageSize(v int64) *DescribeLangfuseUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseUsersResponseBodyData) SetTotalCount(v int64) *DescribeLangfuseUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeLangfuseUsersResponseBodyData) SetUsers(v []*DescribeLangfuseUsersResponseBodyDataUsers) *DescribeLangfuseUsersResponseBodyData {
	s.Users = v
	return s
}

func (s *DescribeLangfuseUsersResponseBodyData) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseUsersResponseBodyDataUsers struct {
	// The time when the user was created.
	//
	// example:
	//
	// 2026-06-01T10:03:05Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username.
	//
	// example:
	//
	// john
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeLangfuseUsersResponseBodyDataUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseUsersResponseBodyDataUsers) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeLangfuseUsersResponseBodyDataUsers) GetEmail() *string {
	return s.Email
}

func (s *DescribeLangfuseUsersResponseBodyDataUsers) GetName() *string {
	return s.Name
}

func (s *DescribeLangfuseUsersResponseBodyDataUsers) SetCreatedAt(v string) *DescribeLangfuseUsersResponseBodyDataUsers {
	s.CreatedAt = &v
	return s
}

func (s *DescribeLangfuseUsersResponseBodyDataUsers) SetEmail(v string) *DescribeLangfuseUsersResponseBodyDataUsers {
	s.Email = &v
	return s
}

func (s *DescribeLangfuseUsersResponseBodyDataUsers) SetName(v string) *DescribeLangfuseUsersResponseBodyDataUsers {
	s.Name = &v
	return s
}

func (s *DescribeLangfuseUsersResponseBodyDataUsers) Validate() error {
	return dara.Validate(s)
}
