// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAuthorizationLetterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorization(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorization() *string
	SetAuthorizationLetterExpDate(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorizationLetterExpDate() *string
	SetAuthorizationLetterName(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorizationLetterName() *string
	SetAuthorizationLetterPic(v string) *CreateSmsAuthorizationLetterRequest
	GetAuthorizationLetterPic() *string
	SetOrganizationCode(v string) *CreateSmsAuthorizationLetterRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *CreateSmsAuthorizationLetterRequest
	GetOwnerId() *int64
	SetProxyAuthorization(v string) *CreateSmsAuthorizationLetterRequest
	GetProxyAuthorization() *string
	SetResourceOwnerAccount(v string) *CreateSmsAuthorizationLetterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsAuthorizationLetterRequest
	GetResourceOwnerId() *int64
	SetSignList(v []*string) *CreateSmsAuthorizationLetterRequest
	GetSignList() []*string
}

type CreateSmsAuthorizationLetterRequest struct {
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
	SignList []*string `json:"SignList,omitempty" xml:"SignList,omitempty" type:"Repeated"`
}

func (s CreateSmsAuthorizationLetterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAuthorizationLetterRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorizationLetterExpDate() *string {
	return s.AuthorizationLetterExpDate
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorizationLetterName() *string {
	return s.AuthorizationLetterName
}

func (s *CreateSmsAuthorizationLetterRequest) GetAuthorizationLetterPic() *string {
	return s.AuthorizationLetterPic
}

func (s *CreateSmsAuthorizationLetterRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *CreateSmsAuthorizationLetterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsAuthorizationLetterRequest) GetProxyAuthorization() *string {
	return s.ProxyAuthorization
}

func (s *CreateSmsAuthorizationLetterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsAuthorizationLetterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsAuthorizationLetterRequest) GetSignList() []*string {
	return s.SignList
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorization(v string) *CreateSmsAuthorizationLetterRequest {
	s.Authorization = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorizationLetterExpDate(v string) *CreateSmsAuthorizationLetterRequest {
	s.AuthorizationLetterExpDate = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorizationLetterName(v string) *CreateSmsAuthorizationLetterRequest {
	s.AuthorizationLetterName = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetAuthorizationLetterPic(v string) *CreateSmsAuthorizationLetterRequest {
	s.AuthorizationLetterPic = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetOrganizationCode(v string) *CreateSmsAuthorizationLetterRequest {
	s.OrganizationCode = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetOwnerId(v int64) *CreateSmsAuthorizationLetterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetProxyAuthorization(v string) *CreateSmsAuthorizationLetterRequest {
	s.ProxyAuthorization = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetResourceOwnerAccount(v string) *CreateSmsAuthorizationLetterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetResourceOwnerId(v int64) *CreateSmsAuthorizationLetterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) SetSignList(v []*string) *CreateSmsAuthorizationLetterRequest {
	s.SignList = v
	return s
}

func (s *CreateSmsAuthorizationLetterRequest) Validate() error {
	return dara.Validate(s)
}
