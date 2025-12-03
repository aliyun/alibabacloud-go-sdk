// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *DescribeClusterAttributeResponseBody
	GetAutoRenew() *string
	SetClusterId(v string) *DescribeClusterAttributeResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeClusterAttributeResponseBody
	GetClusterName() *string
	SetClusterType(v string) *DescribeClusterAttributeResponseBody
	GetClusterType() *string
	SetColdStorageStatus(v string) *DescribeClusterAttributeResponseBody
	GetColdStorageStatus() *string
	SetConnectionInfo(v *DescribeClusterAttributeResponseBodyConnectionInfo) *DescribeClusterAttributeResponseBody
	GetConnectionInfo() *DescribeClusterAttributeResponseBodyConnectionInfo
	SetCoreDiskQuantity(v int32) *DescribeClusterAttributeResponseBody
	GetCoreDiskQuantity() *int32
	SetCoreDiskSize(v string) *DescribeClusterAttributeResponseBody
	GetCoreDiskSize() *string
	SetCoreDiskType(v string) *DescribeClusterAttributeResponseBody
	GetCoreDiskType() *string
	SetCoreInstanceQuantity(v int32) *DescribeClusterAttributeResponseBody
	GetCoreInstanceQuantity() *int32
	SetCoreInstanceType(v string) *DescribeClusterAttributeResponseBody
	GetCoreInstanceType() *string
	SetCreateTime(v string) *DescribeClusterAttributeResponseBody
	GetCreateTime() *string
	SetExpireTime(v string) *DescribeClusterAttributeResponseBody
	GetExpireTime() *string
	SetHaType(v string) *DescribeClusterAttributeResponseBody
	GetHaType() *string
	SetHasUser(v string) *DescribeClusterAttributeResponseBody
	GetHasUser() *string
	SetLockMode(v string) *DescribeClusterAttributeResponseBody
	GetLockMode() *string
	SetMainVersion(v string) *DescribeClusterAttributeResponseBody
	GetMainVersion() *string
	SetMasterDiskSize(v int32) *DescribeClusterAttributeResponseBody
	GetMasterDiskSize() *int32
	SetMasterDiskType(v string) *DescribeClusterAttributeResponseBody
	GetMasterDiskType() *string
	SetMasterInstanceType(v string) *DescribeClusterAttributeResponseBody
	GetMasterInstanceType() *string
	SetMinorVersion(v string) *DescribeClusterAttributeResponseBody
	GetMinorVersion() *string
	SetNetInfo(v *DescribeClusterAttributeResponseBodyNetInfo) *DescribeClusterAttributeResponseBody
	GetNetInfo() *DescribeClusterAttributeResponseBodyNetInfo
	SetNodeList(v *DescribeClusterAttributeResponseBodyNodeList) *DescribeClusterAttributeResponseBody
	GetNodeList() *DescribeClusterAttributeResponseBodyNodeList
	SetPayType(v string) *DescribeClusterAttributeResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeClusterAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeClusterAttributeResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeClusterAttributeResponseBody
	GetStatus() *string
	SetUpdateStatus(v string) *DescribeClusterAttributeResponseBody
	GetUpdateStatus() *string
	SetZoneId(v string) *DescribeClusterAttributeResponseBody
	GetZoneId() *string
}

