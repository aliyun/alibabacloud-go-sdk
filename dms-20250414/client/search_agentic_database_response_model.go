// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAgenticDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAgenticDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAgenticDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *SearchAgenticDatabaseResponseBody) *SearchAgenticDatabaseResponse
	GetBody() *SearchAgenticDatabaseResponseBody
}

type SearchAgenticDatabaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAgenticDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAgenticDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAgenticDatabaseResponse) GoString() string {
	return s.String()
}

func (s *SearchAgenticDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAgenticDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAgenticDatabaseResponse) GetBody() *SearchAgenticDatabaseResponseBody {
	return s.Body
}

func (s *SearchAgenticDatabaseResponse) SetHeaders(v map[string]*string) *SearchAgenticDatabaseResponse {
	s.Headers = v
	return s
}

func (s *SearchAgenticDatabaseResponse) SetStatusCode(v int32) *SearchAgenticDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAgenticDatabaseResponse) SetBody(v *SearchAgenticDatabaseResponseBody) *SearchAgenticDatabaseResponse {
	s.Body = v
	return s
}

func (s *SearchAgenticDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
