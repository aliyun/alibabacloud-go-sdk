// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7UsKeepaliveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeL7UsKeepaliveRequest
	GetDomain() *string
}

type DescribeL7UsKeepaliveRequest struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeL7UsKeepaliveRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7UsKeepaliveRequest) GoString() string {
	return s.String()
}

func (s *DescribeL7UsKeepaliveRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeL7UsKeepaliveRequest) SetDomain(v string) *DescribeL7UsKeepaliveRequest {
	s.Domain = &v
	return s
}

func (s *DescribeL7UsKeepaliveRequest) Validate() error {
	return dara.Validate(s)
}
