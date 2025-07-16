// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteUsersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InviteUsersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InviteUsersShrinkHeaders
	GetAccountContextShrink() *string
}

type InviteUsersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InviteUsersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InviteUsersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InviteUsersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InviteUsersShrinkHeaders) SetCommonHeaders(v map[string]*string) *InviteUsersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InviteUsersShrinkHeaders) SetAccountContextShrink(v string) *InviteUsersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InviteUsersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
