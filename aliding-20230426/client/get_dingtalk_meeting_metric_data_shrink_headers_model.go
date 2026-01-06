// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMetricDataShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMetricDataShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetDingtalkMeetingMetricDataShrinkHeaders
	GetAccountContextShrink() *string
}

type GetDingtalkMeetingMetricDataShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetDingtalkMeetingMetricDataShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingMetricDataShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetDingtalkMeetingMetricDataShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMetricDataShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkHeaders) SetAccountContextShrink(v string) *GetDingtalkMeetingMetricDataShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
