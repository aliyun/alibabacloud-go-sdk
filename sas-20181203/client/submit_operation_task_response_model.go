// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitOperationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitOperationTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitOperationTaskResponseBody) *SubmitOperationTaskResponse
	GetBody() *SubmitOperationTaskResponseBody
}

type SubmitOperationTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitOperationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitOperationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitOperationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitOperationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitOperationTaskResponse) GetBody() *SubmitOperationTaskResponseBody {
	return s.Body
}

func (s *SubmitOperationTaskResponse) SetHeaders(v map[string]*string) *SubmitOperationTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitOperationTaskResponse) SetStatusCode(v int32) *SubmitOperationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitOperationTaskResponse) SetBody(v *SubmitOperationTaskResponseBody) *SubmitOperationTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitOperationTaskResponse) Validate() error {
	return dara.Validate(s)
}
