// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedirectTaskShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RedirectTaskShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *RedirectTaskShrinkHeaders
	GetAccountContextShrink() *string
}

type RedirectTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s RedirectTaskShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s RedirectTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *RedirectTaskShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RedirectTaskShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *RedirectTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *RedirectTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RedirectTaskShrinkHeaders) SetAccountContextShrink(v string) *RedirectTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *RedirectTaskShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
