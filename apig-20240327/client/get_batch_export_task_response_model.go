// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchExportTaskResponseBody) *GetBatchExportTaskResponse
	GetBody() *GetBatchExportTaskResponseBody
}

type GetBatchExportTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchExportTaskResponse) GoString() string {
	return s.String()
}

func (s *GetBatchExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchExportTaskResponse) GetBody() *GetBatchExportTaskResponseBody {
	return s.Body
}

func (s *GetBatchExportTaskResponse) SetHeaders(v map[string]*string) *GetBatchExportTaskResponse {
	s.Headers = v
	return s
}

func (s *GetBatchExportTaskResponse) SetStatusCode(v int32) *GetBatchExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchExportTaskResponse) SetBody(v *GetBatchExportTaskResponseBody) *GetBatchExportTaskResponse {
	s.Body = v
	return s
}

func (s *GetBatchExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
