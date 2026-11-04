// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFaceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFaceUserResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFaceUserResponseBody) *UpdateFaceUserResponse
	GetBody() *UpdateFaceUserResponseBody
}

type UpdateFaceUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFaceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFaceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFaceUserResponse) GetBody() *UpdateFaceUserResponseBody {
	return s.Body
}

func (s *UpdateFaceUserResponse) SetHeaders(v map[string]*string) *UpdateFaceUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateFaceUserResponse) SetStatusCode(v int32) *UpdateFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFaceUserResponse) SetBody(v *UpdateFaceUserResponseBody) *UpdateFaceUserResponse {
	s.Body = v
	return s
}

func (s *UpdateFaceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
