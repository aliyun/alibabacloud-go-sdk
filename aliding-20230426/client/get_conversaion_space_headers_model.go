// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversaionSpaceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetConversaionSpaceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetConversaionSpaceHeadersAccountContext) *GetConversaionSpaceHeaders
	GetAccountContext() *GetConversaionSpaceHeadersAccountContext
}

type GetConversaionSpaceHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetConversaionSpaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetConversaionSpaceHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceHeaders) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetConversaionSpaceHeaders) GetAccountContext() *GetConversaionSpaceHeadersAccountContext {
	return s.AccountContext
}

func (s *GetConversaionSpaceHeaders) SetCommonHeaders(v map[string]*string) *GetConversaionSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversaionSpaceHeaders) SetAccountContext(v *GetConversaionSpaceHeadersAccountContext) *GetConversaionSpaceHeaders {
	s.AccountContext = v
	return s
}

func (s *GetConversaionSpaceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConversaionSpaceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetConversaionSpaceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetConversaionSpaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetConversaionSpaceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetConversaionSpaceHeadersAccountContext) SetAccountId(v string) *GetConversaionSpaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetConversaionSpaceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
