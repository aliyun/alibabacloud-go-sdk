// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlidingAssistantHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateAlidingAssistantHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateAlidingAssistantHeadersAccountContext) *UpdateAlidingAssistantHeaders
	GetAccountContext() *UpdateAlidingAssistantHeadersAccountContext
}

type UpdateAlidingAssistantHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateAlidingAssistantHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateAlidingAssistantHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateAlidingAssistantHeaders) GetAccountContext() *UpdateAlidingAssistantHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateAlidingAssistantHeaders) SetCommonHeaders(v map[string]*string) *UpdateAlidingAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAlidingAssistantHeaders) SetAccountContext(v *UpdateAlidingAssistantHeadersAccountContext) *UpdateAlidingAssistantHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateAlidingAssistantHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateAlidingAssistantHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateAlidingAssistantHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateAlidingAssistantHeadersAccountContext) SetAccountId(v string) *UpdateAlidingAssistantHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateAlidingAssistantHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
