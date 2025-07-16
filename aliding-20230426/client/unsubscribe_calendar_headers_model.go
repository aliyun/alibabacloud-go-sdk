// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeCalendarHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UnsubscribeCalendarHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UnsubscribeCalendarHeadersAccountContext) *UnsubscribeCalendarHeaders
	GetAccountContext() *UnsubscribeCalendarHeadersAccountContext
}

type UnsubscribeCalendarHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UnsubscribeCalendarHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UnsubscribeCalendarHeaders) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeCalendarHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeCalendarHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UnsubscribeCalendarHeaders) GetAccountContext() *UnsubscribeCalendarHeadersAccountContext {
	return s.AccountContext
}

func (s *UnsubscribeCalendarHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeCalendarHeaders) SetAccountContext(v *UnsubscribeCalendarHeadersAccountContext) *UnsubscribeCalendarHeaders {
	s.AccountContext = v
	return s
}

func (s *UnsubscribeCalendarHeaders) Validate() error {
	return dara.Validate(s)
}

type UnsubscribeCalendarHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UnsubscribeCalendarHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeCalendarHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UnsubscribeCalendarHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UnsubscribeCalendarHeadersAccountContext) SetAccountId(v string) *UnsubscribeCalendarHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UnsubscribeCalendarHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
