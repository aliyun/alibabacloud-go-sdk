// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceResponseBody) *ListDeviceResponse
	GetBody() *ListDeviceResponseBody
}

type ListDeviceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceResponse) GetBody() *ListDeviceResponseBody {
	return s.Body
}

func (s *ListDeviceResponse) SetHeaders(v map[string]*string) *ListDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceResponse) SetStatusCode(v int32) *ListDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceResponse) SetBody(v *ListDeviceResponseBody) *ListDeviceResponse {
	s.Body = v
	return s
}

func (s *ListDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
