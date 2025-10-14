// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetExportTaskResponseBody) *GetExportTaskResponse
	GetBody() *GetExportTaskResponseBody
}

type GetExportTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExportTaskResponse) GoString() string {
	return s.String()
}

func (s *GetExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExportTaskResponse) GetBody() *GetExportTaskResponseBody {
	return s.Body
}

func (s *GetExportTaskResponse) SetHeaders(v map[string]*string) *GetExportTaskResponse {
	s.Headers = v
	return s
}

func (s *GetExportTaskResponse) SetStatusCode(v int32) *GetExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExportTaskResponse) SetBody(v *GetExportTaskResponseBody) *GetExportTaskResponse {
	s.Body = v
	return s
}

func (s *GetExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
