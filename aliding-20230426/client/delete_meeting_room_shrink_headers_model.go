// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteMeetingRoomShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteMeetingRoomShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteMeetingRoomShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMeetingRoomShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteMeetingRoomShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMeetingRoomShrinkHeaders) SetAccountContextShrink(v string) *DeleteMeetingRoomShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteMeetingRoomShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
