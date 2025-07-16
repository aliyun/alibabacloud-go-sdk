// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryMeetingRoomGroupListShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryMeetingRoomGroupListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryMeetingRoomGroupListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomGroupListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryMeetingRoomGroupListShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomGroupListShrinkHeaders) SetAccountContextShrink(v string) *QueryMeetingRoomGroupListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryMeetingRoomGroupListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
