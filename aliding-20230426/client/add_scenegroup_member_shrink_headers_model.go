// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScenegroupMemberShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddScenegroupMemberShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddScenegroupMemberShrinkHeaders
	GetAccountContextShrink() *string
}

type AddScenegroupMemberShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddScenegroupMemberShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddScenegroupMemberShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddScenegroupMemberShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddScenegroupMemberShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddScenegroupMemberShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddScenegroupMemberShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddScenegroupMemberShrinkHeaders) SetAccountContextShrink(v string) *AddScenegroupMemberShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddScenegroupMemberShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
