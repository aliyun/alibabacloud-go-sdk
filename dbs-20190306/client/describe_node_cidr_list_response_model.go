// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeCidrListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeCidrListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeCidrListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeCidrListResponseBody) *DescribeNodeCidrListResponse
	GetBody() *DescribeNodeCidrListResponseBody
}

type DescribeNodeCidrListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeCidrListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeCidrListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeCidrListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeCidrListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeCidrListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeCidrListResponse) GetBody() *DescribeNodeCidrListResponseBody {
	return s.Body
}

func (s *DescribeNodeCidrListResponse) SetHeaders(v map[string]*string) *DescribeNodeCidrListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeCidrListResponse) SetStatusCode(v int32) *DescribeNodeCidrListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeCidrListResponse) SetBody(v *DescribeNodeCidrListResponseBody) *DescribeNodeCidrListResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeCidrListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
