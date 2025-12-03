// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *DescribeClusterModelResponseBody
	GetAutoRenew() *string
	SetBackupStatus(v string) *DescribeClusterModelResponseBody
	GetBackupStatus() *string
	SetClusterId(v string) *DescribeClusterModelResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeClusterModelResponseBody
	GetClusterName() *string
	SetClusterType(v string) *DescribeClusterModelResponseBody
	GetClusterType() *string
	SetColdStorageStatus(v string) *DescribeClusterModelResponseBody
	GetColdStorageStatus() *string
	SetCoreDiskQuantity(v int32) *DescribeClusterModelResponseBody
	GetCoreDiskQuantity() *int32
	SetCoreDiskSize(v string) *DescribeClusterModelResponseBody
	GetCoreDiskSize() *string
	SetCoreDiskType(v string) *DescribeClusterModelResponseBody
	GetCoreDiskType() *string
	SetCoreInstanceQuantity(v int32) *DescribeClusterModelResponseBody
	GetCoreInstanceQuantity() *int32
	SetCoreInstanceType(v string) *DescribeClusterModelResponseBody
	GetCoreInstanceType() *string
	SetCreateTime(v string) *DescribeClusterModelResponseBody
	GetCreateTime() *string
	SetDbType(v string) *DescribeClusterModelResponseBody
	GetDbType() *string
	SetExpireTime(v string) *DescribeClusterModelResponseBody
	GetExpireTime() *string
	SetHaType(v string) *DescribeClusterModelResponseBody
	GetHaType() *string
	SetHasUser(v string) *DescribeClusterModelResponseBody
	GetHasUser() *string
	SetIsMultimod(v string) *DescribeClusterModelResponseBody
	GetIsMultimod() *string
	SetLockMode(v string) *DescribeClusterModelResponseBody
	GetLockMode() *string
	SetMainVersion(v string) *DescribeClusterModelResponseBody
	GetMainVersion() *string
	SetMasterDiskSize(v int32) *DescribeClusterModelResponseBody
	GetMasterDiskSize() *int32
	SetMasterDiskType(v string) *DescribeClusterModelResponseBody
	GetMasterDiskType() *string
	SetMasterInstanceType(v string) *DescribeClusterModelResponseBody
	GetMasterInstanceType() *string
	SetMinorVersion(v string) *DescribeClusterModelResponseBody
	GetMinorVersion() *string
	SetPayType(v string) *DescribeClusterModelResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeClusterModelResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeClusterModelResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeClusterModelResponseBody
	GetStatus() *string
	SetTags(v *DescribeClusterModelResponseBodyTags) *DescribeClusterModelResponseBody
	GetTags() *DescribeClusterModelResponseBodyTags
	SetUpdateStatus(v string) *DescribeClusterModelResponseBody
	GetUpdateStatus() *string
	SetZoneId(v string) *DescribeClusterModelResponseBody
	GetZoneId() *string
}

