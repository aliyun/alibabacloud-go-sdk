// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListTemplateHeadersAccountContext) *ListTemplateHeaders
	GetAccountContext() *ListTemplateHeadersAccountContext
}

type ListTemplateHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListTemplateHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateHeaders) GoString() string {
	return s.String()
}

func (s *ListTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTemplateHeaders) GetAccountContext() *ListTemplateHeadersAccountContext {
	return s.AccountContext
}

func (s *ListTemplateHeaders) SetCommonHeaders(v map[string]*string) *ListTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTemplateHeaders) SetAccountContext(v *ListTemplateHeadersAccountContext) *ListTemplateHeaders {
	s.AccountContext = v
	return s
}

func (s *ListTemplateHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTemplateHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListTemplateHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListTemplateHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListTemplateHeadersAccountContext) SetAccountId(v string) *ListTemplateHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListTemplateHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
