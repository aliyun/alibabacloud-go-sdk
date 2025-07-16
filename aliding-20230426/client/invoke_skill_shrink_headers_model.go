// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSkillShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokeSkillShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InvokeSkillShrinkHeaders
	GetAccountContextShrink() *string
}

type InvokeSkillShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InvokeSkillShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokeSkillShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InvokeSkillShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokeSkillShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InvokeSkillShrinkHeaders) SetCommonHeaders(v map[string]*string) *InvokeSkillShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeSkillShrinkHeaders) SetAccountContextShrink(v string) *InvokeSkillShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InvokeSkillShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
