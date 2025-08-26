// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataLakeDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataLakeDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataLakeDatabaseResponseBody) *CreateDataLakeDatabaseResponse
	GetBody() *CreateDataLakeDatabaseResponseBody
}

type CreateDataLakeDatabaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataLakeDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataLakeDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateDataLakeDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataLakeDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataLakeDatabaseResponse) GetBody() *CreateDataLakeDatabaseResponseBody {
	return s.Body
}

func (s *CreateDataLakeDatabaseResponse) SetHeaders(v map[string]*string) *CreateDataLakeDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CreateDataLakeDatabaseResponse) SetStatusCode(v int32) *CreateDataLakeDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataLakeDatabaseResponse) SetBody(v *CreateDataLakeDatabaseResponseBody) *CreateDataLakeDatabaseResponse {
	s.Body = v
	return s
}

func (s *CreateDataLakeDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
