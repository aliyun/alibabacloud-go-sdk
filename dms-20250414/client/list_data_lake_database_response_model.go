// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeDatabaseResponseBody) *ListDataLakeDatabaseResponse
	GetBody() *ListDataLakeDatabaseResponseBody
}

type ListDataLakeDatabaseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeDatabaseResponse) GetBody() *ListDataLakeDatabaseResponseBody {
	return s.Body
}

func (s *ListDataLakeDatabaseResponse) SetHeaders(v map[string]*string) *ListDataLakeDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeDatabaseResponse) SetStatusCode(v int32) *ListDataLakeDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeDatabaseResponse) SetBody(v *ListDataLakeDatabaseResponseBody) *ListDataLakeDatabaseResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
