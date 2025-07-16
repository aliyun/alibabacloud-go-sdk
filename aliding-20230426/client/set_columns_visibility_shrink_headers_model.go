// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetColumnsVisibilityShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetColumnsVisibilityShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SetColumnsVisibilityShrinkHeaders
	GetAccountContextShrink() *string
}

type SetColumnsVisibilityShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SetColumnsVisibilityShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetColumnsVisibilityShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SetColumnsVisibilityShrinkHeaders) SetCommonHeaders(v map[string]*string) *SetColumnsVisibilityShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetColumnsVisibilityShrinkHeaders) SetAccountContextShrink(v string) *SetColumnsVisibilityShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SetColumnsVisibilityShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
