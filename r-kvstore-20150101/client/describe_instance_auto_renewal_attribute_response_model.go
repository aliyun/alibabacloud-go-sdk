// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewalAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAutoRenewalAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAutoRenewalAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAutoRenewalAttributeResponseBody) *DescribeInstanceAutoRenewalAttributeResponse
	GetBody() *DescribeInstanceAutoRenewalAttributeResponseBody
}

type DescribeInstanceAutoRenewalAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAutoRenewalAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) GetBody() *DescribeInstanceAutoRenewalAttributeResponseBody {
	return s.Body
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAutoRenewalAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) SetStatusCode(v int32) *DescribeInstanceAutoRenewalAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) SetBody(v *DescribeInstanceAutoRenewalAttributeResponseBody) *DescribeInstanceAutoRenewalAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeResponse) Validate() error {
	return dara.Validate(s)
}
