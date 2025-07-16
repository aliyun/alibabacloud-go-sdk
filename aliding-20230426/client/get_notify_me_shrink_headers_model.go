// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyMeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNotifyMeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetNotifyMeShrinkHeaders
	GetAccountContextShrink() *string
}

type GetNotifyMeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetNotifyMeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyMeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetNotifyMeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNotifyMeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetNotifyMeShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetNotifyMeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNotifyMeShrinkHeaders) SetAccountContextShrink(v string) *GetNotifyMeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetNotifyMeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
