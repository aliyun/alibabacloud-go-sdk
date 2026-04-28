// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtra(v string) *GetLinkInfoRequest
	GetExtra() *string
	SetIdentity(v string) *GetLinkInfoRequest
	GetIdentity() *string
	SetType(v string) *GetLinkInfoRequest
	GetType() *string
}

type GetLinkInfoRequest struct {
	// The additional information about the unique identifier of the account. For example, if type is set to mobile, set the value of extra to a country code.
	//
	// example:
	//
	// 1
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The unique identifier of the account, such as a mobile number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130***
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
	// mobile
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetLinkInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLinkInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLinkInfoRequest) GetExtra() *string {
	return s.Extra
}

func (s *GetLinkInfoRequest) GetIdentity() *string {
	return s.Identity
}

func (s *GetLinkInfoRequest) GetType() *string {
	return s.Type
}

func (s *GetLinkInfoRequest) SetExtra(v string) *GetLinkInfoRequest {
	s.Extra = &v
	return s
}

func (s *GetLinkInfoRequest) SetIdentity(v string) *GetLinkInfoRequest {
	s.Identity = &v
	return s
}

func (s *GetLinkInfoRequest) SetType(v string) *GetLinkInfoRequest {
	s.Type = &v
	return s
}

func (s *GetLinkInfoRequest) Validate() error {
	return dara.Validate(s)
}
