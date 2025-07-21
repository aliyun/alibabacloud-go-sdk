// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceOtaAutoStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDeviceOtaAutoStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDeviceOtaAutoStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDeviceOtaAutoStatusResponseBody) *SetDeviceOtaAutoStatusResponse
	GetBody() *SetDeviceOtaAutoStatusResponseBody
}

type SetDeviceOtaAutoStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeviceOtaAutoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeviceOtaAutoStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceOtaAutoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaAutoStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDeviceOtaAutoStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDeviceOtaAutoStatusResponse) GetBody() *SetDeviceOtaAutoStatusResponseBody {
	return s.Body
}

func (s *SetDeviceOtaAutoStatusResponse) SetHeaders(v map[string]*string) *SetDeviceOtaAutoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceOtaAutoStatusResponse) SetStatusCode(v int32) *SetDeviceOtaAutoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceOtaAutoStatusResponse) SetBody(v *SetDeviceOtaAutoStatusResponseBody) *SetDeviceOtaAutoStatusResponse {
	s.Body = v
	return s
}

func (s *SetDeviceOtaAutoStatusResponse) Validate() error {
	return dara.Validate(s)
}
