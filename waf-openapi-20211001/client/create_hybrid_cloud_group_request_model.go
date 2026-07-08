// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackSourceMark(v string) *CreateHybridCloudGroupRequest
	GetBackSourceMark() *string
	SetClusterId(v int64) *CreateHybridCloudGroupRequest
	GetClusterId() *int64
	SetGroupName(v string) *CreateHybridCloudGroupRequest
	GetGroupName() *string
	SetGroupType(v string) *CreateHybridCloudGroupRequest
	GetGroupType() *string
	SetInstanceId(v string) *CreateHybridCloudGroupRequest
	GetInstanceId() *string
	SetLoadBalanceIp(v string) *CreateHybridCloudGroupRequest
	GetLoadBalanceIp() *string
	SetLocationCode(v string) *CreateHybridCloudGroupRequest
	GetLocationCode() *string
	SetRegionId(v string) *CreateHybridCloudGroupRequest
	GetRegionId() *string
	SetRemark(v string) *CreateHybridCloudGroupRequest
	GetRemark() *string
	SetResourceManagerResourceGroupId(v string) *CreateHybridCloudGroupRequest
	GetResourceManagerResourceGroupId() *string
}

type CreateHybridCloudGroupRequest struct {
	// The back-to-origin mark of the node group. The value is in the format of Carrier-Continent-City, which is used to identify the origin of back-to-origin requests.
	//
	// example:
	//
	// aliyun-asiapacific-beijing
	BackSourceMark *string `json:"BackSourceMark,omitempty" xml:"BackSourceMark,omitempty"`
	// The ID of the Hybrid Cloud WAF cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// - **protect**: a protection node group that processes traffic filtering.
	//
	// - **control**: a control node group that manages cluster configurations.
	//
	// - **storage**: a storage node group that stores logs and data.
	//
	// - **controlStorage**: a node group that serves as both control and storage.
	//
	// This parameter is required.
	//
	// example:
	//
	// protect
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-n6w***x52m
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the load balancer that is associated with the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.3.3.3
	LoadBalanceIp *string `json:"LoadBalanceIp,omitempty" xml:"LoadBalanceIp,omitempty"`
	// The location code of the region where the node group resides. The value is in the format of Carrier-Continent-City.
	//
	// example:
	//
	// 0-410-0
	LocationCode *string `json:"LocationCode,omitempty" xml:"LocationCode,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks on the node group. You can use this parameter to add a brief description for the node group.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s CreateHybridCloudGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudGroupRequest) GetBackSourceMark() *string {
	return s.BackSourceMark
}

func (s *CreateHybridCloudGroupRequest) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *CreateHybridCloudGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateHybridCloudGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateHybridCloudGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHybridCloudGroupRequest) GetLoadBalanceIp() *string {
	return s.LoadBalanceIp
}

func (s *CreateHybridCloudGroupRequest) GetLocationCode() *string {
	return s.LocationCode
}

func (s *CreateHybridCloudGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHybridCloudGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateHybridCloudGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateHybridCloudGroupRequest) SetBackSourceMark(v string) *CreateHybridCloudGroupRequest {
	s.BackSourceMark = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetClusterId(v int64) *CreateHybridCloudGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetGroupName(v string) *CreateHybridCloudGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetGroupType(v string) *CreateHybridCloudGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetInstanceId(v string) *CreateHybridCloudGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetLoadBalanceIp(v string) *CreateHybridCloudGroupRequest {
	s.LoadBalanceIp = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetLocationCode(v string) *CreateHybridCloudGroupRequest {
	s.LocationCode = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetRegionId(v string) *CreateHybridCloudGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetRemark(v string) *CreateHybridCloudGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) SetResourceManagerResourceGroupId(v string) *CreateHybridCloudGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateHybridCloudGroupRequest) Validate() error {
	return dara.Validate(s)
}
