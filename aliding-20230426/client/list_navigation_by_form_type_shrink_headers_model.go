// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNavigationByFormTypeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListNavigationByFormTypeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListNavigationByFormTypeShrinkHeaders
	GetAccountContextShrink() *string
}

type ListNavigationByFormTypeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListNavigationByFormTypeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListNavigationByFormTypeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListNavigationByFormTypeShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListNavigationByFormTypeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListNavigationByFormTypeShrinkHeaders) SetAccountContextShrink(v string) *ListNavigationByFormTypeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListNavigationByFormTypeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
