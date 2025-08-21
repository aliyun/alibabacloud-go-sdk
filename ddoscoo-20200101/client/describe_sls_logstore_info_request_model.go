// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogstoreInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest
	GetResourceGroupId() *string
}

type DescribeSlsLogstoreInfoRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsLogstoreInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogstoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlsLogstoreInfoRequest) SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) Validate() error {
	return dara.Validate(s)
}
