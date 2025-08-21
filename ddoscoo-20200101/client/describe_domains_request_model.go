// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeDomainsRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribeDomainsRequest
	GetResourceGroupId() *string
}

type DescribeDomainsRequest struct {
	// The ID of the instance that you want to query.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainsRequest) SetInstanceIds(v []*string) *DescribeDomainsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) Validate() error {
	return dara.Validate(s)
}
