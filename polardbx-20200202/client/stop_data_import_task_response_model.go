// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDataImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDataImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopDataImportTaskResponseBody) *StopDataImportTaskResponse
	GetBody() *StopDataImportTaskResponseBody
}

type StopDataImportTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDataImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDataImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDataImportTaskResponse) GoString() string {
	return s.String()
}

func (s *StopDataImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDataImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDataImportTaskResponse) GetBody() *StopDataImportTaskResponseBody {
	return s.Body
}

func (s *StopDataImportTaskResponse) SetHeaders(v map[string]*string) *StopDataImportTaskResponse {
	s.Headers = v
	return s
}

func (s *StopDataImportTaskResponse) SetStatusCode(v int32) *StopDataImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDataImportTaskResponse) SetBody(v *StopDataImportTaskResponseBody) *StopDataImportTaskResponse {
	s.Body = v
	return s
}

func (s *StopDataImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
