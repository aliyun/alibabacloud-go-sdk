// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersWithBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBClustersWithBackupsResponseBodyItems) *DescribeDBClustersWithBackupsResponseBody
	GetItems() *DescribeDBClustersWithBackupsResponseBodyItems
	SetPageNumber(v int32) *DescribeDBClustersWithBackupsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBClustersWithBackupsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBClustersWithBackupsResponseBody struct {
	Items *DescribeDBClustersWithBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of clusters returned per page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F8529AA2-522F-4B30-B80B-8F7D39******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBClustersWithBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetItems() *DescribeDBClustersWithBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClustersWithBackupsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetItems(v *DescribeDBClustersWithBackupsResponseBodyItems) *DescribeDBClustersWithBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetPageNumber(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetPageRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetRequestId(v string) *DescribeDBClustersWithBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetTotalRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClustersWithBackupsResponseBodyItems struct {
	DBCluster []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersWithBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBodyItems) GetDBCluster() []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	return s.DBCluster
}

func (s *DescribeDBClustersWithBackupsResponseBodyItems) SetDBCluster(v []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) *DescribeDBClustersWithBackupsResponseBodyItems {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItems) Validate() error {
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

type DescribeDBClustersWithBackupsResponseBodyItemsDBCluster struct {
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterStatus      *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBNodeClass          *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DeletedTime          *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	DeletionLock         *int32  `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	ExpireTime           *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	IsDeleted            *int32  `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetDeletionLock() *int32 {
	return s.DeletionLock
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDeletedTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DeletedTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDeletionLock(v int32) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetIsDeleted(v int32) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetVpcId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) Validate() error {
	return dara.Validate(s)
}
