// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscribedCalendarHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSubscribedCalendarHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetSubscribedCalendarHeadersAccountContext) *GetSubscribedCalendarHeaders
	GetAccountContext() *GetSubscribedCalendarHeadersAccountContext
}

type GetSubscribedCalendarHeaders struct {
	CommonHeaders  map[string]*string                          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetSubscribedCalendarHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetSubscribedCalendarHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSubscribedCalendarHeaders) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSubscribedCalendarHeaders) GetAccountContext() *GetSubscribedCalendarHeadersAccountContext {
	return s.AccountContext
}

func (s *GetSubscribedCalendarHeaders) SetCommonHeaders(v map[string]*string) *GetSubscribedCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSubscribedCalendarHeaders) SetAccountContext(v *GetSubscribedCalendarHeadersAccountContext) *GetSubscribedCalendarHeaders {
	s.AccountContext = v
	return s
}

func (s *GetSubscribedCalendarHeaders) Validate() error {
	return dara.Validate(s)
}

type GetSubscribedCalendarHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetSubscribedCalendarHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetSubscribedCalendarHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSubscribedCalendarHeadersAccountContext) SetAccountId(v string) *GetSubscribedCalendarHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetSubscribedCalendarHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
