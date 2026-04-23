// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostOverviewMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryCostOverviewMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryCostOverviewMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryCostOverviewMetricsResponseBody) *ModelRouterQueryCostOverviewMetricsResponse
	GetBody() *ModelRouterQueryCostOverviewMetricsResponseBody
}

type ModelRouterQueryCostOverviewMetricsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryCostOverviewMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryCostOverviewMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostOverviewMetricsResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostOverviewMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryCostOverviewMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryCostOverviewMetricsResponse) GetBody() *ModelRouterQueryCostOverviewMetricsResponseBody {
	return s.Body
}

func (s *ModelRouterQueryCostOverviewMetricsResponse) SetHeaders(v map[string]*string) *ModelRouterQueryCostOverviewMetricsResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponse) SetStatusCode(v int32) *ModelRouterQueryCostOverviewMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponse) SetBody(v *ModelRouterQueryCostOverviewMetricsResponseBody) *ModelRouterQueryCostOverviewMetricsResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
