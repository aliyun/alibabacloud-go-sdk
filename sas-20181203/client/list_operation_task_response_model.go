// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationTaskResponseBody) *ListOperationTaskResponse
	GetBody() *ListOperationTaskResponseBody
}

type ListOperationTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOperationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationTaskResponse) GetBody() *ListOperationTaskResponseBody {
	return s.Body
}

func (s *ListOperationTaskResponse) SetHeaders(v map[string]*string) *ListOperationTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOperationTaskResponse) SetStatusCode(v int32) *ListOperationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationTaskResponse) SetBody(v *ListOperationTaskResponseBody) *ListOperationTaskResponse {
	s.Body = v
	return s
}

func (s *ListOperationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
