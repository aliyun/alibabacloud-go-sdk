// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindParentPlatformDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUnbindParentPlatformDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUnbindParentPlatformDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BatchUnbindParentPlatformDevicesResponseBody) *BatchUnbindParentPlatformDevicesResponse
	GetBody() *BatchUnbindParentPlatformDevicesResponseBody
}

type BatchUnbindParentPlatformDevicesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUnbindParentPlatformDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUnbindParentPlatformDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUnbindParentPlatformDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUnbindParentPlatformDevicesResponse) GetBody() *BatchUnbindParentPlatformDevicesResponseBody {
	return s.Body
}

func (s *BatchUnbindParentPlatformDevicesResponse) SetHeaders(v map[string]*string) *BatchUnbindParentPlatformDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponse) SetStatusCode(v int32) *BatchUnbindParentPlatformDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponse) SetBody(v *BatchUnbindParentPlatformDevicesResponseBody) *BatchUnbindParentPlatformDevicesResponse {
	s.Body = v
	return s
}

func (s *BatchUnbindParentPlatformDevicesResponse) Validate() error {
	return dara.Validate(s)
}
