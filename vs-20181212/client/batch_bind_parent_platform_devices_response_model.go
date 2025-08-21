// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindParentPlatformDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchBindParentPlatformDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchBindParentPlatformDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BatchBindParentPlatformDevicesResponseBody) *BatchBindParentPlatformDevicesResponse
	GetBody() *BatchBindParentPlatformDevicesResponseBody
}

type BatchBindParentPlatformDevicesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchBindParentPlatformDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchBindParentPlatformDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchBindParentPlatformDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindParentPlatformDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchBindParentPlatformDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchBindParentPlatformDevicesResponse) GetBody() *BatchBindParentPlatformDevicesResponseBody {
	return s.Body
}

func (s *BatchBindParentPlatformDevicesResponse) SetHeaders(v map[string]*string) *BatchBindParentPlatformDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindParentPlatformDevicesResponse) SetStatusCode(v int32) *BatchBindParentPlatformDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchBindParentPlatformDevicesResponse) SetBody(v *BatchBindParentPlatformDevicesResponseBody) *BatchBindParentPlatformDevicesResponse {
	s.Body = v
	return s
}

func (s *BatchBindParentPlatformDevicesResponse) Validate() error {
	return dara.Validate(s)
}
