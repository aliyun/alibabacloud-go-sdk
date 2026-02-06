// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobGroupExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJobGroupExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJobGroupExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateJobGroupExportTaskResponseBody) *CreateJobGroupExportTaskResponse
	GetBody() *CreateJobGroupExportTaskResponseBody
}

type CreateJobGroupExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobGroupExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobGroupExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateJobGroupExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJobGroupExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJobGroupExportTaskResponse) GetBody() *CreateJobGroupExportTaskResponseBody {
	return s.Body
}

func (s *CreateJobGroupExportTaskResponse) SetHeaders(v map[string]*string) *CreateJobGroupExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateJobGroupExportTaskResponse) SetStatusCode(v int32) *CreateJobGroupExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobGroupExportTaskResponse) SetBody(v *CreateJobGroupExportTaskResponseBody) *CreateJobGroupExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateJobGroupExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
