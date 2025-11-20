// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlidingAssistantHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateAlidingAssistantHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateAlidingAssistantHeadersAccountContext) *CreateAlidingAssistantHeaders
	GetAccountContext() *CreateAlidingAssistantHeadersAccountContext
}

type CreateAlidingAssistantHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateAlidingAssistantHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateAlidingAssistantHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantHeaders) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateAlidingAssistantHeaders) GetAccountContext() *CreateAlidingAssistantHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateAlidingAssistantHeaders) SetCommonHeaders(v map[string]*string) *CreateAlidingAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAlidingAssistantHeaders) SetAccountContext(v *CreateAlidingAssistantHeadersAccountContext) *CreateAlidingAssistantHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateAlidingAssistantHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAlidingAssistantHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateAlidingAssistantHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateAlidingAssistantHeadersAccountContext) SetAccountId(v string) *CreateAlidingAssistantHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateAlidingAssistantHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
