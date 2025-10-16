// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskSecurityGroupDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskSecurityGroupDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskSecurityGroupDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskSecurityGroupDetailResponseBody) *DescribeRiskSecurityGroupDetailResponse
	GetBody() *DescribeRiskSecurityGroupDetailResponseBody
}

type DescribeRiskSecurityGroupDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskSecurityGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskSecurityGroupDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskSecurityGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskSecurityGroupDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskSecurityGroupDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskSecurityGroupDetailResponse) GetBody() *DescribeRiskSecurityGroupDetailResponseBody {
	return s.Body
}

func (s *DescribeRiskSecurityGroupDetailResponse) SetHeaders(v map[string]*string) *DescribeRiskSecurityGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponse) SetStatusCode(v int32) *DescribeRiskSecurityGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponse) SetBody(v *DescribeRiskSecurityGroupDetailResponseBody) *DescribeRiskSecurityGroupDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskSecurityGroupDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
