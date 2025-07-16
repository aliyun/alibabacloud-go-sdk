// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryMeetingRoomGroupHeadersAccountContext) *QueryMeetingRoomGroupHeaders
	GetAccountContext() *QueryMeetingRoomGroupHeadersAccountContext
}

type QueryMeetingRoomGroupHeaders struct {
	CommonHeaders  map[string]*string                          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryMeetingRoomGroupHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomGroupHeaders) GetAccountContext() *QueryMeetingRoomGroupHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomGroupHeaders) SetAccountContext(v *QueryMeetingRoomGroupHeadersAccountContext) *QueryMeetingRoomGroupHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryMeetingRoomGroupHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomGroupHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryMeetingRoomGroupHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryMeetingRoomGroupHeadersAccountContext) SetAccountId(v string) *QueryMeetingRoomGroupHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryMeetingRoomGroupHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
