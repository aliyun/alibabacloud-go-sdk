// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeContainerShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokeContainerShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InvokeContainerShrinkHeaders
	GetAccountContextShrink() *string
}

type InvokeContainerShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"accountContext,omitempty" xml:"accountContext,omitempty"`
}

func (s InvokeContainerShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokeContainerShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InvokeContainerShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokeContainerShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InvokeContainerShrinkHeaders) SetCommonHeaders(v map[string]*string) *InvokeContainerShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeContainerShrinkHeaders) SetAccountContextShrink(v string) *InvokeContainerShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InvokeContainerShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
