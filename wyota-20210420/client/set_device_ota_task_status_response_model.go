// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceOtaTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDeviceOtaTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDeviceOtaTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDeviceOtaTaskStatusResponseBody) *SetDeviceOtaTaskStatusResponse
	GetBody() *SetDeviceOtaTaskStatusResponseBody
}

type SetDeviceOtaTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDeviceOtaTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDeviceOtaTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceOtaTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDeviceOtaTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDeviceOtaTaskStatusResponse) GetBody() *SetDeviceOtaTaskStatusResponseBody {
	return s.Body
}

func (s *SetDeviceOtaTaskStatusResponse) SetHeaders(v map[string]*string) *SetDeviceOtaTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDeviceOtaTaskStatusResponse) SetStatusCode(v int32) *SetDeviceOtaTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeviceOtaTaskStatusResponse) SetBody(v *SetDeviceOtaTaskStatusResponseBody) *SetDeviceOtaTaskStatusResponse {
	s.Body = v
	return s
}

func (s *SetDeviceOtaTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
