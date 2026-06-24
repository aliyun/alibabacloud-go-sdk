// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOssObjectDetail(v *DescribeOssObjectDetailResponseBodyOssObjectDetail) *DescribeOssObjectDetailResponseBody
	GetOssObjectDetail() *DescribeOssObjectDetailResponseBodyOssObjectDetail
	SetRequestId(v string) *DescribeOssObjectDetailResponseBody
	GetRequestId() *string
}

type DescribeOssObjectDetailResponseBody struct {
	// The details of the OSS object.
	OssObjectDetail *DescribeOssObjectDetailResponseBodyOssObjectDetail `json:"OssObjectDetail,omitempty" xml:"OssObjectDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssObjectDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBody) GetOssObjectDetail() *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	return s.OssObjectDetail
}

func (s *DescribeOssObjectDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssObjectDetailResponseBody) SetOssObjectDetail(v *DescribeOssObjectDetailResponseBodyOssObjectDetail) *DescribeOssObjectDetailResponseBody {
	s.OssObjectDetail = v
	return s
}

func (s *DescribeOssObjectDetailResponseBody) SetRequestId(v string) *DescribeOssObjectDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBody) Validate() error {
	if s.OssObjectDetail != nil {
		if err := s.OssObjectDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOssObjectDetailResponseBodyOssObjectDetail struct {
	// The name of the bucket to which the OSS object belongs.
	//
	// example:
	//
	// bucke***
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The name of the OSS object type.
	//
	// example:
	//
	// Excel
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The name of the OSS object.
	//
	// example:
	//
	// obj_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the OSS object is stored.
	//
	// example:
	//
	// cn-***
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the risk level for the OSS object.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// A list of sensitive data detection rules that the OSS object hits.
	RuleList []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetail) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) GetName() *string {
	return s.Name
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) GetRuleList() []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	return s.RuleList
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetBucketName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.BucketName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetCategoryName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRegionId(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RegionId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRiskLevelName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) SetRuleList(v []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) *DescribeOssObjectDetailResponseBodyOssObjectDetail {
	s.RuleList = v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetail) Validate() error {
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

type DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList struct {
	// The name of the OSS object type.
	//
	// example:
	//
	// Excel
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The number of times the sensitive data detection rule was hit.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A list of data tags that are hit by the detection model.
	ModelTags []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags `json:"ModelTags,omitempty" xml:"ModelTags,omitempty" type:"Repeated"`
	// The ID of the risk level for the OSS object.
	//
	// - **1**: No sensitive data is detected.
	//
	// - **2**: Level 1 sensitive data.
	//
	// - **3**: Level 2 sensitive data.
	//
	// - **4**: Level 3 sensitive data.
	//
	// - **5**: Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the risk level for the OSS object.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The name of the sensitive data detection rule that was hit.
	//
	// example:
	//
	// name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GetModelTags() []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags {
	return s.ModelTags
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetCategoryName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetCount(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.Count = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetModelTags(v []*DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.ModelTags = v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRiskLevelId(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRiskLevelName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) SetRuleName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList {
	s.RuleName = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleList) Validate() error {
	if s.ModelTags != nil {
		for _, item := range s.ModelTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags struct {
	// The ID of the data tag for the detection model.
	//
	// - **101**: Personal sensitive information.
	//
	// - **102**: Personal information.
	//
	// - **103**: Important data.
	//
	// example:
	//
	// 101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data tag for the detection model.
	//
	// - Personal sensitive information.
	//
	// - Personal information.
	//
	// - Important data.
	//
	// example:
	//
	// personal sensitive data
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) GetId() *int64 {
	return s.Id
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) GetName() *string {
	return s.Name
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) SetId(v int64) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) SetName(v string) *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags {
	s.Name = &v
	return s
}

func (s *DescribeOssObjectDetailResponseBodyOssObjectDetailRuleListModelTags) Validate() error {
	return dara.Validate(s)
}
