// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCsrResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCsrResponseBody) *UpdateCsrResponse
	GetBody() *UpdateCsrResponseBody
}

type UpdateCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCsrResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCsrResponse) GoString() string {
	return s.String()
}

func (s *UpdateCsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCsrResponse) GetBody() *UpdateCsrResponseBody {
	return s.Body
}

func (s *UpdateCsrResponse) SetHeaders(v map[string]*string) *UpdateCsrResponse {
	s.Headers = v
	return s
}

func (s *UpdateCsrResponse) SetStatusCode(v int32) *UpdateCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCsrResponse) SetBody(v *UpdateCsrResponseBody) *UpdateCsrResponse {
	s.Body = v
	return s
}

func (s *UpdateCsrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
