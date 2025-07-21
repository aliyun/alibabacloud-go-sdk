// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaTaskVersionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceOtaTaskVersionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceOtaTaskVersionInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceOtaTaskVersionInfoResponseBody) *GetDeviceOtaTaskVersionInfoResponse
	GetBody() *GetDeviceOtaTaskVersionInfoResponseBody
}

type GetDeviceOtaTaskVersionInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceOtaTaskVersionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceOtaTaskVersionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceOtaTaskVersionInfoResponse) GetBody() *GetDeviceOtaTaskVersionInfoResponseBody {
	return s.Body
}

func (s *GetDeviceOtaTaskVersionInfoResponse) SetHeaders(v map[string]*string) *GetDeviceOtaTaskVersionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponse) SetStatusCode(v int32) *GetDeviceOtaTaskVersionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponse) SetBody(v *GetDeviceOtaTaskVersionInfoResponseBody) *GetDeviceOtaTaskVersionInfoResponse {
	s.Body = v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoResponse) Validate() error {
	return dara.Validate(s)
}
