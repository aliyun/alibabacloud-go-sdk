// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceResponseBody) *ListDataSourceResponse
	GetBody() *ListDataSourceResponseBody
}

type ListDataSourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceResponse) GetBody() *ListDataSourceResponseBody {
	return s.Body
}

func (s *ListDataSourceResponse) SetHeaders(v map[string]*string) *ListDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceResponse) SetStatusCode(v int32) *ListDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceResponse) SetBody(v *ListDataSourceResponseBody) *ListDataSourceResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
