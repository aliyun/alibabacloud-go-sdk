// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TerminateInstanceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *TerminateInstanceShrinkHeaders
	GetAccountContextShrink() *string
}

type TerminateInstanceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s TerminateInstanceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *TerminateInstanceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TerminateInstanceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *TerminateInstanceShrinkHeaders) SetCommonHeaders(v map[string]*string) *TerminateInstanceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TerminateInstanceShrinkHeaders) SetAccountContextShrink(v string) *TerminateInstanceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *TerminateInstanceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
