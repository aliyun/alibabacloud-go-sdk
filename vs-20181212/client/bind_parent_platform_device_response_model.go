// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindParentPlatformDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindParentPlatformDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindParentPlatformDeviceResponse
	GetStatusCode() *int32
	SetBody(v *BindParentPlatformDeviceResponseBody) *BindParentPlatformDeviceResponse
	GetBody() *BindParentPlatformDeviceResponseBody
}

type BindParentPlatformDeviceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindParentPlatformDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindParentPlatformDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s BindParentPlatformDeviceResponse) GoString() string {
	return s.String()
}

func (s *BindParentPlatformDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindParentPlatformDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindParentPlatformDeviceResponse) GetBody() *BindParentPlatformDeviceResponseBody {
	return s.Body
}

func (s *BindParentPlatformDeviceResponse) SetHeaders(v map[string]*string) *BindParentPlatformDeviceResponse {
	s.Headers = v
	return s
}

func (s *BindParentPlatformDeviceResponse) SetStatusCode(v int32) *BindParentPlatformDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindParentPlatformDeviceResponse) SetBody(v *BindParentPlatformDeviceResponseBody) *BindParentPlatformDeviceResponse {
	s.Body = v
	return s
}

func (s *BindParentPlatformDeviceResponse) Validate() error {
	return dara.Validate(s)
}
