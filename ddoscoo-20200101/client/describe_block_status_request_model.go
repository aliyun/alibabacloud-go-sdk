// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlockStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeBlockStatusRequest
	GetInstanceIds() []*string
	SetResourceGroupId(v string) *DescribeBlockStatusRequest
	GetResourceGroupId() *string
}

type DescribeBlockStatusRequest struct {
	// An array consisting of information about the IDs of the instances that you want to query.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeBlockStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeBlockStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBlockStatusRequest) SetInstanceIds(v []*string) *DescribeBlockStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeBlockStatusRequest) SetResourceGroupId(v string) *DescribeBlockStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBlockStatusRequest) Validate() error {
	return dara.Validate(s)
}
