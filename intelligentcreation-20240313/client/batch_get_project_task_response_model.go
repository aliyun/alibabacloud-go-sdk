// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetProjectTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetProjectTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetProjectTaskResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetProjectTaskResponseBody) *BatchGetProjectTaskResponse
	GetBody() *BatchGetProjectTaskResponseBody
}

type BatchGetProjectTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetProjectTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchGetProjectTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetProjectTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetProjectTaskResponse) GetBody() *BatchGetProjectTaskResponseBody {
	return s.Body
}

func (s *BatchGetProjectTaskResponse) SetHeaders(v map[string]*string) *BatchGetProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchGetProjectTaskResponse) SetStatusCode(v int32) *BatchGetProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetProjectTaskResponse) SetBody(v *BatchGetProjectTaskResponseBody) *BatchGetProjectTaskResponse {
	s.Body = v
	return s
}

func (s *BatchGetProjectTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
