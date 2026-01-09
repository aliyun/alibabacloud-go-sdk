// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeTaskResponseBody) *ListComputeTaskResponse
	GetBody() *ListComputeTaskResponseBody
}

type ListComputeTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTaskResponse) GoString() string {
	return s.String()
}

func (s *ListComputeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeTaskResponse) GetBody() *ListComputeTaskResponseBody {
	return s.Body
}

func (s *ListComputeTaskResponse) SetHeaders(v map[string]*string) *ListComputeTaskResponse {
	s.Headers = v
	return s
}

func (s *ListComputeTaskResponse) SetStatusCode(v int32) *ListComputeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeTaskResponse) SetBody(v *ListComputeTaskResponseBody) *ListComputeTaskResponse {
	s.Body = v
	return s
}

func (s *ListComputeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
