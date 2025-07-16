// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryMeetingRoomShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryMeetingRoomShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryMeetingRoomShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryMeetingRoomShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomShrinkHeaders) SetAccountContextShrink(v string) *QueryMeetingRoomShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryMeetingRoomShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
