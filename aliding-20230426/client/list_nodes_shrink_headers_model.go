// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListNodesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListNodesShrinkHeaders
	GetAccountContextShrink() *string
}

type ListNodesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListNodesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListNodesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListNodesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListNodesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListNodesShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListNodesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListNodesShrinkHeaders) SetAccountContextShrink(v string) *ListNodesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListNodesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
