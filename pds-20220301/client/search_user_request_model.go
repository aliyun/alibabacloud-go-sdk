// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *SearchUserRequest
	GetEmail() *string
	SetLimit(v int32) *SearchUserRequest
	GetLimit() *int32
	SetMarker(v string) *SearchUserRequest
	GetMarker() *string
	SetNickName(v string) *SearchUserRequest
	GetNickName() *string
	SetNickNameForFuzzy(v string) *SearchUserRequest
	GetNickNameForFuzzy() *string
	SetPhone(v string) *SearchUserRequest
	GetPhone() *string
	SetRole(v string) *SearchUserRequest
	GetRole() *string
	SetStatus(v string) *SearchUserRequest
	GetStatus() *string
	SetUserName(v string) *SearchUserRequest
	GetUserName() *string
}

type SearchUserRequest struct {
	// The email address of the user.
	//
	// example:
	//
	// 123@pds.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The nickname of the user. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The nickname used for fuzzy searches. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// la
	NickNameForFuzzy *string `json:"nick_name_for_fuzzy,omitempty" xml:"nick_name_for_fuzzy,omitempty"`
	// The mobile number of the user.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Valid values:
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// 	- enabled: The user is in a normal state.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The name of the user. The name can be up to 128 characters in length.
	//
	// example:
	//
	// pds
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s SearchUserRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchUserRequest) GoString() string {
	return s.String()
}

func (s *SearchUserRequest) GetEmail() *string {
	return s.Email
}

func (s *SearchUserRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchUserRequest) GetMarker() *string {
	return s.Marker
}

func (s *SearchUserRequest) GetNickName() *string {
	return s.NickName
}

func (s *SearchUserRequest) GetNickNameForFuzzy() *string {
	return s.NickNameForFuzzy
}

func (s *SearchUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *SearchUserRequest) GetRole() *string {
	return s.Role
}

func (s *SearchUserRequest) GetStatus() *string {
	return s.Status
}

func (s *SearchUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *SearchUserRequest) SetEmail(v string) *SearchUserRequest {
	s.Email = &v
	return s
}

func (s *SearchUserRequest) SetLimit(v int32) *SearchUserRequest {
	s.Limit = &v
	return s
}

func (s *SearchUserRequest) SetMarker(v string) *SearchUserRequest {
	s.Marker = &v
	return s
}

func (s *SearchUserRequest) SetNickName(v string) *SearchUserRequest {
	s.NickName = &v
	return s
}

func (s *SearchUserRequest) SetNickNameForFuzzy(v string) *SearchUserRequest {
	s.NickNameForFuzzy = &v
	return s
}

func (s *SearchUserRequest) SetPhone(v string) *SearchUserRequest {
	s.Phone = &v
	return s
}

func (s *SearchUserRequest) SetRole(v string) *SearchUserRequest {
	s.Role = &v
	return s
}

func (s *SearchUserRequest) SetStatus(v string) *SearchUserRequest {
	s.Status = &v
	return s
}

func (s *SearchUserRequest) SetUserName(v string) *SearchUserRequest {
	s.UserName = &v
	return s
}

func (s *SearchUserRequest) Validate() error {
	return dara.Validate(s)
}
