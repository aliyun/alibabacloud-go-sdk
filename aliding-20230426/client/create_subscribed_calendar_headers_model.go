// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscribedCalendarHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSubscribedCalendarHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateSubscribedCalendarHeadersAccountContext) *CreateSubscribedCalendarHeaders
	GetAccountContext() *CreateSubscribedCalendarHeadersAccountContext
}

type CreateSubscribedCalendarHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateSubscribedCalendarHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateSubscribedCalendarHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarHeaders) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSubscribedCalendarHeaders) GetAccountContext() *CreateSubscribedCalendarHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateSubscribedCalendarHeaders) SetCommonHeaders(v map[string]*string) *CreateSubscribedCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSubscribedCalendarHeaders) SetAccountContext(v *CreateSubscribedCalendarHeadersAccountContext) *CreateSubscribedCalendarHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateSubscribedCalendarHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSubscribedCalendarHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateSubscribedCalendarHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateSubscribedCalendarHeadersAccountContext) SetAccountId(v string) *CreateSubscribedCalendarHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateSubscribedCalendarHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
