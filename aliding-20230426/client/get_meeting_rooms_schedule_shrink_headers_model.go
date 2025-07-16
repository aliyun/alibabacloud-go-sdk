// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeetingRoomsScheduleShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMeetingRoomsScheduleShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetMeetingRoomsScheduleShrinkHeaders
	GetAccountContextShrink() *string
}

type GetMeetingRoomsScheduleShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetMeetingRoomsScheduleShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMeetingRoomsScheduleShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetMeetingRoomsScheduleShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetMeetingRoomsScheduleShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMeetingRoomsScheduleShrinkHeaders) SetAccountContextShrink(v string) *GetMeetingRoomsScheduleShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetMeetingRoomsScheduleShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
