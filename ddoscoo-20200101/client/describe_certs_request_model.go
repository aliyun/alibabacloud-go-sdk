// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeCertsRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DescribeCertsRequest
	GetResourceGroupId() *string
}

type DescribeCertsRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
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

func (s DescribeCertsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCertsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCertsRequest) SetDomain(v string) *DescribeCertsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeCertsRequest) SetResourceGroupId(v string) *DescribeCertsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCertsRequest) Validate() error {
	return dara.Validate(s)
}
