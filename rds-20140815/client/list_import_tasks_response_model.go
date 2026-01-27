// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImportTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListImportTasksResponseBody) *ListImportTasksResponse
	GetBody() *ListImportTasksResponseBody
}

type ListImportTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListImportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImportTasksResponse) GetBody() *ListImportTasksResponseBody {
	return s.Body
}

func (s *ListImportTasksResponse) SetHeaders(v map[string]*string) *ListImportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListImportTasksResponse) SetStatusCode(v int32) *ListImportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImportTasksResponse) SetBody(v *ListImportTasksResponseBody) *ListImportTasksResponse {
	s.Body = v
	return s
}

func (s *ListImportTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
