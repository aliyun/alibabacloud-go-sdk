// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessDefinitionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetProcessDefinitionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetProcessDefinitionHeadersAccountContext) *GetProcessDefinitionHeaders
	GetAccountContext() *GetProcessDefinitionHeadersAccountContext
}

type GetProcessDefinitionHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetProcessDefinitionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetProcessDefinitionHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetProcessDefinitionHeaders) GetAccountContext() *GetProcessDefinitionHeadersAccountContext {
	return s.AccountContext
}

func (s *GetProcessDefinitionHeaders) SetCommonHeaders(v map[string]*string) *GetProcessDefinitionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessDefinitionHeaders) SetAccountContext(v *GetProcessDefinitionHeadersAccountContext) *GetProcessDefinitionHeaders {
	s.AccountContext = v
	return s
}

func (s *GetProcessDefinitionHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProcessDefinitionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetProcessDefinitionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetProcessDefinitionHeadersAccountContext) SetAccountId(v string) *GetProcessDefinitionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetProcessDefinitionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
