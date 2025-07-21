// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBoundDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBoundDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBoundDevicesResponse
	GetStatusCode() *int32
	SetBody(v *ListBoundDevicesResponseBody) *ListBoundDevicesResponse
	GetBody() *ListBoundDevicesResponseBody
}

type ListBoundDevicesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBoundDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBoundDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBoundDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListBoundDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBoundDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBoundDevicesResponse) GetBody() *ListBoundDevicesResponseBody {
	return s.Body
}

func (s *ListBoundDevicesResponse) SetHeaders(v map[string]*string) *ListBoundDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListBoundDevicesResponse) SetStatusCode(v int32) *ListBoundDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBoundDevicesResponse) SetBody(v *ListBoundDevicesResponseBody) *ListBoundDevicesResponse {
	s.Body = v
	return s
}

func (s *ListBoundDevicesResponse) Validate() error {
	return dara.Validate(s)
}
