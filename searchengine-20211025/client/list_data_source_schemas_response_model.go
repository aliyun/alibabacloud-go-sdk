// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceSchemasResponseBody) *ListDataSourceSchemasResponse
	GetBody() *ListDataSourceSchemasResponseBody
}

type ListDataSourceSchemasResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceSchemasResponse) GetBody() *ListDataSourceSchemasResponseBody {
	return s.Body
}

func (s *ListDataSourceSchemasResponse) SetHeaders(v map[string]*string) *ListDataSourceSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceSchemasResponse) SetStatusCode(v int32) *ListDataSourceSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceSchemasResponse) SetBody(v *ListDataSourceSchemasResponseBody) *ListDataSourceSchemasResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceSchemasResponse) Validate() error {
	return dara.Validate(s)
}
