// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExportTaskVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceExportTaskVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceExportTaskVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceExportTaskVersionsResponseBody) *ListResourceExportTaskVersionsResponse
	GetBody() *ListResourceExportTaskVersionsResponseBody
}

type ListResourceExportTaskVersionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceExportTaskVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceExportTaskVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceExportTaskVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceExportTaskVersionsResponse) GetBody() *ListResourceExportTaskVersionsResponseBody {
	return s.Body
}

func (s *ListResourceExportTaskVersionsResponse) SetHeaders(v map[string]*string) *ListResourceExportTaskVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceExportTaskVersionsResponse) SetStatusCode(v int32) *ListResourceExportTaskVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponse) SetBody(v *ListResourceExportTaskVersionsResponseBody) *ListResourceExportTaskVersionsResponse {
	s.Body = v
	return s
}

func (s *ListResourceExportTaskVersionsResponse) Validate() error {
	return dara.Validate(s)
}
