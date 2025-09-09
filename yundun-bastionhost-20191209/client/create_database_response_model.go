// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatabaseResponseBody) *CreateDatabaseResponse
	GetBody() *CreateDatabaseResponseBody
}

type CreateDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatabaseResponse) GetBody() *CreateDatabaseResponseBody {
	return s.Body
}

func (s *CreateDatabaseResponse) SetHeaders(v map[string]*string) *CreateDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseResponse) SetStatusCode(v int32) *CreateDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatabaseResponse) SetBody(v *CreateDatabaseResponseBody) *CreateDatabaseResponse {
	s.Body = v
	return s
}

func (s *CreateDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
