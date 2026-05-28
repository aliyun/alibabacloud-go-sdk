// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateModelApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateModelApplicationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateModelApplicationResponseBody) *UpdateModelApplicationResponse
	GetBody() *UpdateModelApplicationResponseBody
}

type UpdateModelApplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateModelApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateModelApplicationResponse) GetBody() *UpdateModelApplicationResponseBody {
	return s.Body
}

func (s *UpdateModelApplicationResponse) SetHeaders(v map[string]*string) *UpdateModelApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelApplicationResponse) SetStatusCode(v int32) *UpdateModelApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelApplicationResponse) SetBody(v *UpdateModelApplicationResponseBody) *UpdateModelApplicationResponse {
	s.Body = v
	return s
}

func (s *UpdateModelApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
