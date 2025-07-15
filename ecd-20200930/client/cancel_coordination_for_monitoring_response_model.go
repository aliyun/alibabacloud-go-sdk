// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCoordinationForMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCoordinationForMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCoordinationForMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *CancelCoordinationForMonitoringResponseBody) *CancelCoordinationForMonitoringResponse
	GetBody() *CancelCoordinationForMonitoringResponseBody
}

type CancelCoordinationForMonitoringResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCoordinationForMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCoordinationForMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCoordinationForMonitoringResponse) GoString() string {
	return s.String()
}

func (s *CancelCoordinationForMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCoordinationForMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCoordinationForMonitoringResponse) GetBody() *CancelCoordinationForMonitoringResponseBody {
	return s.Body
}

func (s *CancelCoordinationForMonitoringResponse) SetHeaders(v map[string]*string) *CancelCoordinationForMonitoringResponse {
	s.Headers = v
	return s
}

func (s *CancelCoordinationForMonitoringResponse) SetStatusCode(v int32) *CancelCoordinationForMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCoordinationForMonitoringResponse) SetBody(v *CancelCoordinationForMonitoringResponseBody) *CancelCoordinationForMonitoringResponse {
	s.Body = v
	return s
}

func (s *CancelCoordinationForMonitoringResponse) Validate() error {
	return dara.Validate(s)
}
