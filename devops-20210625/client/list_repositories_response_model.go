// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoriesResponseBody) *ListRepositoriesResponse
	GetBody() *ListRepositoriesResponseBody
}

type ListRepositoriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoriesResponse) GetBody() *ListRepositoriesResponseBody {
	return s.Body
}

func (s *ListRepositoriesResponse) SetHeaders(v map[string]*string) *ListRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoriesResponse) SetStatusCode(v int32) *ListRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoriesResponse) SetBody(v *ListRepositoriesResponseBody) *ListRepositoriesResponse {
	s.Body = v
	return s
}

func (s *ListRepositoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
