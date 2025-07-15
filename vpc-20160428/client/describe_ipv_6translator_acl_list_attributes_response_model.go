// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorAclListAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIPv6TranslatorAclListAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIPv6TranslatorAclListAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIPv6TranslatorAclListAttributesResponseBody) *DescribeIPv6TranslatorAclListAttributesResponse
	GetBody() *DescribeIPv6TranslatorAclListAttributesResponseBody
}

type DescribeIPv6TranslatorAclListAttributesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIPv6TranslatorAclListAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIPv6TranslatorAclListAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIPv6TranslatorAclListAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIPv6TranslatorAclListAttributesResponse) GetBody() *DescribeIPv6TranslatorAclListAttributesResponseBody {
	return s.Body
}

func (s *DescribeIPv6TranslatorAclListAttributesResponse) SetHeaders(v map[string]*string) *DescribeIPv6TranslatorAclListAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponse) SetStatusCode(v int32) *DescribeIPv6TranslatorAclListAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponse) SetBody(v *DescribeIPv6TranslatorAclListAttributesResponseBody) *DescribeIPv6TranslatorAclListAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeIPv6TranslatorAclListAttributesResponse) Validate() error {
	return dara.Validate(s)
}
