// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceConfigsResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceConfigsResponseBody) *GetDeviceConfigsResponse
	GetBody() *GetDeviceConfigsResponseBody
}

type GetDeviceConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceConfigsResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceConfigsResponse) GetBody() *GetDeviceConfigsResponseBody {
	return s.Body
}

func (s *GetDeviceConfigsResponse) SetHeaders(v map[string]*string) *GetDeviceConfigsResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceConfigsResponse) SetStatusCode(v int32) *GetDeviceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceConfigsResponse) SetBody(v *GetDeviceConfigsResponseBody) *GetDeviceConfigsResponse {
	s.Body = v
	return s
}

func (s *GetDeviceConfigsResponse) Validate() error {
	return dara.Validate(s)
}
