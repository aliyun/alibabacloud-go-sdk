// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDingtalkMeetingMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDingtalkMeetingMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *GetDingtalkMeetingMetricDataResponseBody) *GetDingtalkMeetingMetricDataResponse
	GetBody() *GetDingtalkMeetingMetricDataResponseBody
}

type GetDingtalkMeetingMetricDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingtalkMeetingMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingtalkMeetingMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataResponse) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDingtalkMeetingMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDingtalkMeetingMetricDataResponse) GetBody() *GetDingtalkMeetingMetricDataResponseBody {
	return s.Body
}

func (s *GetDingtalkMeetingMetricDataResponse) SetHeaders(v map[string]*string) *GetDingtalkMeetingMetricDataResponse {
	s.Headers = v
	return s
}

func (s *GetDingtalkMeetingMetricDataResponse) SetStatusCode(v int32) *GetDingtalkMeetingMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataResponse) SetBody(v *GetDingtalkMeetingMetricDataResponseBody) *GetDingtalkMeetingMetricDataResponse {
	s.Body = v
	return s
}

func (s *GetDingtalkMeetingMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
