// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAuthKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAuthKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAuthKeyResponseBody) *CreateAuthKeyResponse
	GetBody() *CreateAuthKeyResponseBody
}

type CreateAuthKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuthKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAuthKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAuthKeyResponse) GetBody() *CreateAuthKeyResponseBody {
	return s.Body
}

func (s *CreateAuthKeyResponse) SetHeaders(v map[string]*string) *CreateAuthKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthKeyResponse) SetStatusCode(v int32) *CreateAuthKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthKeyResponse) SetBody(v *CreateAuthKeyResponseBody) *CreateAuthKeyResponse {
	s.Body = v
	return s
}

func (s *CreateAuthKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
