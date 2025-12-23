// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExampleQueriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExampleQueriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExampleQueriesResponse
	GetStatusCode() *int32
	SetBody(v *ListExampleQueriesResponseBody) *ListExampleQueriesResponse
	GetBody() *ListExampleQueriesResponseBody
}

type ListExampleQueriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExampleQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExampleQueriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExampleQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExampleQueriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExampleQueriesResponse) GetBody() *ListExampleQueriesResponseBody {
	return s.Body
}

func (s *ListExampleQueriesResponse) SetHeaders(v map[string]*string) *ListExampleQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListExampleQueriesResponse) SetStatusCode(v int32) *ListExampleQueriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExampleQueriesResponse) SetBody(v *ListExampleQueriesResponseBody) *ListExampleQueriesResponse {
	s.Body = v
	return s
}

func (s *ListExampleQueriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
