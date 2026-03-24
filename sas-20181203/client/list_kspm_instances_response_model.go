// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKspmInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKspmInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKspmInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListKspmInstancesResponseBody) *ListKspmInstancesResponse
	GetBody() *ListKspmInstancesResponseBody
}

type ListKspmInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKspmInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKspmInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKspmInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListKspmInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKspmInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKspmInstancesResponse) GetBody() *ListKspmInstancesResponseBody {
	return s.Body
}

func (s *ListKspmInstancesResponse) SetHeaders(v map[string]*string) *ListKspmInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListKspmInstancesResponse) SetStatusCode(v int32) *ListKspmInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKspmInstancesResponse) SetBody(v *ListKspmInstancesResponseBody) *ListKspmInstancesResponse {
	s.Body = v
	return s
}

func (s *ListKspmInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
