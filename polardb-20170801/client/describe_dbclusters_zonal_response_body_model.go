// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBClustersZonalResponseBodyItems) *DescribeDBClustersZonalResponseBody
	GetItems() []*DescribeDBClustersZonalResponseBodyItems
	SetMaxResults(v int32) *DescribeDBClustersZonalResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDBClustersZonalResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDBClustersZonalResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBClustersZonalResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBClustersZonalResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBClustersZonalResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBClustersZonalResponseBody struct {
	Items []*DescribeDBClustersZonalResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 12
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 16
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBClustersZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersZonalResponseBody) GetItems() []*DescribeDBClustersZonalResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClustersZonalResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDBClustersZonalResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDBClustersZonalResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersZonalResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBClustersZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClustersZonalResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBClustersZonalResponseBody) SetItems(v []*DescribeDBClustersZonalResponseBodyItems) *DescribeDBClustersZonalResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClustersZonalResponseBody) SetMaxResults(v int32) *DescribeDBClustersZonalResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBody) SetNextToken(v string) *DescribeDBClustersZonalResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBody) SetPageNumber(v int32) *DescribeDBClustersZonalResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBody) SetPageRecordCount(v int32) *DescribeDBClustersZonalResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBody) SetRequestId(v string) *DescribeDBClustersZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBody) SetTotalRecordCount(v int32) *DescribeDBClustersZonalResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersZonalResponseBodyItems struct {
	// example:
	//
	// SearchNode
	AiType *string `json:"AiType,omitempty" xml:"AiType,omitempty"`
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// cn-beijing
	CentralControlRegionId *string `json:"CentralControlRegionId,omitempty" xml:"CentralControlRegionId,omitempty"`
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// example:
	//
	// 1
	CpuCores *string `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// GDN-1
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// example:
	//
	// polar.mysql.g1.tiny.c
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeNumber *int32 `json:"DBNodeNumber,omitempty" xml:"DBNodeNumber,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// sg-singapore-9
	ENSRegionId *string `json:"ENSRegionId,omitempty" xml:"ENSRegionId,omitempty"`
	// example:
	//
	// 2022-09-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// example:
	//
	// 10
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// example:
	//
	// essdautopl
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// 3009413120
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// example:
	//
	// Exclusive
	SubCategory *string                                         `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
	Tags        []*DescribeDBClustersZonalResponseBodyItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-****************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-***************
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClustersZonalResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersZonalResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetAiType() *string {
	return s.AiType
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetCentralControlRegionId() *string {
	return s.CentralControlRegionId
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetCpuCores() *string {
	return s.CpuCores
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetDBNodeNumber() *int32 {
	return s.DBNodeNumber
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetENSRegionId() *string {
	return s.ENSRegionId
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetStrictConsistency() *string {
	return s.StrictConsistency
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetSubCategory() *string {
	return s.SubCategory
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetTags() []*DescribeDBClustersZonalResponseBodyItemsTags {
	return s.Tags
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDBClustersZonalResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetAiType(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.AiType = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetCategory(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetCentralControlRegionId(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.CentralControlRegionId = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetCloudProvider(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.CloudProvider = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetCpuCores(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetCreateTime(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetDBClusterDescription(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetDBClusterId(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetDBClusterStatus(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetDBNodeClass(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetDBNodeNumber(v int32) *DescribeDBClustersZonalResponseBodyItems {
	s.DBNodeNumber = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetDBType(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetDBVersion(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetENSRegionId(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.ENSRegionId = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetExpireTime(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetExpired(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetLockMode(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetPayType(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetServerlessType(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetStorageSpace(v int64) *DescribeDBClustersZonalResponseBodyItems {
	s.StorageSpace = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetStorageType(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetStorageUsed(v int64) *DescribeDBClustersZonalResponseBodyItems {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetStrictConsistency(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.StrictConsistency = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetSubCategory(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.SubCategory = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetTags(v []*DescribeDBClustersZonalResponseBodyItemsTags) *DescribeDBClustersZonalResponseBodyItems {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetVpcId(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetVswitchId(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.VswitchId = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) SetZoneId(v string) *DescribeDBClustersZonalResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItems) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClustersZonalResponseBodyItemsTags struct {
	// example:
	//
	// MySQL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 5.6
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersZonalResponseBodyItemsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersZonalResponseBodyItemsTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersZonalResponseBodyItemsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersZonalResponseBodyItemsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersZonalResponseBodyItemsTags) SetKey(v string) *DescribeDBClustersZonalResponseBodyItemsTags {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItemsTags) SetValue(v string) *DescribeDBClustersZonalResponseBodyItemsTags {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersZonalResponseBodyItemsTags) Validate() error {
	return dara.Validate(s)
}
