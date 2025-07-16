// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateListByUserIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTemplateListByUserIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetTemplateListByUserIdHeadersAccountContext) *GetTemplateListByUserIdHeaders
	GetAccountContext() *GetTemplateListByUserIdHeadersAccountContext
}

type GetTemplateListByUserIdHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetTemplateListByUserIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetTemplateListByUserIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTemplateListByUserIdHeaders) GetAccountContext() *GetTemplateListByUserIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetTemplateListByUserIdHeaders) SetCommonHeaders(v map[string]*string) *GetTemplateListByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTemplateListByUserIdHeaders) SetAccountContext(v *GetTemplateListByUserIdHeadersAccountContext) *GetTemplateListByUserIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetTemplateListByUserIdHeaders) Validate() error {
	return dara.Validate(s)
}

type GetTemplateListByUserIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetTemplateListByUserIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTemplateListByUserIdHeadersAccountContext) SetAccountId(v string) *GetTemplateListByUserIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetTemplateListByUserIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
