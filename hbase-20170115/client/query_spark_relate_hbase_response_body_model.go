// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySparkRelateHBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v *QuerySparkRelateHBaseResponseBodyClusterList) *QuerySparkRelateHBaseResponseBody
	GetClusterList() *QuerySparkRelateHBaseResponseBodyClusterList
	SetRequestId(v string) *QuerySparkRelateHBaseResponseBody
	GetRequestId() *string
}

type QuerySparkRelateHBaseResponseBody struct {
	ClusterList *QuerySparkRelateHBaseResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySparkRelateHBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySparkRelateHBaseResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponseBody) GetClusterList() *QuerySparkRelateHBaseResponseBodyClusterList {
	return s.ClusterList
}

func (s *QuerySparkRelateHBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySparkRelateHBaseResponseBody) SetClusterList(v *QuerySparkRelateHBaseResponseBodyClusterList) *QuerySparkRelateHBaseResponseBody {
	s.ClusterList = v
	return s
}

func (s *QuerySparkRelateHBaseResponseBody) SetRequestId(v string) *QuerySparkRelateHBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBody) Validate() error {
	if s.ClusterList != nil {
		if err := s.ClusterList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySparkRelateHBaseResponseBodyClusterList struct {
	Cluster []*QuerySparkRelateHBaseResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QuerySparkRelateHBaseResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s QuerySparkRelateHBaseResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponseBodyClusterList) GetCluster() []*QuerySparkRelateHBaseResponseBodyClusterListCluster {
	return s.Cluster
}

func (s *QuerySparkRelateHBaseResponseBodyClusterList) SetCluster(v []*QuerySparkRelateHBaseResponseBodyClusterListCluster) *QuerySparkRelateHBaseResponseBodyClusterList {
	s.Cluster = v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterList) Validate() error {
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

type QuerySparkRelateHBaseResponseBodyClusterListCluster struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType          *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CoreDiskType         *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	CoreInstanceQuantity *int32  `json:"CoreInstanceQuantity,omitempty" xml:"CoreInstanceQuantity,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DbType               *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	ExpireTime           *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IsRelated            *bool   `json:"IsRelated,omitempty" xml:"IsRelated,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MainVersion          *string `json:"MainVersion,omitempty" xml:"MainVersion,omitempty"`
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Reason               *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QuerySparkRelateHBaseResponseBodyClusterListCluster) String() string {
	return dara.Prettify(s)
}

func (s QuerySparkRelateHBaseResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetClusterType() *string {
	return s.ClusterType
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetCoreInstanceQuantity() *int32 {
	return s.CoreInstanceQuantity
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetDbType() *string {
	return s.DbType
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetIsRelated() *bool {
	return s.IsRelated
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetMainVersion() *string {
	return s.MainVersion
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetNetType() *string {
	return s.NetType
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetPayType() *string {
	return s.PayType
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetReason() *string {
	return s.Reason
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetStatus() *string {
	return s.Status
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetUserId() *string {
	return s.UserId
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetClusterId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetClusterName(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetClusterType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetCoreDiskType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.CoreDiskType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetCoreInstanceQuantity(v int32) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.CoreInstanceQuantity = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetCreateTime(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.CreateTime = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetDbType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.DbType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetExpireTime(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ExpireTime = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetIsRelated(v bool) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.IsRelated = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetLockMode(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.LockMode = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetMainVersion(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.MainVersion = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetNetType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.NetType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetPayType(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.PayType = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetReason(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.Reason = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetRegionId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetStatus(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetUserId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.UserId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) SetZoneId(v string) *QuerySparkRelateHBaseResponseBodyClusterListCluster {
	s.ZoneId = &v
	return s
}

func (s *QuerySparkRelateHBaseResponseBodyClusterListCluster) Validate() error {
	return dara.Validate(s)
}
