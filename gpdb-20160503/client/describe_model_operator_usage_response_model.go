// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelOperatorUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelOperatorUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelOperatorUsageResponseBody) *DescribeModelOperatorUsageResponse
	GetBody() *DescribeModelOperatorUsageResponseBody
}

type DescribeModelOperatorUsageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelOperatorUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelOperatorUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelOperatorUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelOperatorUsageResponse) GetBody() *DescribeModelOperatorUsageResponseBody {
	return s.Body
}

func (s *DescribeModelOperatorUsageResponse) SetHeaders(v map[string]*string) *DescribeModelOperatorUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelOperatorUsageResponse) SetStatusCode(v int32) *DescribeModelOperatorUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelOperatorUsageResponse) SetBody(v *DescribeModelOperatorUsageResponseBody) *DescribeModelOperatorUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeModelOperatorUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
