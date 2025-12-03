// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiIpControlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiIpControlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiIpControlsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiIpControlsResponseBody) *DescribeApiIpControlsResponse
	GetBody() *DescribeApiIpControlsResponseBody
}

type DescribeApiIpControlsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiIpControlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiIpControlsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiIpControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiIpControlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiIpControlsResponse) GetBody() *DescribeApiIpControlsResponseBody {
	return s.Body
}

func (s *DescribeApiIpControlsResponse) SetHeaders(v map[string]*string) *DescribeApiIpControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiIpControlsResponse) SetStatusCode(v int32) *DescribeApiIpControlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiIpControlsResponse) SetBody(v *DescribeApiIpControlsResponseBody) *DescribeApiIpControlsResponse {
	s.Body = v
	return s
}

func (s *DescribeApiIpControlsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
