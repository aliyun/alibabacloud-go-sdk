// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAgenticDBClustersResponseBodyItems) *DescribeAgenticDBClustersResponseBody
	GetItems() []*DescribeAgenticDBClustersResponseBodyItems
	SetMaxResults(v int32) *DescribeAgenticDBClustersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAgenticDBClustersResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeAgenticDBClustersResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeAgenticDBClustersResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeAgenticDBClustersResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeAgenticDBClustersResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAgenticDBClustersResponseBody struct {
	// The cluster list.
	Items []*DescribeAgenticDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries to return. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call. If there is no next query, do not pass this parameter.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of clusters on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAgenticDBClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBClustersResponseBody) GetItems() []*DescribeAgenticDBClustersResponseBodyItems {
	return s.Items
}

func (s *DescribeAgenticDBClustersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAgenticDBClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAgenticDBClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBClustersResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAgenticDBClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBClustersResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAgenticDBClustersResponseBody) SetItems(v []*DescribeAgenticDBClustersResponseBodyItems) *DescribeAgenticDBClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAgenticDBClustersResponseBody) SetMaxResults(v int32) *DescribeAgenticDBClustersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBody) SetNextToken(v string) *DescribeAgenticDBClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBody) SetPageNumber(v int32) *DescribeAgenticDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBody) SetPageRecordCount(v int32) *DescribeAgenticDBClustersResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBody) SetRequestId(v string) *DescribeAgenticDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBody) SetTotalRecordCount(v int32) *DescribeAgenticDBClustersResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBody) Validate() error {
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

type DescribeAgenticDBClustersResponseBodyItems struct {
	// The Agentic cluster description.
	//
	// example:
	//
	// pagc-**************
	AgenticDbClusterDescription *string `json:"AgenticDbClusterDescription,omitempty" xml:"AgenticDbClusterDescription,omitempty"`
	// The Agentic cluster ID.
	//
	// example:
	//
	// pagc-**************
	AgenticDbClusterId *string `json:"AgenticDbClusterId,omitempty" xml:"AgenticDbClusterId,omitempty"`
	// The cluster edition. Valid values:
	//
	// - **Normal**: Cluster Edition
	//
	// - **Basic**: Single Node Edition
	//
	// - **Archive**: X-Engine Edition
	//
	// - **NormalMultimaster**: Multi-master Cluster (Database/Table)
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// ocpx
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-xxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The expiration time of the cluster.
	//
	// > A specific value is returned only for clusters whose billing method is **Prepaid*	- (subscription). An empty value is returned for **Postpaid*	- (pay-as-you-go) clusters.
	//
	// example:
	//
	// 2025-06-25T09:37:10Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// > This parameter is returned only for clusters whose billing method is **Prepaid*	- (subscription).
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The lock status of the cluster. Valid values:
	//
	// - **Unlock**: Normal.
	//
	// - **ManualLock**: Manually locked.
	//
	// - **LockByExpiration**: Automatically locked due to cluster expiration.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum value: 1 to 32 PCUs. Unit: PCU.
	//
	// example:
	//
	// 9
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum value: 0 to 32 PCUs. The minimum value must be less than or equal to the maximum value. Unit: PCU.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The serverless type. A value of **AgileServerless*	- indicates that the cluster is a serverless cluster.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The storage type. Valid values:
	//
	// 	- **essdpl0**
	//
	// 	- **essdpl1**
	//
	// 	- **essdpl2**
	//
	// 	- **essdpl3**
	//
	// example:
	//
	// city_redundancy
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The storage usage of the cluster. Unit: bytes.
	//
	// example:
	//
	// 3009413120
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// The tag key. You can filter the cluster list by tag. You can specify up to 20 tag pairs. The number n for each tag pair must be unique and must be a consecutive integer starting from 1. The value of Tag.n.Key corresponds to Tag.n.Value.
	Tags []*DescribeAgenticDBClustersResponseBodyItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) in which the endpoint resides.
	//
	// example:
	//
	// vpc-****************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-**************
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAgenticDBClustersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetAgenticDbClusterDescription() *string {
	return s.AgenticDbClusterDescription
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetAgenticDbClusterId() *string {
	return s.AgenticDbClusterId
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetDBType() *string {
	return s.DBType
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetExpired() *string {
	return s.Expired
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetPayType() *string {
	return s.PayType
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetTags() []*DescribeAgenticDBClustersResponseBodyItemsTags {
	return s.Tags
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeAgenticDBClustersResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetAgenticDbClusterDescription(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.AgenticDbClusterDescription = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetAgenticDbClusterId(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.AgenticDbClusterId = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetCategory(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetCreateTime(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetDBClusterDescription(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetDBClusterId(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetDBClusterStatus(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetDBType(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetDBVersion(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetExpireTime(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetExpired(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.Expired = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetLockMode(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.LockMode = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetPayType(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.PayType = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetRegionId(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetScaleMax(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.ScaleMax = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetScaleMin(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.ScaleMin = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetServerlessType(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.ServerlessType = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetStorageType(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetStorageUsed(v int64) *DescribeAgenticDBClustersResponseBodyItems {
	s.StorageUsed = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetTags(v []*DescribeAgenticDBClustersResponseBodyItemsTags) *DescribeAgenticDBClustersResponseBodyItems {
	s.Tags = v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetVpcId(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetVswitchId(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.VswitchId = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) SetZoneId(v string) *DescribeAgenticDBClustersResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItems) Validate() error {
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

type DescribeAgenticDBClustersResponseBodyItemsTags struct {
	// The tag key.
	//
	// example:
	//
	// MySQL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValueData
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAgenticDBClustersResponseBodyItemsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBClustersResponseBodyItemsTags) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBClustersResponseBodyItemsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAgenticDBClustersResponseBodyItemsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAgenticDBClustersResponseBodyItemsTags) SetKey(v string) *DescribeAgenticDBClustersResponseBodyItemsTags {
	s.Key = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItemsTags) SetValue(v string) *DescribeAgenticDBClustersResponseBodyItemsTags {
	s.Value = &v
	return s
}

func (s *DescribeAgenticDBClustersResponseBodyItemsTags) Validate() error {
	return dara.Validate(s)
}
