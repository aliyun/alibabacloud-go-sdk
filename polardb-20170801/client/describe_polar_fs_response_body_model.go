// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribePolarFsResponseBodyItems) *DescribePolarFsResponseBody
	GetItems() *DescribePolarFsResponseBodyItems
	SetPageNumber(v int32) *DescribePolarFsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribePolarFsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribePolarFsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribePolarFsResponseBody
	GetTotalRecordCount() *int32
}

type DescribePolarFsResponseBody struct {
	Items *DescribePolarFsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 13
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 8CD1852F-8176-55E4-891C-4478DC00CC7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 16
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePolarFsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBody) GetItems() *DescribePolarFsResponseBodyItems {
	return s.Items
}

func (s *DescribePolarFsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarFsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribePolarFsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarFsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribePolarFsResponseBody) SetItems(v *DescribePolarFsResponseBodyItems) *DescribePolarFsResponseBody {
	s.Items = v
	return s
}

func (s *DescribePolarFsResponseBody) SetPageNumber(v int32) *DescribePolarFsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarFsResponseBody) SetPageRecordCount(v int32) *DescribePolarFsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePolarFsResponseBody) SetRequestId(v string) *DescribePolarFsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarFsResponseBody) SetTotalRecordCount(v int32) *DescribePolarFsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePolarFsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarFsResponseBodyItems struct {
	PolarFsPaths []*DescribePolarFsResponseBodyItemsPolarFsPaths `json:"PolarFsPaths,omitempty" xml:"PolarFsPaths,omitempty" type:"Repeated"`
}

func (s DescribePolarFsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBodyItems) GetPolarFsPaths() []*DescribePolarFsResponseBodyItemsPolarFsPaths {
	return s.PolarFsPaths
}

func (s *DescribePolarFsResponseBodyItems) SetPolarFsPaths(v []*DescribePolarFsResponseBodyItemsPolarFsPaths) *DescribePolarFsResponseBodyItems {
	s.PolarFsPaths = v
	return s
}

