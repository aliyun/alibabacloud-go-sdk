// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryMeetingRoomListShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryMeetingRoomListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryMeetingRoomListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryMeetingRoomListShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomListShrinkHeaders) SetAccountContextShrink(v string) *QueryMeetingRoomListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryMeetingRoomListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
