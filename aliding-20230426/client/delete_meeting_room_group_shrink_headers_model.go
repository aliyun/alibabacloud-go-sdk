// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomGroupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomGroupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteMeetingRoomGroupShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteMeetingRoomGroupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteMeetingRoomGroupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMeetingRoomGroupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteMeetingRoomGroupShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomGroupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMeetingRoomGroupShrinkHeaders) SetAccountContextShrink(v string) *DeleteMeetingRoomGroupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteMeetingRoomGroupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