func (s *DescribePolarFsResponseBodyItems) Validate() error {
	if s.PolarFsPaths != nil {
		for _, item := range s.PolarFsPaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsResponseBodyItemsPolarFsPaths struct {
	AccelerateType             *string                                                          `json:"AccelerateType,omitempty" xml:"AccelerateType,omitempty"`
	AcceleratedStorageSpace    *string                                                          `json:"AcceleratedStorageSpace,omitempty" xml:"AcceleratedStorageSpace,omitempty"`
	AcceleratingEnable         *string                                                          `json:"AcceleratingEnable,omitempty" xml:"AcceleratingEnable,omitempty"`
	Bandwidth                  *int64                                                           `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Category                   *string                                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	CreateTime                 *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime                 *string                                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired                    *string                                                          `json:"Expired,omitempty" xml:"Expired,omitempty"`
	MountInfo                  *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo           `json:"MountInfo,omitempty" xml:"MountInfo,omitempty" type:"Struct"`
	MountedAIDBClusters        *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters `json:"MountedAIDBClusters,omitempty" xml:"MountedAIDBClusters,omitempty" type:"Struct"`
	PayType                    *string                                                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PolarFsInstanceDescription *string                                                          `json:"PolarFsInstanceDescription,omitempty" xml:"PolarFsInstanceDescription,omitempty"`
	PolarFsInstanceId          *string                                                          `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	PolarFsPath                *string                                                          `json:"PolarFsPath,omitempty" xml:"PolarFsPath,omitempty"`
	PolarFsStatus              *string                                                          `json:"PolarFsStatus,omitempty" xml:"PolarFsStatus,omitempty"`
	PolarFsType                *string                                                          `json:"PolarFsType,omitempty" xml:"PolarFsType,omitempty"`
	RegionId                   *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelativeDbClusterId        *string                                                          `json:"RelativeDbClusterId,omitempty" xml:"RelativeDbClusterId,omitempty"`
	SecurityGroupId            *string                                                          `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	StorageSpace               *int64                                                           `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	StorageType                *string                                                          `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags                       *DescribePolarFsResponseBodyItemsPolarFsPathsTags                `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VPCId                      *string                                                          `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId                  *string                                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                     *string                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePolarFsResponseBodyItemsPolarFsPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBodyItemsPolarFsPaths) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetAccelerateType() *string {
	return s.AccelerateType
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetAcceleratedStorageSpace() *string {
	return s.AcceleratedStorageSpace
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetAcceleratingEnable() *string {
	return s.AcceleratingEnable
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetCategory() *string {
	return s.Category
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetExpired() *string {
	return s.Expired
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetMountInfo() *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo {
	return s.MountInfo
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetMountedAIDBClusters() *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters {
	return s.MountedAIDBClusters
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetPayType() *string {
	return s.PayType
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetPolarFsInstanceDescription() *string {
	return s.PolarFsInstanceDescription
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetPolarFsPath() *string {
	return s.PolarFsPath
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetPolarFsStatus() *string {
	return s.PolarFsStatus
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetPolarFsType() *string {
	return s.PolarFsType
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetRelativeDbClusterId() *string {
	return s.RelativeDbClusterId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetTags() *DescribePolarFsResponseBodyItemsPolarFsPathsTags {
	return s.Tags
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetAccelerateType(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.AccelerateType = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetAcceleratedStorageSpace(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.AcceleratedStorageSpace = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetAcceleratingEnable(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.AcceleratingEnable = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetBandwidth(v int64) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.Bandwidth = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetCategory(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.Category = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetCreateTime(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.CreateTime = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetExpireTime(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.ExpireTime = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetExpired(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.Expired = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetMountInfo(v *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.MountInfo = v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetMountedAIDBClusters(v *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.MountedAIDBClusters = v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetPayType(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.PayType = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetPolarFsInstanceDescription(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.PolarFsInstanceDescription = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetPolarFsInstanceId(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetPolarFsPath(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.PolarFsPath = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetPolarFsStatus(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.PolarFsStatus = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetPolarFsType(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.PolarFsType = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetRegionId(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.RegionId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetRelativeDbClusterId(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.RelativeDbClusterId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetSecurityGroupId(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetStorageSpace(v int64) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.StorageSpace = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetStorageType(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.StorageType = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetTags(v *DescribePolarFsResponseBodyItemsPolarFsPathsTags) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.Tags = v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetVPCId(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.VPCId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetVSwitchId(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.VSwitchId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) SetZoneId(v string) *DescribePolarFsResponseBodyItemsPolarFsPaths {
	s.ZoneId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPaths) Validate() error {
	if s.MountInfo != nil {
		if err := s.MountInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MountedAIDBClusters != nil {
		if err := s.MountedAIDBClusters.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo struct {
	PolarDbProxy   *string `json:"PolarDbProxy,omitempty" xml:"PolarDbProxy,omitempty"`
	PolarFsCluster *string `json:"PolarFsCluster,omitempty" xml:"PolarFsCluster,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) GetPolarDbProxy() *string {
	return s.PolarDbProxy
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) GetPolarFsCluster() *string {
	return s.PolarFsCluster
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) GetToken() *string {
	return s.Token
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) SetPolarDbProxy(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo {
	s.PolarDbProxy = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) SetPolarFsCluster(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo {
	s.PolarFsCluster = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) SetToken(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo {
	s.Token = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters struct {
	MountedAIDBClusters []*DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters `json:"MountedAIDBClusters,omitempty" xml:"MountedAIDBClusters,omitempty" type:"Repeated"`
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters) GetMountedAIDBClusters() []*DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters {
	return s.MountedAIDBClusters
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters) SetMountedAIDBClusters(v []*DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters {
	s.MountedAIDBClusters = v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClusters) Validate() error {
	if s.MountedAIDBClusters != nil {
		for _, item := range s.MountedAIDBClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	MountDir    *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	MountStatus *string `json:"MountStatus,omitempty" xml:"MountStatus,omitempty"`
	MountedTime *string `json:"MountedTime,omitempty" xml:"MountedTime,omitempty"`
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) GetMountDir() *string {
	return s.MountDir
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) GetMountStatus() *string {
	return s.MountStatus
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) GetMountedTime() *string {
	return s.MountedTime
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) SetDBClusterId(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) SetMountDir(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters {
	s.MountDir = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) SetMountStatus(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters {
	s.MountStatus = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) SetMountedTime(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters {
	s.MountedTime = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsMountedAIDBClustersMountedAIDBClusters) Validate() error {
	return dara.Validate(s)
}

type DescribePolarFsResponseBodyItemsPolarFsPathsTags struct {
	Tag []*DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsTags) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTags) GetTag() []*DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag {
	return s.Tag
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTags) SetTag(v []*DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) *DescribePolarFsResponseBodyItemsPolarFsPathsTags {
	s.Tag = v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTags) Validate() error {
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

type DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag struct {
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) GoString() string {
	return s.String()
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) SetKey(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag {
	s.Key = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) SetValue(v string) *DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag {
	s.Value = &v
	return s
}

func (s *DescribePolarFsResponseBodyItemsPolarFsPathsTagsTag) Validate() error {
	return dara.Validate(s)
}
