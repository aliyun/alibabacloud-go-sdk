// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgenticDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgenticDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgenticDataSourcesResponseBody) *ListAgenticDataSourcesResponse
	GetBody() *ListAgenticDataSourcesResponseBody
}

type ListAgenticDataSourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgenticDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgenticDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAgenticDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgenticDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgenticDataSourcesResponse) GetBody() *ListAgenticDataSourcesResponseBody {
	return s.Body
}

func (s *ListAgenticDataSourcesResponse) SetHeaders(v map[string]*string) *ListAgenticDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAgenticDataSourcesResponse) SetStatusCode(v int32) *ListAgenticDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgenticDataSourcesResponse) SetBody(v *ListAgenticDataSourcesResponseBody) *ListAgenticDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ListAgenticDataSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
