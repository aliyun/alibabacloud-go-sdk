// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerlessInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *DescribeServerlessInstanceResponseBody
	GetAutoRenew() *string
	SetClusterType(v string) *DescribeServerlessInstanceResponseBody
	GetClusterType() *string
	SetCreateTime(v string) *DescribeServerlessInstanceResponseBody
	GetCreateTime() *string
	SetCuSize(v string) *DescribeServerlessInstanceResponseBody
	GetCuSize() *string
	SetDiskSize(v string) *DescribeServerlessInstanceResponseBody
	GetDiskSize() *string
	SetExpireTime(v string) *DescribeServerlessInstanceResponseBody
	GetExpireTime() *string
	SetHaType(v string) *DescribeServerlessInstanceResponseBody
	GetHaType() *string
	SetHasUser(v string) *DescribeServerlessInstanceResponseBody
	GetHasUser() *string
	SetInnerEndpoint(v string) *DescribeServerlessInstanceResponseBody
	GetInnerEndpoint() *string
	SetInstanceId(v string) *DescribeServerlessInstanceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeServerlessInstanceResponseBody
	GetInstanceName() *string
	SetIsDeletionProtection(v string) *DescribeServerlessInstanceResponseBody
	GetIsDeletionProtection() *string
	SetLockMode(v string) *DescribeServerlessInstanceResponseBody
	GetLockMode() *string
	SetMainVersion(v string) *DescribeServerlessInstanceResponseBody
	GetMainVersion() *string
	SetOuterEndpoint(v string) *DescribeServerlessInstanceResponseBody
	GetOuterEndpoint() *string
	SetPayType(v string) *DescribeServerlessInstanceResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeServerlessInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeServerlessInstanceResponseBody
	GetRequestId() *string
	SetReserverMaxQpsNum(v string) *DescribeServerlessInstanceResponseBody
	GetReserverMaxQpsNum() *string
	SetReserverMinQpsNum(v string) *DescribeServerlessInstanceResponseBody
	GetReserverMinQpsNum() *string
	SetStatus(v string) *DescribeServerlessInstanceResponseBody
	GetStatus() *string
	SetUpdateStatus(v string) *DescribeServerlessInstanceResponseBody
	GetUpdateStatus() *string
	SetVSwitchId(v string) *DescribeServerlessInstanceResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeServerlessInstanceResponseBody
	GetVpcId() *string
	SetZoneId(v string) *DescribeServerlessInstanceResponseBody
	GetZoneId() *string
}

type DescribeServerlessInstanceResponseBody struct {
	AutoRenew            *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClusterType          *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CuSize               *string `json:"CuSize,omitempty" xml:"CuSize,omitempty"`
	DiskSize             *string `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	ExpireTime           *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HaType               *string `json:"HaType,omitempty" xml:"HaType,omitempty"`
	HasUser              *string `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	InnerEndpoint        *string `json:"InnerEndpoint,omitempty" xml:"InnerEndpoint,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsDeletionProtection *string `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	OuterEndpoint        *string `json:"OuterEndpoint,omitempty" xml:"OuterEndpoint,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReserverMaxQpsNum    *string `json:"ReserverMaxQpsNum,omitempty" xml:"ReserverMaxQpsNum,omitempty"`
	ReserverMinQpsNum    *string `json:"ReserverMinQpsNum,omitempty" xml:"ReserverMinQpsNum,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateStatus         *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerlessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerlessInstanceResponseBody) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *DescribeServerlessInstanceResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeServerlessInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeServerlessInstanceResponseBody) GetCuSize() *string {
	return s.CuSize
}

func (s *DescribeServerlessInstanceResponseBody) GetDiskSize() *string {
	return s.DiskSize
}

func (s *DescribeServerlessInstanceResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeServerlessInstanceResponseBody) GetHaType() *string {
	return s.HaType
}

func (s *DescribeServerlessInstanceResponseBody) GetHasUser() *string {
	return s.HasUser
}

func (s *DescribeServerlessInstanceResponseBody) GetInnerEndpoint() *string {
	return s.InnerEndpoint
}

func (s *DescribeServerlessInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeServerlessInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeServerlessInstanceResponseBody) GetIsDeletionProtection() *string {
	return s.IsDeletionProtection
}

func (s *DescribeServerlessInstanceResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeServerlessInstanceResponseBody) GetMainVersion() *string {
	return s.MainVersion
}

func (s *DescribeServerlessInstanceResponseBody) GetOuterEndpoint() *string {
	return s.OuterEndpoint
}

func (s *DescribeServerlessInstanceResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeServerlessInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServerlessInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServerlessInstanceResponseBody) GetReserverMaxQpsNum() *string {
	return s.ReserverMaxQpsNum
}

func (s *DescribeServerlessInstanceResponseBody) GetReserverMinQpsNum() *string {
	return s.ReserverMinQpsNum
}

func (s *DescribeServerlessInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeServerlessInstanceResponseBody) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeServerlessInstanceResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeServerlessInstanceResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeServerlessInstanceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeServerlessInstanceResponseBody) SetAutoRenew(v string) *DescribeServerlessInstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetClusterType(v string) *DescribeServerlessInstanceResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetCreateTime(v string) *DescribeServerlessInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetCuSize(v string) *DescribeServerlessInstanceResponseBody {
	s.CuSize = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetDiskSize(v string) *DescribeServerlessInstanceResponseBody {
	s.DiskSize = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetExpireTime(v string) *DescribeServerlessInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetHaType(v string) *DescribeServerlessInstanceResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetHasUser(v string) *DescribeServerlessInstanceResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetInnerEndpoint(v string) *DescribeServerlessInstanceResponseBody {
	s.InnerEndpoint = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetInstanceId(v string) *DescribeServerlessInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetInstanceName(v string) *DescribeServerlessInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetIsDeletionProtection(v string) *DescribeServerlessInstanceResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetLockMode(v string) *DescribeServerlessInstanceResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetMainVersion(v string) *DescribeServerlessInstanceResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetOuterEndpoint(v string) *DescribeServerlessInstanceResponseBody {
	s.OuterEndpoint = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetPayType(v string) *DescribeServerlessInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetRegionId(v string) *DescribeServerlessInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetRequestId(v string) *DescribeServerlessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetReserverMaxQpsNum(v string) *DescribeServerlessInstanceResponseBody {
	s.ReserverMaxQpsNum = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetReserverMinQpsNum(v string) *DescribeServerlessInstanceResponseBody {
	s.ReserverMinQpsNum = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetStatus(v string) *DescribeServerlessInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetUpdateStatus(v string) *DescribeServerlessInstanceResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetVSwitchId(v string) *DescribeServerlessInstanceResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetVpcId(v string) *DescribeServerlessInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) SetZoneId(v string) *DescribeServerlessInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeServerlessInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
