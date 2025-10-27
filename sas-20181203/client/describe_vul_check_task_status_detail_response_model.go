// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulCheckTaskStatusDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulCheckTaskStatusDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulCheckTaskStatusDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulCheckTaskStatusDetailResponseBody) *DescribeVulCheckTaskStatusDetailResponse
	GetBody() *DescribeVulCheckTaskStatusDetailResponseBody
}

type DescribeVulCheckTaskStatusDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulCheckTaskStatusDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulCheckTaskStatusDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulCheckTaskStatusDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulCheckTaskStatusDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulCheckTaskStatusDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulCheckTaskStatusDetailResponse) GetBody() *DescribeVulCheckTaskStatusDetailResponseBody {
	return s.Body
}

func (s *DescribeVulCheckTaskStatusDetailResponse) SetHeaders(v map[string]*string) *DescribeVulCheckTaskStatusDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponse) SetStatusCode(v int32) *DescribeVulCheckTaskStatusDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponse) SetBody(v *DescribeVulCheckTaskStatusDetailResponseBody) *DescribeVulCheckTaskStatusDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
