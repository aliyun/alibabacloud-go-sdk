// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCcProtectSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeWebCcProtectSwitchRequest
	GetDomains() []*string
	SetResourceGroupId(v string) *DescribeWebCcProtectSwitchRequest
	GetResourceGroupId() *string
}

type DescribeWebCcProtectSwitchRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebCcProtectSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCcProtectSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeWebCcProtectSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebCcProtectSwitchRequest) SetDomains(v []*string) *DescribeWebCcProtectSwitchRequest {
	s.Domains = v
	return s
}

func (s *DescribeWebCcProtectSwitchRequest) SetResourceGroupId(v string) *DescribeWebCcProtectSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebCcProtectSwitchRequest) Validate() error {
	return dara.Validate(s)
}
