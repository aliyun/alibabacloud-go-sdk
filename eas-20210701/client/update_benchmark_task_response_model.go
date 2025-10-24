// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBenchmarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBenchmarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBenchmarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBenchmarkTaskResponseBody) *UpdateBenchmarkTaskResponse
	GetBody() *UpdateBenchmarkTaskResponseBody
}

type UpdateBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBenchmarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBenchmarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBenchmarkTaskResponse) GetBody() *UpdateBenchmarkTaskResponseBody {
	return s.Body
}

func (s *UpdateBenchmarkTaskResponse) SetHeaders(v map[string]*string) *UpdateBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateBenchmarkTaskResponse) SetStatusCode(v int32) *UpdateBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBenchmarkTaskResponse) SetBody(v *UpdateBenchmarkTaskResponseBody) *UpdateBenchmarkTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateBenchmarkTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
