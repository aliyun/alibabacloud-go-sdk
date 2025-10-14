// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateExportTaskResponseBody) *CreateExportTaskResponse
	GetBody() *CreateExportTaskResponseBody
}

type CreateExportTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExportTaskResponse) GetBody() *CreateExportTaskResponseBody {
	return s.Body
}

func (s *CreateExportTaskResponse) SetHeaders(v map[string]*string) *CreateExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateExportTaskResponse) SetStatusCode(v int32) *CreateExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExportTaskResponse) SetBody(v *CreateExportTaskResponseBody) *CreateExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
