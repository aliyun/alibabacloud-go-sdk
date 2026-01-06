// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMemberEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDingtalkMeetingMemberEventShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDingtalkMeetingMemberEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDingtalkMeetingMemberEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingMemberEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDingtalkMeetingMemberEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMemberEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkHeaders) SetAccountContextShrink(v string) *GetDingtalkMeetingMemberEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
