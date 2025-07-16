// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBannerShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendBannerShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SendBannerShrinkHeaders
	GetAccountContextShrink() *string
}

type SendBannerShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SendBannerShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendBannerShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SendBannerShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendBannerShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SendBannerShrinkHeaders) SetCommonHeaders(v map[string]*string) *SendBannerShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendBannerShrinkHeaders) SetAccountContextShrink(v string) *SendBannerShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SendBannerShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
