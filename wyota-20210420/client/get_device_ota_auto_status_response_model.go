// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaAutoStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceOtaAutoStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceOtaAutoStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceOtaAutoStatusResponseBody) *GetDeviceOtaAutoStatusResponse
	GetBody() *GetDeviceOtaAutoStatusResponseBody
}

type GetDeviceOtaAutoStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaAutoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaAutoStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaAutoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaAutoStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceOtaAutoStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceOtaAutoStatusResponse) GetBody() *GetDeviceOtaAutoStatusResponseBody {
	return s.Body
}

func (s *GetDeviceOtaAutoStatusResponse) SetHeaders(v map[string]*string) *GetDeviceOtaAutoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaAutoStatusResponse) SetStatusCode(v int32) *GetDeviceOtaAutoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaAutoStatusResponse) SetBody(v *GetDeviceOtaAutoStatusResponseBody) *GetDeviceOtaAutoStatusResponse {
	s.Body = v
	return s
}

func (s *GetDeviceOtaAutoStatusResponse) Validate() error {
	return dara.Validate(s)
}
