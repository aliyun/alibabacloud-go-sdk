// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopRiskyResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeTopRiskyResourcesResponseBodyData) *DescribeTopRiskyResourcesResponseBody
	GetData() *DescribeTopRiskyResourcesResponseBodyData
	SetRequestId(v string) *DescribeTopRiskyResourcesResponseBody
	GetRequestId() *string
}

type DescribeTopRiskyResourcesResponseBody struct {
	// The data returned in the response.
	Data *DescribeTopRiskyResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 34081B20-C4C0-514F-93F6-8EEC3D1A587E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTopRiskyResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponseBody) GetData() *DescribeTopRiskyResourcesResponseBodyData {
	return s.Data
}

func (s *DescribeTopRiskyResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTopRiskyResourcesResponseBody) SetData(v *DescribeTopRiskyResourcesResponseBodyData) *DescribeTopRiskyResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBody) SetRequestId(v string) *DescribeTopRiskyResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTopRiskyResourcesResponseBodyData struct {
	// A list of resource objects.
	Content []*DescribeTopRiskyResourcesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If this parameter is not returned, it indicates that all results have been returned.
	//
	// example:
	//
	// e557bc9a65fe22cb5e2a3b240f06b0de
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of matching entries. This parameter is optional and might not be returned in the response.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTopRiskyResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponseBodyData) GetContent() []*DescribeTopRiskyResourcesResponseBodyDataContent {
	return s.Content
}

func (s *DescribeTopRiskyResourcesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTopRiskyResourcesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTopRiskyResourcesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetContent(v []*DescribeTopRiskyResourcesResponseBodyDataContent) *DescribeTopRiskyResourcesResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetMaxResults(v int32) *DescribeTopRiskyResourcesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetNextToken(v string) *DescribeTopRiskyResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetTotalCount(v int64) *DescribeTopRiskyResourcesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTopRiskyResourcesResponseBodyDataContent struct {
	// The size of data in the Archive storage class, in bytes.
	//
	// example:
	//
	// 0
	ArchiveDataSize *int64 `json:"ArchiveDataSize,omitempty" xml:"ArchiveDataSize,omitempty"`
	// The number of checks that failed.
	//
	// example:
	//
	// 0
	CheckFailedCount *int64 `json:"CheckFailedCount,omitempty" xml:"CheckFailedCount,omitempty"`
	// The size of data in the Cold Archive storage class, in bytes.
	//
	// example:
	//
	// 0
	ColdArchiveDataSize *int64 `json:"ColdArchiveDataSize,omitempty" xml:"ColdArchiveDataSize,omitempty"`
	// The timestamp indicating when the resource was created.
	//
	// example:
	//
	// 1697798340
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data redundancy type. For example, \\"LRS\\" (locally redundant storage).
	//
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	// Indicates whether the data protection score assessment is enabled.
	//
	// example:
	//
	// true
	EnableCheck *bool `json:"EnableCheck,omitempty" xml:"EnableCheck,omitempty"`
	// The size of data in the Infrequent Access (IA) storage class, in bytes.
	//
	// example:
	//
	// 0
	IaDataSize *int64 `json:"IaDataSize,omitempty" xml:"IaDataSize,omitempty"`
	// The product type.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The data protection score of the resource.
	//
	// example:
	//
	// 90
	ProtectionScore *int32 `json:"ProtectionScore,omitempty" xml:"ProtectionScore,omitempty"`
	// The timestamp when the protection score was last updated.
	//
	// example:
	//
	// 1726036498
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	//
	// example:
	//
	// acs:ecs:cn-hangzhou:xxxxxxxx:instance/xxxxx
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-xxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// test-server
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the Alibaba Cloud account that owns the resource.
	//
	// example:
	//
	// 123***7890
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of identified risks.
	//
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The size of data in the Standard storage class, in bytes.
	//
	// example:
	//
	// 0
	StandardDataSize *int64 `json:"StandardDataSize,omitempty" xml:"StandardDataSize,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the data. For example, \\"Standard\\".
	//
	// example:
	//
	// STANDARD
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The total data size, in bytes.
	//
	// example:
	//
	// 0
	TotalDataSize *int64 `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-xxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-xxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTopRiskyResourcesResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetArchiveDataSize() *int64 {
	return s.ArchiveDataSize
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetCheckFailedCount() *int64 {
	return s.CheckFailedCount
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetColdArchiveDataSize() *int64 {
	return s.ColdArchiveDataSize
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetEnableCheck() *bool {
	return s.EnableCheck
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetIaDataSize() *int64 {
	return s.IaDataSize
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetProtectionScore() *int32 {
	return s.ProtectionScore
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetProtectionScoreUpdatedTime() *int64 {
	return s.ProtectionScoreUpdatedTime
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetStandardDataSize() *int64 {
	return s.StandardDataSize
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetStatus() *string {
	return s.Status
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetTotalDataSize() *int64 {
	return s.TotalDataSize
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetArchiveDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ArchiveDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetCheckFailedCount(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.CheckFailedCount = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetColdArchiveDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ColdArchiveDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetCreateTime(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.CreateTime = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetDataRedundancyType(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.DataRedundancyType = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetEnableCheck(v bool) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.EnableCheck = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetIaDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.IaDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetProductType(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetProtectionScore(v int32) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ProtectionScore = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetProtectionScoreUpdatedTime(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetRegionId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.RegionId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceArn(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceArn = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceName(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceName = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceOwnerId(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceType(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetRiskCount(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.RiskCount = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetStandardDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.StandardDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetStatus(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetStorageClass(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.StorageClass = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetTotalDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.TotalDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetVSwitchId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetVpcId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.VpcId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetZoneId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ZoneId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
