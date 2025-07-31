// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBatchTaskResponseBody) *UpdateBatchTaskResponse
	GetBody() *UpdateBatchTaskResponseBody
}

type UpdateBatchTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBatchTaskResponse) GetBody() *UpdateBatchTaskResponseBody {
	return s.Body
}

func (s *UpdateBatchTaskResponse) SetHeaders(v map[string]*string) *UpdateBatchTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateBatchTaskResponse) SetStatusCode(v int32) *UpdateBatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBatchTaskResponse) SetBody(v *UpdateBatchTaskResponseBody) *UpdateBatchTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateBatchTaskResponse) Validate() error {
	return dara.Validate(s)
}
