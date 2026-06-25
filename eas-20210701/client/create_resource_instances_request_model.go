// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewal(v bool) *CreateResourceInstancesRequest
	GetAutoRenewal() *bool
	SetChargeType(v string) *CreateResourceInstancesRequest
	GetChargeType() *string
	SetEcsInstanceCount(v int32) *CreateResourceInstancesRequest
	GetEcsInstanceCount() *int32
	SetEcsInstanceType(v string) *CreateResourceInstancesRequest
	GetEcsInstanceType() *string
	SetLabels(v map[string]*string) *CreateResourceInstancesRequest
	GetLabels() map[string]*string
	SetSystemDiskSize(v int32) *CreateResourceInstancesRequest
	GetSystemDiskSize() *int32
	SetUserData(v string) *CreateResourceInstancesRequest
	GetUserData() *string
	SetZone(v string) *CreateResourceInstancesRequest
	GetZone() *string
}

type CreateResourceInstancesRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - false (default): Auto-renewal is disabled.
	//
	// - true: Auto-renewal is enabled.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The billing method. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of new instances to create. The value must be between 1 and 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	EcsInstanceCount *int32 `json:"EcsInstanceCount,omitempty" xml:"EcsInstanceCount,omitempty"`
	// The instance type. This corresponds to an ECS instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.s6-c1m2.xlarge
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// The user-defined tags.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The size of the system disk, in GiB. The value must be between 200 and 2,000. If you do not configure this parameter, the default value is 200 GiB.
	//
	// example:
	//
	// 200
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// Deprecated
	//
	// The custom user data. This parameter is not currently used.
	//
	// example:
	//
	// x112223333
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The zone where the instance belongs.
	//
	// example:
	//
	// cn-shanghai-f
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s CreateResourceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesRequest) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *CreateResourceInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateResourceInstancesRequest) GetEcsInstanceCount() *int32 {
	return s.EcsInstanceCount
}

func (s *CreateResourceInstancesRequest) GetEcsInstanceType() *string {
	return s.EcsInstanceType
}

func (s *CreateResourceInstancesRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *CreateResourceInstancesRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateResourceInstancesRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateResourceInstancesRequest) GetZone() *string {
	return s.Zone
}

func (s *CreateResourceInstancesRequest) SetAutoRenewal(v bool) *CreateResourceInstancesRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetChargeType(v string) *CreateResourceInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetEcsInstanceCount(v int32) *CreateResourceInstancesRequest {
	s.EcsInstanceCount = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetEcsInstanceType(v string) *CreateResourceInstancesRequest {
	s.EcsInstanceType = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetLabels(v map[string]*string) *CreateResourceInstancesRequest {
	s.Labels = v
	return s
}

func (s *CreateResourceInstancesRequest) SetSystemDiskSize(v int32) *CreateResourceInstancesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetUserData(v string) *CreateResourceInstancesRequest {
	s.UserData = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetZone(v string) *CreateResourceInstancesRequest {
	s.Zone = &v
	return s
}

func (s *CreateResourceInstancesRequest) Validate() error {
	return dara.Validate(s)
}
