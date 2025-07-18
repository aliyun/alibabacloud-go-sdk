// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetTotalNum(v string) *ListUsersResponseBody
	GetTotalNum() *string
	SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody
	GetUsers() []*ListUsersResponseBodyUsers
}

type ListUsersResponseBody struct {
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *string                       `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	Users    []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetTotalNum() *string {
	return s.TotalNum
}

func (s *ListUsersResponseBody) GetUsers() []*ListUsersResponseBodyUsers {
	return s.Users
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalNum(v string) *ListUsersResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyUsers struct {
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// a***@example.net
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IdpName *string `json:"IdpName,omitempty" xml:"IdpName,omitempty"`
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// Enabled
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) GetDepartment() *string {
	return s.Department
}

func (s *ListUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyUsers) GetIdpName() *string {
	return s.IdpName
}

func (s *ListUsersResponseBodyUsers) GetPhone() *string {
	return s.Phone
}

func (s *ListUsersResponseBodyUsers) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListUsersResponseBodyUsers) GetStatus() *string {
	return s.Status
}

func (s *ListUsersResponseBodyUsers) GetUsername() *string {
	return s.Username
}

func (s *ListUsersResponseBodyUsers) SetDepartment(v string) *ListUsersResponseBodyUsers {
	s.Department = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetIdpName(v string) *ListUsersResponseBodyUsers {
	s.IdpName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhone(v string) *ListUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSaseUserId(v string) *ListUsersResponseBodyUsers {
	s.SaseUserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetStatus(v string) *ListUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUsername(v string) *ListUsersResponseBodyUsers {
	s.Username = &v
	return s
}

func (s *ListUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
