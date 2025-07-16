// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMeetingRoomsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddMeetingRoomsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddMeetingRoomsHeadersAccountContext) *AddMeetingRoomsHeaders
	GetAccountContext() *AddMeetingRoomsHeadersAccountContext
}

type AddMeetingRoomsHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddMeetingRoomsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddMeetingRoomsHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsHeaders) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddMeetingRoomsHeaders) GetAccountContext() *AddMeetingRoomsHeadersAccountContext {
	return s.AccountContext
}

func (s *AddMeetingRoomsHeaders) SetCommonHeaders(v map[string]*string) *AddMeetingRoomsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMeetingRoomsHeaders) SetAccountContext(v *AddMeetingRoomsHeadersAccountContext) *AddMeetingRoomsHeaders {
	s.AccountContext = v
	return s
}

func (s *AddMeetingRoomsHeaders) Validate() error {
	return dara.Validate(s)
}

type AddMeetingRoomsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddMeetingRoomsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddMeetingRoomsHeadersAccountContext) SetAccountId(v string) *AddMeetingRoomsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddMeetingRoomsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
