// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserStreamMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveUserStreamMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveUserStreamMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveUserStreamMetricDataResponseBody) *DescribeLiveUserStreamMetricDataResponse
	GetBody() *DescribeLiveUserStreamMetricDataResponseBody
}

type DescribeLiveUserStreamMetricDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveUserStreamMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveUserStreamMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserStreamMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserStreamMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveUserStreamMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveUserStreamMetricDataResponse) GetBody() *DescribeLiveUserStreamMetricDataResponseBody {
	return s.Body
}

func (s *DescribeLiveUserStreamMetricDataResponse) SetHeaders(v map[string]*string) *DescribeLiveUserStreamMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponse) SetStatusCode(v int32) *DescribeLiveUserStreamMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponse) SetBody(v *DescribeLiveUserStreamMetricDataResponseBody) *DescribeLiveUserStreamMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
