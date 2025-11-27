// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSecurityGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCSecurityGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCSecurityGroupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCSecurityGroupListResponseBody) *DescribeRCSecurityGroupListResponse
	GetBody() *DescribeRCSecurityGroupListResponseBody
}

type DescribeRCSecurityGroupListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCSecurityGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCSecurityGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCSecurityGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCSecurityGroupListResponse) GetBody() *DescribeRCSecurityGroupListResponseBody {
	return s.Body
}

func (s *DescribeRCSecurityGroupListResponse) SetHeaders(v map[string]*string) *DescribeRCSecurityGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCSecurityGroupListResponse) SetStatusCode(v int32) *DescribeRCSecurityGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCSecurityGroupListResponse) SetBody(v *DescribeRCSecurityGroupListResponseBody) *DescribeRCSecurityGroupListResponse {
	s.Body = v
	return s
}

func (s *DescribeRCSecurityGroupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
