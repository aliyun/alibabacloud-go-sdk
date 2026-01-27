// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImportTaskResponseBody) *CreateImportTaskResponse
	GetBody() *CreateImportTaskResponseBody
}

type CreateImportTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImportTaskResponse) GetBody() *CreateImportTaskResponseBody {
	return s.Body
}

func (s *CreateImportTaskResponse) SetHeaders(v map[string]*string) *CreateImportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImportTaskResponse) SetStatusCode(v int32) *CreateImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImportTaskResponse) SetBody(v *CreateImportTaskResponseBody) *CreateImportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
