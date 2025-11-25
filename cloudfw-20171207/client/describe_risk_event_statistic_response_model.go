// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskEventStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskEventStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskEventStatisticResponseBody) *DescribeRiskEventStatisticResponse
	GetBody() *DescribeRiskEventStatisticResponseBody
}

type DescribeRiskEventStatisticResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskEventStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskEventStatisticResponse) GetBody() *DescribeRiskEventStatisticResponseBody {
	return s.Body
}

func (s *DescribeRiskEventStatisticResponse) SetHeaders(v map[string]*string) *DescribeRiskEventStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventStatisticResponse) SetStatusCode(v int32) *DescribeRiskEventStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventStatisticResponse) SetBody(v *DescribeRiskEventStatisticResponseBody) *DescribeRiskEventStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskEventStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
