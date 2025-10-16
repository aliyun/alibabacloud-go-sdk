// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventTopAttackAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskEventTopAttackAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskEventTopAttackAssetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskEventTopAttackAssetResponseBody) *DescribeRiskEventTopAttackAssetResponse
	GetBody() *DescribeRiskEventTopAttackAssetResponseBody
}

type DescribeRiskEventTopAttackAssetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventTopAttackAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventTopAttackAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventTopAttackAssetResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventTopAttackAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskEventTopAttackAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskEventTopAttackAssetResponse) GetBody() *DescribeRiskEventTopAttackAssetResponseBody {
	return s.Body
}

func (s *DescribeRiskEventTopAttackAssetResponse) SetHeaders(v map[string]*string) *DescribeRiskEventTopAttackAssetResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponse) SetStatusCode(v int32) *DescribeRiskEventTopAttackAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponse) SetBody(v *DescribeRiskEventTopAttackAssetResponseBody) *DescribeRiskEventTopAttackAssetResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskEventTopAttackAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
