// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListEventsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListEventsHeadersAccountContext) *ListEventsHeaders
	GetAccountContext() *ListEventsHeadersAccountContext
}

type ListEventsHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListEventsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListEventsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListEventsHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListEventsHeaders) GetAccountContext() *ListEventsHeadersAccountContext {
	return s.AccountContext
}

func (s *ListEventsHeaders) SetCommonHeaders(v map[string]*string) *ListEventsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsHeaders) SetAccountContext(v *ListEventsHeadersAccountContext) *ListEventsHeaders {
	s.AccountContext = v
	return s
}

func (s *ListEventsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 208579
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListEventsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListEventsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListEventsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListEventsHeadersAccountContext) SetAccountId(v string) *ListEventsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListEventsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
