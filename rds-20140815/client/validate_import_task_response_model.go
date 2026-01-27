// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *ValidateImportTaskResponseBody) *ValidateImportTaskResponse
	GetBody() *ValidateImportTaskResponseBody
}

type ValidateImportTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateImportTaskResponse) GoString() string {
	return s.String()
}

func (s *ValidateImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateImportTaskResponse) GetBody() *ValidateImportTaskResponseBody {
	return s.Body
}

func (s *ValidateImportTaskResponse) SetHeaders(v map[string]*string) *ValidateImportTaskResponse {
	s.Headers = v
	return s
}

func (s *ValidateImportTaskResponse) SetStatusCode(v int32) *ValidateImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateImportTaskResponse) SetBody(v *ValidateImportTaskResponseBody) *ValidateImportTaskResponse {
	s.Body = v
	return s
}

func (s *ValidateImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
