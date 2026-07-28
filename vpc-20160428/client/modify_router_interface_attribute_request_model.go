// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteHealthCheckIp(v bool) *ModifyRouterInterfaceAttributeRequest
	GetDeleteHealthCheckIp() *bool
	SetDescription(v string) *ModifyRouterInterfaceAttributeRequest
	GetDescription() *string
	SetHcRate(v int32) *ModifyRouterInterfaceAttributeRequest
	GetHcRate() *int32
	SetHcThreshold(v int32) *ModifyRouterInterfaceAttributeRequest
	GetHcThreshold() *int32
	SetHealthCheckSourceIp(v string) *ModifyRouterInterfaceAttributeRequest
	GetHealthCheckSourceIp() *string
	SetHealthCheckTargetIp(v string) *ModifyRouterInterfaceAttributeRequest
	GetHealthCheckTargetIp() *string
	SetName(v string) *ModifyRouterInterfaceAttributeRequest
	GetName() *string
	SetOppositeInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest
	GetOppositeInterfaceId() *string
	SetOppositeInterfaceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest
	GetOppositeInterfaceOwnerId() *int64
	SetOppositeRouterId(v string) *ModifyRouterInterfaceAttributeRequest
	GetOppositeRouterId() *string
	SetOppositeRouterType(v string) *ModifyRouterInterfaceAttributeRequest
	GetOppositeRouterType() *string
	SetOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRouterInterfaceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRouterInterfaceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest
	GetRouterInterfaceId() *string
}

type ModifyRouterInterfaceAttributeRequest struct {
	// Specifies whether to delete the health check IP addresses configured on the router interface. Valid values:
	//
	// - **true**: Deletes the health check IP addresses.
	//
	//
	//
	// - **false*	- (default): Does not delete the health check IP addresses.
	//
	// example:
	//
	// false
	DeleteHealthCheckIp *bool `json:"DeleteHealthCheckIp,omitempty" xml:"DeleteHealthCheckIp,omitempty"`
	// The description of the router interface.
	//
	// The description must be 2 to 256 characters in length and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// 路由器接口
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The health check rate. Unit: milliseconds. Recommended value: **2000**. This parameter specifies the interval between consecutive probe packets sent during a health check.
	//
	// In this example, **HcThreshold*	- is set to **8*	- and **HcRate*	- is set to **2000**. This means that during a health check, a probe packet is sent from **HealthCheckSourceIp*	- (the source IP address for health checks) to **HealthCheckTargetIp*	- (the destination IP address for health checks) every 2000 milliseconds. If 8 consecutive probe packets receive no response, the health check fails.
	//
	// example:
	//
	// 2000
	HcRate *int32 `json:"HcRate,omitempty" xml:"HcRate,omitempty"`
	// The health check threshold. Unit: packets. Recommended value: **8**. This parameter specifies the number of probe packets sent during a health check.
	//
	// example:
	//
	// 8
	HcThreshold *int32 `json:"HcThreshold,omitempty" xml:"HcThreshold,omitempty"`
	// The source IP address for health checks. The IP address must be an unused IP address in the local VPC.
	//
	// > You can specify this parameter in Express Connect circuit scenarios.
	//
	// example:
	//
	// 116.62.XX.XX
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	// The destination IP address for health checks.
	//
	// > This parameter is required if **HealthCheckSourceIp*	- is specified.
	//
	// example:
	//
	// 116.62.XX.XX
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	// The name of the router interface.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// TEST
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the peer router interface.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urz****
	OppositeInterfaceId *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	// The ID of the account to which the peer router interface belongs.
	//
	// example:
	//
	// 28768383240243****
	OppositeInterfaceOwnerId *int64 `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	// The ID of the peer router.
	//
	// example:
	//
	// vrt-bp1jcg5cmxjbl9xgc****
	OppositeRouterId *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	// The type of the router to which the peer router interface belongs. Valid values:
	//
	// - **VRouter**: vRouter.
	//
	// - **VBR*	- (default): Virtual Border Router.
	//
	// example:
	//
	// VBR
	OppositeRouterType *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the router interface.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urz****
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
}

func (s ModifyRouterInterfaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceAttributeRequest) GetDeleteHealthCheckIp() *bool {
	return s.DeleteHealthCheckIp
}

func (s *ModifyRouterInterfaceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyRouterInterfaceAttributeRequest) GetHcRate() *int32 {
	return s.HcRate
}

func (s *ModifyRouterInterfaceAttributeRequest) GetHcThreshold() *int32 {
	return s.HcThreshold
}

func (s *ModifyRouterInterfaceAttributeRequest) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *ModifyRouterInterfaceAttributeRequest) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *ModifyRouterInterfaceAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeInterfaceId() *string {
	return s.OppositeInterfaceId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeInterfaceOwnerId() *int64 {
	return s.OppositeInterfaceOwnerId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeRouterId() *string {
	return s.OppositeRouterId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOppositeRouterType() *string {
	return s.OppositeRouterType
}

func (s *ModifyRouterInterfaceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRouterInterfaceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRouterInterfaceAttributeRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *ModifyRouterInterfaceAttributeRequest) SetDeleteHealthCheckIp(v bool) *ModifyRouterInterfaceAttributeRequest {
	s.DeleteHealthCheckIp = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetDescription(v string) *ModifyRouterInterfaceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetHcRate(v int32) *ModifyRouterInterfaceAttributeRequest {
	s.HcRate = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetHcThreshold(v int32) *ModifyRouterInterfaceAttributeRequest {
	s.HcThreshold = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetHealthCheckSourceIp(v string) *ModifyRouterInterfaceAttributeRequest {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetHealthCheckTargetIp(v string) *ModifyRouterInterfaceAttributeRequest {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetName(v string) *ModifyRouterInterfaceAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeInterfaceId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeInterfaceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeInterfaceOwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeRouterId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeRouterId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOppositeRouterType(v string) *ModifyRouterInterfaceAttributeRequest {
	s.OppositeRouterType = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetRegionId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyRouterInterfaceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetResourceOwnerId(v int64) *ModifyRouterInterfaceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) SetRouterInterfaceId(v string) *ModifyRouterInterfaceAttributeRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *ModifyRouterInterfaceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
