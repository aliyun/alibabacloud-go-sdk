// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSkillsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetSkillsShrinkHeaders
	GetAccountContextShrink() *string
}

type GetSkillsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetSkillsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSkillsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetSkillsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSkillsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetSkillsShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetSkillsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSkillsShrinkHeaders) SetAccountContextShrink(v string) *GetSkillsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetSkillsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
