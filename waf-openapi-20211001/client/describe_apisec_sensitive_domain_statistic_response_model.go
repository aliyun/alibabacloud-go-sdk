// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSensitiveDomainStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecSensitiveDomainStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecSensitiveDomainStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecSensitiveDomainStatisticResponseBody) *DescribeApisecSensitiveDomainStatisticResponse
	GetBody() *DescribeApisecSensitiveDomainStatisticResponseBody
}

type DescribeApisecSensitiveDomainStatisticResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecSensitiveDomainStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecSensitiveDomainStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSensitiveDomainStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecSensitiveDomainStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecSensitiveDomainStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecSensitiveDomainStatisticResponse) GetBody() *DescribeApisecSensitiveDomainStatisticResponseBody {
	return s.Body
}

func (s *DescribeApisecSensitiveDomainStatisticResponse) SetHeaders(v map[string]*string) *DescribeApisecSensitiveDomainStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponse) SetStatusCode(v int32) *DescribeApisecSensitiveDomainStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponse) SetBody(v *DescribeApisecSensitiveDomainStatisticResponseBody) *DescribeApisecSensitiveDomainStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecSensitiveDomainStatisticResponse) Validate() error {
	return dara.Validate(s)
}
