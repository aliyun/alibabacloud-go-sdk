// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPopupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendPopupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SendPopupShrinkHeaders
	GetAccountContextShrink() *string
}

type SendPopupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SendPopupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendPopupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SendPopupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendPopupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SendPopupShrinkHeaders) SetCommonHeaders(v map[string]*string) *SendPopupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendPopupShrinkHeaders) SetAccountContextShrink(v string) *SendPopupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SendPopupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
