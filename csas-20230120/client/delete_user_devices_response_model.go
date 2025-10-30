// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserDevicesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserDevicesResponseBody) *DeleteUserDevicesResponse
	GetBody() *DeleteUserDevicesResponseBody
}

type DeleteUserDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserDevicesResponse) GetBody() *DeleteUserDevicesResponseBody {
	return s.Body
}

func (s *DeleteUserDevicesResponse) SetHeaders(v map[string]*string) *DeleteUserDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDevicesResponse) SetStatusCode(v int32) *DeleteUserDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDevicesResponse) SetBody(v *DeleteUserDevicesResponseBody) *DeleteUserDevicesResponse {
	s.Body = v
	return s
}

func (s *DeleteUserDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
