// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserPoolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserPoolResponseBody) *UpdateUserPoolResponse
	GetBody() *UpdateUserPoolResponseBody
}

type UpdateUserPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserPoolResponse) GetBody() *UpdateUserPoolResponseBody {
	return s.Body
}

func (s *UpdateUserPoolResponse) SetHeaders(v map[string]*string) *UpdateUserPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserPoolResponse) SetStatusCode(v int32) *UpdateUserPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserPoolResponse) SetBody(v *UpdateUserPoolResponseBody) *UpdateUserPoolResponse {
	s.Body = v
	return s
}

func (s *UpdateUserPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
