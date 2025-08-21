// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsOpenStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest
	GetResourceGroupId() *string
}

type DescribeSlsOpenStatusRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsOpenStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlsOpenStatusRequest) SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) Validate() error {
	return dara.Validate(s)
}
