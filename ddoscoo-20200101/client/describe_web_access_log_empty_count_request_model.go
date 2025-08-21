// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogEmptyCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeWebAccessLogEmptyCountRequest
	GetResourceGroupId() *string
}

type DescribeWebAccessLogEmptyCountRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebAccessLogEmptyCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogEmptyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogEmptyCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebAccessLogEmptyCountRequest) SetResourceGroupId(v string) *DescribeWebAccessLogEmptyCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebAccessLogEmptyCountRequest) Validate() error {
	return dara.Validate(s)
}
