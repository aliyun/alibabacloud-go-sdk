// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountLinkInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v string) *AccountLinkInfo
	GetAuthenticationType() *string
	SetCreatedAt(v int64) *AccountLinkInfo
	GetCreatedAt() *int64
	SetDisplayName(v string) *AccountLinkInfo
	GetDisplayName() *string
	SetDomainId(v string) *AccountLinkInfo
	GetDomainId() *string
	SetExtra(v string) *AccountLinkInfo
	GetExtra() *string
	SetIdentity(v string) *AccountLinkInfo
	GetIdentity() *string
	SetLastLoginTime(v int64) *AccountLinkInfo
	GetLastLoginTime() *int64
	SetStatus(v string) *AccountLinkInfo
	GetStatus() *string
	SetUserId(v string) *AccountLinkInfo
	GetUserId() *string
}

type AccountLinkInfo struct {
	// The type of the account. Valid values:
	//
	// 	- mobile: mobile number
	//
	// 	- email: email address
	//
	// 	- ding: DingTalk account
	//
	// 	- ram: Alibaba Cloud Resource Access Management (RAM) user
	//
	// 	- wechat: WeCom account
	//
	// 	- ldap: Lightweight Directory Access Protocol (LDAP) account
	//
	// 	- custom: custom account
	//
	// example:
	//
	// mobile
	AuthenticationType *string `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1639762579768
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The display name of the account. For example, the unique identifier of an LDAP account is its UID, but the account display name can be the job number or other information.
	//
	// example:
	//
	// 001
	DisplayName *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The additional information about the account. If the account type is a mobile number, the value of this parameter indicates the country code. For example, the country code in the Chinese mainland is 86 and a value of 86 is returned only if authentication_type is set to mobile.
	//
	// example:
	//
	// 86
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The unique identifier of the account.
	//
	// example:
	//
	// 136***000
	Identity      *string `json:"identity,omitempty" xml:"identity,omitempty"`
	LastLoginTime *int64  `json:"last_login_time,omitempty" xml:"last_login_time,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	// The user ID of the account.
	//
	// example:
	//
	// 00016a587b62b50003deea299a4f5b50
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s AccountLinkInfo) String() string {
	return dara.Prettify(s)
}

func (s AccountLinkInfo) GoString() string {
	return s.String()
}

func (s *AccountLinkInfo) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *AccountLinkInfo) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *AccountLinkInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *AccountLinkInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *AccountLinkInfo) GetExtra() *string {
	return s.Extra
}

func (s *AccountLinkInfo) GetIdentity() *string {
	return s.Identity
}

func (s *AccountLinkInfo) GetLastLoginTime() *int64 {
	return s.LastLoginTime
}

func (s *AccountLinkInfo) GetStatus() *string {
	return s.Status
}

func (s *AccountLinkInfo) GetUserId() *string {
	return s.UserId
}

func (s *AccountLinkInfo) SetAuthenticationType(v string) *AccountLinkInfo {
	s.AuthenticationType = &v
	return s
}

func (s *AccountLinkInfo) SetCreatedAt(v int64) *AccountLinkInfo {
	s.CreatedAt = &v
	return s
}

func (s *AccountLinkInfo) SetDisplayName(v string) *AccountLinkInfo {
	s.DisplayName = &v
	return s
}

func (s *AccountLinkInfo) SetDomainId(v string) *AccountLinkInfo {
	s.DomainId = &v
	return s
}

func (s *AccountLinkInfo) SetExtra(v string) *AccountLinkInfo {
	s.Extra = &v
	return s
}

func (s *AccountLinkInfo) SetIdentity(v string) *AccountLinkInfo {
	s.Identity = &v
	return s
}

func (s *AccountLinkInfo) SetLastLoginTime(v int64) *AccountLinkInfo {
	s.LastLoginTime = &v
	return s
}

func (s *AccountLinkInfo) SetStatus(v string) *AccountLinkInfo {
	s.Status = &v
	return s
}

func (s *AccountLinkInfo) SetUserId(v string) *AccountLinkInfo {
	s.UserId = &v
	return s
}

func (s *AccountLinkInfo) Validate() error {
	return dara.Validate(s)
}
