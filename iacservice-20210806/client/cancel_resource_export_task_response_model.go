// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelResourceExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelResourceExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelResourceExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelResourceExportTaskResponseBody) *CancelResourceExportTaskResponse
	GetBody() *CancelResourceExportTaskResponseBody
}

type CancelResourceExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelResourceExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelResourceExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelResourceExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelResourceExportTaskResponse) GetBody() *CancelResourceExportTaskResponseBody {
	return s.Body
}

func (s *CancelResourceExportTaskResponse) SetHeaders(v map[string]*string) *CancelResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelResourceExportTaskResponse) SetStatusCode(v int32) *CancelResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelResourceExportTaskResponse) SetBody(v *CancelResourceExportTaskResponseBody) *CancelResourceExportTaskResponse {
	s.Body = v
	return s
}

func (s *CancelResourceExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
