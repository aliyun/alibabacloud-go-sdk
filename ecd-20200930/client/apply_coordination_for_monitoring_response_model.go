// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationForMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyCoordinationForMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyCoordinationForMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *ApplyCoordinationForMonitoringResponseBody) *ApplyCoordinationForMonitoringResponse
	GetBody() *ApplyCoordinationForMonitoringResponseBody
}

type ApplyCoordinationForMonitoringResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCoordinationForMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCoordinationForMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationForMonitoringResponse) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationForMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyCoordinationForMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyCoordinationForMonitoringResponse) GetBody() *ApplyCoordinationForMonitoringResponseBody {
	return s.Body
}

func (s *ApplyCoordinationForMonitoringResponse) SetHeaders(v map[string]*string) *ApplyCoordinationForMonitoringResponse {
	s.Headers = v
	return s
}

func (s *ApplyCoordinationForMonitoringResponse) SetStatusCode(v int32) *ApplyCoordinationForMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCoordinationForMonitoringResponse) SetBody(v *ApplyCoordinationForMonitoringResponseBody) *ApplyCoordinationForMonitoringResponse {
	s.Body = v
	return s
}

func (s *ApplyCoordinationForMonitoringResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
