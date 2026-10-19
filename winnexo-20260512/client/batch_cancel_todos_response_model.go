// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTodosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCancelTodosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCancelTodosResponse
	GetStatusCode() *int32
	SetBody(v *BatchCancelTodosResponseBody) *BatchCancelTodosResponse
	GetBody() *BatchCancelTodosResponseBody
}

type BatchCancelTodosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCancelTodosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCancelTodosResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTodosResponse) GoString() string {
	return s.String()
}

func (s *BatchCancelTodosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCancelTodosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCancelTodosResponse) GetBody() *BatchCancelTodosResponseBody {
	return s.Body
}

func (s *BatchCancelTodosResponse) SetHeaders(v map[string]*string) *BatchCancelTodosResponse {
	s.Headers = v
	return s
}

func (s *BatchCancelTodosResponse) SetStatusCode(v int32) *BatchCancelTodosResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCancelTodosResponse) SetBody(v *BatchCancelTodosResponseBody) *BatchCancelTodosResponse {
	s.Body = v
	return s
}

func (s *BatchCancelTodosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
