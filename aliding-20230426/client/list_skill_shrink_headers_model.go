// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSkillShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListSkillShrinkHeaders
	GetAccountContextShrink() *string
}

type ListSkillShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListSkillShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSkillShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListSkillShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSkillShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListSkillShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListSkillShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSkillShrinkHeaders) SetAccountContextShrink(v string) *ListSkillShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListSkillShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
