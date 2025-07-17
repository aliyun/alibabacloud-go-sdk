// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *SearchDatabaseResponseBody) *SearchDatabaseResponse
	GetBody() *SearchDatabaseResponseBody
}

type SearchDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchDatabaseResponse) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchDatabaseResponse) GetBody() *SearchDatabaseResponseBody {
	return s.Body
}

func (s *SearchDatabaseResponse) SetHeaders(v map[string]*string) *SearchDatabaseResponse {
	s.Headers = v
	return s
}

func (s *SearchDatabaseResponse) SetStatusCode(v int32) *SearchDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDatabaseResponse) SetBody(v *SearchDatabaseResponseBody) *SearchDatabaseResponse {
	s.Body = v
	return s
}

func (s *SearchDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
