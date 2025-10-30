// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserDevicesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserDevicesResponseBody) *ListUserDevicesResponse
	GetBody() *ListUserDevicesResponseBody
}

type ListUserDevicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListUserDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserDevicesResponse) GetBody() *ListUserDevicesResponseBody {
	return s.Body
}

func (s *ListUserDevicesResponse) SetHeaders(v map[string]*string) *ListUserDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListUserDevicesResponse) SetStatusCode(v int32) *ListUserDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDevicesResponse) SetBody(v *ListUserDevicesResponseBody) *ListUserDevicesResponse {
	s.Body = v
	return s
}

func (s *ListUserDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
