// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePracticeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeletePracticeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeletePracticeTaskResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeletePracticeTaskResponseBody) *BatchDeletePracticeTaskResponse
	GetBody() *BatchDeletePracticeTaskResponseBody
}

type BatchDeletePracticeTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeletePracticeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeletePracticeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePracticeTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchDeletePracticeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeletePracticeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeletePracticeTaskResponse) GetBody() *BatchDeletePracticeTaskResponseBody {
	return s.Body
}

func (s *BatchDeletePracticeTaskResponse) SetHeaders(v map[string]*string) *BatchDeletePracticeTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchDeletePracticeTaskResponse) SetStatusCode(v int32) *BatchDeletePracticeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeletePracticeTaskResponse) SetBody(v *BatchDeletePracticeTaskResponseBody) *BatchDeletePracticeTaskResponse {
	s.Body = v
	return s
}

func (s *BatchDeletePracticeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
