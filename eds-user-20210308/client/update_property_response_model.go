// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePropertyResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePropertyResponseBody) *UpdatePropertyResponse
	GetBody() *UpdatePropertyResponseBody
}

type UpdatePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePropertyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePropertyResponse) GetBody() *UpdatePropertyResponseBody {
	return s.Body
}

func (s *UpdatePropertyResponse) SetHeaders(v map[string]*string) *UpdatePropertyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePropertyResponse) SetStatusCode(v int32) *UpdatePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePropertyResponse) SetBody(v *UpdatePropertyResponseBody) *UpdatePropertyResponse {
	s.Body = v
	return s
}

func (s *UpdatePropertyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
