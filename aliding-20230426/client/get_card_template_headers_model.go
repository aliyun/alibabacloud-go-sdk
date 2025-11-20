// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCardTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetCardTemplateHeadersAccountContext) *GetCardTemplateHeaders
	GetAccountContext() *GetCardTemplateHeadersAccountContext
}

type GetCardTemplateHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetCardTemplateHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetCardTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateHeaders) GoString() string {
	return s.String()
}

func (s *GetCardTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCardTemplateHeaders) GetAccountContext() *GetCardTemplateHeadersAccountContext {
	return s.AccountContext
}

func (s *GetCardTemplateHeaders) SetCommonHeaders(v map[string]*string) *GetCardTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCardTemplateHeaders) SetAccountContext(v *GetCardTemplateHeadersAccountContext) *GetCardTemplateHeaders {
	s.AccountContext = v
	return s
}

func (s *GetCardTemplateHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCardTemplateHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetCardTemplateHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetCardTemplateHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetCardTemplateHeadersAccountContext) SetAccountId(v string) *GetCardTemplateHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetCardTemplateHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
