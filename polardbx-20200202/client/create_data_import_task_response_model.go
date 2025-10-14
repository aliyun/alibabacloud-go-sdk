// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataImportTaskResponseBody) *CreateDataImportTaskResponse
	GetBody() *CreateDataImportTaskResponseBody
}

type CreateDataImportTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDataImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataImportTaskResponse) GetBody() *CreateDataImportTaskResponseBody {
	return s.Body
}

func (s *CreateDataImportTaskResponse) SetHeaders(v map[string]*string) *CreateDataImportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDataImportTaskResponse) SetStatusCode(v int32) *CreateDataImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataImportTaskResponse) SetBody(v *CreateDataImportTaskResponseBody) *CreateDataImportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDataImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
