// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlidingAssistantInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAlidingAssistantInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetAlidingAssistantInfoHeadersAccountContext) *GetAlidingAssistantInfoHeaders
	GetAccountContext() *GetAlidingAssistantInfoHeadersAccountContext
}

type GetAlidingAssistantInfoHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetAlidingAssistantInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetAlidingAssistantInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAlidingAssistantInfoHeaders) GetAccountContext() *GetAlidingAssistantInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *GetAlidingAssistantInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAlidingAssistantInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlidingAssistantInfoHeaders) SetAccountContext(v *GetAlidingAssistantInfoHeadersAccountContext) *GetAlidingAssistantInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *GetAlidingAssistantInfoHeaders) Validate() error {
	return dara.Validate(s)
}

type GetAlidingAssistantInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetAlidingAssistantInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAlidingAssistantInfoHeadersAccountContext) SetAccountId(v string) *GetAlidingAssistantInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetAlidingAssistantInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
