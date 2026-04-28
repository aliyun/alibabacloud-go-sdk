// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtra(v string) *LinkAccountRequest
	GetExtra() *string
	SetIdentity(v string) *LinkAccountRequest
	GetIdentity() *string
	SetType(v string) *LinkAccountRequest
	GetType() *string
	SetUserId(v string) *LinkAccountRequest
	GetUserId() *string
}

type LinkAccountRequest struct {
	// The additional information about the unique identifier of the account. For example, if type is set to mobile, set the value of extra to a country code. For example, a value of 86 specifies a mobile number in the Chinese mainland. If you do not specify this parameter, 86 is used by default.
	//
	// example:
	//
	// 86
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The unique identifier of the account, such as a mobile number.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyy***
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// The account type. Valid values:
	//
	// 	- mobile: a mobile number.
	//
	// 	- email: an email address.
	//
	// 	- ding: a DingTalk account.
	//
	// 	- ram: an Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: a WeCom account.
	//
	// 	- ldap: a Lightweight Directory Access Protocol (LDAP) account.
	//
	// 	- custom: a custom account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ding
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the user with which you want to associate an account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s LinkAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s LinkAccountRequest) GoString() string {
	return s.String()
}

func (s *LinkAccountRequest) GetExtra() *string {
	return s.Extra
}

func (s *LinkAccountRequest) GetIdentity() *string {
	return s.Identity
}

func (s *LinkAccountRequest) GetType() *string {
	return s.Type
}

func (s *LinkAccountRequest) GetUserId() *string {
	return s.UserId
}

func (s *LinkAccountRequest) SetExtra(v string) *LinkAccountRequest {
	s.Extra = &v
	return s
}

func (s *LinkAccountRequest) SetIdentity(v string) *LinkAccountRequest {
	s.Identity = &v
	return s
}

func (s *LinkAccountRequest) SetType(v string) *LinkAccountRequest {
	s.Type = &v
	return s
}

func (s *LinkAccountRequest) SetUserId(v string) *LinkAccountRequest {
	s.UserId = &v
	return s
}

func (s *LinkAccountRequest) Validate() error {
	return dara.Validate(s)
}
