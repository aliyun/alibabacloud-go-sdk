// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAuthorizationLetterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationLetterIdList(v []*int64) *QuerySmsAuthorizationLetterRequest
	GetAuthorizationLetterIdList() []*int64
	SetOrganizationCode(v string) *QuerySmsAuthorizationLetterRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *QuerySmsAuthorizationLetterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QuerySmsAuthorizationLetterRequest
	GetSignName() *string
	SetState(v string) *QuerySmsAuthorizationLetterRequest
	GetState() *string
	SetStatus(v string) *QuerySmsAuthorizationLetterRequest
	GetStatus() *string
}

type QuerySmsAuthorizationLetterRequest struct {
	// 委托授权书id列表
	AuthorizationLetterIdList []*int64 `json:"AuthorizationLetterIdList,omitempty" xml:"AuthorizationLetterIdList,omitempty" type:"Repeated"`
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

func (s QuerySmsAuthorizationLetterRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAuthorizationLetterRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsAuthorizationLetterRequest) GetAuthorizationLetterIdList() []*int64 {
	return s.AuthorizationLetterIdList
}

func (s *QuerySmsAuthorizationLetterRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *QuerySmsAuthorizationLetterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsAuthorizationLetterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsAuthorizationLetterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsAuthorizationLetterRequest) GetSignName() *string {
	return s.SignName
}

func (s *QuerySmsAuthorizationLetterRequest) GetState() *string {
	return s.State
}

func (s *QuerySmsAuthorizationLetterRequest) GetStatus() *string {
	return s.Status
}

func (s *QuerySmsAuthorizationLetterRequest) SetAuthorizationLetterIdList(v []*int64) *QuerySmsAuthorizationLetterRequest {
	s.AuthorizationLetterIdList = v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetOrganizationCode(v string) *QuerySmsAuthorizationLetterRequest {
	s.OrganizationCode = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetOwnerId(v int64) *QuerySmsAuthorizationLetterRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetSignName(v string) *QuerySmsAuthorizationLetterRequest {
	s.SignName = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetState(v string) *QuerySmsAuthorizationLetterRequest {
	s.State = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetStatus(v string) *QuerySmsAuthorizationLetterRequest {
	s.Status = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) Validate() error {
	return dara.Validate(s)
}
