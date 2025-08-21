// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDeviceSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDeviceSettingResponse
	GetStatusCode() *int32
	SetBody(v *SetDeviceSettingResponseBody) *SetDeviceSettingResponse
	GetBody() *SetDeviceSettingResponseBody
}

type SetDeviceSettingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeviceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeviceSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceSettingResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDeviceSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDeviceSettingResponse) GetBody() *SetDeviceSettingResponseBody {
	return s.Body
}

func (s *SetDeviceSettingResponse) SetHeaders(v map[string]*string) *SetDeviceSettingResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceSettingResponse) SetStatusCode(v int32) *SetDeviceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceSettingResponse) SetBody(v *SetDeviceSettingResponseBody) *SetDeviceSettingResponse {
	s.Body = v
	return s
}

func (s *SetDeviceSettingResponse) Validate() error {
	return dara.Validate(s)
}
