// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAuthorizationLetterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationLetterIdListShrink(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetAuthorizationLetterIdListShrink() *string
	SetOrganizationCode(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetSignName() *string
	SetState(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetState() *string
	SetStatus(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetStatus() *string
}

type QuerySmsAuthorizationLetterShrinkRequest struct {
	// 委托授权书id列表
	AuthorizationLetterIdListShrink *string `json:"AuthorizationLetterIdList,omitempty" xml:"AuthorizationLetterIdList,omitempty"`
	// 授权方社会统一信用代码
	//
	// example:
	//
	// 9****************A
	OrganizationCode     *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名名称（支持命中签名范围查询）
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 授权书审核状态，INT:审核中，PASSED:审核通过
	//
	// example:
	//
	// PASSED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 授权书可用状态，VALID可用，INVALID不可用
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuerySmsAuthorizationLetterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAuthorizationLetterShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetAuthorizationLetterIdListShrink() *string {
	return s.AuthorizationLetterIdListShrink
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetSignName() *string {
	return s.SignName
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetState() *string {
	return s.State
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetAuthorizationLetterIdListShrink(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.AuthorizationLetterIdListShrink = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetOrganizationCode(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.OrganizationCode = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetSignName(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.SignName = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetState(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.State = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetStatus(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.Status = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
