// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeWebAccessModeRequest
	GetDomains() []*string
}

type DescribeWebAccessModeRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for a domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeWebAccessModeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeWebAccessModeRequest) SetDomains(v []*string) *DescribeWebAccessModeRequest {
	s.Domains = v
	return s
}

func (s *DescribeWebAccessModeRequest) Validate() error {
	return dara.Validate(s)
}
