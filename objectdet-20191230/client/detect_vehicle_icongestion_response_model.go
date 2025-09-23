// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleICongestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectVehicleICongestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectVehicleICongestionResponse
	GetStatusCode() *int32
	SetBody(v *DetectVehicleICongestionResponseBody) *DetectVehicleICongestionResponse
	GetBody() *DetectVehicleICongestionResponseBody
}

type DetectVehicleICongestionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectVehicleICongestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectVehicleICongestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponse) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectVehicleICongestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectVehicleICongestionResponse) GetBody() *DetectVehicleICongestionResponseBody {
	return s.Body
}

func (s *DetectVehicleICongestionResponse) SetHeaders(v map[string]*string) *DetectVehicleICongestionResponse {
	s.Headers = v
	return s
}

func (s *DetectVehicleICongestionResponse) SetStatusCode(v int32) *DetectVehicleICongestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVehicleICongestionResponse) SetBody(v *DetectVehicleICongestionResponseBody) *DetectVehicleICongestionResponse {
	s.Body = v
	return s
}

func (s *DetectVehicleICongestionResponse) Validate() error {
	return dara.Validate(s)
}
