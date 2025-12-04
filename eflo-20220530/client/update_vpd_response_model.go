// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpdResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVpdResponseBody) *UpdateVpdResponse
	GetBody() *UpdateVpdResponseBody
}

type UpdateVpdResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpdResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpdResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpdResponse) GetBody() *UpdateVpdResponseBody {
	return s.Body
}

func (s *UpdateVpdResponse) SetHeaders(v map[string]*string) *UpdateVpdResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpdResponse) SetStatusCode(v int32) *UpdateVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpdResponse) SetBody(v *UpdateVpdResponseBody) *UpdateVpdResponse {
	s.Body = v
	return s
}

func (s *UpdateVpdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
