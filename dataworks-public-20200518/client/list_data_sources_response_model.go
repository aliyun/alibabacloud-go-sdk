// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse
	GetBody() *ListDataSourcesResponseBody
}

type ListDataSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourcesResponse) GetBody() *ListDataSourcesResponseBody {
	return s.Body
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetStatusCode(v int32) *ListDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ListDataSourcesResponse) Validate() error {
	return dara.Validate(s)
}
