// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebInstanceRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeWebInstanceRelationsRequest
	GetDomains() []*string
	SetResourceGroupId(v string) *DescribeWebInstanceRelationsRequest
	GetResourceGroupId() *string
}

type DescribeWebInstanceRelationsRequest struct {
	// The domain names of the website.
	//
	// >  A forwarding rule must be configured for the domain names. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
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

func (s DescribeWebInstanceRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebInstanceRelationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeWebInstanceRelationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebInstanceRelationsRequest) SetDomains(v []*string) *DescribeWebInstanceRelationsRequest {
	s.Domains = v
	return s
}

func (s *DescribeWebInstanceRelationsRequest) SetResourceGroupId(v string) *DescribeWebInstanceRelationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebInstanceRelationsRequest) Validate() error {
	return dara.Validate(s)
}
