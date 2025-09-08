// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceTypesResponseBody) *ListDataSourceTypesResponse
	GetBody() *ListDataSourceTypesResponseBody
}

type ListDataSourceTypesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceTypesResponse) GetBody() *ListDataSourceTypesResponseBody {
	return s.Body
}

func (s *ListDataSourceTypesResponse) SetHeaders(v map[string]*string) *ListDataSourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTypesResponse) SetStatusCode(v int32) *ListDataSourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTypesResponse) SetBody(v *ListDataSourceTypesResponseBody) *ListDataSourceTypesResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceTypesResponse) Validate() error {
	return dara.Validate(s)
}
