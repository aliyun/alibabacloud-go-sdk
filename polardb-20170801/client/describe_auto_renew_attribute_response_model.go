// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoRenewAttributeResponseBody) *DescribeAutoRenewAttributeResponse
	GetBody() *DescribeAutoRenewAttributeResponseBody
}

type DescribeAutoRenewAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoRenewAttributeResponse) GetBody() *DescribeAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *DescribeAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetBody(v *DescribeAutoRenewAttributeResponseBody) *DescribeAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) Validate() error {
	return dara.Validate(s)
}
