// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *MeetingFlashMinutesShrinkHeaders
	GetAccountContextShrink() *string
}

type MeetingFlashMinutesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s MeetingFlashMinutesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MeetingFlashMinutesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *MeetingFlashMinutesShrinkHeaders) SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MeetingFlashMinutesShrinkHeaders) SetAccountContextShrink(v string) *MeetingFlashMinutesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *MeetingFlashMinutesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
