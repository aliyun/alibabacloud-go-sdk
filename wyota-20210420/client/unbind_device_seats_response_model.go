// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindDeviceSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindDeviceSeatsResponse
	GetStatusCode() *int32
	SetBody(v *UnbindDeviceSeatsResponseBody) *UnbindDeviceSeatsResponse
	GetBody() *UnbindDeviceSeatsResponseBody
}

type UnbindDeviceSeatsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDeviceSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *UnbindDeviceSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindDeviceSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindDeviceSeatsResponse) GetBody() *UnbindDeviceSeatsResponseBody {
	return s.Body
}

func (s *UnbindDeviceSeatsResponse) SetHeaders(v map[string]*string) *UnbindDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *UnbindDeviceSeatsResponse) SetStatusCode(v int32) *UnbindDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDeviceSeatsResponse) SetBody(v *UnbindDeviceSeatsResponseBody) *UnbindDeviceSeatsResponse {
	s.Body = v
	return s
}

func (s *UnbindDeviceSeatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
