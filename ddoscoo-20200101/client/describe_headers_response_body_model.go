// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHeadersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomHeader(v *DescribeHeadersResponseBodyCustomHeader) *DescribeHeadersResponseBody
	GetCustomHeader() *DescribeHeadersResponseBodyCustomHeader
	SetRequestId(v string) *DescribeHeadersResponseBody
	GetRequestId() *string
}

type DescribeHeadersResponseBody struct {
	// The information about the custom header.
	CustomHeader *DescribeHeadersResponseBodyCustomHeader `json:"CustomHeader,omitempty" xml:"CustomHeader,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 48BC7BA5-69BE-5C31-A080-AFF2431AE48D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHeadersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHeadersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHeadersResponseBody) GetCustomHeader() *DescribeHeadersResponseBodyCustomHeader {
	return s.CustomHeader
}

func (s *DescribeHeadersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHeadersResponseBody) SetCustomHeader(v *DescribeHeadersResponseBodyCustomHeader) *DescribeHeadersResponseBody {
	s.CustomHeader = v
	return s
}

func (s *DescribeHeadersResponseBody) SetRequestId(v string) *DescribeHeadersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHeadersResponseBody) Validate() error {
	if s.CustomHeader != nil {
		if err := s.CustomHeader.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHeadersResponseBodyCustomHeader struct {
	// The domain name of the website.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The header of the response.
	//
	// example:
	//
	// {"X-Forwarded-ClientSrcPort":"","header1":"hLeLele"}
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
}

func (s DescribeHeadersResponseBodyCustomHeader) String() string {
	return dara.Prettify(s)
}

func (s DescribeHeadersResponseBodyCustomHeader) GoString() string {
	return s.String()
}

func (s *DescribeHeadersResponseBodyCustomHeader) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHeadersResponseBodyCustomHeader) GetHeaders() *string {
	return s.Headers
}

func (s *DescribeHeadersResponseBodyCustomHeader) SetDomain(v string) *DescribeHeadersResponseBodyCustomHeader {
	s.Domain = &v
	return s
}

func (s *DescribeHeadersResponseBodyCustomHeader) SetHeaders(v string) *DescribeHeadersResponseBodyCustomHeader {
	s.Headers = &v
	return s
}

func (s *DescribeHeadersResponseBodyCustomHeader) Validate() error {
	return dara.Validate(s)
}
