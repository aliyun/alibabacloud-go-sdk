// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveInteractionMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveInteractionMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveInteractionMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveInteractionMetricDataResponseBody) *DescribeLiveInteractionMetricDataResponse
	GetBody() *DescribeLiveInteractionMetricDataResponseBody
}

type DescribeLiveInteractionMetricDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveInteractionMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveInteractionMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveInteractionMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveInteractionMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveInteractionMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveInteractionMetricDataResponse) GetBody() *DescribeLiveInteractionMetricDataResponseBody {
	return s.Body
}

func (s *DescribeLiveInteractionMetricDataResponse) SetHeaders(v map[string]*string) *DescribeLiveInteractionMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponse) SetStatusCode(v int32) *DescribeLiveInteractionMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponse) SetBody(v *DescribeLiveInteractionMetricDataResponseBody) *DescribeLiveInteractionMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveInteractionMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
