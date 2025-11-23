// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataLakeDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataLakeDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataLakeDatabaseResponseBody) *UpdateDataLakeDatabaseResponse
	GetBody() *UpdateDataLakeDatabaseResponseBody
}

type UpdateDataLakeDatabaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataLakeDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataLakeDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeDatabaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataLakeDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataLakeDatabaseResponse) GetBody() *UpdateDataLakeDatabaseResponseBody {
	return s.Body
}

func (s *UpdateDataLakeDatabaseResponse) SetHeaders(v map[string]*string) *UpdateDataLakeDatabaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataLakeDatabaseResponse) SetStatusCode(v int32) *UpdateDataLakeDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataLakeDatabaseResponse) SetBody(v *UpdateDataLakeDatabaseResponseBody) *UpdateDataLakeDatabaseResponse {
	s.Body = v
	return s
}

func (s *UpdateDataLakeDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
