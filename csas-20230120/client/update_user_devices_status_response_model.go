// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDevicesStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserDevicesStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserDevicesStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserDevicesStatusResponseBody) *UpdateUserDevicesStatusResponse
	GetBody() *UpdateUserDevicesStatusResponseBody
}

type UpdateUserDevicesStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDevicesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDevicesStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserDevicesStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserDevicesStatusResponse) GetBody() *UpdateUserDevicesStatusResponseBody {
	return s.Body
}

func (s *UpdateUserDevicesStatusResponse) SetHeaders(v map[string]*string) *UpdateUserDevicesStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDevicesStatusResponse) SetStatusCode(v int32) *UpdateUserDevicesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDevicesStatusResponse) SetBody(v *UpdateUserDevicesStatusResponseBody) *UpdateUserDevicesStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateUserDevicesStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
