// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamPushMetricDetailDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamPushMetricDetailDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamPushMetricDetailDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamPushMetricDetailDataResponseBody) *DescribeLiveStreamPushMetricDetailDataResponse
	GetBody() *DescribeLiveStreamPushMetricDetailDataResponseBody
}

type DescribeLiveStreamPushMetricDetailDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamPushMetricDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamPushMetricDetailDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamPushMetricDetailDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamPushMetricDetailDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamPushMetricDetailDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamPushMetricDetailDataResponse) GetBody() *DescribeLiveStreamPushMetricDetailDataResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamPushMetricDetailDataResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamPushMetricDetailDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponse) SetStatusCode(v int32) *DescribeLiveStreamPushMetricDetailDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponse) SetBody(v *DescribeLiveStreamPushMetricDetailDataResponseBody) *DescribeLiveStreamPushMetricDetailDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamPushMetricDetailDataResponse) Validate() error {
	return dara.Validate(s)
}
