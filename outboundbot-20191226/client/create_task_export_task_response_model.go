// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTaskExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTaskExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTaskExportTaskResponseBody) *CreateTaskExportTaskResponse
	GetBody() *CreateTaskExportTaskResponseBody
}

type CreateTaskExportTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTaskExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTaskExportTaskResponse) GetBody() *CreateTaskExportTaskResponseBody {
	return s.Body
}

func (s *CreateTaskExportTaskResponse) SetHeaders(v map[string]*string) *CreateTaskExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskExportTaskResponse) SetStatusCode(v int32) *CreateTaskExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskExportTaskResponse) SetBody(v *CreateTaskExportTaskResponseBody) *CreateTaskExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTaskExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
