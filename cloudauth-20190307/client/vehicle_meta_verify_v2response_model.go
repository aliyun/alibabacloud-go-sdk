// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicleMetaVerifyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VehicleMetaVerifyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VehicleMetaVerifyV2Response
	GetStatusCode() *int32
	SetBody(v *VehicleMetaVerifyV2ResponseBody) *VehicleMetaVerifyV2Response
	GetBody() *VehicleMetaVerifyV2ResponseBody
}

type VehicleMetaVerifyV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VehicleMetaVerifyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VehicleMetaVerifyV2Response) String() string {
	return dara.Prettify(s)
}

func (s VehicleMetaVerifyV2Response) GoString() string {
	return s.String()
}

func (s *VehicleMetaVerifyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VehicleMetaVerifyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VehicleMetaVerifyV2Response) GetBody() *VehicleMetaVerifyV2ResponseBody {
	return s.Body
}

func (s *VehicleMetaVerifyV2Response) SetHeaders(v map[string]*string) *VehicleMetaVerifyV2Response {
	s.Headers = v
	return s
}

func (s *VehicleMetaVerifyV2Response) SetStatusCode(v int32) *VehicleMetaVerifyV2Response {
	s.StatusCode = &v
	return s
}

func (s *VehicleMetaVerifyV2Response) SetBody(v *VehicleMetaVerifyV2ResponseBody) *VehicleMetaVerifyV2Response {
	s.Body = v
	return s
}

func (s *VehicleMetaVerifyV2Response) Validate() error {
	return dara.Validate(s)
}
