// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnBlockCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeUnBlockCountRequest
	GetResourceGroupId() *string
}

type DescribeUnBlockCountRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeUnBlockCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnBlockCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnBlockCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeUnBlockCountRequest) SetResourceGroupId(v string) *DescribeUnBlockCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUnBlockCountRequest) Validate() error {
	return dara.Validate(s)
}
