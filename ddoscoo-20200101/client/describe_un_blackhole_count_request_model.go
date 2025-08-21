// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnBlackholeCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeUnBlackholeCountRequest
	GetResourceGroupId() *string
}

type DescribeUnBlackholeCountRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeUnBlackholeCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnBlackholeCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnBlackholeCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeUnBlackholeCountRequest) SetResourceGroupId(v string) *DescribeUnBlackholeCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUnBlackholeCountRequest) Validate() error {
	return dara.Validate(s)
}
