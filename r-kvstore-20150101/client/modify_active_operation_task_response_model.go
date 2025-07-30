// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyActiveOperationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyActiveOperationTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyActiveOperationTaskResponseBody) *ModifyActiveOperationTaskResponse
	GetBody() *ModifyActiveOperationTaskResponseBody
}

type ModifyActiveOperationTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActiveOperationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActiveOperationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyActiveOperationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyActiveOperationTaskResponse) GetBody() *ModifyActiveOperationTaskResponseBody {
	return s.Body
}

func (s *ModifyActiveOperationTaskResponse) SetHeaders(v map[string]*string) *ModifyActiveOperationTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyActiveOperationTaskResponse) SetStatusCode(v int32) *ModifyActiveOperationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActiveOperationTaskResponse) SetBody(v *ModifyActiveOperationTaskResponseBody) *ModifyActiveOperationTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyActiveOperationTaskResponse) Validate() error {
	return dara.Validate(s)
}
