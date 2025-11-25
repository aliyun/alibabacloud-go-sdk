// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingRiskDomainAndIpCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingRiskDomainAndIpCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingRiskDomainAndIpCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingRiskDomainAndIpCountResponseBody) *DescribeOutgoingRiskDomainAndIpCountResponse
	GetBody() *DescribeOutgoingRiskDomainAndIpCountResponseBody
}

type DescribeOutgoingRiskDomainAndIpCountResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingRiskDomainAndIpCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingRiskDomainAndIpCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingRiskDomainAndIpCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponse) GetBody() *DescribeOutgoingRiskDomainAndIpCountResponseBody {
	return s.Body
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponse) SetHeaders(v map[string]*string) *DescribeOutgoingRiskDomainAndIpCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponse) SetStatusCode(v int32) *DescribeOutgoingRiskDomainAndIpCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponse) SetBody(v *DescribeOutgoingRiskDomainAndIpCountResponseBody) *DescribeOutgoingRiskDomainAndIpCountResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
