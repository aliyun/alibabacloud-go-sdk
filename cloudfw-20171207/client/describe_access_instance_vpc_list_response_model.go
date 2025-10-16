// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceVpcListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessInstanceVpcListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessInstanceVpcListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessInstanceVpcListResponseBody) *DescribeAccessInstanceVpcListResponse
	GetBody() *DescribeAccessInstanceVpcListResponseBody
}

type DescribeAccessInstanceVpcListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessInstanceVpcListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessInstanceVpcListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVpcListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVpcListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessInstanceVpcListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessInstanceVpcListResponse) GetBody() *DescribeAccessInstanceVpcListResponseBody {
	return s.Body
}

func (s *DescribeAccessInstanceVpcListResponse) SetHeaders(v map[string]*string) *DescribeAccessInstanceVpcListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessInstanceVpcListResponse) SetStatusCode(v int32) *DescribeAccessInstanceVpcListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessInstanceVpcListResponse) SetBody(v *DescribeAccessInstanceVpcListResponseBody) *DescribeAccessInstanceVpcListResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessInstanceVpcListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
