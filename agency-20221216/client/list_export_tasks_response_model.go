// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExportTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListExportTasksResponseBody) *ListExportTasksResponse
	GetBody() *ListExportTasksResponseBody
}

type ListExportTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListExportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExportTasksResponse) GetBody() *ListExportTasksResponseBody {
	return s.Body
}

func (s *ListExportTasksResponse) SetHeaders(v map[string]*string) *ListExportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListExportTasksResponse) SetStatusCode(v int32) *ListExportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExportTasksResponse) SetBody(v *ListExportTasksResponseBody) *ListExportTasksResponse {
	s.Body = v
	return s
}

func (s *ListExportTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
