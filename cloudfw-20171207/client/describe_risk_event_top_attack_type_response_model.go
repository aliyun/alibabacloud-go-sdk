// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskEventTopAttackTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskEventTopAttackTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskEventTopAttackTypeResponseBody) *DescribeRiskEventTopAttackTypeResponse
	GetBody() *DescribeRiskEventTopAttackTypeResponseBody
}

type DescribeRiskEventTopAttackTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventTopAttackTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventTopAttackTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskEventTopAttackTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskEventTopAttackTypeResponse) GetBody() *DescribeRiskEventTopAttackTypeResponseBody {
	return s.Body
}

func (s *DescribeRiskEventTopAttackTypeResponse) SetHeaders(v map[string]*string) *DescribeRiskEventTopAttackTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponse) SetStatusCode(v int32) *DescribeRiskEventTopAttackTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponse) SetBody(v *DescribeRiskEventTopAttackTypeResponseBody) *DescribeRiskEventTopAttackTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskEventTopAttackTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
