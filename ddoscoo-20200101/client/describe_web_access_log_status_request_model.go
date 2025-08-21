// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeWebAccessLogStatusRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DescribeWebAccessLogStatusRequest
	GetResourceGroupId() *string
}

type DescribeWebAccessLogStatusRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebAccessLogStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogStatusRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebAccessLogStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebAccessLogStatusRequest) SetDomain(v string) *DescribeWebAccessLogStatusRequest {
	s.Domain = &v
	return s
}

func (s *DescribeWebAccessLogStatusRequest) SetResourceGroupId(v string) *DescribeWebAccessLogStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebAccessLogStatusRequest) Validate() error {
	return dara.Validate(s)
}
