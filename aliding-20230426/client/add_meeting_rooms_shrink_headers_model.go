// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMeetingRoomsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddMeetingRoomsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddMeetingRoomsShrinkHeaders
	GetAccountContextShrink() *string
}

type AddMeetingRoomsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddMeetingRoomsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddMeetingRoomsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddMeetingRoomsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddMeetingRoomsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddMeetingRoomsShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddMeetingRoomsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMeetingRoomsShrinkHeaders) SetAccountContextShrink(v string) *AddMeetingRoomsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddMeetingRoomsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
