// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMetricDetailDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamMetricDetailDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamMetricDetailDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamMetricDetailDataResponseBody) *DescribeLiveStreamMetricDetailDataResponse
	GetBody() *DescribeLiveStreamMetricDetailDataResponseBody
}

type DescribeLiveStreamMetricDetailDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamMetricDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamMetricDetailDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMetricDetailDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMetricDetailDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamMetricDetailDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamMetricDetailDataResponse) GetBody() *DescribeLiveStreamMetricDetailDataResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamMetricDetailDataResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamMetricDetailDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponse) SetStatusCode(v int32) *DescribeLiveStreamMetricDetailDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponse) SetBody(v *DescribeLiveStreamMetricDetailDataResponseBody) *DescribeLiveStreamMetricDetailDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
