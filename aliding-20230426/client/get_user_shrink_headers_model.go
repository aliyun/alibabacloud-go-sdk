// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetUserShrinkHeaders
	GetAccountContextShrink() *string
}

type GetUserShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetUserShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetUserShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetUserShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetUserShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserShrinkHeaders) SetAccountContextShrink(v string) *GetUserShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetUserShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
