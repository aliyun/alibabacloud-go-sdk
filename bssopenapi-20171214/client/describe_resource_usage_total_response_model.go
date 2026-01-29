// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceUsageTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceUsageTotalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceUsageTotalResponseBody) *DescribeResourceUsageTotalResponse
	GetBody() *DescribeResourceUsageTotalResponseBody
}

type DescribeResourceUsageTotalResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceUsageTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceUsageTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceUsageTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceUsageTotalResponse) GetBody() *DescribeResourceUsageTotalResponseBody {
	return s.Body
}

func (s *DescribeResourceUsageTotalResponse) SetHeaders(v map[string]*string) *DescribeResourceUsageTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceUsageTotalResponse) SetStatusCode(v int32) *DescribeResourceUsageTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceUsageTotalResponse) SetBody(v *DescribeResourceUsageTotalResponseBody) *DescribeResourceUsageTotalResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceUsageTotalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
