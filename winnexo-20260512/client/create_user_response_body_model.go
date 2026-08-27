// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateUserResponseBody
	GetAccountId() *string
	SetCode(v string) *CreateUserResponseBody
	GetCode() *string
	SetDisplayName(v string) *CreateUserResponseBody
	GetDisplayName() *string
	SetIsNewUser(v bool) *CreateUserResponseBody
	GetIsNewUser() *bool
	SetMessage(v string) *CreateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUserResponseBody
	GetRequestId() *string
	SetWnUserId(v string) *CreateUserResponseBody
	GetWnUserId() *string
}

type CreateUserResponseBody struct {
	// The ID of your Alibaba Cloud account.
	//
	// example:
	//
	// exampleAccountId
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Indicates whether the user is newly created. A value of false indicates that an existing user is added to the tenant.
	//
	// example:
	//
	// true
	IsNewUser *bool `json:"isNewUser,omitempty" xml:"isNewUser,omitempty"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
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

func (s CreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateUserResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserResponseBody) GetIsNewUser() *bool {
	return s.IsNewUser
}

func (s *CreateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserResponseBody) GetWnUserId() *string {
	return s.WnUserId
}

func (s *CreateUserResponseBody) SetAccountId(v string) *CreateUserResponseBody {
	s.AccountId = &v
	return s
}

func (s *CreateUserResponseBody) SetCode(v string) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetDisplayName(v string) *CreateUserResponseBody {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseBody) SetIsNewUser(v bool) *CreateUserResponseBody {
	s.IsNewUser = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetWnUserId(v string) *CreateUserResponseBody {
	s.WnUserId = &v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
