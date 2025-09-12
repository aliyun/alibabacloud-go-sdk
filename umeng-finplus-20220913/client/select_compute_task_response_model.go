// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectComputeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectComputeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectComputeTaskResponse
	GetStatusCode() *int32
	SetBody(v *SelectComputeTaskResponseBody) *SelectComputeTaskResponse
	GetBody() *SelectComputeTaskResponseBody
}

type SelectComputeTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectComputeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectComputeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTaskResponse) GoString() string {
	return s.String()
}

func (s *SelectComputeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectComputeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectComputeTaskResponse) GetBody() *SelectComputeTaskResponseBody {
	return s.Body
}

func (s *SelectComputeTaskResponse) SetHeaders(v map[string]*string) *SelectComputeTaskResponse {
	s.Headers = v
	return s
}

func (s *SelectComputeTaskResponse) SetStatusCode(v int32) *SelectComputeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectComputeTaskResponse) SetBody(v *SelectComputeTaskResponseBody) *SelectComputeTaskResponse {
	s.Body = v
	return s
}

func (s *SelectComputeTaskResponse) Validate() error {
	return dara.Validate(s)
}
