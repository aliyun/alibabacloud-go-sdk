// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAbnormalDomainStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecAbnormalDomainStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecAbnormalDomainStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecAbnormalDomainStatisticResponseBody) *DescribeApisecAbnormalDomainStatisticResponse
	GetBody() *DescribeApisecAbnormalDomainStatisticResponseBody
}

type DescribeApisecAbnormalDomainStatisticResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecAbnormalDomainStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecAbnormalDomainStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalDomainStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalDomainStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecAbnormalDomainStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecAbnormalDomainStatisticResponse) GetBody() *DescribeApisecAbnormalDomainStatisticResponseBody {
	return s.Body
}

func (s *DescribeApisecAbnormalDomainStatisticResponse) SetHeaders(v map[string]*string) *DescribeApisecAbnormalDomainStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponse) SetStatusCode(v int32) *DescribeApisecAbnormalDomainStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponse) SetBody(v *DescribeApisecAbnormalDomainStatisticResponseBody) *DescribeApisecAbnormalDomainStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecAbnormalDomainStatisticResponse) Validate() error {
	return dara.Validate(s)
}
