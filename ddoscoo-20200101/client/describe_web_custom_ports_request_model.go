// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCustomPortsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeWebCustomPortsRequest
	GetResourceGroupId() *string
}

type DescribeWebCustomPortsRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebCustomPortsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCustomPortsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebCustomPortsRequest) SetResourceGroupId(v string) *DescribeWebCustomPortsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebCustomPortsRequest) Validate() error {
	return dara.Validate(s)
}
