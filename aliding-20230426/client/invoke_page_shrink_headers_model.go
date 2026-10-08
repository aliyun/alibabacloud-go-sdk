// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePageShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokePageShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InvokePageShrinkHeaders
	GetAccountContextShrink() *string
}

type InvokePageShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"accountContext,omitempty" xml:"accountContext,omitempty"`
}

func (s InvokePageShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokePageShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InvokePageShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokePageShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InvokePageShrinkHeaders) SetCommonHeaders(v map[string]*string) *InvokePageShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokePageShrinkHeaders) SetAccountContextShrink(v string) *InvokePageShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InvokePageShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
