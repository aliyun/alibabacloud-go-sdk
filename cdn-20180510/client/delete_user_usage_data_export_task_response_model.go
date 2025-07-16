// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserUsageDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserUsageDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserUsageDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserUsageDataExportTaskResponseBody) *DeleteUserUsageDataExportTaskResponse
	GetBody() *DeleteUserUsageDataExportTaskResponseBody
}

type DeleteUserUsageDataExportTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserUsageDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserUsageDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserUsageDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserUsageDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserUsageDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserUsageDataExportTaskResponse) GetBody() *DeleteUserUsageDataExportTaskResponseBody {
	return s.Body
}

func (s *DeleteUserUsageDataExportTaskResponse) SetHeaders(v map[string]*string) *DeleteUserUsageDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserUsageDataExportTaskResponse) SetStatusCode(v int32) *DeleteUserUsageDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserUsageDataExportTaskResponse) SetBody(v *DeleteUserUsageDataExportTaskResponseBody) *DeleteUserUsageDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteUserUsageDataExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
