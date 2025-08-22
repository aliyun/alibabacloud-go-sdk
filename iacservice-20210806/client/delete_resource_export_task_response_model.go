// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceExportTaskResponseBody) *DeleteResourceExportTaskResponse
	GetBody() *DeleteResourceExportTaskResponseBody
}

type DeleteResourceExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceExportTaskResponse) GetBody() *DeleteResourceExportTaskResponseBody {
	return s.Body
}

func (s *DeleteResourceExportTaskResponse) SetHeaders(v map[string]*string) *DeleteResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceExportTaskResponse) SetStatusCode(v int32) *DeleteResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceExportTaskResponse) SetBody(v *DeleteResourceExportTaskResponseBody) *DeleteResourceExportTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
