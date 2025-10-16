// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDevicesResponse
	GetStatusCode() *int32
	SetBody(v *AddDevicesResponseBody) *AddDevicesResponse
	GetBody() *AddDevicesResponseBody
}

type AddDevicesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDevicesResponse) GoString() string {
	return s.String()
}

func (s *AddDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDevicesResponse) GetBody() *AddDevicesResponseBody {
	return s.Body
}

func (s *AddDevicesResponse) SetHeaders(v map[string]*string) *AddDevicesResponse {
	s.Headers = v
	return s
}

func (s *AddDevicesResponse) SetStatusCode(v int32) *AddDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDevicesResponse) SetBody(v *AddDevicesResponseBody) *AddDevicesResponse {
	s.Body = v
	return s
}

func (s *AddDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
