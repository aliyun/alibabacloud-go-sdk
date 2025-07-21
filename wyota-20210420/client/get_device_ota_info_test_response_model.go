// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaInfoTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceOtaInfoTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceOtaInfoTestResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceOtaInfoTestResponseBody) *GetDeviceOtaInfoTestResponse
	GetBody() *GetDeviceOtaInfoTestResponseBody
}

type GetDeviceOtaInfoTestResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaInfoTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaInfoTestResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoTestResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceOtaInfoTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceOtaInfoTestResponse) GetBody() *GetDeviceOtaInfoTestResponseBody {
	return s.Body
}

func (s *GetDeviceOtaInfoTestResponse) SetHeaders(v map[string]*string) *GetDeviceOtaInfoTestResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaInfoTestResponse) SetStatusCode(v int32) *GetDeviceOtaInfoTestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaInfoTestResponse) SetBody(v *GetDeviceOtaInfoTestResponseBody) *GetDeviceOtaInfoTestResponse {
	s.Body = v
	return s
}

func (s *GetDeviceOtaInfoTestResponse) Validate() error {
	return dara.Validate(s)
}
