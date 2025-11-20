// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetEventHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetEventHeadersAccountContext) *GetEventHeaders
	GetAccountContext() *GetEventHeadersAccountContext
}

type GetEventHeaders struct {
	CommonHeaders  map[string]*string             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetEventHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetEventHeaders) GoString() string {
	return s.String()
}

func (s *GetEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetEventHeaders) GetAccountContext() *GetEventHeadersAccountContext {
	return s.AccountContext
}

func (s *GetEventHeaders) SetCommonHeaders(v map[string]*string) *GetEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEventHeaders) SetAccountContext(v *GetEventHeadersAccountContext) *GetEventHeaders {
	s.AccountContext = v
	return s
}

func (s *GetEventHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetEventHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetEventHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetEventHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetEventHeadersAccountContext) SetAccountId(v string) *GetEventHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetEventHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
