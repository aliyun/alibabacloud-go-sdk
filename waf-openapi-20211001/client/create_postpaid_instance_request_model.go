// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostpaidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreatePostpaidInstanceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreatePostpaidInstanceRequest
	GetResourceManagerResourceGroupId() *string
}

type CreatePostpaidInstanceRequest struct {
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm4gh****wela
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s CreatePostpaidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePostpaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePostpaidInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePostpaidInstanceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreatePostpaidInstanceRequest) SetRegionId(v string) *CreatePostpaidInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostpaidInstanceRequest) SetResourceManagerResourceGroupId(v string) *CreatePostpaidInstanceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreatePostpaidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
