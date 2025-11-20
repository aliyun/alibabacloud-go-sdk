// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscribedCalendarHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSubscribedCalendarHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteSubscribedCalendarHeadersAccountContext) *DeleteSubscribedCalendarHeaders
	GetAccountContext() *DeleteSubscribedCalendarHeadersAccountContext
}

type DeleteSubscribedCalendarHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteSubscribedCalendarHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteSubscribedCalendarHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscribedCalendarHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSubscribedCalendarHeaders) GetAccountContext() *DeleteSubscribedCalendarHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteSubscribedCalendarHeaders) SetCommonHeaders(v map[string]*string) *DeleteSubscribedCalendarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSubscribedCalendarHeaders) SetAccountContext(v *DeleteSubscribedCalendarHeadersAccountContext) *DeleteSubscribedCalendarHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteSubscribedCalendarHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSubscribedCalendarHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteSubscribedCalendarHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscribedCalendarHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteSubscribedCalendarHeadersAccountContext) SetAccountId(v string) *DeleteSubscribedCalendarHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteSubscribedCalendarHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
