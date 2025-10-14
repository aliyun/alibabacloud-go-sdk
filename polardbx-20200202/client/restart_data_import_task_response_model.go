// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDataImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDataImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *RestartDataImportTaskResponseBody) *RestartDataImportTaskResponse
	GetBody() *RestartDataImportTaskResponseBody
}

type RestartDataImportTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDataImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDataImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDataImportTaskResponse) GoString() string {
	return s.String()
}

func (s *RestartDataImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDataImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDataImportTaskResponse) GetBody() *RestartDataImportTaskResponseBody {
	return s.Body
}

func (s *RestartDataImportTaskResponse) SetHeaders(v map[string]*string) *RestartDataImportTaskResponse {
	s.Headers = v
	return s
}

func (s *RestartDataImportTaskResponse) SetStatusCode(v int32) *RestartDataImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDataImportTaskResponse) SetBody(v *RestartDataImportTaskResponseBody) *RestartDataImportTaskResponse {
	s.Body = v
	return s
}

func (s *RestartDataImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
