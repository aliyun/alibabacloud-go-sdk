// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAttendeeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveAttendeeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *RemoveAttendeeHeadersAccountContext) *RemoveAttendeeHeaders
	GetAccountContext() *RemoveAttendeeHeadersAccountContext
}

type RemoveAttendeeHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *RemoveAttendeeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s RemoveAttendeeHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeHeaders) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveAttendeeHeaders) GetAccountContext() *RemoveAttendeeHeadersAccountContext {
	return s.AccountContext
}

func (s *RemoveAttendeeHeaders) SetCommonHeaders(v map[string]*string) *RemoveAttendeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveAttendeeHeaders) SetAccountContext(v *RemoveAttendeeHeadersAccountContext) *RemoveAttendeeHeaders {
	s.AccountContext = v
	return s
}

func (s *RemoveAttendeeHeaders) Validate() error {
	return dara.Validate(s)
}

type RemoveAttendeeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s RemoveAttendeeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *RemoveAttendeeHeadersAccountContext) SetAccountId(v string) *RemoveAttendeeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *RemoveAttendeeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
