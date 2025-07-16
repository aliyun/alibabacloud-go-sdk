// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryMeetingRoomGroupShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryMeetingRoomGroupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryMeetingRoomGroupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomGroupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryMeetingRoomGroupShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomGroupShrinkHeaders) SetAccountContextShrink(v string) *QueryMeetingRoomGroupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryMeetingRoomGroupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
