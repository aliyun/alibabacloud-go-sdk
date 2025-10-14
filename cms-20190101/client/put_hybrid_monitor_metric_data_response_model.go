// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutHybridMonitorMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutHybridMonitorMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutHybridMonitorMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *PutHybridMonitorMetricDataResponseBody) *PutHybridMonitorMetricDataResponse
	GetBody() *PutHybridMonitorMetricDataResponseBody
}

type PutHybridMonitorMetricDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutHybridMonitorMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutHybridMonitorMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s PutHybridMonitorMetricDataResponse) GoString() string {
	return s.String()
}

func (s *PutHybridMonitorMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutHybridMonitorMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutHybridMonitorMetricDataResponse) GetBody() *PutHybridMonitorMetricDataResponseBody {
	return s.Body
}

func (s *PutHybridMonitorMetricDataResponse) SetHeaders(v map[string]*string) *PutHybridMonitorMetricDataResponse {
	s.Headers = v
	return s
}

func (s *PutHybridMonitorMetricDataResponse) SetStatusCode(v int32) *PutHybridMonitorMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PutHybridMonitorMetricDataResponse) SetBody(v *PutHybridMonitorMetricDataResponseBody) *PutHybridMonitorMetricDataResponse {
	s.Body = v
	return s
}

func (s *PutHybridMonitorMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