type DescribeClusterModelResponseBody struct {
	AutoRenew            *string                               `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BackupStatus         *string                               `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	ClusterId            *string                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                               `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string                               `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ColdStorageStatus    *string                               `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	CoreDiskQuantity     *int32                                `json:"CoreDiskQuantity,omitempty" xml:"CoreDiskQuantity,omitempty"`
	CoreDiskSize         *string                               `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string                               `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32                                `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CoreInstanceType     *string                               `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	CreateTime           *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbType               *string                               `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ExpireTime           *string                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	HaType               *string                               `json:"HaType,omitempty" xml:"HaType,omitempty"`
	HasUser              *string                               `json:"HasUser,omitempty" xml:"HasUser,omitempty"`
	IsMultimod           *string                               `json:"IsMultimod,omitempty" xml:"IsMultimod,omitempty"`
	LockMode             *string                               `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string                               `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	MasterDiskSize       *int32                                `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	MasterDiskType       *string                               `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	MasterInstanceType   *string                               `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	MinorVersion         *string                               `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	PayType              *string                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status               *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                 *DescribeClusterModelResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateStatus         *string                               `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
	ZoneId               *string                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterModelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponseBody) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *DescribeClusterModelResponseBody) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeClusterModelResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterModelResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeClusterModelResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClusterModelResponseBody) GetColdStorageStatus() *string {
	return s.ColdStorageStatus
}

func (s *DescribeClusterModelResponseBody) GetCoreDiskQuantity() *int32 {
	return s.CoreDiskQuantity
}

func (s *DescribeClusterModelResponseBody) GetCoreDiskSize() *string {
	return s.CoreDiskSize
}

func (s *DescribeClusterModelResponseBody) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *DescribeClusterModelResponseBody) GetCoreInstanceQuantity() *int32 {
	return s.CoreInstanceQuantity
}

func (s *DescribeClusterModelResponseBody) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *DescribeClusterModelResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeClusterModelResponseBody) GetDbType() *string {
	return s.DbType
}

func (s *DescribeClusterModelResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeClusterModelResponseBody) GetHaType() *string {
	return s.HaType
}

func (s *DescribeClusterModelResponseBody) GetHasUser() *string {
	return s.HasUser
}

func (s *DescribeClusterModelResponseBody) GetIsMultimod() *string {
	return s.IsMultimod
}

func (s *DescribeClusterModelResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeClusterModelResponseBody) GetMainVersion() *string {
	return s.MainVersion
}

func (s *DescribeClusterModelResponseBody) GetMasterDiskSize() *int32 {
	return s.MasterDiskSize
}

func (s *DescribeClusterModelResponseBody) GetMasterDiskType() *string {
	return s.MasterDiskType
}

func (s *DescribeClusterModelResponseBody) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *DescribeClusterModelResponseBody) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeClusterModelResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeClusterModelResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterModelResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeClusterModelResponseBody) GetTags() *DescribeClusterModelResponseBodyTags {
	return s.Tags
}

func (s *DescribeClusterModelResponseBody) GetUpdateStatus() *string {
	return s.UpdateStatus
}

func (s *DescribeClusterModelResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterModelResponseBody) SetAutoRenew(v string) *DescribeClusterModelResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetBackupStatus(v string) *DescribeClusterModelResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetClusterId(v string) *DescribeClusterModelResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetClusterName(v string) *DescribeClusterModelResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetClusterType(v string) *DescribeClusterModelResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetColdStorageStatus(v string) *DescribeClusterModelResponseBody {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreDiskQuantity(v int32) *DescribeClusterModelResponseBody {
	s.CoreDiskQuantity = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreDiskSize(v string) *DescribeClusterModelResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreDiskType(v string) *DescribeClusterModelResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreInstanceQuantity(v int32) *DescribeClusterModelResponseBody {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCoreInstanceType(v string) *DescribeClusterModelResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetCreateTime(v string) *DescribeClusterModelResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetDbType(v string) *DescribeClusterModelResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetExpireTime(v string) *DescribeClusterModelResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetHaType(v string) *DescribeClusterModelResponseBody {
	s.HaType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetHasUser(v string) *DescribeClusterModelResponseBody {
	s.HasUser = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetIsMultimod(v string) *DescribeClusterModelResponseBody {
	s.IsMultimod = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetLockMode(v string) *DescribeClusterModelResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMainVersion(v string) *DescribeClusterModelResponseBody {
	s.MainVersion = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMasterDiskSize(v int32) *DescribeClusterModelResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMasterDiskType(v string) *DescribeClusterModelResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMasterInstanceType(v string) *DescribeClusterModelResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetMinorVersion(v string) *DescribeClusterModelResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetPayType(v string) *DescribeClusterModelResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetRegionId(v string) *DescribeClusterModelResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetRequestId(v string) *DescribeClusterModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetStatus(v string) *DescribeClusterModelResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetTags(v *DescribeClusterModelResponseBodyTags) *DescribeClusterModelResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeClusterModelResponseBody) SetUpdateStatus(v string) *DescribeClusterModelResponseBody {
	s.UpdateStatus = &v
	return s
}

func (s *DescribeClusterModelResponseBody) SetZoneId(v string) *DescribeClusterModelResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterModelResponseBody) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterModelResponseBodyTags struct {
	Tag []*DescribeClusterModelResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClusterModelResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterModelResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponseBodyTags) GetTag() []*DescribeClusterModelResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeClusterModelResponseBodyTags) SetTag(v []*DescribeClusterModelResponseBodyTagsTag) *DescribeClusterModelResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeClusterModelResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterModelResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClusterModelResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterModelResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterModelResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeClusterModelResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeClusterModelResponseBodyTagsTag) SetKey(v string) *DescribeClusterModelResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeClusterModelResponseBodyTagsTag) SetValue(v string) *DescribeClusterModelResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeClusterModelResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
