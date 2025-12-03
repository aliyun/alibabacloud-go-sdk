// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcAccessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcAccessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcAccessesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcAccessesResponseBody) *DescribeVpcAccessesResponse
	GetBody() *DescribeVpcAccessesResponseBody
}

type DescribeVpcAccessesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcAccessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcAccessesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcAccessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcAccessesResponse) GetBody() *DescribeVpcAccessesResponseBody {
	return s.Body
}

func (s *DescribeVpcAccessesResponse) SetHeaders(v map[string]*string) *DescribeVpcAccessesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcAccessesResponse) SetStatusCode(v int32) *DescribeVpcAccessesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcAccessesResponse) SetBody(v *DescribeVpcAccessesResponseBody) *DescribeVpcAccessesResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcAccessesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
