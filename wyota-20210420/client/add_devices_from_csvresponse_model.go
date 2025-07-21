// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDevicesFromCSVResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDevicesFromCSVResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDevicesFromCSVResponse
	GetStatusCode() *int32
	SetBody(v *AddDevicesFromCSVResponseBody) *AddDevicesFromCSVResponse
	GetBody() *AddDevicesFromCSVResponseBody
}

type AddDevicesFromCSVResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDevicesFromCSVResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDevicesFromCSVResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDevicesFromCSVResponse) GoString() string {
	return s.String()
}

func (s *AddDevicesFromCSVResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDevicesFromCSVResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDevicesFromCSVResponse) GetBody() *AddDevicesFromCSVResponseBody {
	return s.Body
}

func (s *AddDevicesFromCSVResponse) SetHeaders(v map[string]*string) *AddDevicesFromCSVResponse {
	s.Headers = v
	return s
}

func (s *AddDevicesFromCSVResponse) SetStatusCode(v int32) *AddDevicesFromCSVResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDevicesFromCSVResponse) SetBody(v *AddDevicesFromCSVResponseBody) *AddDevicesFromCSVResponse {
	s.Body = v
	return s
}

func (s *AddDevicesFromCSVResponse) Validate() error {
	return dara.Validate(s)
}
