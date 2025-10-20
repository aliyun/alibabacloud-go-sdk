// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabaseDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabaseDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabaseDetailsResponseBody) *ListDatabaseDetailsResponse
	GetBody() *ListDatabaseDetailsResponseBody
}

type ListDatabaseDetailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabaseDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabaseDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabaseDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabaseDetailsResponse) GetBody() *ListDatabaseDetailsResponseBody {
	return s.Body
}

func (s *ListDatabaseDetailsResponse) SetHeaders(v map[string]*string) *ListDatabaseDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseDetailsResponse) SetStatusCode(v int32) *ListDatabaseDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseDetailsResponse) SetBody(v *ListDatabaseDetailsResponseBody) *ListDatabaseDetailsResponse {
	s.Body = v
	return s
}

func (s *ListDatabaseDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
