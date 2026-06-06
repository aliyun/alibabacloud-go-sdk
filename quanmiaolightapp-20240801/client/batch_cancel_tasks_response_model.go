// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCancelTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCancelTasksResponse
	GetStatusCode() *int32
	SetBody(v *BatchCancelTasksResponseBody) *BatchCancelTasksResponse
	GetBody() *BatchCancelTasksResponseBody
}

type BatchCancelTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCancelTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCancelTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTasksResponse) GoString() string {
	return s.String()
}

func (s *BatchCancelTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCancelTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCancelTasksResponse) GetBody() *BatchCancelTasksResponseBody {
	return s.Body
}

func (s *BatchCancelTasksResponse) SetHeaders(v map[string]*string) *BatchCancelTasksResponse {
	s.Headers = v
	return s
}

func (s *BatchCancelTasksResponse) SetStatusCode(v int32) *BatchCancelTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCancelTasksResponse) SetBody(v *BatchCancelTasksResponseBody) *BatchCancelTasksResponse {
	s.Body = v
	return s
}

func (s *BatchCancelTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
