// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecAbnormalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalIds(v []*string) *ModifyApisecAbnormalsRequest
	GetAbnormalIds() []*string
	SetClusterId(v string) *ModifyApisecAbnormalsRequest
	GetClusterId() *string
	SetInstanceId(v string) *ModifyApisecAbnormalsRequest
	GetInstanceId() *string
	SetNote(v string) *ModifyApisecAbnormalsRequest
	GetNote() *string
	SetRegionId(v string) *ModifyApisecAbnormalsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyApisecAbnormalsRequest
	GetResourceManagerResourceGroupId() *string
	SetUserStatus(v string) *ModifyApisecAbnormalsRequest
	GetUserStatus() *string
}

type ModifyApisecAbnormalsRequest struct {
	// The risk IDs.
	//
	// This parameter is required.
	AbnormalIds []*string `json:"AbnormalIds,omitempty" xml:"AbnormalIds,omitempty" type:"Repeated"`
	// The ID of the hybrid cloud cluster.
	//
	// >  This parameter is available only in hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The description.
	//
	// example:
	//
	// already fixed.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
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
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The risk status. Valid values:
	//
	// 	- **toBeConfirmed**
	//
	// 	- **confirmed**
	//
	// 	- **toBeFixed**
	//
	// 	- **fixed**
	//
	// 	- **ignored**
	//
	// This parameter is required.
	//
	// example:
	//
	// fixed
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s ModifyApisecAbnormalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecAbnormalsRequest) GoString() string {
	return s.String()
}

func (s *ModifyApisecAbnormalsRequest) GetAbnormalIds() []*string {
	return s.AbnormalIds
}

func (s *ModifyApisecAbnormalsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyApisecAbnormalsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApisecAbnormalsRequest) GetNote() *string {
	return s.Note
}

func (s *ModifyApisecAbnormalsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApisecAbnormalsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyApisecAbnormalsRequest) GetUserStatus() *string {
	return s.UserStatus
}

func (s *ModifyApisecAbnormalsRequest) SetAbnormalIds(v []*string) *ModifyApisecAbnormalsRequest {
	s.AbnormalIds = v
	return s
}

func (s *ModifyApisecAbnormalsRequest) SetClusterId(v string) *ModifyApisecAbnormalsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyApisecAbnormalsRequest) SetInstanceId(v string) *ModifyApisecAbnormalsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyApisecAbnormalsRequest) SetNote(v string) *ModifyApisecAbnormalsRequest {
	s.Note = &v
	return s
}

func (s *ModifyApisecAbnormalsRequest) SetRegionId(v string) *ModifyApisecAbnormalsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApisecAbnormalsRequest) SetResourceManagerResourceGroupId(v string) *ModifyApisecAbnormalsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyApisecAbnormalsRequest) SetUserStatus(v string) *ModifyApisecAbnormalsRequest {
	s.UserStatus = &v
	return s
}

func (s *ModifyApisecAbnormalsRequest) Validate() error {
	return dara.Validate(s)
}
