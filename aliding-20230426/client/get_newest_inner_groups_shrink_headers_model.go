// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNewestInnerGroupsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNewestInnerGroupsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetNewestInnerGroupsShrinkHeaders
	GetAccountContextShrink() *string
}

type GetNewestInnerGroupsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetNewestInnerGroupsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNewestInnerGroupsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetNewestInnerGroupsShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetNewestInnerGroupsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNewestInnerGroupsShrinkHeaders) SetAccountContextShrink(v string) *GetNewestInnerGroupsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetNewestInnerGroupsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
