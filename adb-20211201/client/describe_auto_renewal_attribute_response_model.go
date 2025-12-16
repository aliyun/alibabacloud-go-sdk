// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewalAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoRenewalAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoRenewalAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoRenewalAttributeResponseBody) *DescribeAutoRenewalAttributeResponse
	GetBody() *DescribeAutoRenewalAttributeResponseBody
}

type DescribeAutoRenewalAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoRenewalAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoRenewalAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewalAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewalAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoRenewalAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoRenewalAttributeResponse) GetBody() *DescribeAutoRenewalAttributeResponseBody {
	return s.Body
}

func (s *DescribeAutoRenewalAttributeResponse) SetHeaders(v map[string]*string) *DescribeAutoRenewalAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoRenewalAttributeResponse) SetStatusCode(v int32) *DescribeAutoRenewalAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponse) SetBody(v *DescribeAutoRenewalAttributeResponseBody) *DescribeAutoRenewalAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoRenewalAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
