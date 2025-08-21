// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceSettingResponseBody) *GetDeviceSettingResponse
	GetBody() *GetDeviceSettingResponseBody
}

type GetDeviceSettingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceSettingResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceSettingResponse) GetBody() *GetDeviceSettingResponseBody {
	return s.Body
}

func (s *GetDeviceSettingResponse) SetHeaders(v map[string]*string) *GetDeviceSettingResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceSettingResponse) SetStatusCode(v int32) *GetDeviceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceSettingResponse) SetBody(v *GetDeviceSettingResponseBody) *GetDeviceSettingResponse {
	s.Body = v
	return s
}

func (s *GetDeviceSettingResponse) Validate() error {
	return dara.Validate(s)
}
