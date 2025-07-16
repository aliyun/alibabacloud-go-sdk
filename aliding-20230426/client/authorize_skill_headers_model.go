// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSkillHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AuthorizeSkillHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AuthorizeSkillHeadersAccountContext) *AuthorizeSkillHeaders
	GetAccountContext() *AuthorizeSkillHeadersAccountContext
}

type AuthorizeSkillHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AuthorizeSkillHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AuthorizeSkillHeaders) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSkillHeaders) GoString() string {
	return s.String()
}

func (s *AuthorizeSkillHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AuthorizeSkillHeaders) GetAccountContext() *AuthorizeSkillHeadersAccountContext {
	return s.AccountContext
}

func (s *AuthorizeSkillHeaders) SetCommonHeaders(v map[string]*string) *AuthorizeSkillHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthorizeSkillHeaders) SetAccountContext(v *AuthorizeSkillHeadersAccountContext) *AuthorizeSkillHeaders {
	s.AccountContext = v
	return s
}

func (s *AuthorizeSkillHeaders) Validate() error {
	return dara.Validate(s)
}

type AuthorizeSkillHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AuthorizeSkillHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSkillHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AuthorizeSkillHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AuthorizeSkillHeadersAccountContext) SetAccountId(v string) *AuthorizeSkillHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AuthorizeSkillHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
