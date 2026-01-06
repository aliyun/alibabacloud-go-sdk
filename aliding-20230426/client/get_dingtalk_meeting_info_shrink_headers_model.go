// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDingtalkMeetingInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDingtalkMeetingInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDingtalkMeetingInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDingtalkMeetingInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingInfoShrinkHeaders) SetAccountContextShrink(v string) *GetDingtalkMeetingInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
