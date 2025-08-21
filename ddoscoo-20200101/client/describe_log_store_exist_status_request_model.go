// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreExistStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest
	GetResourceGroupId() *string
}

type DescribeLogStoreExistStatusRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeLogStoreExistStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreExistStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLogStoreExistStatusRequest) SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) Validate() error {
	return dara.Validate(s)
}
