// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnbindDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnbindDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnbindDevicesResponse
	GetStatusCode() *int32
	SetBody(v *ListUnbindDevicesResponseBody) *ListUnbindDevicesResponse
	GetBody() *ListUnbindDevicesResponseBody
}

type ListUnbindDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnbindDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnbindDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnbindDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListUnbindDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnbindDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnbindDevicesResponse) GetBody() *ListUnbindDevicesResponseBody {
	return s.Body
}

func (s *ListUnbindDevicesResponse) SetHeaders(v map[string]*string) *ListUnbindDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListUnbindDevicesResponse) SetStatusCode(v int32) *ListUnbindDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnbindDevicesResponse) SetBody(v *ListUnbindDevicesResponseBody) *ListUnbindDevicesResponse {
	s.Body = v
	return s
}

func (s *ListUnbindDevicesResponse) Validate() error {
	return dara.Validate(s)
}
