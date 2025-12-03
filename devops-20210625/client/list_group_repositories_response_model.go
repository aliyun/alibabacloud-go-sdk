// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupRepositoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupRepositoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupRepositoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupRepositoriesResponseBody) *ListGroupRepositoriesResponse
	GetBody() *ListGroupRepositoriesResponseBody
}

type ListGroupRepositoriesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupRepositoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupRepositoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupRepositoriesResponse) GetBody() *ListGroupRepositoriesResponseBody {
	return s.Body
}

func (s *ListGroupRepositoriesResponse) SetHeaders(v map[string]*string) *ListGroupRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListGroupRepositoriesResponse) SetStatusCode(v int32) *ListGroupRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupRepositoriesResponse) SetBody(v *ListGroupRepositoriesResponseBody) *ListGroupRepositoriesResponse {
	s.Body = v
	return s
}

func (s *ListGroupRepositoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
