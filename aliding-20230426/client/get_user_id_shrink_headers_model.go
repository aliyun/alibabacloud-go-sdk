// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetUserIdShrinkHeaders
	GetAccountContextShrink() *string
}

type GetUserIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetUserIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetUserIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdShrinkHeaders) SetAccountContextShrink(v string) *GetUserIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetUserIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
