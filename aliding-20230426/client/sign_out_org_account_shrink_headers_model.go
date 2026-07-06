// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutOrgAccountShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SignOutOrgAccountShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SignOutOrgAccountShrinkHeaders
	GetAccountContextShrink() *string
}

type SignOutOrgAccountShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SignOutOrgAccountShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SignOutOrgAccountShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SignOutOrgAccountShrinkHeaders) SetCommonHeaders(v map[string]*string) *SignOutOrgAccountShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SignOutOrgAccountShrinkHeaders) SetAccountContextShrink(v string) *SignOutOrgAccountShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SignOutOrgAccountShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
