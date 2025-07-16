// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserUsageDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserUsageDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserUsageDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserUsageDataExportTaskResponseBody) *CreateUserUsageDataExportTaskResponse
	GetBody() *CreateUserUsageDataExportTaskResponseBody
}

type CreateUserUsageDataExportTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserUsageDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserUsageDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserUsageDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUserUsageDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserUsageDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserUsageDataExportTaskResponse) GetBody() *CreateUserUsageDataExportTaskResponseBody {
	return s.Body
}

func (s *CreateUserUsageDataExportTaskResponse) SetHeaders(v map[string]*string) *CreateUserUsageDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUserUsageDataExportTaskResponse) SetStatusCode(v int32) *CreateUserUsageDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponse) SetBody(v *CreateUserUsageDataExportTaskResponseBody) *CreateUserUsageDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateUserUsageDataExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
