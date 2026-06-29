// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVCUInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *CreateVCUInstanceRequest
	GetAliasName() *string
	SetAutoRenewPeriodInMonth(v int64) *CreateVCUInstanceRequest
	GetAutoRenewPeriodInMonth() *int64
	SetClusterType(v string) *CreateVCUInstanceRequest
	GetClusterType() *string
	SetDryRun(v bool) *CreateVCUInstanceRequest
	GetDryRun() *bool
	SetEnableAutoRenew(v bool) *CreateVCUInstanceRequest
	GetEnableAutoRenew() *bool
	SetEnableElasticVCU(v bool) *CreateVCUInstanceRequest
	GetEnableElasticVCU() *bool
	SetInstanceDescription(v string) *CreateVCUInstanceRequest
	GetInstanceDescription() *string
	SetPeriodInMonth(v int64) *CreateVCUInstanceRequest
	GetPeriodInMonth() *int64
	SetResourceGroupId(v string) *CreateVCUInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateVCUInstanceRequestTags) *CreateVCUInstanceRequest
	GetTags() []*CreateVCUInstanceRequestTags
	SetVCU(v int64) *CreateVCUInstanceRequest
	GetVCU() *int64
}

type CreateVCUInstanceRequest struct {
	// The alias of the instance.
	//
	// example:
	//
	// test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The auto-renewal cycle. This parameter is required if auto-renewal is enabled.
	//
	// example:
	//
	// 1
	AutoRenewPeriodInMonth *int64 `json:"AutoRenewPeriodInMonth,omitempty" xml:"AutoRenewPeriodInMonth,omitempty"`
	// The cluster type.
	//
	// This parameter is required.
	//
	// example:
	//
	// SSD vs HYBRID
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Specifies whether to perform a dry run. If you perform a dry run, no instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable auto-renewal.
	EnableAutoRenew *bool `json:"EnableAutoRenew,omitempty" xml:"EnableAutoRenew,omitempty"`
	// Specifies whether to enable instance elasticity. If you enable instance elasticity, the peak VCU usage can exceed the purchased VCU amount, but additional fees are incurred.
	EnableElasticVCU *bool `json:"EnableElasticVCU,omitempty" xml:"EnableElasticVCU,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// testvcu
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The subscription duration. Unit: month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PeriodInMonth *int64 `json:"PeriodInMonth,omitempty" xml:"PeriodInMonth,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxh4em5jnbcd
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*CreateVCUInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance type: the number of VCUs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	VCU *int64 `json:"VCU,omitempty" xml:"VCU,omitempty"`
}

func (s CreateVCUInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVCUInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateVCUInstanceRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateVCUInstanceRequest) GetAutoRenewPeriodInMonth() *int64 {
	return s.AutoRenewPeriodInMonth
}

func (s *CreateVCUInstanceRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateVCUInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVCUInstanceRequest) GetEnableAutoRenew() *bool {
	return s.EnableAutoRenew
}

func (s *CreateVCUInstanceRequest) GetEnableElasticVCU() *bool {
	return s.EnableElasticVCU
}

func (s *CreateVCUInstanceRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateVCUInstanceRequest) GetPeriodInMonth() *int64 {
	return s.PeriodInMonth
}

func (s *CreateVCUInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVCUInstanceRequest) GetTags() []*CreateVCUInstanceRequestTags {
	return s.Tags
}

func (s *CreateVCUInstanceRequest) GetVCU() *int64 {
	return s.VCU
}

func (s *CreateVCUInstanceRequest) SetAliasName(v string) *CreateVCUInstanceRequest {
	s.AliasName = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetAutoRenewPeriodInMonth(v int64) *CreateVCUInstanceRequest {
	s.AutoRenewPeriodInMonth = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetClusterType(v string) *CreateVCUInstanceRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetDryRun(v bool) *CreateVCUInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetEnableAutoRenew(v bool) *CreateVCUInstanceRequest {
	s.EnableAutoRenew = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetEnableElasticVCU(v bool) *CreateVCUInstanceRequest {
	s.EnableElasticVCU = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetInstanceDescription(v string) *CreateVCUInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetPeriodInMonth(v int64) *CreateVCUInstanceRequest {
	s.PeriodInMonth = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetResourceGroupId(v string) *CreateVCUInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVCUInstanceRequest) SetTags(v []*CreateVCUInstanceRequestTags) *CreateVCUInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateVCUInstanceRequest) SetVCU(v int64) *CreateVCUInstanceRequest {
	s.VCU = &v
	return s
}

func (s *CreateVCUInstanceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVCUInstanceRequestTags struct {
	// The key of the tag. The key can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// created
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The value can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0woauavextilfqr61
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVCUInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateVCUInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateVCUInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateVCUInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateVCUInstanceRequestTags) SetKey(v string) *CreateVCUInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateVCUInstanceRequestTags) SetValue(v string) *CreateVCUInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateVCUInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
