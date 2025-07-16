// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMeetingRoomShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateMeetingRoomShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateMeetingRoomShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateMeetingRoomShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMeetingRoomShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateMeetingRoomShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateMeetingRoomShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMeetingRoomShrinkHeaders) SetAccountContextShrink(v string) *CreateMeetingRoomShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateMeetingRoomShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
