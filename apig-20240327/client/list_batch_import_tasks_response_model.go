// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchImportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBatchImportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBatchImportTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListBatchImportTasksResponseBody) *ListBatchImportTasksResponse
	GetBody() *ListBatchImportTasksResponseBody
}

type ListBatchImportTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBatchImportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBatchImportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBatchImportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListBatchImportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBatchImportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBatchImportTasksResponse) GetBody() *ListBatchImportTasksResponseBody {
	return s.Body
}

func (s *ListBatchImportTasksResponse) SetHeaders(v map[string]*string) *ListBatchImportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListBatchImportTasksResponse) SetStatusCode(v int32) *ListBatchImportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchImportTasksResponse) SetBody(v *ListBatchImportTasksResponseBody) *ListBatchImportTasksResponse {
	s.Body = v
	return s
}

func (s *ListBatchImportTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
