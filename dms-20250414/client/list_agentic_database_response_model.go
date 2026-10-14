// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgenticDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgenticDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *ListAgenticDatabaseResponseBody) *ListAgenticDatabaseResponse
	GetBody() *ListAgenticDatabaseResponseBody
}

type ListAgenticDatabaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgenticDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgenticDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ListAgenticDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgenticDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgenticDatabaseResponse) GetBody() *ListAgenticDatabaseResponseBody {
	return s.Body
}

func (s *ListAgenticDatabaseResponse) SetHeaders(v map[string]*string) *ListAgenticDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ListAgenticDatabaseResponse) SetStatusCode(v int32) *ListAgenticDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgenticDatabaseResponse) SetBody(v *ListAgenticDatabaseResponseBody) *ListAgenticDatabaseResponse {
	s.Body = v
	return s
}

func (s *ListAgenticDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
