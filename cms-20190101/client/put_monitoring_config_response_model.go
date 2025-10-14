// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMonitoringConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutMonitoringConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutMonitoringConfigResponse
	GetStatusCode() *int32
	SetBody(v *PutMonitoringConfigResponseBody) *PutMonitoringConfigResponse
	GetBody() *PutMonitoringConfigResponseBody
}

type PutMonitoringConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMonitoringConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMonitoringConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutMonitoringConfigResponse) GoString() string {
	return s.String()
}

func (s *PutMonitoringConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutMonitoringConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutMonitoringConfigResponse) GetBody() *PutMonitoringConfigResponseBody {
	return s.Body
}

func (s *PutMonitoringConfigResponse) SetHeaders(v map[string]*string) *PutMonitoringConfigResponse {
	s.Headers = v
	return s
}

func (s *PutMonitoringConfigResponse) SetStatusCode(v int32) *PutMonitoringConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMonitoringConfigResponse) SetBody(v *PutMonitoringConfigResponseBody) *PutMonitoringConfigResponse {
	s.Body = v
	return s
}

func (s *PutMonitoringConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
