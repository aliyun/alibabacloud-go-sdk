// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPredictiveMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPredictiveMetricResponse
	GetStatusCode() *int32
	SetBody(v *QueryPredictiveMetricResponseBody) *QueryPredictiveMetricResponse
	GetBody() *QueryPredictiveMetricResponseBody
}

type QueryPredictiveMetricResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPredictiveMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPredictiveMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryPredictiveMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPredictiveMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPredictiveMetricResponse) GetBody() *QueryPredictiveMetricResponseBody {
	return s.Body
}

func (s *QueryPredictiveMetricResponse) SetHeaders(v map[string]*string) *QueryPredictiveMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryPredictiveMetricResponse) SetStatusCode(v int32) *QueryPredictiveMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPredictiveMetricResponse) SetBody(v *QueryPredictiveMetricResponseBody) *QueryPredictiveMetricResponse {
	s.Body = v
	return s
}

func (s *QueryPredictiveMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
