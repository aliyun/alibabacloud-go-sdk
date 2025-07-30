// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobDedicatedClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterId(v string) *ModifyDtsJobDedicatedClusterRequest
	GetDedicatedClusterId() *string
	SetDtsJobIds(v string) *ModifyDtsJobDedicatedClusterRequest
	GetDtsJobIds() *string
	SetOwnerId(v string) *ModifyDtsJobDedicatedClusterRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyDtsJobDedicatedClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDtsJobDedicatedClusterRequest
	GetResourceGroupId() *string
}

type ModifyDtsJobDedicatedClusterRequest struct {
	// The dedicated cluster ID.
	//
	// example:
	//
	// dtsxxxx
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The DTS task IDs. The value can be a JSON array that consists of multiple DTS task IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// ["dtsxxxx01", "dtsxxx02"]
	DtsJobIds *string `json:"DtsJobIds,omitempty" xml:"DtsJobIds,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the DTS instance resides.
	//
	// > For information about the regions that support dedicated clusters, see [DTS dedicated cluster](https://help.aliyun.com/document_detail/417481.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyDtsJobDedicatedClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobDedicatedClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobDedicatedClusterRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *ModifyDtsJobDedicatedClusterRequest) GetDtsJobIds() *string {
	return s.DtsJobIds
}

func (s *ModifyDtsJobDedicatedClusterRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyDtsJobDedicatedClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobDedicatedClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobDedicatedClusterRequest) SetDedicatedClusterId(v string) *ModifyDtsJobDedicatedClusterRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterRequest) SetDtsJobIds(v string) *ModifyDtsJobDedicatedClusterRequest {
	s.DtsJobIds = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterRequest) SetOwnerId(v string) *ModifyDtsJobDedicatedClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterRequest) SetRegionId(v string) *ModifyDtsJobDedicatedClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterRequest) SetResourceGroupId(v string) *ModifyDtsJobDedicatedClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterRequest) Validate() error {
	return dara.Validate(s)
}
