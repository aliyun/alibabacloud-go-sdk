// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSkillShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AuthorizeSkillShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AuthorizeSkillShrinkHeaders
	GetAccountContextShrink() *string
}

type AuthorizeSkillShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AuthorizeSkillShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSkillShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AuthorizeSkillShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AuthorizeSkillShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AuthorizeSkillShrinkHeaders) SetCommonHeaders(v map[string]*string) *AuthorizeSkillShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthorizeSkillShrinkHeaders) SetAccountContextShrink(v string) *AuthorizeSkillShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AuthorizeSkillShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
