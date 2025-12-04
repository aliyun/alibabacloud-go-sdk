// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DescribeSlrRequest
	GetResourceGroupId() *string
}

type DescribeSlrRequest struct {
	// The ID of the resource group to which the RAM instance belongs.
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlrRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlrRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlrRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlrRequest) SetResourceGroupId(v string) *DescribeSlrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlrRequest) Validate() error {
	return dara.Validate(s)
}
