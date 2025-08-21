// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceByUserIdResponseBody) *ListDeviceByUserIdResponse
	GetBody() *ListDeviceByUserIdResponseBody
}

type ListDeviceByUserIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceByUserIdResponse) GetBody() *ListDeviceByUserIdResponseBody {
	return s.Body
}

func (s *ListDeviceByUserIdResponse) SetHeaders(v map[string]*string) *ListDeviceByUserIdResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceByUserIdResponse) SetStatusCode(v int32) *ListDeviceByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceByUserIdResponse) SetBody(v *ListDeviceByUserIdResponseBody) *ListDeviceByUserIdResponse {
	s.Body = v
	return s
}

func (s *ListDeviceByUserIdResponse) Validate() error {
	return dara.Validate(s)
}
