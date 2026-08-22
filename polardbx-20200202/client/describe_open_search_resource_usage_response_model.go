// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchResourceUsageResponseBody) *DescribeOpenSearchResourceUsageResponse
	GetBody() *DescribeOpenSearchResourceUsageResponseBody
}

type DescribeOpenSearchResourceUsageResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchResourceUsageResponse) GetBody() *DescribeOpenSearchResourceUsageResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchResourceUsageResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponse) SetStatusCode(v int32) *DescribeOpenSearchResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponse) SetBody(v *DescribeOpenSearchResourceUsageResponseBody) *DescribeOpenSearchResourceUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
