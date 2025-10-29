// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamTranscodeMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamTranscodeMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamTranscodeMetricDataResponseBody) *DescribeLiveStreamTranscodeMetricDataResponse
	GetBody() *DescribeLiveStreamTranscodeMetricDataResponseBody
}

type DescribeLiveStreamTranscodeMetricDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamTranscodeMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamTranscodeMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamTranscodeMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamTranscodeMetricDataResponse) GetBody() *DescribeLiveStreamTranscodeMetricDataResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamTranscodeMetricDataResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamTranscodeMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponse) SetStatusCode(v int32) *DescribeLiveStreamTranscodeMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponse) SetBody(v *DescribeLiveStreamTranscodeMetricDataResponseBody) *DescribeLiveStreamTranscodeMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamTranscodeMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
