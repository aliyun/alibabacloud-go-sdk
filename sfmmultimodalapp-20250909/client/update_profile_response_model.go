// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProfileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProfileResponseBody) *UpdateProfileResponse
	GetBody() *UpdateProfileResponseBody
}

type UpdateProfileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProfileResponse) GetBody() *UpdateProfileResponseBody {
	return s.Body
}

func (s *UpdateProfileResponse) SetHeaders(v map[string]*string) *UpdateProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateProfileResponse) SetStatusCode(v int32) *UpdateProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProfileResponse) SetBody(v *UpdateProfileResponseBody) *UpdateProfileResponse {
	s.Body = v
	return s
}

func (s *UpdateProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
