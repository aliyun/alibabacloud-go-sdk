// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeAIDBClustersResponseBodyItems) *DescribeAIDBClustersResponseBody
	GetItems() *DescribeAIDBClustersResponseBodyItems
	SetPageNumber(v int32) *DescribeAIDBClustersResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeAIDBClustersResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeAIDBClustersResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeAIDBClustersResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAIDBClustersResponseBody struct {
	Items *DescribeAIDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 7
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9D6CE7C6-1C52-5BF6-B3D7-10977D44542C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAIDBClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersResponseBody) GetItems() *DescribeAIDBClustersResponseBodyItems {
	return s.Items
}

func (s *DescribeAIDBClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAIDBClustersResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAIDBClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClustersResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAIDBClustersResponseBody) SetItems(v *DescribeAIDBClustersResponseBodyItems) *DescribeAIDBClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAIDBClustersResponseBody) SetPageNumber(v int32) *DescribeAIDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClustersResponseBody) SetPageRecordCount(v int32) *DescribeAIDBClustersResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAIDBClustersResponseBody) SetRequestId(v string) *DescribeAIDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClustersResponseBody) SetTotalRecordCount(v int32) *DescribeAIDBClustersResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAIDBClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClustersResponseBodyItems struct {
	DBCluster []*DescribeAIDBClustersResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeAIDBClustersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersResponseBodyItems) GetDBCluster() []*DescribeAIDBClustersResponseBodyItemsDBCluster {
	return s.DBCluster
}

func (s *DescribeAIDBClustersResponseBodyItems) SetDBCluster(v []*DescribeAIDBClustersResponseBodyItemsDBCluster) *DescribeAIDBClustersResponseBodyItems {
	s.DBCluster = v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClustersResponseBodyItemsDBCluster struct {
	// example:
	//
	// vnode
	AiNodeType *string `json:"AiNodeType,omitempty" xml:"AiNodeType,omitempty"`
	// example:
	//
	// 2021-09-13T15:45:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ocpx
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Creating
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// example:
	//
	// polar.pg.g8.8xlarge.gu30
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// polardb_ai
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 2028-09-01T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// False
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// xxxxxxxxxxxx
	KubeClusterId *string `json:"KubeClusterId,omitempty" xml:"KubeClusterId,omitempty"`
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// container
	RunType *string `json:"RunType,omitempty" xml:"RunType,omitempty"`
	// example:
	//
	// 10
	StorageSpace *int32 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// example:
	//
	// essdpl0
	StorageType *string                                             `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags        *DescribeAIDBClustersResponseBodyItemsDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// vpc-***************
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAIDBClustersResponseBodyItemsDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetAiNodeType() *string {
	return s.AiNodeType
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetDBType() *string {
	return s.DBType
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetKubeClusterId() *string {
	return s.KubeClusterId
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetRunType() *string {
	return s.RunType
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetStorageSpace() *int32 {
	return s.StorageSpace
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetTags() *DescribeAIDBClustersResponseBodyItemsDBClusterTags {
	return s.Tags
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetAiNodeType(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.AiNodeType = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetDBType(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetExpired(v bool) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetKubeClusterId(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.KubeClusterId = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetRunType(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.RunType = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetStorageSpace(v int32) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.StorageSpace = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetStorageType(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.StorageType = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeAIDBClustersResponseBodyItemsDBClusterTags) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetVpcId(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetVswitchId(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.VswitchId = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeAIDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBCluster) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClustersResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAIDBClustersResponseBodyItemsDBClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTags) GetTag() []*DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag {
	return s.Tag
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTags) SetTag(v []*DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) *DescribeAIDBClustersResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTags) Validate() error {
	return dara.Validate(s)
}

type DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeAIDBClustersResponseBodyItemsDBClusterTagsTag) Validate() error {
	return dara.Validate(s)
}
