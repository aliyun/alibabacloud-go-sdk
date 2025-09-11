// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody
	GetInstance() *GetInstanceResponseBodyInstance
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
}

type GetInstanceResponseBody struct {
	// The information about the HSM.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetInstance() *GetInstanceResponseBodyInstance {
	return s.Instance
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstance struct {
	// The ID of the cluster to which the HSM belongs.
	//
	// example:
	//
	// cluster-w3G9vOJI2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// cluster_online
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The time when the HSM was created.
	//
	// example:
	//
	// 1699515963000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the device.
	//
	// example:
	//
	// jnta.SJJ1528-G
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The time when the HSM expired.
	//
	// example:
	//
	// 1699496389720
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the HSM.
	//
	// example:
	//
	// hsm-cn-g4t3jwsc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the HSM in the VPC.
	//
	// example:
	//
	// 10.192.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Indicates whether the HSM is for trial use. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsTrial *bool `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	// Indicates whether the HSM is a master HSM. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Master *bool `json:"Master,omitempty" xml:"Master,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 23576634952****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1
	PqcEnabled *int32 `json:"PqcEnabled,omitempty" xml:"PqcEnabled,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the HSM.
	//
	// example:
	//
	// hsmOnline
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the HSM. Valid values:
	//
	// 	- PENDING: The HSM is disabled.
	//
	// 	- ACTIVE: The HSM is enabled.
	//
	// 	- EXPIRED: The HSM expired.
	//
	// 	- INVALID: The HSM is invalid.
	//
	// 	- FAILURE: The HSM failed to be created.
	//
	// 	- RESET: The HSM is being reset.
	//
	// 	- PAUSED: The HSM is paused.
	//
	// 	- MODIFYING: The HSM is being modified.
	//
	// example:
	//
	// EXPIRED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of HSMs that is classified based on resource isolation. Valid values:
	//
	// - vsm: Virtual security modules (VSMs)
	//
	// - hostedHsm: Dedicated HSMs.
	//
	// example:
	//
	// vsm
	TenantIsolationType *string `json:"TenantIsolationType,omitempty" xml:"TenantIsolationType,omitempty"`
	// The ID of the vSwitch that is configured for the HSM.
	//
	// example:
	//
	// vsw-bp1mvfs31ltt0wyhf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The information about the vendor.
	//
	// example:
	//
	// jnta
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the HSM belongs.
	//
	// example:
	//
	// vpc-uf69i66j9kmoko52p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// 18.68.XX.XX
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetInstanceResponseBodyInstance) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetInstanceResponseBodyInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceResponseBodyInstance) GetDeviceType() *string {
	return s.DeviceType
}

func (s *GetInstanceResponseBodyInstance) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyInstance) GetIp() *string {
	return s.Ip
}

func (s *GetInstanceResponseBodyInstance) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *GetInstanceResponseBodyInstance) GetMaster() *bool {
	return s.Master
}

func (s *GetInstanceResponseBodyInstance) GetOrderId() *string {
	return s.OrderId
}

func (s *GetInstanceResponseBodyInstance) GetPqcEnabled() *int32 {
	return s.PqcEnabled
}

func (s *GetInstanceResponseBodyInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceResponseBodyInstance) GetRemark() *string {
	return s.Remark
}

func (s *GetInstanceResponseBodyInstance) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyInstance) GetTenantIsolationType() *string {
	return s.TenantIsolationType
}

func (s *GetInstanceResponseBodyInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceResponseBodyInstance) GetVendor() *string {
	return s.Vendor
}

func (s *GetInstanceResponseBodyInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyInstance) GetWhitelist() *string {
	return s.Whitelist
}

func (s *GetInstanceResponseBodyInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetInstanceResponseBodyInstance) SetClusterId(v string) *GetInstanceResponseBodyInstance {
	s.ClusterId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetClusterName(v string) *GetInstanceResponseBodyInstance {
	s.ClusterName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCreateTime(v int64) *GetInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDeviceType(v string) *GetInstanceResponseBodyInstance {
	s.DeviceType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetExpireTime(v int64) *GetInstanceResponseBodyInstance {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetIp(v string) *GetInstanceResponseBodyInstance {
	s.Ip = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetIsTrial(v bool) *GetInstanceResponseBodyInstance {
	s.IsTrial = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetMaster(v bool) *GetInstanceResponseBodyInstance {
	s.Master = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetOrderId(v string) *GetInstanceResponseBodyInstance {
	s.OrderId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetPqcEnabled(v int32) *GetInstanceResponseBodyInstance {
	s.PqcEnabled = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetRegionId(v string) *GetInstanceResponseBodyInstance {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetRemark(v string) *GetInstanceResponseBodyInstance {
	s.Remark = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetStatus(v string) *GetInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetTenantIsolationType(v string) *GetInstanceResponseBodyInstance {
	s.TenantIsolationType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVSwitchId(v string) *GetInstanceResponseBodyInstance {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVendor(v string) *GetInstanceResponseBodyInstance {
	s.Vendor = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVpcId(v string) *GetInstanceResponseBodyInstance {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetWhitelist(v string) *GetInstanceResponseBodyInstance {
	s.Whitelist = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetZoneId(v string) *GetInstanceResponseBodyInstance {
	s.ZoneId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
