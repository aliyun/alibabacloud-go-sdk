// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSearchShadeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendSearchShadeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SendSearchShadeShrinkHeaders
	GetAccountContextShrink() *string
}

type SendSearchShadeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SendSearchShadeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SendSearchShadeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendSearchShadeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SendSearchShadeShrinkHeaders) SetCommonHeaders(v map[string]*string) *SendSearchShadeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendSearchShadeShrinkHeaders) SetAccountContextShrink(v string) *SendSearchShadeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SendSearchShadeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
