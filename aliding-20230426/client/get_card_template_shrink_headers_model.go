// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardTemplateShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCardTemplateShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetCardTemplateShrinkHeaders
	GetAccountContextShrink() *string
}

type GetCardTemplateShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetCardTemplateShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetCardTemplateShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCardTemplateShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetCardTemplateShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetCardTemplateShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCardTemplateShrinkHeaders) SetAccountContextShrink(v string) *GetCardTemplateShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetCardTemplateShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
