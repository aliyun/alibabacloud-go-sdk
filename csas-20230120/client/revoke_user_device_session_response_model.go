// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserDeviceSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeUserDeviceSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeUserDeviceSessionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeUserDeviceSessionResponseBody) *RevokeUserDeviceSessionResponse
	GetBody() *RevokeUserDeviceSessionResponseBody
}

type RevokeUserDeviceSessionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeUserDeviceSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeUserDeviceSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserDeviceSessionResponse) GoString() string {
	return s.String()
}

func (s *RevokeUserDeviceSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeUserDeviceSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeUserDeviceSessionResponse) GetBody() *RevokeUserDeviceSessionResponseBody {
	return s.Body
}

func (s *RevokeUserDeviceSessionResponse) SetHeaders(v map[string]*string) *RevokeUserDeviceSessionResponse {
	s.Headers = v
	return s
}

func (s *RevokeUserDeviceSessionResponse) SetStatusCode(v int32) *RevokeUserDeviceSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeUserDeviceSessionResponse) SetBody(v *RevokeUserDeviceSessionResponseBody) *RevokeUserDeviceSessionResponse {
	s.Body = v
	return s
}

func (s *RevokeUserDeviceSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
