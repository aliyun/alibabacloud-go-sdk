// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInstanceTypesRequest
	GetClusterId() *string
	SetClusterType(v string) *ListInstanceTypesRequest
	GetClusterType() *string
	SetDeployMode(v string) *ListInstanceTypesRequest
	GetDeployMode() *string
	SetInstanceType(v string) *ListInstanceTypesRequest
	GetInstanceType() *string
	SetIsModification(v bool) *ListInstanceTypesRequest
	GetIsModification() *bool
	SetNodeGroupId(v string) *ListInstanceTypesRequest
	GetNodeGroupId() *string
	SetNodeGroupType(v string) *ListInstanceTypesRequest
	GetNodeGroupType() *string
	SetPaymentType(v string) *ListInstanceTypesRequest
	GetPaymentType() *string
	SetRegionId(v string) *ListInstanceTypesRequest
	GetRegionId() *string
	SetReleaseVersion(v string) *ListInstanceTypesRequest
	GetReleaseVersion() *string
	SetZoneId(v string) *ListInstanceTypesRequest
	GetZoneId() *string
}

type ListInstanceTypesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATALAKE
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Specifies the deployment mode.
	//
	// example:
	//
	// HA
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.g6.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the instance type is for an instance type change. A value of true indicates an instance type change.
	//
	// example:
	//
	// false
	IsModification *bool `json:"IsModification,omitempty" xml:"IsModification,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// G-F06C4B47966A****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node group type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// The billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The EMR release version.
	//
	// example:
	//
	// Released version EMR-5.8.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListInstanceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstanceTypesRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListInstanceTypesRequest) GetDeployMode() *string {
	return s.DeployMode
}

func (s *ListInstanceTypesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstanceTypesRequest) GetIsModification() *bool {
	return s.IsModification
}

func (s *ListInstanceTypesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListInstanceTypesRequest) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *ListInstanceTypesRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstanceTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceTypesRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListInstanceTypesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListInstanceTypesRequest) SetClusterId(v string) *ListInstanceTypesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstanceTypesRequest) SetClusterType(v string) *ListInstanceTypesRequest {
	s.ClusterType = &v
	return s
}

func (s *ListInstanceTypesRequest) SetDeployMode(v string) *ListInstanceTypesRequest {
	s.DeployMode = &v
	return s
}

func (s *ListInstanceTypesRequest) SetInstanceType(v string) *ListInstanceTypesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListInstanceTypesRequest) SetIsModification(v bool) *ListInstanceTypesRequest {
	s.IsModification = &v
	return s
}

func (s *ListInstanceTypesRequest) SetNodeGroupId(v string) *ListInstanceTypesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListInstanceTypesRequest) SetNodeGroupType(v string) *ListInstanceTypesRequest {
	s.NodeGroupType = &v
	return s
}

func (s *ListInstanceTypesRequest) SetPaymentType(v string) *ListInstanceTypesRequest {
	s.PaymentType = &v
	return s
}

func (s *ListInstanceTypesRequest) SetRegionId(v string) *ListInstanceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceTypesRequest) SetReleaseVersion(v string) *ListInstanceTypesRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *ListInstanceTypesRequest) SetZoneId(v string) *ListInstanceTypesRequest {
	s.ZoneId = &v
	return s
}

func (s *ListInstanceTypesRequest) Validate() error {
	return dara.Validate(s)
}
