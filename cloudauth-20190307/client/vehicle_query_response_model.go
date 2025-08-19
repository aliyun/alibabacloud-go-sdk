// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VehicleQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VehicleQueryResponse
	GetStatusCode() *int32
	SetBody(v *VehicleQueryResponseBody) *VehicleQueryResponse
	GetBody() *VehicleQueryResponseBody
}

type VehicleQueryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s VehicleQueryResponse) GoString() string {
	return s.String()
}

func (s *VehicleQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VehicleQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VehicleQueryResponse) GetBody() *VehicleQueryResponseBody {
	return s.Body
}

func (s *VehicleQueryResponse) SetHeaders(v map[string]*string) *VehicleQueryResponse {
	s.Headers = v
	return s
}

func (s *VehicleQueryResponse) SetStatusCode(v int32) *VehicleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *VehicleQueryResponse) SetBody(v *VehicleQueryResponseBody) *VehicleQueryResponse {
	s.Body = v
	return s
}

func (s *VehicleQueryResponse) Validate() error {
	return dara.Validate(s)
}
