// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ImportDataSourcesResponseBody) *ImportDataSourcesResponse
	GetBody() *ImportDataSourcesResponseBody
}

type ImportDataSourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ImportDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportDataSourcesResponse) GetBody() *ImportDataSourcesResponseBody {
	return s.Body
}

func (s *ImportDataSourcesResponse) SetHeaders(v map[string]*string) *ImportDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ImportDataSourcesResponse) SetStatusCode(v int32) *ImportDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportDataSourcesResponse) SetBody(v *ImportDataSourcesResponseBody) *ImportDataSourcesResponse {
	s.Body = v
	return s
}

func (s *ImportDataSourcesResponse) Validate() error {
	return dara.Validate(s)
}
