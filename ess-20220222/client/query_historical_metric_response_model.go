// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoricalMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHistoricalMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHistoricalMetricResponse
	GetStatusCode() *int32
	SetBody(v *QueryHistoricalMetricResponseBody) *QueryHistoricalMetricResponse
	GetBody() *QueryHistoricalMetricResponseBody
}

type QueryHistoricalMetricResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHistoricalMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHistoricalMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoricalMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryHistoricalMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHistoricalMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHistoricalMetricResponse) GetBody() *QueryHistoricalMetricResponseBody {
	return s.Body
}

func (s *QueryHistoricalMetricResponse) SetHeaders(v map[string]*string) *QueryHistoricalMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryHistoricalMetricResponse) SetStatusCode(v int32) *QueryHistoricalMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHistoricalMetricResponse) SetBody(v *QueryHistoricalMetricResponseBody) *QueryHistoricalMetricResponse {
	s.Body = v
	return s
}

func (s *QueryHistoricalMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
