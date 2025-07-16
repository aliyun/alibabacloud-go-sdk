// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryMeetingRoomListHeadersAccountContext) *QueryMeetingRoomListHeaders
	GetAccountContext() *QueryMeetingRoomListHeadersAccountContext
}

type QueryMeetingRoomListHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryMeetingRoomListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomListHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomListHeaders) GetAccountContext() *QueryMeetingRoomListHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryMeetingRoomListHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomListHeaders) SetAccountContext(v *QueryMeetingRoomListHeadersAccountContext) *QueryMeetingRoomListHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryMeetingRoomListHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryMeetingRoomListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryMeetingRoomListHeadersAccountContext) SetAccountId(v string) *QueryMeetingRoomListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryMeetingRoomListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
