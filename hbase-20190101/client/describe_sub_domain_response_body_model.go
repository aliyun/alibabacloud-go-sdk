// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSubDomainResponseBody
	GetRequestId() *string
	SetSubDomain(v string) *DescribeSubDomainResponseBody
	GetSubDomain() *string
}

type DescribeSubDomainResponseBody struct {
	// example:
	//
	// F4208C83-B9BC-4A64-A739-8F88E98DA469
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// cn-hangzhou-h-aliyun
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeSubDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubDomainResponseBody) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeSubDomainResponseBody) SetRequestId(v string) *DescribeSubDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubDomainResponseBody) SetSubDomain(v string) *DescribeSubDomainResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeSubDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
