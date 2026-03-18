// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryObservationMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryObservationMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryObservationMetricsResponseBody) *ModelRouterQueryObservationMetricsResponse
	GetBody() *ModelRouterQueryObservationMetricsResponseBody
}

type ModelRouterQueryObservationMetricsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryObservationMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryObservationMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationMetricsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryObservationMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryObservationMetricsResponse) GetBody() *ModelRouterQueryObservationMetricsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryObservationMetricsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryObservationMetricsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponse) SetStatusCode(v int32) *ModelRouterQueryObservationMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponse) SetBody(v *ModelRouterQueryObservationMetricsResponseBody) *ModelRouterQueryObservationMetricsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
