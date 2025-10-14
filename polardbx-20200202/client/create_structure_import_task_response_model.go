// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStructureImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStructureImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStructureImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateStructureImportTaskResponseBody) *CreateStructureImportTaskResponse
	GetBody() *CreateStructureImportTaskResponseBody
}

type CreateStructureImportTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStructureImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStructureImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStructureImportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateStructureImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStructureImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStructureImportTaskResponse) GetBody() *CreateStructureImportTaskResponseBody {
	return s.Body
}

func (s *CreateStructureImportTaskResponse) SetHeaders(v map[string]*string) *CreateStructureImportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateStructureImportTaskResponse) SetStatusCode(v int32) *CreateStructureImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStructureImportTaskResponse) SetBody(v *CreateStructureImportTaskResponseBody) *CreateStructureImportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateStructureImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
