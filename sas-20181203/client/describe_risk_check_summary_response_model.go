// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskCheckSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskCheckSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskCheckSummaryResponseBody) *DescribeRiskCheckSummaryResponse
	GetBody() *DescribeRiskCheckSummaryResponseBody
}

type DescribeRiskCheckSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskCheckSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskCheckSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskCheckSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskCheckSummaryResponse) GetBody() *DescribeRiskCheckSummaryResponseBody {
	return s.Body
}

func (s *DescribeRiskCheckSummaryResponse) SetHeaders(v map[string]*string) *DescribeRiskCheckSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskCheckSummaryResponse) SetStatusCode(v int32) *DescribeRiskCheckSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponse) SetBody(v *DescribeRiskCheckSummaryResponseBody) *DescribeRiskCheckSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskCheckSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
