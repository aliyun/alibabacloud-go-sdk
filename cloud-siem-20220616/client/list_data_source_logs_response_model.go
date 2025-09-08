// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceLogsResponseBody) *ListDataSourceLogsResponse
	GetBody() *ListDataSourceLogsResponseBody
}

type ListDataSourceLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceLogsResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceLogsResponse) GetBody() *ListDataSourceLogsResponseBody {
	return s.Body
}

func (s *ListDataSourceLogsResponse) SetHeaders(v map[string]*string) *ListDataSourceLogsResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceLogsResponse) SetStatusCode(v int32) *ListDataSourceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceLogsResponse) SetBody(v *ListDataSourceLogsResponseBody) *ListDataSourceLogsResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceLogsResponse) Validate() error {
	return dara.Validate(s)
}
