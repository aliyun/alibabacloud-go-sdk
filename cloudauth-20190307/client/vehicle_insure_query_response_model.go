// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleInsureQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VehicleInsureQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VehicleInsureQueryResponse
	GetStatusCode() *int32
	SetBody(v *VehicleInsureQueryResponseBody) *VehicleInsureQueryResponse
	GetBody() *VehicleInsureQueryResponseBody
}

type VehicleInsureQueryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleInsureQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleInsureQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s VehicleInsureQueryResponse) GoString() string {
	return s.String()
}

func (s *VehicleInsureQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VehicleInsureQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VehicleInsureQueryResponse) GetBody() *VehicleInsureQueryResponseBody {
	return s.Body
}

func (s *VehicleInsureQueryResponse) SetHeaders(v map[string]*string) *VehicleInsureQueryResponse {
	s.Headers = v
	return s
}

func (s *VehicleInsureQueryResponse) SetStatusCode(v int32) *VehicleInsureQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *VehicleInsureQueryResponse) SetBody(v *VehicleInsureQueryResponseBody) *VehicleInsureQueryResponse {
	s.Body = v
	return s
}

func (s *VehicleInsureQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
