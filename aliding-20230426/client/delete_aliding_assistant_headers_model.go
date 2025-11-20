// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlidingAssistantHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteAlidingAssistantHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteAlidingAssistantHeadersAccountContext) *DeleteAlidingAssistantHeaders
	GetAccountContext() *DeleteAlidingAssistantHeadersAccountContext
}

type DeleteAlidingAssistantHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteAlidingAssistantHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteAlidingAssistantHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteAlidingAssistantHeaders) GetAccountContext() *DeleteAlidingAssistantHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteAlidingAssistantHeaders) SetCommonHeaders(v map[string]*string) *DeleteAlidingAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAlidingAssistantHeaders) SetAccountContext(v *DeleteAlidingAssistantHeadersAccountContext) *DeleteAlidingAssistantHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteAlidingAssistantHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAlidingAssistantHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteAlidingAssistantHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteAlidingAssistantHeadersAccountContext) SetAccountId(v string) *DeleteAlidingAssistantHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteAlidingAssistantHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
