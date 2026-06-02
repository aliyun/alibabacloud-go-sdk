// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesTextShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesTextShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *MeetingFlashMinutesTextShrinkHeaders
	GetAccountContextShrink() *string
}

type MeetingFlashMinutesTextShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s MeetingFlashMinutesTextShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextShrinkHeaders) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MeetingFlashMinutesTextShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *MeetingFlashMinutesTextShrinkHeaders) SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesTextShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MeetingFlashMinutesTextShrinkHeaders) SetAccountContextShrink(v string) *MeetingFlashMinutesTextShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *MeetingFlashMinutesTextShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
