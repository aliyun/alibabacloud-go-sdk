// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceSeatsResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceSeatsResponseBody) *ListDeviceSeatsResponse
	GetBody() *ListDeviceSeatsResponseBody
}

type ListDeviceSeatsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceSeatsResponse) GetBody() *ListDeviceSeatsResponseBody {
	return s.Body
}

func (s *ListDeviceSeatsResponse) SetHeaders(v map[string]*string) *ListDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceSeatsResponse) SetStatusCode(v int32) *ListDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceSeatsResponse) SetBody(v *ListDeviceSeatsResponseBody) *ListDeviceSeatsResponse {
	s.Body = v
	return s
}

func (s *ListDeviceSeatsResponse) Validate() error {
	return dara.Validate(s)
}
