// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomGroupShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMeetingRoomGroupShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateMeetingRoomGroupShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateMeetingRoomGroupShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateMeetingRoomGroupShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMeetingRoomGroupShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateMeetingRoomGroupShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateMeetingRoomGroupShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMeetingRoomGroupShrinkHeaders) SetAccountContextShrink(v string) *CreateMeetingRoomGroupShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateMeetingRoomGroupShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
