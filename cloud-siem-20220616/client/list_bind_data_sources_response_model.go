// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBindDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBindDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListBindDataSourcesResponseBody) *ListBindDataSourcesResponse
	GetBody() *ListBindDataSourcesResponseBody
}

type ListBindDataSourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBindDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListBindDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBindDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBindDataSourcesResponse) GetBody() *ListBindDataSourcesResponseBody {
	return s.Body
}

func (s *ListBindDataSourcesResponse) SetHeaders(v map[string]*string) *ListBindDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListBindDataSourcesResponse) SetStatusCode(v int32) *ListBindDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindDataSourcesResponse) SetBody(v *ListBindDataSourcesResponseBody) *ListBindDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ListBindDataSourcesResponse) Validate() error {
	return dara.Validate(s)
}
