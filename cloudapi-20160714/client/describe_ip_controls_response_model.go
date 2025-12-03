// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpControlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpControlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpControlsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpControlsResponseBody) *DescribeIpControlsResponse
	GetBody() *DescribeIpControlsResponseBody
}

type DescribeIpControlsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpControlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpControlsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpControlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpControlsResponse) GetBody() *DescribeIpControlsResponseBody {
	return s.Body
}

func (s *DescribeIpControlsResponse) SetHeaders(v map[string]*string) *DescribeIpControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpControlsResponse) SetStatusCode(v int32) *DescribeIpControlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpControlsResponse) SetBody(v *DescribeIpControlsResponseBody) *DescribeIpControlsResponse {
	s.Body = v
	return s
}

func (s *DescribeIpControlsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
