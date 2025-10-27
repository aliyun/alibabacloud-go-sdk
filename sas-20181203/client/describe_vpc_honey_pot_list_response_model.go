// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcHoneyPotListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcHoneyPotListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcHoneyPotListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcHoneyPotListResponseBody) *DescribeVpcHoneyPotListResponse
	GetBody() *DescribeVpcHoneyPotListResponseBody
}

type DescribeVpcHoneyPotListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcHoneyPotListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcHoneyPotListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcHoneyPotListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcHoneyPotListResponse) GetBody() *DescribeVpcHoneyPotListResponseBody {
	return s.Body
}

func (s *DescribeVpcHoneyPotListResponse) SetHeaders(v map[string]*string) *DescribeVpcHoneyPotListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcHoneyPotListResponse) SetStatusCode(v int32) *DescribeVpcHoneyPotListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponse) SetBody(v *DescribeVpcHoneyPotListResponseBody) *DescribeVpcHoneyPotListResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcHoneyPotListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
