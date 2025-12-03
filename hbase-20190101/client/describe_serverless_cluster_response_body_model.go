// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerlessClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *DescribeServerlessClusterResponseBody
	GetAutoRenew() *string
	SetClusterType(v string) *DescribeServerlessClusterResponseBody
	GetClusterType() *string
	SetCreateTime(v string) *DescribeServerlessClusterResponseBody
	GetCreateTime() *string
	SetCuSize(v string) *DescribeServerlessClusterResponseBody
	GetCuSize() *string
	SetDiskSize(v string) *DescribeServerlessClusterResponseBody
	GetDiskSize() *string
	SetExpireTime(v string) *DescribeServerlessClusterResponseBody
	GetExpireTime() *string
	SetHaType(v string) *DescribeServerlessClusterResponseBody
	GetHaType() *string
	SetHasUser(v string) *DescribeServerlessClusterResponseBody
	GetHasUser() *string
	SetInnerEndpoint(v string) *DescribeServerlessClusterResponseBody
	GetInnerEndpoint() *string
	SetInstanceId(v string) *DescribeServerlessClusterResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeServerlessClusterResponseBody
	GetInstanceName() *string
	SetIsDeletionProtection(v string) *DescribeServerlessClusterResponseBody
	GetIsDeletionProtection() *string
	SetLockMode(v string) *DescribeServerlessClusterResponseBody
	GetLockMode() *string
	SetMainVersion(v string) *DescribeServerlessClusterResponseBody
	GetMainVersion() *string
	SetOuterEndpoint(v string) *DescribeServerlessClusterResponseBody
	GetOuterEndpoint() *string
	SetPayType(v string) *DescribeServerlessClusterResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeServerlessClusterResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeServerlessClusterResponseBody
	GetRequestId() *string
	SetReserverMaxQpsNum(v string) *DescribeServerlessClusterResponseBody
	GetReserverMaxQpsNum() *string
	SetReserverMinQpsNum(v string) *DescribeServerlessClusterResponseBody
	GetReserverMinQpsNum() *string
	SetResourceGroupId(v string) *DescribeServerlessClusterResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeServerlessClusterResponseBody
	GetStatus() *string
	SetUpdateStatus(v string) *DescribeServerlessClusterResponseBody
	GetUpdateStatus() *string
	SetVSwitchId(v string) *DescribeServerlessClusterResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeServerlessClusterResponseBody
	GetVpcId() *string
	SetZoneId(v string) *DescribeServerlessClusterResponseBody
	GetZoneId() *string
}

type DescribeServerlessClusterResponseBody struct {
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// single
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 150
	CuSize *string `json:"CuSize,omitempty" xml:"CuSize,omitempty"`
	// example:
	//
	// 200
	DiskSize *string `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	HaType *string `json:"HaType,omitempty" xml:"HaType,omitempty"`
	// example:
	//
	// false
	HasUser *string `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	// example:
	//
	// https://sh-wz91452kg946i****-lindorm-serverless-in.lindorm.rds.aliyuncs.com:443
	InnerEndpoint *string `json:"InnerEndpoint,omitempty" xml:"InnerEndpoint,omitempty"`
	// example:
	//
	// hb-bp16f1441y6p2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// true
	IsDeletionProtection *string `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 2.0.8
	MainVersion *string `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	// example:
	//
	// https://sh-wz91452kg946i****-lindorm-serverless.lindorm.rds.aliyuncs.com:443
	OuterEndpoint *string `json:"OuterEndpoint,omitempty" xml:"OuterEndpoint,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 89F81C30-320B-4550-91DB-C37C81D2358F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ReserverMaxQpsNum *string `json:"ReserverMaxQpsNum,omitempty" xml:"ReserverMaxQpsNum,omitempty"`
	// example:
	//
	// 50
	ReserverMinQpsNum *string `json:"ReserverMinQpsNum,omitempty" xml:"ReserverMinQpsNum,omitempty"`
	// example:
	//
	// rg-fjm2d4v7sf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// NO
	UpdateStatus *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	// example:
	//
	// vsw-bp191ipotqf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp120k6ixs4eoghz****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerlessClusterResponseBody) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *DescribeServerlessClusterResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeServerlessClusterResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeServerlessClusterResponseBody) GetCuSize() *string {
	return s.CuSize
}

func (s *DescribeServerlessClusterResponseBody) GetDiskSize() *string {
	return s.DiskSize
}

func (s *DescribeServerlessClusterResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeServerlessClusterResponseBody) GetHaType() *string {
	return s.HaType
}

func (s *DescribeServerlessClusterResponseBody) GetHasUser() *string {
	return s.HasUser
}

func (s *DescribeServerlessClusterResponseBody) GetInnerEndpoint() *string {
	return s.InnerEndpoint
}

func (s *DescribeServerlessClusterResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeServerlessClusterResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeServerlessClusterResponseBody) GetIsDeletionProtection() *string {
	return s.IsDeletionProtection
}

func (s *DescribeServerlessClusterResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeServerlessClusterResponseBody) GetMainVersion() *string {
	return s.MainVersion
}

func (s *DescribeServerlessClusterResponseBody) GetOuterEndpoint() *string {
	return s.OuterEndpoint
}

func (s *DescribeServerlessClusterResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeServerlessClusterResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServerlessClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServerlessClusterResponseBody) GetReserverMaxQpsNum() *string {
	return s.ReserverMaxQpsNum
}

func (s *DescribeServerlessClusterResponseBody) GetReserverMinQpsNum() *string {
	return s.ReserverMinQpsNum
}

func (s *DescribeServerlessClusterResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeServerlessClusterResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeServerlessClusterResponseBody) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeServerlessClusterResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeServerlessClusterResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeServerlessClusterResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeServerlessClusterResponseBody) SetAutoRenew(v string) *DescribeServerlessClusterResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetClusterType(v string) *DescribeServerlessClusterResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetCreateTime(v string) *DescribeServerlessClusterResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetCuSize(v string) *DescribeServerlessClusterResponseBody {
	s.CuSize = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetDiskSize(v string) *DescribeServerlessClusterResponseBody {
	s.DiskSize = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetExpireTime(v string) *DescribeServerlessClusterResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetHaType(v string) *DescribeServerlessClusterResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetHasUser(v string) *DescribeServerlessClusterResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetInnerEndpoint(v string) *DescribeServerlessClusterResponseBody {
	s.InnerEndpoint = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetInstanceId(v string) *DescribeServerlessClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetInstanceName(v string) *DescribeServerlessClusterResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetIsDeletionProtection(v string) *DescribeServerlessClusterResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetLockMode(v string) *DescribeServerlessClusterResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetMainVersion(v string) *DescribeServerlessClusterResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetOuterEndpoint(v string) *DescribeServerlessClusterResponseBody {
	s.OuterEndpoint = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetPayType(v string) *DescribeServerlessClusterResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetRegionId(v string) *DescribeServerlessClusterResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetRequestId(v string) *DescribeServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetReserverMaxQpsNum(v string) *DescribeServerlessClusterResponseBody {
	s.ReserverMaxQpsNum = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetReserverMinQpsNum(v string) *DescribeServerlessClusterResponseBody {
	s.ReserverMinQpsNum = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetResourceGroupId(v string) *DescribeServerlessClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetStatus(v string) *DescribeServerlessClusterResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetUpdateStatus(v string) *DescribeServerlessClusterResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetVSwitchId(v string) *DescribeServerlessClusterResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetVpcId(v string) *DescribeServerlessClusterResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) SetZoneId(v string) *DescribeServerlessClusterResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeServerlessClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
