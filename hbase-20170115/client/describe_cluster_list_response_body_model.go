// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v *DescribeClusterListResponseBodyClusterList) *DescribeClusterListResponseBody
	GetClusterList() *DescribeClusterListResponseBodyClusterList
	SetPageNumber(v int32) *DescribeClusterListResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeClusterListResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeClusterListResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeClusterListResponseBody
	GetTotalRecordCount() *int32
}

type DescribeClusterListResponseBody struct {
	ClusterList      *DescribeClusterListResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	PageNumber       *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                      `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalRecordCount *int32                                      `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeClusterListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBody) GetClusterList() *DescribeClusterListResponseBodyClusterList {
	return s.ClusterList
}

func (s *DescribeClusterListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterListResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeClusterListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterListResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeClusterListResponseBody) SetClusterList(v *DescribeClusterListResponseBodyClusterList) *DescribeClusterListResponseBody {
	s.ClusterList = v
	return s
}

func (s *DescribeClusterListResponseBody) SetPageNumber(v int32) *DescribeClusterListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterListResponseBody) SetPageRecordCount(v int32) *DescribeClusterListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeClusterListResponseBody) SetRequestId(v string) *DescribeClusterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterListResponseBody) SetTotalRecordCount(v int32) *DescribeClusterListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeClusterListResponseBody) Validate() error {
	if s.ClusterList != nil {
		if err := s.ClusterList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterListResponseBodyClusterList struct {
	Cluster []*DescribeClusterListResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s DescribeClusterListResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterList) GetCluster() []*DescribeClusterListResponseBodyClusterListCluster {
	return s.Cluster
}

func (s *DescribeClusterListResponseBodyClusterList) SetCluster(v []*DescribeClusterListResponseBodyClusterListCluster) *DescribeClusterListResponseBodyClusterList {
	s.Cluster = v
	return s
}

func (s *DescribeClusterListResponseBodyClusterList) Validate() error {
	if s.Cluster != nil {
		for _, item := range s.Cluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterListResponseBodyClusterListCluster struct {
	ClusterId            *string                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                                                `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string                                                `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CoreDiskSize         *string                                                `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	CoreDiskType         *string                                                `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32                                                 `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CreateTime           *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbType               *string                                                `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ExpireTime           *string                                                `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	LockMode             *string                                                `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string                                                `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	NetType              *string                                                `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PayType              *string                                                `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Reason               *string                                                `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RegionId             *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                 *DescribeClusterListResponseBodyClusterListClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserId               *string                                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcId                *string                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeClusterListResponseBodyClusterListCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetCoreDiskSize() *string {
	return s.CoreDiskSize
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetCoreInstanceQuantity() *int32 {
	return s.CoreInstanceQuantity
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetDbType() *string {
	return s.DbType
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetMainVersion() *string {
	return s.MainVersion
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetReason() *string {
	return s.Reason
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetStatus() *string {
	return s.Status
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetTags() *DescribeClusterListResponseBodyClusterListClusterTags {
	return s.Tags
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetUserId() *string {
	return s.UserId
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClusterListResponseBodyClusterListCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetClusterId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetClusterName(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetClusterType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCoreDiskSize(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCoreDiskType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCoreInstanceQuantity(v int32) *DescribeClusterListResponseBodyClusterListCluster {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetCreateTime(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetDbType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.DbType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetExpireTime(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetLockMode(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetMainVersion(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.MainVersion = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetNetType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.NetType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetPayType(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.PayType = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetReason(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.Reason = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetRegionId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetStatus(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetTags(v *DescribeClusterListResponseBodyClusterListClusterTags) *DescribeClusterListResponseBodyClusterListCluster {
	s.Tags = v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetUserId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.UserId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetVpcId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) SetZoneId(v string) *DescribeClusterListResponseBodyClusterListCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListCluster) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterListResponseBodyClusterListClusterTags struct {
	Tag []*DescribeClusterListResponseBodyClusterListClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClusterListResponseBodyClusterListClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterListClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterListClusterTags) GetTag() []*DescribeClusterListResponseBodyClusterListClusterTagsTag {
	return s.Tag
}

func (s *DescribeClusterListResponseBodyClusterListClusterTags) SetTag(v []*DescribeClusterListResponseBodyClusterListClusterTagsTag) *DescribeClusterListResponseBodyClusterListClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListClusterTags) Validate() error {
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

type DescribeClusterListResponseBodyClusterListClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeClusterListResponseBodyClusterListClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListResponseBodyClusterListClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponseBodyClusterListClusterTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeClusterListResponseBodyClusterListClusterTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeClusterListResponseBodyClusterListClusterTagsTag) SetKey(v string) *DescribeClusterListResponseBodyClusterListClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListClusterTagsTag) SetValue(v string) *DescribeClusterListResponseBodyClusterListClusterTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeClusterListResponseBodyClusterListClusterTagsTag) Validate() error {
	return dara.Validate(s)
}
