// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatabaseAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatabaseAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatabaseAccountResponseBody) *CreateDatabaseAccountResponse
	GetBody() *CreateDatabaseAccountResponseBody
}

type CreateDatabaseAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatabaseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatabaseAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatabaseAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatabaseAccountResponse) GetBody() *CreateDatabaseAccountResponseBody {
	return s.Body
}

func (s *CreateDatabaseAccountResponse) SetHeaders(v map[string]*string) *CreateDatabaseAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseAccountResponse) SetStatusCode(v int32) *CreateDatabaseAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatabaseAccountResponse) SetBody(v *CreateDatabaseAccountResponseBody) *CreateDatabaseAccountResponse {
	s.Body = v
	return s
}

func (s *CreateDatabaseAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
