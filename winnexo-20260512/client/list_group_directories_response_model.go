// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupDirectoriesResponseBody) *ListGroupDirectoriesResponse
	GetBody() *ListGroupDirectoriesResponseBody
}

type ListGroupDirectoriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupDirectoriesResponse) GetBody() *ListGroupDirectoriesResponseBody {
	return s.Body
}

func (s *ListGroupDirectoriesResponse) SetHeaders(v map[string]*string) *ListGroupDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListGroupDirectoriesResponse) SetStatusCode(v int32) *ListGroupDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupDirectoriesResponse) SetBody(v *ListGroupDirectoriesResponseBody) *ListGroupDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListGroupDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
