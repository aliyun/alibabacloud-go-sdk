// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeResourcesResponseBodyData) *DescribeResourcesResponseBody
	GetData() *DescribeResourcesResponseBodyData
	SetRequestId(v string) *DescribeResourcesResponseBody
	GetRequestId() *string
}

type DescribeResourcesResponseBody struct {
	// The response data.
	Data *DescribeResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 700683DE-0154-56D4-8D76-3B7A2C2C7DF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponseBody) GetData() *DescribeResourcesResponseBodyData {
	return s.Data
}

func (s *DescribeResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourcesResponseBody) SetData(v *DescribeResourcesResponseBodyData) *DescribeResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourcesResponseBody) SetRequestId(v string) *DescribeResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcesResponseBodyData struct {
	// The list of returned resources.
	Content []*DescribeResourcesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of resources to return on each page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token to retrieve the next page of results. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// fb836242f4225fa0f0e0257362dfc6dd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of resources that match the query criteria. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 149
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponseBodyData) GetContent() []*DescribeResourcesResponseBodyDataContent {
	return s.Content
}

func (s *DescribeResourcesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourcesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourcesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeResourcesResponseBodyData) SetContent(v []*DescribeResourcesResponseBodyDataContent) *DescribeResourcesResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeResourcesResponseBodyData) SetMaxResults(v int32) *DescribeResourcesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourcesResponseBodyData) SetNextToken(v string) *DescribeResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeResourcesResponseBodyData) SetTotalCount(v int64) *DescribeResourcesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourcesResponseBodyData) Validate() error {
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

type DescribeResourcesResponseBodyDataContent struct {
	// The amount of data in the Archive storage class.
	//
	// example:
	//
	// 0
	ArchiveDataSize *int64 `json:"ArchiveDataSize,omitempty" xml:"ArchiveDataSize,omitempty"`
	// The number of failed check items.
	//
	// example:
	//
	// 0
	CheckFailedCount *int64 `json:"CheckFailedCount,omitempty" xml:"CheckFailedCount,omitempty"`
	// The amount of data in the Cold Archive storage class.
	//
	// example:
	//
	// 0
	ColdArchiveDataSize *int64 `json:"ColdArchiveDataSize,omitempty" xml:"ColdArchiveDataSize,omitempty"`
	// The UNIX timestamp that indicates when the resource was created.
	//
	// example:
	//
	// 1697798340
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data redundancy type.
	//
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	// Indicates whether data protection scoring is enabled for the resource.
	//
	// example:
	//
	// 0
	EnableCheck *bool `json:"EnableCheck,omitempty" xml:"EnableCheck,omitempty"`
	// The amount of data in the Infrequent Access (IA) storage class.
	//
	// example:
	//
	// 0
	IaDataSize *int64 `json:"IaDataSize,omitempty" xml:"IaDataSize,omitempty"`
	// The type of the cloud service.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The data protection score of the resource.
	//
	// example:
	//
	// 0
	ProtectionScore *int32 `json:"ProtectionScore,omitempty" xml:"ProtectionScore,omitempty"`
	// The UNIX timestamp that indicates when the score was last updated.
	//
	// example:
	//
	// 0
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The unique resource ARN.
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
	// The name of the resource.
	//
	// example:
	//
	// test server
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource owner ID.
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
	// The number of check items with potential risks.
	//
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The amount of data in the Standard storage class.
	//
	// example:
	//
	// 0
	StandardDataSize *int64 `json:"StandardDataSize,omitempty" xml:"StandardDataSize,omitempty"`
	// The resource status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the resource.
	//
	// example:
	//
	// STANDARD
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The total amount of data.
	//
	// example:
	//
	// 0
	TotalDataSize *int64 `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-xxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeResourcesResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponseBodyDataContent) GetArchiveDataSize() *int64 {
	return s.ArchiveDataSize
}

func (s *DescribeResourcesResponseBodyDataContent) GetCheckFailedCount() *int64 {
	return s.CheckFailedCount
}

func (s *DescribeResourcesResponseBodyDataContent) GetColdArchiveDataSize() *int64 {
	return s.ColdArchiveDataSize
}

func (s *DescribeResourcesResponseBodyDataContent) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeResourcesResponseBodyDataContent) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *DescribeResourcesResponseBodyDataContent) GetEnableCheck() *bool {
	return s.EnableCheck
}

func (s *DescribeResourcesResponseBodyDataContent) GetIaDataSize() *int64 {
	return s.IaDataSize
}

func (s *DescribeResourcesResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeResourcesResponseBodyDataContent) GetProtectionScore() *int32 {
	return s.ProtectionScore
}

func (s *DescribeResourcesResponseBodyDataContent) GetProtectionScoreUpdatedTime() *int64 {
	return s.ProtectionScoreUpdatedTime
}

func (s *DescribeResourcesResponseBodyDataContent) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourcesResponseBodyDataContent) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DescribeResourcesResponseBodyDataContent) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourcesResponseBodyDataContent) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeResourcesResponseBodyDataContent) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeResourcesResponseBodyDataContent) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourcesResponseBodyDataContent) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeResourcesResponseBodyDataContent) GetStandardDataSize() *int64 {
	return s.StandardDataSize
}

func (s *DescribeResourcesResponseBodyDataContent) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourcesResponseBodyDataContent) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeResourcesResponseBodyDataContent) GetTotalDataSize() *int64 {
	return s.TotalDataSize
}

func (s *DescribeResourcesResponseBodyDataContent) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeResourcesResponseBodyDataContent) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeResourcesResponseBodyDataContent) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeResourcesResponseBodyDataContent) SetArchiveDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.ArchiveDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetCheckFailedCount(v int64) *DescribeResourcesResponseBodyDataContent {
	s.CheckFailedCount = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetColdArchiveDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.ColdArchiveDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetCreateTime(v int64) *DescribeResourcesResponseBodyDataContent {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetDataRedundancyType(v string) *DescribeResourcesResponseBodyDataContent {
	s.DataRedundancyType = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetEnableCheck(v bool) *DescribeResourcesResponseBodyDataContent {
	s.EnableCheck = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetIaDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.IaDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetProductType(v string) *DescribeResourcesResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetProtectionScore(v int32) *DescribeResourcesResponseBodyDataContent {
	s.ProtectionScore = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetProtectionScoreUpdatedTime(v int64) *DescribeResourcesResponseBodyDataContent {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetRegionId(v string) *DescribeResourcesResponseBodyDataContent {
	s.RegionId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceArn(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceArn = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceId(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceName(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceName = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceOwnerId(v int64) *DescribeResourcesResponseBodyDataContent {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceType(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetRiskCount(v int64) *DescribeResourcesResponseBodyDataContent {
	s.RiskCount = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetStandardDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.StandardDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetStatus(v string) *DescribeResourcesResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetStorageClass(v string) *DescribeResourcesResponseBodyDataContent {
	s.StorageClass = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetTotalDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.TotalDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetVSwitchId(v string) *DescribeResourcesResponseBodyDataContent {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetVpcId(v string) *DescribeResourcesResponseBodyDataContent {
	s.VpcId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetZoneId(v string) *DescribeResourcesResponseBodyDataContent {
	s.ZoneId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
