// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewal(v bool) *CreateResourceRequest
	GetAutoRenewal() *bool
	SetChargeType(v string) *CreateResourceRequest
	GetChargeType() *string
	SetEcsInstanceCount(v int32) *CreateResourceRequest
	GetEcsInstanceCount() *int32
	SetEcsInstanceType(v string) *CreateResourceRequest
	GetEcsInstanceType() *string
	SetLabels(v map[string]*string) *CreateResourceRequest
	GetLabels() map[string]*string
	SetResourceName(v string) *CreateResourceRequest
	GetResourceName() *string
	SetResourceType(v string) *CreateResourceRequest
	GetResourceType() *string
	SetSelfManagedResourceOptions(v *CreateResourceRequestSelfManagedResourceOptions) *CreateResourceRequest
	GetSelfManagedResourceOptions() *CreateResourceRequestSelfManagedResourceOptions
	SetSystemDiskSize(v int32) *CreateResourceRequest
	GetSystemDiskSize() *int32
	SetZone(v string) *CreateResourceRequest
	GetZone() *string
}

type CreateResourceRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: the subscription billing method.
	//
	// 	- PostPaid: the pay-as-you-go billing method.
	//
	// >  This parameter is required when the ResourceType parameter is set to Dedicated.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of ECS instances.
	//
	// >  This parameter is required when the ResourceType parameter is set to Dedicated.
	//
	// example:
	//
	// 5
	EcsInstanceCount *int32 `json:"EcsInstanceCount,omitempty" xml:"EcsInstanceCount,omitempty"`
	// The type of the Elastic Compute Service (ECS) instance.
	//
	// >  This parameter is required when the ResourceType parameter is set to Dedicated.
	//
	// example:
	//
	// ecs.c6.8xlarge
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// The labels.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// MyResource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- Dedicated: the dedicated resource group.
	//
	// 	- SelfManaged: the self-managed resource group.
	//
	// >  If you use a self-managed resource group, you must configure a whitelist.
	//
	// example:
	//
	// Dedicated
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The configurations of the self-managed resource group.
	SelfManagedResourceOptions *CreateResourceRequestSelfManagedResourceOptions `json:"SelfManagedResourceOptions,omitempty" xml:"SelfManagedResourceOptions,omitempty" type:"Struct"`
	// The size of the system disk. Unit: GiB. Valid values: 200 to 2000. Default value: 200.
	//
	// example:
	//
	// 200
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The ID of the zone in which the instance resides.
	//
	// example:
	//
	// cn-shanghai-f
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *CreateResourceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateResourceRequest) GetEcsInstanceCount() *int32 {
	return s.EcsInstanceCount
}

func (s *CreateResourceRequest) GetEcsInstanceType() *string {
	return s.EcsInstanceType
}

func (s *CreateResourceRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *CreateResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateResourceRequest) GetSelfManagedResourceOptions() *CreateResourceRequestSelfManagedResourceOptions {
	return s.SelfManagedResourceOptions
}

func (s *CreateResourceRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateResourceRequest) GetZone() *string {
	return s.Zone
}

func (s *CreateResourceRequest) SetAutoRenewal(v bool) *CreateResourceRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateResourceRequest) SetChargeType(v string) *CreateResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateResourceRequest) SetEcsInstanceCount(v int32) *CreateResourceRequest {
	s.EcsInstanceCount = &v
	return s
}

func (s *CreateResourceRequest) SetEcsInstanceType(v string) *CreateResourceRequest {
	s.EcsInstanceType = &v
	return s
}

func (s *CreateResourceRequest) SetLabels(v map[string]*string) *CreateResourceRequest {
	s.Labels = v
	return s
}

func (s *CreateResourceRequest) SetResourceName(v string) *CreateResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateResourceRequest) SetResourceType(v string) *CreateResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceRequest) SetSelfManagedResourceOptions(v *CreateResourceRequestSelfManagedResourceOptions) *CreateResourceRequest {
	s.SelfManagedResourceOptions = v
	return s
}

func (s *CreateResourceRequest) SetSystemDiskSize(v int32) *CreateResourceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateResourceRequest) SetZone(v string) *CreateResourceRequest {
	s.Zone = &v
	return s
}

func (s *CreateResourceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateResourceRequestSelfManagedResourceOptions struct {
	// The ID of the self-managed cluster.
	//
	// example:
	//
	// cf0386f250f2545689ca7fdd1cd******
	ExternalClusterId *string `json:"ExternalClusterId,omitempty" xml:"ExternalClusterId,omitempty"`
	// The tag key-value pairs of the node.
	NodeMatchLabels map[string]*string `json:"NodeMatchLabels,omitempty" xml:"NodeMatchLabels,omitempty"`
	// The tolerations for the node taint.
	NodeTolerations []*CreateResourceRequestSelfManagedResourceOptionsNodeTolerations `json:"NodeTolerations,omitempty" xml:"NodeTolerations,omitempty" type:"Repeated"`
	// The name of the RAM user to which the permissions on Elastic Algorithm Service (EAS) of Platform for AI (PAI) are granted.
	//
	// example:
	//
	// clusterrole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateResourceRequestSelfManagedResourceOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequestSelfManagedResourceOptions) GoString() string {
	return s.String()
}

func (s *CreateResourceRequestSelfManagedResourceOptions) GetExternalClusterId() *string {
	return s.ExternalClusterId
}

func (s *CreateResourceRequestSelfManagedResourceOptions) GetNodeMatchLabels() map[string]*string {
	return s.NodeMatchLabels
}

func (s *CreateResourceRequestSelfManagedResourceOptions) GetNodeTolerations() []*CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	return s.NodeTolerations
}

func (s *CreateResourceRequestSelfManagedResourceOptions) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetExternalClusterId(v string) *CreateResourceRequestSelfManagedResourceOptions {
	s.ExternalClusterId = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetNodeMatchLabels(v map[string]*string) *CreateResourceRequestSelfManagedResourceOptions {
	s.NodeMatchLabels = v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetNodeTolerations(v []*CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) *CreateResourceRequestSelfManagedResourceOptions {
	s.NodeTolerations = v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetRoleName(v string) *CreateResourceRequestSelfManagedResourceOptions {
	s.RoleName = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptions) Validate() error {
	return dara.Validate(s)
}

type CreateResourceRequestSelfManagedResourceOptionsNodeTolerations struct {
	// The effect.
	//
	// Valid values:
	//
	// 	- PreferNoSchedule
	//
	// 	- NoSchedule
	//
	// 	- NoExecute
	//
	// example:
	//
	// NoSchedule
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The key name.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The relationship between key names and key values.
	//
	// Valid values:
	//
	// 	- Equal
	//
	// 	- Exists
	//
	// example:
	//
	// Equal
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The key value.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) GoString() string {
	return s.String()
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetEffect() *string {
	return s.Effect
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetKey() *string {
	return s.Key
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetOperator() *string {
	return s.Operator
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetValue() *string {
	return s.Value
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetEffect(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Effect = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetKey(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Key = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetOperator(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Operator = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetValue(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Value = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) Validate() error {
	return dara.Validate(s)
}
