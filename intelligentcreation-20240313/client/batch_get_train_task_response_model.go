// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetTrainTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetTrainTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetTrainTaskResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetTrainTaskResponseBody) *BatchGetTrainTaskResponse
	GetBody() *BatchGetTrainTaskResponseBody
}

type BatchGetTrainTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetTrainTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *BatchGetTrainTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetTrainTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetTrainTaskResponse) GetBody() *BatchGetTrainTaskResponseBody {
	return s.Body
}

func (s *BatchGetTrainTaskResponse) SetHeaders(v map[string]*string) *BatchGetTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *BatchGetTrainTaskResponse) SetStatusCode(v int32) *BatchGetTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetTrainTaskResponse) SetBody(v *BatchGetTrainTaskResponseBody) *BatchGetTrainTaskResponse {
	s.Body = v
	return s
}

func (s *BatchGetTrainTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
