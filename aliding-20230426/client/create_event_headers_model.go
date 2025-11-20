// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateEventHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateEventHeadersAccountContext) *CreateEventHeaders
	GetAccountContext() *CreateEventHeadersAccountContext
}

type CreateEventHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateEventHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateEventHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateEventHeaders) GoString() string {
	return s.String()
}

func (s *CreateEventHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateEventHeaders) GetAccountContext() *CreateEventHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateEventHeaders) SetCommonHeaders(v map[string]*string) *CreateEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEventHeaders) SetAccountContext(v *CreateEventHeadersAccountContext) *CreateEventHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateEventHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 208579
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateEventHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateEventHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateEventHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateEventHeadersAccountContext) SetAccountId(v string) *CreateEventHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateEventHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
