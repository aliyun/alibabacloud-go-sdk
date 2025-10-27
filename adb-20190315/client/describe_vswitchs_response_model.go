// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVSwitchsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVSwitchsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVSwitchsResponseBody) *DescribeVSwitchsResponse
	GetBody() *DescribeVSwitchsResponseBody
}

type DescribeVSwitchsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVSwitchsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVSwitchsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVSwitchsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVSwitchsResponse) GetBody() *DescribeVSwitchsResponseBody {
	return s.Body
}

func (s *DescribeVSwitchsResponse) SetHeaders(v map[string]*string) *DescribeVSwitchsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchsResponse) SetStatusCode(v int32) *DescribeVSwitchsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVSwitchsResponse) SetBody(v *DescribeVSwitchsResponseBody) *DescribeVSwitchsResponse {
	s.Body = v
	return s
}

func (s *DescribeVSwitchsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
