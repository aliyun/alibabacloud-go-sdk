// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVSwitchsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsVSwitchsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsVSwitchsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsVSwitchsResponseBody) *DescribeRdsVSwitchsResponse
	GetBody() *DescribeRdsVSwitchsResponseBody
}

type DescribeRdsVSwitchsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsVSwitchsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsVSwitchsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVSwitchsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsVSwitchsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsVSwitchsResponse) GetBody() *DescribeRdsVSwitchsResponseBody {
	return s.Body
}

func (s *DescribeRdsVSwitchsResponse) SetHeaders(v map[string]*string) *DescribeRdsVSwitchsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetStatusCode(v int32) *DescribeRdsVSwitchsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetBody(v *DescribeRdsVSwitchsResponseBody) *DescribeRdsVSwitchsResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsVSwitchsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
