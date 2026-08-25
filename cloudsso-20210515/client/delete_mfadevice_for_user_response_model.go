// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMFADeviceForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMFADeviceForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMFADeviceForUserResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMFADeviceForUserResponseBody) *DeleteMFADeviceForUserResponse
	GetBody() *DeleteMFADeviceForUserResponseBody
}

type DeleteMFADeviceForUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMFADeviceForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMFADeviceForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMFADeviceForUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteMFADeviceForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMFADeviceForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMFADeviceForUserResponse) GetBody() *DeleteMFADeviceForUserResponseBody {
	return s.Body
}

func (s *DeleteMFADeviceForUserResponse) SetHeaders(v map[string]*string) *DeleteMFADeviceForUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteMFADeviceForUserResponse) SetStatusCode(v int32) *DeleteMFADeviceForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMFADeviceForUserResponse) SetBody(v *DeleteMFADeviceForUserResponseBody) *DeleteMFADeviceForUserResponse {
	s.Body = v
	return s
}

func (s *DeleteMFADeviceForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
