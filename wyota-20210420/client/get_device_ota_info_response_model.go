// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceOtaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceOtaInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceOtaInfoResponseBody) *GetDeviceOtaInfoResponse
	GetBody() *GetDeviceOtaInfoResponseBody
}

type GetDeviceOtaInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceOtaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceOtaInfoResponse) GetBody() *GetDeviceOtaInfoResponseBody {
	return s.Body
}

func (s *GetDeviceOtaInfoResponse) SetHeaders(v map[string]*string) *GetDeviceOtaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaInfoResponse) SetStatusCode(v int32) *GetDeviceOtaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaInfoResponse) SetBody(v *GetDeviceOtaInfoResponseBody) *GetDeviceOtaInfoResponse {
	s.Body = v
	return s
}

func (s *GetDeviceOtaInfoResponse) Validate() error {
	return dara.Validate(s)
}
