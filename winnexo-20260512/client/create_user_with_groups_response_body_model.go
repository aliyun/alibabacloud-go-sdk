// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWithGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateUserWithGroupsResponseBody
	GetAccountId() *string
	SetCode(v string) *CreateUserWithGroupsResponseBody
	GetCode() *string
	SetDisplayName(v string) *CreateUserWithGroupsResponseBody
	GetDisplayName() *string
	SetIsNewUser(v bool) *CreateUserWithGroupsResponseBody
	GetIsNewUser() *bool
	SetMessage(v string) *CreateUserWithGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUserWithGroupsResponseBody
	GetRequestId() *string
	SetWnUserId(v string) *CreateUserWithGroupsResponseBody
	GetWnUserId() *string
}

type CreateUserWithGroupsResponseBody struct {
	// The WINNEXO logon account.
	//
	// example:
	//
	// exampleAccountId
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- or InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Indicates whether the user is newly created. A value of false indicates that an existing user joined the tenant.
	//
	// example:
	//
	// true
	IsNewUser *bool `json:"isNewUser,omitempty" xml:"isNewUser,omitempty"`
	// The error description. This parameter is empty when the request succeeds.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The WINNEXO platform user ID.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s CreateUserWithGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWithGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserWithGroupsResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateUserWithGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateUserWithGroupsResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserWithGroupsResponseBody) GetIsNewUser() *bool {
	return s.IsNewUser
}

func (s *CreateUserWithGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserWithGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserWithGroupsResponseBody) GetWnUserId() *string {
	return s.WnUserId
}

func (s *CreateUserWithGroupsResponseBody) SetAccountId(v string) *CreateUserWithGroupsResponseBody {
	s.AccountId = &v
	return s
}

func (s *CreateUserWithGroupsResponseBody) SetCode(v string) *CreateUserWithGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserWithGroupsResponseBody) SetDisplayName(v string) *CreateUserWithGroupsResponseBody {
	s.DisplayName = &v
	return s
}

func (s *CreateUserWithGroupsResponseBody) SetIsNewUser(v bool) *CreateUserWithGroupsResponseBody {
	s.IsNewUser = &v
	return s
}

func (s *CreateUserWithGroupsResponseBody) SetMessage(v string) *CreateUserWithGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserWithGroupsResponseBody) SetRequestId(v string) *CreateUserWithGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserWithGroupsResponseBody) SetWnUserId(v string) *CreateUserWithGroupsResponseBody {
	s.WnUserId = &v
	return s
}

func (s *CreateUserWithGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
