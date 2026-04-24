// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusters(v *DescribeDBClustersResponseBodyDBClusters) *DescribeDBClustersResponseBody
	GetDBClusters() *DescribeDBClustersResponseBodyDBClusters
	SetPageNumber(v int32) *DescribeDBClustersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDBClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBClustersResponseBody
	GetTotalCount() *int32
}

type DescribeDBClustersResponseBody struct {
	DBClusters *DescribeDBClustersResponseBodyDBClusters `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Struct"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) GetDBClusters() *DescribeDBClustersResponseBodyDBClusters {
	return s.DBClusters
}

func (s *DescribeDBClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBClustersResponseBody) SetDBClusters(v *DescribeDBClustersResponseBodyDBClusters) *DescribeDBClustersResponseBody {
	s.DBClusters = v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageSize(v int32) *DescribeDBClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetTotalCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBClustersResponseBody) Validate() error {
	if s.DBClusters != nil {
		if err := s.DBClusters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyDBClusters struct {
	DBCluster []*DescribeDBClustersResponseBodyDBClustersDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyDBClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClusters) GetDBCluster() []*DescribeDBClustersResponseBodyDBClustersDBCluster {
	return s.DBCluster
}

func (s *DescribeDBClustersResponseBodyDBClusters) SetDBCluster(v []*DescribeDBClustersResponseBodyDBClustersDBCluster) *DescribeDBClustersResponseBodyDBClusters {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClusters) Validate() error {
	if s.DBCluster != nil {
		for _, item := range s.DBCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersResponseBodyDBClustersDBCluster struct {
	AliUid                      *string                                                          `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid                         *string                                                          `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Category                    *string                                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode               *string                                                          `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnectionString            *string                                                          `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	ControlVersion              *string                                                          `json:"ControlVersion,omitempty" xml:"ControlVersion,omitempty"`
	CreateTime                  *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterDescription        *string                                                          `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId                 *string                                                          `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType        *string                                                          `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus             *string                                                          `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBNodeClass                 *string                                                          `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeCount                 *int64                                                           `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodeStorage               *int64                                                           `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	DbVersion                   *string                                                          `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	ExpireTime                  *string                                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExtStorageSize              *int32                                                           `json:"ExtStorageSize,omitempty" xml:"ExtStorageSize,omitempty"`
	ExtStorageType              *string                                                          `json:"ExtStorageType,omitempty" xml:"ExtStorageType,omitempty"`
	IsExpired                   *string                                                          `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	LockMode                    *string                                                          `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason                  *string                                                          `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	PayType                     *string                                                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Port                        *int32                                                           `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId                    *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScaleOutDisableWriteWindows *string                                                          `json:"ScaleOutDisableWriteWindows,omitempty" xml:"ScaleOutDisableWriteWindows,omitempty"`
	ScaleOutStatus              *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus `json:"ScaleOutStatus,omitempty" xml:"ScaleOutStatus,omitempty" type:"Struct"`
	StorageType                 *string                                                          `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags                        *DescribeDBClustersResponseBodyDBClustersDBClusterTags           `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId                   *string                                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcCloudInstanceId          *string                                                          `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	VpcId                       *string                                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                      *string                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetBid() *string {
	return s.Bid
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetControlVersion() *string {
	return s.ControlVersion
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBNodeCount() *int64 {
	return s.DBNodeCount
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDBNodeStorage() *int64 {
	return s.DBNodeStorage
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetExtStorageSize() *int32 {
	return s.ExtStorageSize
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetExtStorageType() *string {
	return s.ExtStorageType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetIsExpired() *string {
	return s.IsExpired
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetScaleOutDisableWriteWindows() *string {
	return s.ScaleOutDisableWriteWindows
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetScaleOutStatus() *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	return s.ScaleOutStatus
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetTags() *DescribeDBClustersResponseBodyDBClustersDBClusterTags {
	return s.Tags
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetAliUid(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.AliUid = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetBid(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Bid = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCategory(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCommodityCode(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetConnectionString(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetControlVersion(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ControlVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeCount(v int64) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDBNodeStorage(v int64) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetDbVersion(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.DbVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExtStorageSize(v int32) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExtStorageSize = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetExtStorageType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ExtStorageType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetIsExpired(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.IsExpired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetLockReason(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetPort(v int32) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetScaleOutDisableWriteWindows(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ScaleOutDisableWriteWindows = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetScaleOutStatus(v *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ScaleOutStatus = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetStorageType(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetTags(v *DescribeDBClustersResponseBodyDBClustersDBClusterTags) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVSwitchId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVpcCloudInstanceId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetVpcId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyDBClustersDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBCluster) Validate() error {
	if s.ScaleOutStatus != nil {
		if err := s.ScaleOutStatus.Validate(); err != nil {
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

type DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus struct {
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Ratio    *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) GetRatio() *string {
	return s.Ratio
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) SetProgress(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) SetRatio(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus {
	s.Ratio = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterScaleOutStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClustersResponseBodyDBClustersDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTags) GetTag() []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	return s.Tag
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) *DescribeDBClustersResponseBodyDBClustersDBClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTags) Validate() error {
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

type DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersResponseBodyDBClustersDBClusterTagsTag) Validate() error {
	return dara.Validate(s)
}
