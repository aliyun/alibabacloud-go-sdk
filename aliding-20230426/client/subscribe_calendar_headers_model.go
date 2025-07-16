// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeCalendarHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SubscribeCalendarHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SubscribeCalendarHeadersAccountContext) *SubscribeCalendarHeaders
	GetAccountContext() *SubscribeCalendarHeadersAccountContext
}

type SubscribeCalendarHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SubscribeCalendarHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SubscribeCalendarHeaders) String() string {
	return dara.Prettify(s)
}

func (s SubscribeCalendarHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SubscribeCalendarHeaders) GetAccountContext() *SubscribeCalendarHeadersAccountContext {
	return s.AccountContext
}

func (s *SubscribeCalendarHeaders) SetCommonHeaders(v map[string]*string) *SubscribeCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeCalendarHeaders) SetAccountContext(v *SubscribeCalendarHeadersAccountContext) *SubscribeCalendarHeaders {
	s.AccountContext = v
	return s
}

func (s *SubscribeCalendarHeaders) Validate() error {
	return dara.Validate(s)
}

type SubscribeCalendarHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SubscribeCalendarHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SubscribeCalendarHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SubscribeCalendarHeadersAccountContext) SetAccountId(v string) *SubscribeCalendarHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SubscribeCalendarHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
