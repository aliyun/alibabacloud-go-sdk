// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNavigationByFormTypeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListNavigationByFormTypeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListNavigationByFormTypeHeadersAccountContext) *ListNavigationByFormTypeHeaders
	GetAccountContext() *ListNavigationByFormTypeHeadersAccountContext
}

type ListNavigationByFormTypeHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListNavigationByFormTypeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListNavigationByFormTypeHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeHeaders) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListNavigationByFormTypeHeaders) GetAccountContext() *ListNavigationByFormTypeHeadersAccountContext {
	return s.AccountContext
}

func (s *ListNavigationByFormTypeHeaders) SetCommonHeaders(v map[string]*string) *ListNavigationByFormTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListNavigationByFormTypeHeaders) SetAccountContext(v *ListNavigationByFormTypeHeadersAccountContext) *ListNavigationByFormTypeHeaders {
	s.AccountContext = v
	return s
}

func (s *ListNavigationByFormTypeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNavigationByFormTypeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListNavigationByFormTypeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListNavigationByFormTypeHeadersAccountContext) SetAccountId(v string) *ListNavigationByFormTypeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListNavigationByFormTypeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
