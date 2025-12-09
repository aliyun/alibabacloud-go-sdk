// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *CreateApplicationRequest
	GetApplicationType() *string
	SetArchitecture(v string) *CreateApplicationRequest
	GetArchitecture() *string
	SetAutoRenew(v bool) *CreateApplicationRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *CreateApplicationRequest
	GetAutoUseCoupon() *bool
	SetComponents(v []*CreateApplicationRequestComponents) *CreateApplicationRequest
	GetComponents() []*CreateApplicationRequestComponents
	SetDBClusterId(v string) *CreateApplicationRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateApplicationRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateApplicationRequest
	GetDryRun() *bool
	SetEndpoints(v []*CreateApplicationRequestEndpoints) *CreateApplicationRequest
	GetEndpoints() []*CreateApplicationRequestEndpoints
	SetPayType(v string) *CreateApplicationRequest
	GetPayType() *string
	SetPeriod(v string) *CreateApplicationRequest
	GetPeriod() *string
	SetPolarFSInstanceId(v string) *CreateApplicationRequest
	GetPolarFSInstanceId() *string
	SetPromotionCode(v string) *CreateApplicationRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateApplicationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateApplicationRequest
	GetResourceGroupId() *string
	SetUsedTime(v string) *CreateApplicationRequest
	GetUsedTime() *string
	SetVSwitchId(v string) *CreateApplicationRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateApplicationRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateApplicationRequest
	GetZoneId() *string
}

type CreateApplicationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// x86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// true
	AutoRenew     *bool                                 `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoUseCoupon *bool                                 `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	Components    []*CreateApplicationRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// myapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun    *bool                                `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Endpoints []*CreateApplicationRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// pcs-********************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-beijing-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateApplicationRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateApplicationRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateApplicationRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateApplicationRequest) GetComponents() []*CreateApplicationRequestComponents {
	return s.Components
}

func (s *CreateApplicationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateApplicationRequest) GetEndpoints() []*CreateApplicationRequestEndpoints {
	return s.Endpoints
}

func (s *CreateApplicationRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateApplicationRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateApplicationRequest) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *CreateApplicationRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateApplicationRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateApplicationRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateApplicationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateApplicationRequest) SetApplicationType(v string) *CreateApplicationRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateApplicationRequest) SetArchitecture(v string) *CreateApplicationRequest {
	s.Architecture = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoRenew(v bool) *CreateApplicationRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoUseCoupon(v bool) *CreateApplicationRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateApplicationRequest) SetComponents(v []*CreateApplicationRequestComponents) *CreateApplicationRequest {
	s.Components = v
	return s
}

func (s *CreateApplicationRequest) SetDBClusterId(v string) *CreateApplicationRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetDryRun(v bool) *CreateApplicationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateApplicationRequest) SetEndpoints(v []*CreateApplicationRequestEndpoints) *CreateApplicationRequest {
	s.Endpoints = v
	return s
}

func (s *CreateApplicationRequest) SetPayType(v string) *CreateApplicationRequest {
	s.PayType = &v
	return s
}

func (s *CreateApplicationRequest) SetPeriod(v string) *CreateApplicationRequest {
	s.Period = &v
	return s
}

func (s *CreateApplicationRequest) SetPolarFSInstanceId(v string) *CreateApplicationRequest {
	s.PolarFSInstanceId = &v
	return s
}

func (s *CreateApplicationRequest) SetPromotionCode(v string) *CreateApplicationRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateApplicationRequest) SetRegionId(v string) *CreateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationRequest) SetResourceGroupId(v string) *CreateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetUsedTime(v string) *CreateApplicationRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateApplicationRequest) SetVSwitchId(v string) *CreateApplicationRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateApplicationRequest) SetVpcId(v string) *CreateApplicationRequest {
	s.VpcId = &v
	return s
}

func (s *CreateApplicationRequest) SetZoneId(v string) *CreateApplicationRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateApplicationRequestComponents struct {
	// example:
	//
	// polar.app.g2.medium
	ComponentClass *string `json:"ComponentClass,omitempty" xml:"ComponentClass,omitempty"`
	// example:
	//
	// 1
	ComponentMaxReplica *int64 `json:"ComponentMaxReplica,omitempty" xml:"ComponentMaxReplica,omitempty"`
	// example:
	//
	// 1
	ComponentReplica *int64 `json:"ComponentReplica,omitempty" xml:"ComponentReplica,omitempty"`
	// example:
	//
	// gateway
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	ScaleMax      *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	ScaleMin      *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// example:
	//
	// sg-********************
	SecurityGroups *string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty"`
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
}

func (s CreateApplicationRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestComponents) GetComponentClass() *string {
	return s.ComponentClass
}

func (s *CreateApplicationRequestComponents) GetComponentMaxReplica() *int64 {
	return s.ComponentMaxReplica
}

func (s *CreateApplicationRequestComponents) GetComponentReplica() *int64 {
	return s.ComponentReplica
}

func (s *CreateApplicationRequestComponents) GetComponentType() *string {
	return s.ComponentType
}

func (s *CreateApplicationRequestComponents) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *CreateApplicationRequestComponents) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *CreateApplicationRequestComponents) GetSecurityGroups() *string {
	return s.SecurityGroups
}

func (s *CreateApplicationRequestComponents) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *CreateApplicationRequestComponents) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateApplicationRequestComponents) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *CreateApplicationRequestComponents) SetComponentClass(v string) *CreateApplicationRequestComponents {
	s.ComponentClass = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetComponentMaxReplica(v int64) *CreateApplicationRequestComponents {
	s.ComponentMaxReplica = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetComponentReplica(v int64) *CreateApplicationRequestComponents {
	s.ComponentReplica = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetComponentType(v string) *CreateApplicationRequestComponents {
	s.ComponentType = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetScaleMax(v string) *CreateApplicationRequestComponents {
	s.ScaleMax = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetScaleMin(v string) *CreateApplicationRequestComponents {
	s.ScaleMin = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityGroups(v string) *CreateApplicationRequestComponents {
	s.SecurityGroups = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityIPArrayName(v string) *CreateApplicationRequestComponents {
	s.SecurityIPArrayName = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityIPList(v string) *CreateApplicationRequestComponents {
	s.SecurityIPList = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityIPType(v string) *CreateApplicationRequestComponents {
	s.SecurityIPType = &v
	return s
}

func (s *CreateApplicationRequestComponents) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestEndpoints struct {
	// example:
	//
	// my_endpoint
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Primary
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s CreateApplicationRequestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestEndpoints) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationRequestEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateApplicationRequestEndpoints) SetDescription(v string) *CreateApplicationRequestEndpoints {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequestEndpoints) SetEndpointType(v string) *CreateApplicationRequestEndpoints {
	s.EndpointType = &v
	return s
}

func (s *CreateApplicationRequestEndpoints) Validate() error {
	return dara.Validate(s)
}
