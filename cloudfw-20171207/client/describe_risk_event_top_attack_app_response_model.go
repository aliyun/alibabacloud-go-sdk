// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskEventTopAttackAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskEventTopAttackAppResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskEventTopAttackAppResponseBody) *DescribeRiskEventTopAttackAppResponse
	GetBody() *DescribeRiskEventTopAttackAppResponseBody
}

type DescribeRiskEventTopAttackAppResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventTopAttackAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventTopAttackAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskEventTopAttackAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskEventTopAttackAppResponse) GetBody() *DescribeRiskEventTopAttackAppResponseBody {
	return s.Body
}

func (s *DescribeRiskEventTopAttackAppResponse) SetHeaders(v map[string]*string) *DescribeRiskEventTopAttackAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponse) SetStatusCode(v int32) *DescribeRiskEventTopAttackAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponse) SetBody(v *DescribeRiskEventTopAttackAppResponseBody) *DescribeRiskEventTopAttackAppResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskEventTopAttackAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
