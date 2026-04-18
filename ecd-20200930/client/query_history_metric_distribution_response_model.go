// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryMetricDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHistoryMetricDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHistoryMetricDistributionResponse
	GetStatusCode() *int32
	SetBody(v *QueryHistoryMetricDistributionResponseBody) *QueryHistoryMetricDistributionResponse
	GetBody() *QueryHistoryMetricDistributionResponseBody
}

type QueryHistoryMetricDistributionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHistoryMetricDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHistoryMetricDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryMetricDistributionResponse) GoString() string {
	return s.String()
}

func (s *QueryHistoryMetricDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHistoryMetricDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHistoryMetricDistributionResponse) GetBody() *QueryHistoryMetricDistributionResponseBody {
	return s.Body
}

func (s *QueryHistoryMetricDistributionResponse) SetHeaders(v map[string]*string) *QueryHistoryMetricDistributionResponse {
	s.Headers = v
	return s
}

func (s *QueryHistoryMetricDistributionResponse) SetStatusCode(v int32) *QueryHistoryMetricDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHistoryMetricDistributionResponse) SetBody(v *QueryHistoryMetricDistributionResponseBody) *QueryHistoryMetricDistributionResponse {
	s.Body = v
	return s
}

func (s *QueryHistoryMetricDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
