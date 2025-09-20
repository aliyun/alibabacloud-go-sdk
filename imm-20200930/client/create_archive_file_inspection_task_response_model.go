// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArchiveFileInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateArchiveFileInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateArchiveFileInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateArchiveFileInspectionTaskResponseBody) *CreateArchiveFileInspectionTaskResponse
	GetBody() *CreateArchiveFileInspectionTaskResponseBody
}

type CreateArchiveFileInspectionTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArchiveFileInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArchiveFileInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateArchiveFileInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateArchiveFileInspectionTaskResponse) GetBody() *CreateArchiveFileInspectionTaskResponseBody {
	return s.Body
}

func (s *CreateArchiveFileInspectionTaskResponse) SetHeaders(v map[string]*string) *CreateArchiveFileInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponse) SetStatusCode(v int32) *CreateArchiveFileInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponse) SetBody(v *CreateArchiveFileInspectionTaskResponseBody) *CreateArchiveFileInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponse) Validate() error {
	return dara.Validate(s)
}
