// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutOrgAccountHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SignOutOrgAccountHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SignOutOrgAccountHeadersAccountContext) *SignOutOrgAccountHeaders
	GetAccountContext() *SignOutOrgAccountHeadersAccountContext
}

type SignOutOrgAccountHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SignOutOrgAccountHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SignOutOrgAccountHeaders) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountHeaders) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SignOutOrgAccountHeaders) GetAccountContext() *SignOutOrgAccountHeadersAccountContext {
	return s.AccountContext
}

func (s *SignOutOrgAccountHeaders) SetCommonHeaders(v map[string]*string) *SignOutOrgAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SignOutOrgAccountHeaders) SetAccountContext(v *SignOutOrgAccountHeadersAccountContext) *SignOutOrgAccountHeaders {
	s.AccountContext = v
	return s
}

func (s *SignOutOrgAccountHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SignOutOrgAccountHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SignOutOrgAccountHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SignOutOrgAccountHeadersAccountContext) SetAccountId(v string) *SignOutOrgAccountHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SignOutOrgAccountHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
