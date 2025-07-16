// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateMeetingRoomShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateMeetingRoomShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateMeetingRoomShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMeetingRoomShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateMeetingRoomShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMeetingRoomShrinkHeaders) SetAccountContextShrink(v string) *UpdateMeetingRoomShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateMeetingRoomShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
