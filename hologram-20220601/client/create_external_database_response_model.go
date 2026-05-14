// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExternalDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExternalDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateExternalDatabaseResponseBody) *CreateExternalDatabaseResponse
	GetBody() *CreateExternalDatabaseResponseBody
}

type CreateExternalDatabaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExternalDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExternalDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateExternalDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExternalDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExternalDatabaseResponse) GetBody() *CreateExternalDatabaseResponseBody {
	return s.Body
}

func (s *CreateExternalDatabaseResponse) SetHeaders(v map[string]*string) *CreateExternalDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CreateExternalDatabaseResponse) SetStatusCode(v int32) *CreateExternalDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExternalDatabaseResponse) SetBody(v *CreateExternalDatabaseResponseBody) *CreateExternalDatabaseResponse {
	s.Body = v
	return s
}

func (s *CreateExternalDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