type DescribeClusterAttributeResponseBody struct {
	AutoRenew            *string                                             `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ClusterId            *string                                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                                             `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string                                             `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ColdStorageStatus    *string                                             `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	ConnectionInfo       *DescribeClusterAttributeResponseBodyConnectionInfo `json:"ConnectionInfo,omitempty" xml:"ConnectionInfo,omitempty" type:"Struct"`
	CoreDiskQuantity     *int32                                              `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	CoreDiskSize         *string                                             `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string                                             `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32                                              `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CoreInstanceType     *string                                             `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	CreateTime           *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime           *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HaType               *string                                             `json:"HaType,omitempty" xml:"HaType,omitempty"`
	HasUser              *string                                             `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	LockMode             *string                                             `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string                                             `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	MasterDiskSize       *int32                                              `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	MasterDiskType       *string                                             `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	MasterInstanceType   *string                                             `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	MinorVersion         *string                                             `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	NetInfo              *DescribeClusterAttributeResponseBodyNetInfo        `json:"NetInfo,omitempty" xml:"NetInfo,omitempty" type:"Struct"`
	NodeList             *DescribeClusterAttributeResponseBodyNodeList       `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Struct"`
	PayType              *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status               *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateStatus         *string                                             `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	ZoneId               *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBody) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *DescribeClusterAttributeResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterAttributeResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeClusterAttributeResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClusterAttributeResponseBody) GetColdStorageStatus() *string {
	return s.ColdStorageStatus
}

func (s *DescribeClusterAttributeResponseBody) GetConnectionInfo() *DescribeClusterAttributeResponseBodyConnectionInfo {
	return s.ConnectionInfo
}

func (s *DescribeClusterAttributeResponseBody) GetCoreDiskQuantity() *int32 {
	return s.CoreDiskQuantity
}

func (s *DescribeClusterAttributeResponseBody) GetCoreDiskSize() *string {
	return s.CoreDiskSize
}

func (s *DescribeClusterAttributeResponseBody) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *DescribeClusterAttributeResponseBody) GetCoreInstanceQuantity() *int32 {
	return s.CoreInstanceQuantity
}

func (s *DescribeClusterAttributeResponseBody) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *DescribeClusterAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeClusterAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeClusterAttributeResponseBody) GetHaType() *string {
	return s.HaType
}

func (s *DescribeClusterAttributeResponseBody) GetHasUser() *string {
	return s.HasUser
}

func (s *DescribeClusterAttributeResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeClusterAttributeResponseBody) GetMainVersion() *string {
	return s.MainVersion
}

func (s *DescribeClusterAttributeResponseBody) GetMasterDiskSize() *int32 {
	return s.MasterDiskSize
}

func (s *DescribeClusterAttributeResponseBody) GetMasterDiskType() *string {
	return s.MasterDiskType
}

func (s *DescribeClusterAttributeResponseBody) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *DescribeClusterAttributeResponseBody) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeClusterAttributeResponseBody) GetNetInfo() *DescribeClusterAttributeResponseBodyNetInfo {
	return s.NetInfo
}

func (s *DescribeClusterAttributeResponseBody) GetNodeList() *DescribeClusterAttributeResponseBodyNodeList {
	return s.NodeList
}

func (s *DescribeClusterAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeClusterAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeClusterAttributeResponseBody) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeClusterAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterAttributeResponseBody) SetAutoRenew(v string) *DescribeClusterAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetClusterId(v string) *DescribeClusterAttributeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetClusterName(v string) *DescribeClusterAttributeResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetClusterType(v string) *DescribeClusterAttributeResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetColdStorageStatus(v string) *DescribeClusterAttributeResponseBody {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetConnectionInfo(v *DescribeClusterAttributeResponseBodyConnectionInfo) *DescribeClusterAttributeResponseBody {
	s.ConnectionInfo = v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreDiskQuantity(v int32) *DescribeClusterAttributeResponseBody {
	s.CoreDiskQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreDiskSize(v string) *DescribeClusterAttributeResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreDiskType(v string) *DescribeClusterAttributeResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreInstanceQuantity(v int32) *DescribeClusterAttributeResponseBody {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCoreInstanceType(v string) *DescribeClusterAttributeResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetCreateTime(v string) *DescribeClusterAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetExpireTime(v string) *DescribeClusterAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetHaType(v string) *DescribeClusterAttributeResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetHasUser(v string) *DescribeClusterAttributeResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetLockMode(v string) *DescribeClusterAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMainVersion(v string) *DescribeClusterAttributeResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMasterDiskSize(v int32) *DescribeClusterAttributeResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMasterDiskType(v string) *DescribeClusterAttributeResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMasterInstanceType(v string) *DescribeClusterAttributeResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetMinorVersion(v string) *DescribeClusterAttributeResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetNetInfo(v *DescribeClusterAttributeResponseBodyNetInfo) *DescribeClusterAttributeResponseBody {
	s.NetInfo = v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetNodeList(v *DescribeClusterAttributeResponseBodyNodeList) *DescribeClusterAttributeResponseBody {
	s.NodeList = v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetPayType(v string) *DescribeClusterAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetRegionId(v string) *DescribeClusterAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetRequestId(v string) *DescribeClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetStatus(v string) *DescribeClusterAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetUpdateStatus(v string) *DescribeClusterAttributeResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) SetZoneId(v string) *DescribeClusterAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBody) Validate() error {
	if s.ConnectionInfo != nil {
		if err := s.ConnectionInfo.Validate(); err != nil {
			return err
		}
	}
	if s.NetInfo != nil {
		if err := s.NetInfo.Validate(); err != nil {
			return err
		}
	}
	if s.NodeList != nil {
		if err := s.NodeList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterAttributeResponseBodyConnectionInfo struct {
	HaRestConnectionString        *string                                                                          `json:"HaRestConnectionString,omitempty" xml:"HaRestConnectionString,omitempty"`
	HaRestPort                    *string                                                                          `json:"HaRestPort,omitempty" xml:"HaRestPort,omitempty"`
	HaThriftConnectionString      *string                                                                          `json:"HaThriftConnectionString,omitempty" xml:"HaThriftConnectionString,omitempty"`
	HaThriftPort                  *string                                                                          `json:"HaThriftPort,omitempty" xml:"HaThriftPort,omitempty"`
	ThriftConnectionString        *string                                                                          `json:"ThriftConnectionString,omitempty" xml:"ThriftConnectionString,omitempty"`
	ThriftPort                    *string                                                                          `json:"ThriftPort,omitempty" xml:"ThriftPort,omitempty"`
	UIProxyConnectionString       *string                                                                          `json:"UIProxyConnectionString,omitempty" xml:"UIProxyConnectionString,omitempty"`
	ZKClassicConnectionStringList *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList `json:"ZKClassicConnectionStringList,omitempty" xml:"ZKClassicConnectionStringList,omitempty" type:"Struct"`
	ZKConnectionStringList        *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList        `json:"ZKConnectionStringList,omitempty" xml:"ZKConnectionStringList,omitempty" type:"Struct"`
	ZKInnerConnectionStringList   *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList   `json:"ZKInnerConnectionStringList,omitempty" xml:"ZKInnerConnectionStringList,omitempty" type:"Struct"`
	ZKPort                        *int32                                                                           `json:"ZKPort,omitempty" xml:"ZKPort,omitempty"`
	ZKPublicConnectionStringList  *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList  `json:"ZKPublicConnectionStringList,omitempty" xml:"ZKPublicConnectionStringList,omitempty" type:"Struct"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetHaRestConnectionString() *string {
	return s.HaRestConnectionString
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetHaRestPort() *string {
	return s.HaRestPort
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetHaThriftConnectionString() *string {
	return s.HaThriftConnectionString
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetHaThriftPort() *string {
	return s.HaThriftPort
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetThriftConnectionString() *string {
	return s.ThriftConnectionString
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetThriftPort() *string {
	return s.ThriftPort
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetUIProxyConnectionString() *string {
	return s.UIProxyConnectionString
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetZKClassicConnectionStringList() *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList {
	return s.ZKClassicConnectionStringList
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetZKConnectionStringList() *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList {
	return s.ZKConnectionStringList
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetZKInnerConnectionStringList() *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList {
	return s.ZKInnerConnectionStringList
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetZKPort() *int32 {
	return s.ZKPort
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) GetZKPublicConnectionStringList() *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList {
	return s.ZKPublicConnectionStringList
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaRestConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaRestConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaRestPort(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaRestPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaThriftConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaThriftConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetHaThriftPort(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.HaThriftPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetThriftConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ThriftConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetThriftPort(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ThriftPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetUIProxyConnectionString(v string) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.UIProxyConnectionString = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKClassicConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKClassicConnectionStringList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKConnectionStringList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKInnerConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKInnerConnectionStringList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKPort(v int32) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKPort = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) SetZKPublicConnectionStringList(v *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) *DescribeClusterAttributeResponseBodyConnectionInfo {
	s.ZKPublicConnectionStringList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfo) Validate() error {
	if s.ZKClassicConnectionStringList != nil {
		if err := s.ZKClassicConnectionStringList.Validate(); err != nil {
			return err
		}
	}
	if s.ZKConnectionStringList != nil {
		if err := s.ZKConnectionStringList.Validate(); err != nil {
			return err
		}
	}
	if s.ZKInnerConnectionStringList != nil {
		if err := s.ZKInnerConnectionStringList.Validate(); err != nil {
			return err
		}
	}
	if s.ZKPublicConnectionStringList != nil {
		if err := s.ZKPublicConnectionStringList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) GetString_() []*string {
	return s.String_
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList {
	s.String_ = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKClassicConnectionStringList) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) GetString_() []*string {
	return s.String_
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList {
	s.String_ = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKConnectionStringList) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) GetString_() []*string {
	return s.String_
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList {
	s.String_ = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKInnerConnectionStringList) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) GetString_() []*string {
	return s.String_
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) SetString_(v []*string) *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList {
	s.String_ = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyConnectionInfoZKPublicConnectionStringList) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterAttributeResponseBodyNetInfo struct {
	InnerIpAddress  *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	NetType         *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	SecurityIpList  *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClusterAttributeResponseBodyNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) GetInnerIpAddress() *string {
	return s.InnerIpAddress
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetInnerIpAddress(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.InnerIpAddress = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetNetType(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetPublicIpAddress(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetSecurityIpList(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.SecurityIpList = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetVSwitchId(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) SetVpcId(v string) *DescribeClusterAttributeResponseBodyNetInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNetInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterAttributeResponseBodyNodeList struct {
	Node []*DescribeClusterAttributeResponseBodyNodeListNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyNodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeList) GetNode() []*DescribeClusterAttributeResponseBodyNodeListNode {
	return s.Node
}

func (s *DescribeClusterAttributeResponseBodyNodeList) SetNode(v []*DescribeClusterAttributeResponseBodyNodeListNode) *DescribeClusterAttributeResponseBodyNodeList {
	s.Node = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeList) Validate() error {
	if s.Node != nil {
		for _, item := range s.Node {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterAttributeResponseBodyNodeListNode struct {
	DaemonList       *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList `json:"DaemonList,omitempty" xml:"DaemonList,omitempty" type:"Struct"`
	MemStore         *string                                                     `json:"MemStore,omitempty" xml:"MemStore,omitempty"`
	NodeDiskQuantity *string                                                     `json:"NodeDiskQuantity,omitempty" xml:"NodeDiskQuantity,omitempty"`
	NodeDiskSize     *string                                                     `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	NodeDiskType     *string                                                     `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	NodeId           *string                                                     `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeInstanceType *string                                                     `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodeStatus       *string                                                     `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	NodeType         *string                                                     `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RegionQuantity   *string                                                     `json:"RegionQuantity,omitempty" xml:"RegionQuantity,omitempty"`
	ServiceType      *string                                                     `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	StoreFile        *string                                                     `json:"StoreFile,omitempty" xml:"StoreFile,omitempty"`
}

func (s DescribeClusterAttributeResponseBodyNodeListNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeListNode) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetDaemonList() *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList {
	return s.DaemonList
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetMemStore() *string {
	return s.MemStore
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetNodeDiskQuantity() *string {
	return s.NodeDiskQuantity
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetNodeDiskSize() *string {
	return s.NodeDiskSize
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetNodeDiskType() *string {
	return s.NodeDiskType
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetRegionQuantity() *string {
	return s.RegionQuantity
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) GetStoreFile() *string {
	return s.StoreFile
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetDaemonList(v *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.DaemonList = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetMemStore(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.MemStore = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeDiskQuantity(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeDiskQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeDiskSize(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeDiskSize = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeDiskType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeDiskType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeId(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeId = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeInstanceType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeInstanceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeStatus(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetNodeType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.NodeType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetRegionQuantity(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.RegionQuantity = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetServiceType(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.ServiceType = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) SetStoreFile(v string) *DescribeClusterAttributeResponseBodyNodeListNode {
	s.StoreFile = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNode) Validate() error {
	if s.DaemonList != nil {
		if err := s.DaemonList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterAttributeResponseBodyNodeListNodeDaemonList struct {
	Daemon []*DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon `json:"Daemon,omitempty" xml:"Daemon,omitempty" type:"Repeated"`
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) GetDaemon() []*DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon {
	return s.Daemon
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) SetDaemon(v []*DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList {
	s.Daemon = v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonList) Validate() error {
	if s.Daemon != nil {
		for _, item := range s.Daemon {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon struct {
	DaemonName   *string `json:"DaemonName,omitempty" xml:"DaemonName,omitempty"`
	DaemonStatus *string `json:"DaemonStatus,omitempty" xml:"DaemonStatus,omitempty"`
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) GetDaemonName() *string {
	return s.DaemonName
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) GetDaemonStatus() *string {
	return s.DaemonStatus
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) SetDaemonName(v string) *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon {
	s.DaemonName = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) SetDaemonStatus(v string) *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon {
	s.DaemonStatus = &v
	return s
}

func (s *DescribeClusterAttributeResponseBodyNodeListNodeDaemonListDaemon) Validate() error {
	return dara.Validate(s)
}
