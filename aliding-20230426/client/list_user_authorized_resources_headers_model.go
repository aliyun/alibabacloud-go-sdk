// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthorizedResourcesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListUserAuthorizedResourcesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListUserAuthorizedResourcesHeadersAccountContext) *ListUserAuthorizedResourcesHeaders
	GetAccountContext() *ListUserAuthorizedResourcesHeadersAccountContext
}

type ListUserAuthorizedResourcesHeaders struct {
	CommonHeaders  map[string]*string                                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListUserAuthorizedResourcesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListUserAuthorizedResourcesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthorizedResourcesHeaders) GoString() string {
	return s.String()
}

func (s *ListUserAuthorizedResourcesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListUserAuthorizedResourcesHeaders) GetAccountContext() *ListUserAuthorizedResourcesHeadersAccountContext {
	return s.AccountContext
}

func (s *ListUserAuthorizedResourcesHeaders) SetCommonHeaders(v map[string]*string) *ListUserAuthorizedResourcesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserAuthorizedResourcesHeaders) SetAccountContext(v *ListUserAuthorizedResourcesHeadersAccountContext) *ListUserAuthorizedResourcesHeaders {
	s.AccountContext = v
	return s
}

func (s *ListUserAuthorizedResourcesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserAuthorizedResourcesHeadersAccountContext struct {
	AlidingSsoTicket *string `json:"AlidingSsoTicket,omitempty" xml:"AlidingSsoTicket,omitempty"`
	SsoTicket        *string `json:"SsoTicket,omitempty" xml:"SsoTicket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListUserAuthorizedResourcesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthorizedResourcesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListUserAuthorizedResourcesHeadersAccountContext) GetAlidingSsoTicket() *string {
	return s.AlidingSsoTicket
}

func (s *ListUserAuthorizedResourcesHeadersAccountContext) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *ListUserAuthorizedResourcesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListUserAuthorizedResourcesHeadersAccountContext) SetAlidingSsoTicket(v string) *ListUserAuthorizedResourcesHeadersAccountContext {
	s.AlidingSsoTicket = &v
	return s
}

func (s *ListUserAuthorizedResourcesHeadersAccountContext) SetSsoTicket(v string) *ListUserAuthorizedResourcesHeadersAccountContext {
	s.SsoTicket = &v
	return s
}

func (s *ListUserAuthorizedResourcesHeadersAccountContext) SetAccountId(v string) *ListUserAuthorizedResourcesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListUserAuthorizedResourcesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
