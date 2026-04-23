// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostTrendMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryCostTrendMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryCostTrendMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryCostTrendMetricsResponseBody) *ModelRouterQueryCostTrendMetricsResponse
	GetBody() *ModelRouterQueryCostTrendMetricsResponseBody
}

type ModelRouterQueryCostTrendMetricsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryCostTrendMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryCostTrendMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostTrendMetricsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostTrendMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryCostTrendMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryCostTrendMetricsResponse) GetBody() *ModelRouterQueryCostTrendMetricsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryCostTrendMetricsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryCostTrendMetricsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponse) SetStatusCode(v int32) *ModelRouterQueryCostTrendMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponse) SetBody(v *ModelRouterQueryCostTrendMetricsResponseBody) *ModelRouterQueryCostTrendMetricsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
