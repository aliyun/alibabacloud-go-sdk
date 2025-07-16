// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormComponentDefinitionListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFormComponentDefinitionListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetFormComponentDefinitionListHeadersAccountContext) *GetFormComponentDefinitionListHeaders
	GetAccountContext() *GetFormComponentDefinitionListHeadersAccountContext
}

type GetFormComponentDefinitionListHeaders struct {
	CommonHeaders  map[string]*string                                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetFormComponentDefinitionListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetFormComponentDefinitionListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFormComponentDefinitionListHeaders) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFormComponentDefinitionListHeaders) GetAccountContext() *GetFormComponentDefinitionListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetFormComponentDefinitionListHeaders) SetCommonHeaders(v map[string]*string) *GetFormComponentDefinitionListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormComponentDefinitionListHeaders) SetAccountContext(v *GetFormComponentDefinitionListHeadersAccountContext) *GetFormComponentDefinitionListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetFormComponentDefinitionListHeaders) Validate() error {
	return dara.Validate(s)
}

type GetFormComponentDefinitionListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetFormComponentDefinitionListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetFormComponentDefinitionListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetFormComponentDefinitionListHeadersAccountContext) SetAccountId(v string) *GetFormComponentDefinitionListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetFormComponentDefinitionListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
