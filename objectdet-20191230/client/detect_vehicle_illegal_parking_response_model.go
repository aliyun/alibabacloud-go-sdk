// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleIllegalParkingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectVehicleIllegalParkingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectVehicleIllegalParkingResponse
	GetStatusCode() *int32
	SetBody(v *DetectVehicleIllegalParkingResponseBody) *DetectVehicleIllegalParkingResponse
	GetBody() *DetectVehicleIllegalParkingResponseBody
}

type DetectVehicleIllegalParkingResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectVehicleIllegalParkingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectVehicleIllegalParkingResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponse) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectVehicleIllegalParkingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectVehicleIllegalParkingResponse) GetBody() *DetectVehicleIllegalParkingResponseBody {
	return s.Body
}

func (s *DetectVehicleIllegalParkingResponse) SetHeaders(v map[string]*string) *DetectVehicleIllegalParkingResponse {
	s.Headers = v
	return s
}

func (s *DetectVehicleIllegalParkingResponse) SetStatusCode(v int32) *DetectVehicleIllegalParkingResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponse) SetBody(v *DetectVehicleIllegalParkingResponseBody) *DetectVehicleIllegalParkingResponse {
	s.Body = v
	return s
}

func (s *DetectVehicleIllegalParkingResponse) Validate() error {
	return dara.Validate(s)
}
