// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeOssObjectsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeOssObjectsResponseBodyItems) *DescribeOssObjectsResponseBody
	GetItems() []*DescribeOssObjectsResponseBodyItems
	SetMarker(v string) *DescribeOssObjectsResponseBody
	GetMarker() *string
	SetNextMarker(v string) *DescribeOssObjectsResponseBody
	GetNextMarker() *string
	SetPageSize(v int32) *DescribeOssObjectsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeOssObjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOssObjectsResponseBody
	GetTotalCount() *int32
	SetTruncated(v bool) *DescribeOssObjectsResponseBody
	GetTruncated() *bool
}

type DescribeOssObjectsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of OSS objects.
	Items []*DescribeOssObjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// This parameter is deprecated.
	//
	// example:
	//
	// -1
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The token that marks the start of the next page of results.
	//
	// > This parameter is returned only when `Truncated` is `true`.
	//
	// example:
	//
	// 1754786235714378752
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Indicates whether the results are truncated. The default value is false. Valid values:
	//
	// - **true**: The results are truncated.
	//
	// - **false**: The results are not truncated.
	//
	// example:
	//
	// false
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s DescribeOssObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOssObjectsResponseBody) GetItems() []*DescribeOssObjectsResponseBodyItems {
	return s.Items
}

func (s *DescribeOssObjectsResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *DescribeOssObjectsResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *DescribeOssObjectsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOssObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssObjectsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOssObjectsResponseBody) GetTruncated() *bool {
	return s.Truncated
}

func (s *DescribeOssObjectsResponseBody) SetCurrentPage(v int32) *DescribeOssObjectsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetItems(v []*DescribeOssObjectsResponseBodyItems) *DescribeOssObjectsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetMarker(v string) *DescribeOssObjectsResponseBody {
	s.Marker = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetNextMarker(v string) *DescribeOssObjectsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetPageSize(v int32) *DescribeOssObjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetRequestId(v string) *DescribeOssObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetTotalCount(v int32) *DescribeOssObjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) SetTruncated(v bool) *DescribeOssObjectsResponseBody {
	s.Truncated = &v
	return s
}

func (s *DescribeOssObjectsResponseBody) Validate() error {
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

type DescribeOssObjectsResponseBodyItems struct {
	// The bucket name.
	//
	// example:
	//
	// oss-duplicate-***
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The type of the OSS object, such as **900001*	- (MP4 video file), **800015*	- (PDF document), and **800005*	- (OSS configuration file).
	//
	// example:
	//
	// 900001
	Category *int64 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the file type.
	//
	// example:
	//
	// MP4
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The code of the file category.
	//
	// example:
	//
	// 1
	FileCategoryCode *int64 `json:"FileCategoryCode,omitempty" xml:"FileCategoryCode,omitempty"`
	// The name of the file category.
	//
	// example:
	//
	// text file
	FileCategoryName *string `json:"FileCategoryName,omitempty" xml:"FileCategoryName,omitempty"`
	// The ID of the OSS file.
	//
	// example:
	//
	// file-22***
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The unique ID of the OSS object.
	//
	// example:
	//
	// 17383
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the asset instance to which the OSS object belongs.
	//
	// example:
	//
	// 1232122
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the file was last modified.
	//
	// example:
	//
	// 1536751124000
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The name of the OSS object.
	//
	// example:
	//
	// obj_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the OSS object owner.
	//
	// example:
	//
	// cn-***
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The risk level ID of the OSS object. Valid values:
	//
	// - **1**: N/A. No sensitive data is detected.
	//
	// - **2**: S1. Level 1 sensitive data.
	//
	// - **3**: S2. Level 2 sensitive data.
	//
	// - **4**: S3. Level 3 sensitive data.
	//
	// - **5**: S4. Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the risk level for the OSS object.
	//
	// example:
	//
	// High risk
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The number of matched rules.
	//
	// example:
	//
	// 100
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// A list of rules.
	RuleList []*DescribeOssObjectsResponseBodyItemsRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// The number of matched fields.
	//
	// example:
	//
	// 50
	SensitiveCount *int32 `json:"SensitiveCount,omitempty" xml:"SensitiveCount,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeOssObjectsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBodyItems) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeOssObjectsResponseBodyItems) GetCategory() *int64 {
	return s.Category
}

func (s *DescribeOssObjectsResponseBodyItems) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeOssObjectsResponseBodyItems) GetFileCategoryCode() *int64 {
	return s.FileCategoryCode
}

func (s *DescribeOssObjectsResponseBodyItems) GetFileCategoryName() *string {
	return s.FileCategoryName
}

func (s *DescribeOssObjectsResponseBodyItems) GetFileId() *string {
	return s.FileId
}

func (s *DescribeOssObjectsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *DescribeOssObjectsResponseBodyItems) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeOssObjectsResponseBodyItems) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *DescribeOssObjectsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeOssObjectsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOssObjectsResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeOssObjectsResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeOssObjectsResponseBodyItems) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeOssObjectsResponseBodyItems) GetRuleList() []*DescribeOssObjectsResponseBodyItemsRuleList {
	return s.RuleList
}

func (s *DescribeOssObjectsResponseBodyItems) GetSensitiveCount() *int32 {
	return s.SensitiveCount
}

func (s *DescribeOssObjectsResponseBodyItems) GetSize() *int64 {
	return s.Size
}

func (s *DescribeOssObjectsResponseBodyItems) SetBucketName(v string) *DescribeOssObjectsResponseBodyItems {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetCategory(v int64) *DescribeOssObjectsResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetCategoryName(v string) *DescribeOssObjectsResponseBodyItems {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetFileCategoryCode(v int64) *DescribeOssObjectsResponseBodyItems {
	s.FileCategoryCode = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetFileCategoryName(v string) *DescribeOssObjectsResponseBodyItems {
	s.FileCategoryName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetFileId(v string) *DescribeOssObjectsResponseBodyItems {
	s.FileId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetId(v string) *DescribeOssObjectsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetInstanceId(v int64) *DescribeOssObjectsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetLastModifiedTime(v int64) *DescribeOssObjectsResponseBodyItems {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetName(v string) *DescribeOssObjectsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRegionId(v string) *DescribeOssObjectsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRiskLevelId(v int64) *DescribeOssObjectsResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRiskLevelName(v string) *DescribeOssObjectsResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRuleCount(v int32) *DescribeOssObjectsResponseBodyItems {
	s.RuleCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetRuleList(v []*DescribeOssObjectsResponseBodyItemsRuleList) *DescribeOssObjectsResponseBodyItems {
	s.RuleList = v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetSensitiveCount(v int32) *DescribeOssObjectsResponseBodyItems {
	s.SensitiveCount = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) SetSize(v int64) *DescribeOssObjectsResponseBodyItems {
	s.Size = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItems) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOssObjectsResponseBodyItemsRuleList struct {
	// The number of times the rule is matched.
	//
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// ID card
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The risk level ID of the rule. Valid values:
	//
	// - **1**: N/A. No sensitive data is detected.
	//
	// - **2**: S1. Level 1 sensitive data.
	//
	// - **3**: S2. Level 2 sensitive data.
	//
	// - **4**: S3. Level 3 sensitive data.
	//
	// - **5**: S4. Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
}

func (s DescribeOssObjectsResponseBodyItemsRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectsResponseBodyItemsRuleList) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetCount(v int64) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.Count = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetName(v string) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) SetRiskLevelId(v int64) *DescribeOssObjectsResponseBodyItemsRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectsResponseBodyItemsRuleList) Validate() error {
	return dara.Validate(s)
}
