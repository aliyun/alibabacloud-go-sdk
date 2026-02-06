// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSkillDetailShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetSkillDetailShrinkHeaders
	GetAccountContextShrink() *string
}

type GetSkillDetailShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetSkillDetailShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetSkillDetailShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSkillDetailShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetSkillDetailShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetSkillDetailShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSkillDetailShrinkHeaders) SetAccountContextShrink(v string) *GetSkillDetailShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetSkillDetailShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
