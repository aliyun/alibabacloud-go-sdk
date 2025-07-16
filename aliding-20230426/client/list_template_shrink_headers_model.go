// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTemplateShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListTemplateShrinkHeaders
	GetAccountContextShrink() *string
}

type ListTemplateShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListTemplateShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListTemplateShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTemplateShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListTemplateShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListTemplateShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTemplateShrinkHeaders) SetAccountContextShrink(v string) *ListTemplateShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListTemplateShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
