// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindParentPlatformDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindParentPlatformDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindParentPlatformDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnbindParentPlatformDeviceResponseBody) *UnbindParentPlatformDeviceResponse
	GetBody() *UnbindParentPlatformDeviceResponseBody
}

type UnbindParentPlatformDeviceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindParentPlatformDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindParentPlatformDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindParentPlatformDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindParentPlatformDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindParentPlatformDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindParentPlatformDeviceResponse) GetBody() *UnbindParentPlatformDeviceResponseBody {
	return s.Body
}

func (s *UnbindParentPlatformDeviceResponse) SetHeaders(v map[string]*string) *UnbindParentPlatformDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindParentPlatformDeviceResponse) SetStatusCode(v int32) *UnbindParentPlatformDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindParentPlatformDeviceResponse) SetBody(v *UnbindParentPlatformDeviceResponseBody) *UnbindParentPlatformDeviceResponse {
	s.Body = v
	return s
}

func (s *UnbindParentPlatformDeviceResponse) Validate() error {
	return dara.Validate(s)
}
