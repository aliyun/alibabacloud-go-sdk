// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitBatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitBatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitBatchTaskResponseBody) *SubmitBatchTaskResponse
	GetBody() *SubmitBatchTaskResponseBody
}

type SubmitBatchTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitBatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitBatchTaskResponse) GetBody() *SubmitBatchTaskResponseBody {
	return s.Body
}

func (s *SubmitBatchTaskResponse) SetHeaders(v map[string]*string) *SubmitBatchTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitBatchTaskResponse) SetStatusCode(v int32) *SubmitBatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBatchTaskResponse) SetBody(v *SubmitBatchTaskResponseBody) *SubmitBatchTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitBatchTaskResponse) Validate() error {
	return dara.Validate(s)
}
