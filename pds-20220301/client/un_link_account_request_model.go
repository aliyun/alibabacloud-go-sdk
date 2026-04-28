// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnLinkAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtra(v string) *UnLinkAccountRequest
	GetExtra() *string
	SetIdentity(v string) *UnLinkAccountRequest
	GetIdentity() *string
	SetType(v string) *UnLinkAccountRequest
	GetType() *string
	SetUserId(v string) *UnLinkAccountRequest
	GetUserId() *string
}

type UnLinkAccountRequest struct {
	// Additional information for the unique account identifier. For example, when the account is a phone number, this field should be filled with the area code of the phone, such as 86 for Mainland China. If not provided, it defaults to 86.
	//
	// example:
	//
	// 1
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// Unique identifier of the account, such as a phone number
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// Account type
	//
	// mobile: Phone number
	//
	// email: Email address
	//
	// ding: DingTalk
	//
	// ram: Alibaba Cloud RAM User
	//
	// wechat: WeCom
	//
	// ldap: LDAP account
	//
	// custom: Custom account
	//
	// This parameter is required.
	//
	// example:
	//
	// mobile
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// uid1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s UnLinkAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UnLinkAccountRequest) GoString() string {
	return s.String()
}

func (s *UnLinkAccountRequest) GetExtra() *string {
	return s.Extra
}

func (s *UnLinkAccountRequest) GetIdentity() *string {
	return s.Identity
}

func (s *UnLinkAccountRequest) GetType() *string {
	return s.Type
}

func (s *UnLinkAccountRequest) GetUserId() *string {
	return s.UserId
}

func (s *UnLinkAccountRequest) SetExtra(v string) *UnLinkAccountRequest {
	s.Extra = &v
	return s
}

func (s *UnLinkAccountRequest) SetIdentity(v string) *UnLinkAccountRequest {
	s.Identity = &v
	return s
}

func (s *UnLinkAccountRequest) SetType(v string) *UnLinkAccountRequest {
	s.Type = &v
	return s
}

func (s *UnLinkAccountRequest) SetUserId(v string) *UnLinkAccountRequest {
	s.UserId = &v
	return s
}

func (s *UnLinkAccountRequest) Validate() error {
	return dara.Validate(s)
}
