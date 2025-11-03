// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorAclListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIPv6TranslatorAclListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIPv6TranslatorAclListsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIPv6TranslatorAclListsResponseBody) *DescribeIPv6TranslatorAclListsResponse
	GetBody() *DescribeIPv6TranslatorAclListsResponseBody
}

type DescribeIPv6TranslatorAclListsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIPv6TranslatorAclListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIPv6TranslatorAclListsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorAclListsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorAclListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIPv6TranslatorAclListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIPv6TranslatorAclListsResponse) GetBody() *DescribeIPv6TranslatorAclListsResponseBody {
	return s.Body
}

func (s *DescribeIPv6TranslatorAclListsResponse) SetHeaders(v map[string]*string) *DescribeIPv6TranslatorAclListsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponse) SetStatusCode(v int32) *DescribeIPv6TranslatorAclListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponse) SetBody(v *DescribeIPv6TranslatorAclListsResponseBody) *DescribeIPv6TranslatorAclListsResponse {
	s.Body = v
	return s
}

func (s *DescribeIPv6TranslatorAclListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
