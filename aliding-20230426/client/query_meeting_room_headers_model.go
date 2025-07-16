// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryMeetingRoomHeadersAccountContext) *QueryMeetingRoomHeaders
	GetAccountContext() *QueryMeetingRoomHeadersAccountContext
}

type QueryMeetingRoomHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryMeetingRoomHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomHeaders) GetAccountContext() *QueryMeetingRoomHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomHeaders) SetAccountContext(v *QueryMeetingRoomHeadersAccountContext) *QueryMeetingRoomHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryMeetingRoomHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryMeetingRoomHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryMeetingRoomHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryMeetingRoomHeadersAccountContext) SetAccountId(v string) *QueryMeetingRoomHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryMeetingRoomHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
