// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnbindDeviceResponseBody) *UnbindDeviceResponse
	GetBody() *UnbindDeviceResponseBody
}

type UnbindDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindDeviceResponse) GetBody() *UnbindDeviceResponseBody {
	return s.Body
}

func (s *UnbindDeviceResponse) SetHeaders(v map[string]*string) *UnbindDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindDeviceResponse) SetStatusCode(v int32) *UnbindDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDeviceResponse) SetBody(v *UnbindDeviceResponseBody) *UnbindDeviceResponse {
	s.Body = v
	return s
}

func (s *UnbindDeviceResponse) Validate() error {
	return dara.Validate(s)
}
