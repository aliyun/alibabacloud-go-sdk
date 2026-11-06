// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarFsPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarFsPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarFsPerformanceResponseBody) *DescribePolarFsPerformanceResponse
	GetBody() *DescribePolarFsPerformanceResponseBody
}

type DescribePolarFsPerformanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarFsPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarFsPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarFsPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarFsPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarFsPerformanceResponse) GetBody() *DescribePolarFsPerformanceResponseBody {
	return s.Body
}

func (s *DescribePolarFsPerformanceResponse) SetHeaders(v map[string]*string) *DescribePolarFsPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarFsPerformanceResponse) SetStatusCode(v int32) *DescribePolarFsPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarFsPerformanceResponse) SetBody(v *DescribePolarFsPerformanceResponseBody) *DescribePolarFsPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribePolarFsPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
