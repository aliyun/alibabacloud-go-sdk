// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceFromSNResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDeviceFromSNResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDeviceFromSNResponse
	GetStatusCode() *int32
	SetBody(v *AddDeviceFromSNResponseBody) *AddDeviceFromSNResponse
	GetBody() *AddDeviceFromSNResponseBody
}

type AddDeviceFromSNResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDeviceFromSNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDeviceFromSNResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceFromSNResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceFromSNResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDeviceFromSNResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDeviceFromSNResponse) GetBody() *AddDeviceFromSNResponseBody {
	return s.Body
}

func (s *AddDeviceFromSNResponse) SetHeaders(v map[string]*string) *AddDeviceFromSNResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceFromSNResponse) SetStatusCode(v int32) *AddDeviceFromSNResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceFromSNResponse) SetBody(v *AddDeviceFromSNResponseBody) *AddDeviceFromSNResponse {
	s.Body = v
	return s
}

func (s *AddDeviceFromSNResponse) Validate() error {
	return dara.Validate(s)
}
