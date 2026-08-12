// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchExportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBatchExportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBatchExportTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListBatchExportTasksResponseBody) *ListBatchExportTasksResponse
	GetBody() *ListBatchExportTasksResponseBody
}

type ListBatchExportTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBatchExportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBatchExportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBatchExportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListBatchExportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBatchExportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBatchExportTasksResponse) GetBody() *ListBatchExportTasksResponseBody {
	return s.Body
}

func (s *ListBatchExportTasksResponse) SetHeaders(v map[string]*string) *ListBatchExportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListBatchExportTasksResponse) SetStatusCode(v int32) *ListBatchExportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchExportTasksResponse) SetBody(v *ListBatchExportTasksResponseBody) *ListBatchExportTasksResponse {
	s.Body = v
	return s
}

func (s *ListBatchExportTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
