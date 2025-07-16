// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCalendarsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListCalendarsHeadersAccountContext) *ListCalendarsHeaders
	GetAccountContext() *ListCalendarsHeadersAccountContext
}

type ListCalendarsHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListCalendarsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListCalendarsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsHeaders) GoString() string {
	return s.String()
}

func (s *ListCalendarsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCalendarsHeaders) GetAccountContext() *ListCalendarsHeadersAccountContext {
	return s.AccountContext
}

func (s *ListCalendarsHeaders) SetCommonHeaders(v map[string]*string) *ListCalendarsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCalendarsHeaders) SetAccountContext(v *ListCalendarsHeadersAccountContext) *ListCalendarsHeaders {
	s.AccountContext = v
	return s
}

func (s *ListCalendarsHeaders) Validate() error {
	return dara.Validate(s)
}

type ListCalendarsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListCalendarsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListCalendarsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListCalendarsHeadersAccountContext) SetAccountId(v string) *ListCalendarsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListCalendarsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
