// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMigrationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMigrationsResponse
	GetStatusCode() *int32
	SetBody(v *ListMigrationsResponseBody) *ListMigrationsResponse
	GetBody() *ListMigrationsResponseBody
}

type ListMigrationsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMigrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMigrationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponse) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMigrationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMigrationsResponse) GetBody() *ListMigrationsResponseBody {
	return s.Body
}

func (s *ListMigrationsResponse) SetHeaders(v map[string]*string) *ListMigrationsResponse {
	s.Headers = v
	return s
}

func (s *ListMigrationsResponse) SetStatusCode(v int32) *ListMigrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMigrationsResponse) SetBody(v *ListMigrationsResponseBody) *ListMigrationsResponse {
	s.Body = v
	return s
}

func (s *ListMigrationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
