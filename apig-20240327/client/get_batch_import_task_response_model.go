// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchImportTaskResponseBody) *GetBatchImportTaskResponse
	GetBody() *GetBatchImportTaskResponseBody
}

type GetBatchImportTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskResponse) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchImportTaskResponse) GetBody() *GetBatchImportTaskResponseBody {
	return s.Body
}

func (s *GetBatchImportTaskResponse) SetHeaders(v map[string]*string) *GetBatchImportTaskResponse {
	s.Headers = v
	return s
}

func (s *GetBatchImportTaskResponse) SetStatusCode(v int32) *GetBatchImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchImportTaskResponse) SetBody(v *GetBatchImportTaskResponseBody) *GetBatchImportTaskResponse {
	s.Body = v
	return s
}

func (s *GetBatchImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
