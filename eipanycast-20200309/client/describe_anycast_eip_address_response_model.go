// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnycastEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnycastEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnycastEipAddressResponseBody) *DescribeAnycastEipAddressResponse
	GetBody() *DescribeAnycastEipAddressResponseBody
}

type DescribeAnycastEipAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnycastEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnycastEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnycastEipAddressResponse) GetBody() *DescribeAnycastEipAddressResponseBody {
	return s.Body
}

func (s *DescribeAnycastEipAddressResponse) SetHeaders(v map[string]*string) *DescribeAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnycastEipAddressResponse) SetStatusCode(v int32) *DescribeAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnycastEipAddressResponse) SetBody(v *DescribeAnycastEipAddressResponseBody) *DescribeAnycastEipAddressResponse {
	s.Body = v
	return s
}

func (s *DescribeAnycastEipAddressResponse) Validate() error {
	return dara.Validate(s)
}
