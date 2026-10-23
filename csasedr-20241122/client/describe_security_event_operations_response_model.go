// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityEventOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityEventOperationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityEventOperationsResponseBody) *DescribeSecurityEventOperationsResponse
	GetBody() *DescribeSecurityEventOperationsResponseBody
}

type DescribeSecurityEventOperationsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityEventOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityEventOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityEventOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityEventOperationsResponse) GetBody() *DescribeSecurityEventOperationsResponseBody {
	return s.Body
}

func (s *DescribeSecurityEventOperationsResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationsResponse) SetStatusCode(v int32) *DescribeSecurityEventOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponse) SetBody(v *DescribeSecurityEventOperationsResponseBody) *DescribeSecurityEventOperationsResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityEventOperationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
