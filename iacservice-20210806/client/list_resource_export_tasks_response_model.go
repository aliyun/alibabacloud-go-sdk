// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceExportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceExportTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceExportTasksResponseBody) *ListResourceExportTasksResponse
	GetBody() *ListResourceExportTasksResponseBody
}

type ListResourceExportTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceExportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceExportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceExportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceExportTasksResponse) GetBody() *ListResourceExportTasksResponseBody {
	return s.Body
}

func (s *ListResourceExportTasksResponse) SetHeaders(v map[string]*string) *ListResourceExportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListResourceExportTasksResponse) SetStatusCode(v int32) *ListResourceExportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceExportTasksResponse) SetBody(v *ListResourceExportTasksResponseBody) *ListResourceExportTasksResponse {
	s.Body = v
	return s
}

func (s *ListResourceExportTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
