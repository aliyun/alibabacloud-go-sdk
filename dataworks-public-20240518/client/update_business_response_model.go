// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBusinessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBusinessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBusinessResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBusinessResponseBody) *UpdateBusinessResponse
	GetBody() *UpdateBusinessResponseBody
}

type UpdateBusinessResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBusinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBusinessResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBusinessResponse) GoString() string {
	return s.String()
}

func (s *UpdateBusinessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBusinessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBusinessResponse) GetBody() *UpdateBusinessResponseBody {
	return s.Body
}

func (s *UpdateBusinessResponse) SetHeaders(v map[string]*string) *UpdateBusinessResponse {
	s.Headers = v
	return s
}

func (s *UpdateBusinessResponse) SetStatusCode(v int32) *UpdateBusinessResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBusinessResponse) SetBody(v *UpdateBusinessResponseBody) *UpdateBusinessResponse {
	s.Body = v
	return s
}

func (s *UpdateBusinessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
