// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDevicesResponse
	GetStatusCode() *int32
	SetBody(v *ListDevicesResponseBody) *ListDevicesResponse
	GetBody() *ListDevicesResponseBody
}

type ListDevicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDevicesResponse) GetBody() *ListDevicesResponseBody {
	return s.Body
}

func (s *ListDevicesResponse) SetHeaders(v map[string]*string) *ListDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListDevicesResponse) SetStatusCode(v int32) *ListDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
	s.Body = v
	return s
}

func (s *ListDevicesResponse) Validate() error {
	return dara.Validate(s)
}
