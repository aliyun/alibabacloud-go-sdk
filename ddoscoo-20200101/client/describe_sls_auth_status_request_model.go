// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAuthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest
	GetResourceGroupId() *string
}

type DescribeSlsAuthStatusRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsAuthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlsAuthStatusRequest) SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) Validate() error {
	return dara.Validate(s)
}
