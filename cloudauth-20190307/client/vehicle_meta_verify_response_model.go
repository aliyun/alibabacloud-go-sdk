// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleMetaVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VehicleMetaVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VehicleMetaVerifyResponse
	GetStatusCode() *int32
	SetBody(v *VehicleMetaVerifyResponseBody) *VehicleMetaVerifyResponse
	GetBody() *VehicleMetaVerifyResponseBody
}

type VehicleMetaVerifyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleMetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleMetaVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VehicleMetaVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VehicleMetaVerifyResponse) GetBody() *VehicleMetaVerifyResponseBody {
	return s.Body
}

func (s *VehicleMetaVerifyResponse) SetHeaders(v map[string]*string) *VehicleMetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *VehicleMetaVerifyResponse) SetStatusCode(v int32) *VehicleMetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *VehicleMetaVerifyResponse) SetBody(v *VehicleMetaVerifyResponseBody) *VehicleMetaVerifyResponse {
	s.Body = v
	return s
}

func (s *VehicleMetaVerifyResponse) Validate() error {
	return dara.Validate(s)
}
