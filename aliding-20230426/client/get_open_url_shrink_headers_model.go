// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenUrlShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOpenUrlShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetOpenUrlShrinkHeaders
	GetAccountContextShrink() *string
}

type GetOpenUrlShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetOpenUrlShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOpenUrlShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetOpenUrlShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOpenUrlShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetOpenUrlShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetOpenUrlShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOpenUrlShrinkHeaders) SetAccountContextShrink(v string) *GetOpenUrlShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetOpenUrlShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
