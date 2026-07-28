// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationUsageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationUsageResponseBody) *DescribeApplicationUsageResponse
	GetBody() *DescribeApplicationUsageResponseBody
}

type DescribeApplicationUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationUsageResponse) GetBody() *DescribeApplicationUsageResponseBody {
	return s.Body
}

func (s *DescribeApplicationUsageResponse) SetHeaders(v map[string]*string) *DescribeApplicationUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationUsageResponse) SetStatusCode(v int32) *DescribeApplicationUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationUsageResponse) SetBody(v *DescribeApplicationUsageResponseBody) *DescribeApplicationUsageResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
