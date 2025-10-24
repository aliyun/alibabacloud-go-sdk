// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventDomainStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecEventDomainStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecEventDomainStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecEventDomainStatisticResponseBody) *DescribeApisecEventDomainStatisticResponse
	GetBody() *DescribeApisecEventDomainStatisticResponseBody
}

type DescribeApisecEventDomainStatisticResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecEventDomainStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecEventDomainStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventDomainStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventDomainStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecEventDomainStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecEventDomainStatisticResponse) GetBody() *DescribeApisecEventDomainStatisticResponseBody {
	return s.Body
}

func (s *DescribeApisecEventDomainStatisticResponse) SetHeaders(v map[string]*string) *DescribeApisecEventDomainStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponse) SetStatusCode(v int32) *DescribeApisecEventDomainStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponse) SetBody(v *DescribeApisecEventDomainStatisticResponseBody) *DescribeApisecEventDomainStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecEventDomainStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
