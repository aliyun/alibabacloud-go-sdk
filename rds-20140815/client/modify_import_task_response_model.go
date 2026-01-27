// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyImportTaskResponseBody) *ModifyImportTaskResponse
	GetBody() *ModifyImportTaskResponseBody
}

type ModifyImportTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyImportTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyImportTaskResponse) GetBody() *ModifyImportTaskResponseBody {
	return s.Body
}

func (s *ModifyImportTaskResponse) SetHeaders(v map[string]*string) *ModifyImportTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyImportTaskResponse) SetStatusCode(v int32) *ModifyImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImportTaskResponse) SetBody(v *ModifyImportTaskResponseBody) *ModifyImportTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
