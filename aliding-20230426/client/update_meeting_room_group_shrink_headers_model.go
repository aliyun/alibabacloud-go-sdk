// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomGroupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomGroupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateMeetingRoomGroupShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateMeetingRoomGroupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateMeetingRoomGroupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMeetingRoomGroupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateMeetingRoomGroupShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomGroupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMeetingRoomGroupShrinkHeaders) SetAccountContextShrink(v string) *UpdateMeetingRoomGroupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateMeetingRoomGroupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
