// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVerifyContentRequest
	GetDomainName() *string
}

type DescribeVerifyContentRequest struct {
	// The domain name of which the ownership was verified. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeVerifyContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVerifyContentRequest) SetDomainName(v string) *DescribeVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVerifyContentRequest) Validate() error {
	return dara.Validate(s)
}
