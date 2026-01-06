// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDingtalkMeetingListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDingtalkMeetingListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDingtalkMeetingListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDingtalkMeetingListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingListShrinkHeaders) SetAccountContextShrink(v string) *GetDingtalkMeetingListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
