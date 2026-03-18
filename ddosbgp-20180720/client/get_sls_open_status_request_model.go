// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSlsOpenStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSlsOpenStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *GetSlsOpenStatusRequest
	GetResourceGroupId() *string
}

type GetSlsOpenStatusRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// For more information about the valid values of this parameter, see [Regions and zones](https://help.aliyun.com/document_detail/188196.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetSlsOpenStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSlsOpenStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetSlsOpenStatusRequest) SetRegionId(v string) *GetSlsOpenStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetSlsOpenStatusRequest) SetResourceGroupId(v string) *GetSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSlsOpenStatusRequest) Validate() error {
	return dara.Validate(s)
}
