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
	// The list of clusters.
	Items []*DescribeDBClustersZonalResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries returned for the current request. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. If more results are available, this parameter is returned. To retrieve the next page, include this token in your next request. If all results have been returned, this parameter is not returned.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 12
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of clusters on the current page.
	//
	// example:
	//
	// 5
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
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
	// The AI node type. Valid values:
	//
	// - SearchNode: search node
	//
	// - DLNode: AI node
	//
	// example:
	//
	// SearchNode
	AiType *string `json:"AiType,omitempty" xml:"AiType,omitempty"`
	// The Cluster Edition. The following editions are supported:
	//
	// - Normal: Cluster Edition
	//
	// - Basic: single node
	//
	// - Archive: X-Engine
	//
	// - NormalMultimaster: Multi-master Cluster (Database/Table)
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the central control region.
	//
	// example:
	//
	// cn-beijing
	CentralControlRegionId *string `json:"CentralControlRegionId,omitempty" xml:"CentralControlRegionId,omitempty"`
	// The cloud service provider.
	//
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 1
	CpuCores *string `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// GDN-1
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The status of the cluster.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// polar.mysql.g1.tiny.c
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 2
	DBNodeNumber *int32 `json:"DBNodeNumber,omitempty" xml:"DBNodeNumber,omitempty"`
	// The database type.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database version.
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The ENS region ID.
	//
	// example:
	//
	// sg-singapore-9
	ENSRegionId *string `json:"ENSRegionId,omitempty" xml:"ENSRegionId,omitempty"`
	// The expiration time of the cluster.
	//
	// > This parameter is returned only for **Prepaid*	- (subscription) clusters. For **Postpaid*	- (pay-as-you-go) clusters, this parameter is empty.
	//
	// example:
	//
	// 2022-09-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The lock state of the cluster. Valid values:
	//
	// - Unlock: Normal.
	//
	// - ManualLock: The cluster is manually locked.
	//
	// - LockByExpiration: The cluster is automatically locked upon expiration.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The billing method. Valid values:
	//
	// - Postpaid: pay-as-you-go.
	//
	// - Prepaid: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The Serverless type. \\`**AgileServerless**\\` indicates that the cluster is a Serverless cluster. An empty value indicates that the cluster is a common cluster.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The storage capacity of the instance.
	//
	// example:
	//
	// 10
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage class of the Standard Edition cluster. Valid values:
	//
	// - essdpl0
	//
	// - essdpl1
	//
	// - essdpl2
	//
	// - essdpl3
	//
	// - essdautopl
	//
	// example:
	//
	// essdautopl
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The used storage space of the cluster. Unit: bytes.
	//
	// example:
	//
	// 3009413120
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// Indicates whether strong consistency is enabled for data across multiple zones. Valid values:
	//
	// - **ON**: Strong consistency is enabled. This applies to Standard Edition clusters that are deployed in three zones.
	//
	// - **OFF**: Strong consistency is not enabled.
	//
	// example:
	//
	// ON
	StrictConsistency *string `json:"StrictConsistency,omitempty" xml:"StrictConsistency,omitempty"`
	// The specification type of the compute node. Valid values:
	//
	// - **Exclusive**: Dedicated
	//
	// - **General**: General-purpose
	//
	// example:
	//
	// Exclusive
	SubCategory *string `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
	// The list of tags.
	Tags []*DescribeDBClustersZonalResponseBodyItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-****************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The virtual switch ID.
	//
	// example:
	//
	// vsw-***************
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
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
