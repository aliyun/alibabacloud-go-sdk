// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAuthorizationLetterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorization(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetAuthorization() *string
	SetAuthorizationLetterExpDate(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetAuthorizationLetterExpDate() *string
	SetAuthorizationLetterName(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetAuthorizationLetterName() *string
	SetAuthorizationLetterPic(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetAuthorizationLetterPic() *string
	SetOrganizationCode(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *CreateSmsAuthorizationLetterShrinkRequest
	GetOwnerId() *int64
	SetProxyAuthorization(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetProxyAuthorization() *string
	SetResourceOwnerAccount(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsAuthorizationLetterShrinkRequest
	GetResourceOwnerId() *int64
	SetSignListShrink(v string) *CreateSmsAuthorizationLetterShrinkRequest
	GetSignListShrink() *string
}

type CreateSmsAuthorizationLetterShrinkRequest struct {
	// 授权方，授权方命名长度不超过1000个字符，暂不支持包含除中点（·）、空格、中文括号【】、英文括号()外的任何符号或纯数字输入
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云有限公司
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// 委托授权书有效期
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01~2026-01-01
	AuthorizationLetterExpDate *string `json:"AuthorizationLetterExpDate,omitempty" xml:"AuthorizationLetterExpDate,omitempty"`
	// 委托授权书命名非空，不超过100个字符，支持中文、英文或与数字组合进行命名，暂不支持任何符号或纯数字输入
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx公司授权书
	AuthorizationLetterName *string `json:"AuthorizationLetterName,omitempty" xml:"AuthorizationLetterName,omitempty"`
	// 上传oss的委托授权书图片标识
	//
	// This parameter is required.
	AuthorizationLetterPic *string `json:"AuthorizationLetterPic,omitempty" xml:"AuthorizationLetterPic,omitempty"`
	// 授权方社会统一信用代码，长度不超过150个字符
	//
	// This parameter is required.
	//
	// example:
	//
	// 9****************A
	OrganizationCode *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 被授权方，被授权方命名长度不超过1000个字符，暂不支持包含除中点（·）、空格、中文括号【】、英文括号()外的任何符号或纯数字输入
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx公司
	ProxyAuthorization   *string `json:"ProxyAuthorization,omitempty" xml:"ProxyAuthorization,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 委托授权签名列表，签名数量限制100个以内
	//
	// This parameter is required.
	SignListShrink *string `json:"SignList,omitempty" xml:"SignList,omitempty"`
}

func (s CreateSmsAuthorizationLetterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAuthorizationLetterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetAuthorizationLetterExpDate() *string {
	return s.AuthorizationLetterExpDate
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetAuthorizationLetterName() *string {
	return s.AuthorizationLetterName
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetAuthorizationLetterPic() *string {
	return s.AuthorizationLetterPic
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetProxyAuthorization() *string {
	return s.ProxyAuthorization
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) GetSignListShrink() *string {
	return s.SignListShrink
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetAuthorization(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.Authorization = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetAuthorizationLetterExpDate(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.AuthorizationLetterExpDate = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetAuthorizationLetterName(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.AuthorizationLetterName = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetAuthorizationLetterPic(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.AuthorizationLetterPic = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetOrganizationCode(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.OrganizationCode = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetOwnerId(v int64) *CreateSmsAuthorizationLetterShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetProxyAuthorization(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.ProxyAuthorization = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetResourceOwnerAccount(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetResourceOwnerId(v int64) *CreateSmsAuthorizationLetterShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) SetSignListShrink(v string) *CreateSmsAuthorizationLetterShrinkRequest {
	s.SignListShrink = &v
	return s
}

func (s *CreateSmsAuthorizationLetterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
