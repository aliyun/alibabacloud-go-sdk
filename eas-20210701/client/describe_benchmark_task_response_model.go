// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBenchmarkTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBenchmarkTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBenchmarkTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBenchmarkTaskResponseBody) *DescribeBenchmarkTaskResponse
	GetBody() *DescribeBenchmarkTaskResponseBody
}

type DescribeBenchmarkTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBenchmarkTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBenchmarkTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBenchmarkTaskResponse) GetBody() *DescribeBenchmarkTaskResponseBody {
	return s.Body
}

func (s *DescribeBenchmarkTaskResponse) SetHeaders(v map[string]*string) *DescribeBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeBenchmarkTaskResponse) SetStatusCode(v int32) *DescribeBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBenchmarkTaskResponse) SetBody(v *DescribeBenchmarkTaskResponseBody) *DescribeBenchmarkTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeBenchmarkTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
