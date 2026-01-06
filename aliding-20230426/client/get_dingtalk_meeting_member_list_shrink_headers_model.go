// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMemberListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDingtalkMeetingMemberListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDingtalkMeetingMemberListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDingtalkMeetingMemberListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingMemberListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDingtalkMeetingMemberListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMemberListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkHeaders) SetAccountContextShrink(v string) *GetDingtalkMeetingMemberListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
