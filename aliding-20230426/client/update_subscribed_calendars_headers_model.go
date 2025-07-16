// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscribedCalendarsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateSubscribedCalendarsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateSubscribedCalendarsHeadersAccountContext) *UpdateSubscribedCalendarsHeaders
	GetAccountContext() *UpdateSubscribedCalendarsHeadersAccountContext
}

type UpdateSubscribedCalendarsHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateSubscribedCalendarsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateSubscribedCalendarsHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateSubscribedCalendarsHeaders) GetAccountContext() *UpdateSubscribedCalendarsHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateSubscribedCalendarsHeaders) SetCommonHeaders(v map[string]*string) *UpdateSubscribedCalendarsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSubscribedCalendarsHeaders) SetAccountContext(v *UpdateSubscribedCalendarsHeadersAccountContext) *UpdateSubscribedCalendarsHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateSubscribedCalendarsHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateSubscribedCalendarsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateSubscribedCalendarsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateSubscribedCalendarsHeadersAccountContext) SetAccountId(v string) *UpdateSubscribedCalendarsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateSubscribedCalendarsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
