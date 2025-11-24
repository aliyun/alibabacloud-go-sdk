// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterPodsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGuestClusterPodsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGuestClusterPodsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGuestClusterPodsResponseBody) *DescribeGuestClusterPodsResponse
	GetBody() *DescribeGuestClusterPodsResponseBody
}

type DescribeGuestClusterPodsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGuestClusterPodsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGuestClusterPodsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterPodsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterPodsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGuestClusterPodsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGuestClusterPodsResponse) GetBody() *DescribeGuestClusterPodsResponseBody {
	return s.Body
}

func (s *DescribeGuestClusterPodsResponse) SetHeaders(v map[string]*string) *DescribeGuestClusterPodsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGuestClusterPodsResponse) SetStatusCode(v int32) *DescribeGuestClusterPodsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGuestClusterPodsResponse) SetBody(v *DescribeGuestClusterPodsResponseBody) *DescribeGuestClusterPodsResponse {
	s.Body = v
	return s
}

func (s *DescribeGuestClusterPodsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
