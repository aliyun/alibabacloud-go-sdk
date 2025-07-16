// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDentriesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDentriesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListDentriesShrinkHeaders
	GetAccountContextShrink() *string
}

type ListDentriesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListDentriesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListDentriesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDentriesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListDentriesShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListDentriesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDentriesShrinkHeaders) SetAccountContextShrink(v string) *ListDentriesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListDentriesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
