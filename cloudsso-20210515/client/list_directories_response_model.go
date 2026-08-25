// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListDirectoriesResponseBody) *ListDirectoriesResponse
	GetBody() *ListDirectoriesResponseBody
}

type ListDirectoriesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDirectoriesResponse) GetBody() *ListDirectoriesResponseBody {
	return s.Body
}

func (s *ListDirectoriesResponse) SetHeaders(v map[string]*string) *ListDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoriesResponse) SetStatusCode(v int32) *ListDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectoriesResponse) SetBody(v *ListDirectoriesResponseBody) *ListDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
