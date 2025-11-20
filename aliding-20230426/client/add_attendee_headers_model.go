// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAttendeeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddAttendeeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddAttendeeHeadersAccountContext) *AddAttendeeHeaders
	GetAccountContext() *AddAttendeeHeadersAccountContext
}

type AddAttendeeHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddAttendeeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddAttendeeHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeHeaders) GoString() string {
	return s.String()
}

func (s *AddAttendeeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddAttendeeHeaders) GetAccountContext() *AddAttendeeHeadersAccountContext {
	return s.AccountContext
}

func (s *AddAttendeeHeaders) SetCommonHeaders(v map[string]*string) *AddAttendeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAttendeeHeaders) SetAccountContext(v *AddAttendeeHeadersAccountContext) *AddAttendeeHeaders {
	s.AccountContext = v
	return s
}

func (s *AddAttendeeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAttendeeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddAttendeeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddAttendeeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddAttendeeHeadersAccountContext) SetAccountId(v string) *AddAttendeeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddAttendeeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
