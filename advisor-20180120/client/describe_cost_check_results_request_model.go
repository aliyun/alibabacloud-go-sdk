// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunIdList(v []*int64) *DescribeCostCheckResultsRequest
	GetAssumeAliyunIdList() []*int64
	SetCheckIds(v []*string) *DescribeCostCheckResultsRequest
	GetCheckIds() []*string
	SetCheckPlanId(v int64) *DescribeCostCheckResultsRequest
	GetCheckPlanId() *int64
	SetGroupBy(v string) *DescribeCostCheckResultsRequest
	GetGroupBy() *string
	SetLanguage(v string) *DescribeCostCheckResultsRequest
	GetLanguage() *string
	SetProduct(v string) *DescribeCostCheckResultsRequest
	GetProduct() *string
	SetRegionIds(v []*string) *DescribeCostCheckResultsRequest
	GetRegionIds() []*string
	SetResourceGroupIdList(v []*string) *DescribeCostCheckResultsRequest
	GetResourceGroupIdList() []*string
	SetResourceId(v string) *DescribeCostCheckResultsRequest
	GetResourceId() *string
	SetResourceIds(v []*string) *DescribeCostCheckResultsRequest
	GetResourceIds() []*string
	SetResourceName(v string) *DescribeCostCheckResultsRequest
	GetResourceName() *string
	SetSeverity(v int32) *DescribeCostCheckResultsRequest
	GetSeverity() *int32
	SetTagKeys(v []*string) *DescribeCostCheckResultsRequest
	GetTagKeys() []*string
	SetTagList(v []*DescribeCostCheckResultsRequestTagList) *DescribeCostCheckResultsRequest
	GetTagList() []*DescribeCostCheckResultsRequestTagList
	SetTagValues(v []*string) *DescribeCostCheckResultsRequest
	GetTagValues() []*string
}

type DescribeCostCheckResultsRequest struct {
	AssumeAliyunIdList []*int64  `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	CheckIds           []*string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	CheckPlanId        *int64    `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// Category
	GroupBy  *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product             *string   `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionIds           []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	ResourceGroupIdList []*string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty" type:"Repeated"`
	ResourceId          *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIds         []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// SYNC_********_TASK
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity  *int32                                    `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeys   []*string                                 `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	TagList   []*DescribeCostCheckResultsRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	TagValues []*string                                 `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsRequest) GetAssumeAliyunIdList() []*int64 {
	return s.AssumeAliyunIdList
}

func (s *DescribeCostCheckResultsRequest) GetCheckIds() []*string {
	return s.CheckIds
}

func (s *DescribeCostCheckResultsRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeCostCheckResultsRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeCostCheckResultsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeCostCheckResultsRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeCostCheckResultsRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *DescribeCostCheckResultsRequest) GetResourceGroupIdList() []*string {
	return s.ResourceGroupIdList
}

func (s *DescribeCostCheckResultsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCostCheckResultsRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeCostCheckResultsRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCostCheckResultsRequest) GetSeverity() *int32 {
	return s.Severity
}

func (s *DescribeCostCheckResultsRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *DescribeCostCheckResultsRequest) GetTagList() []*DescribeCostCheckResultsRequestTagList {
	return s.TagList
}

func (s *DescribeCostCheckResultsRequest) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeCostCheckResultsRequest) SetAssumeAliyunIdList(v []*int64) *DescribeCostCheckResultsRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetCheckIds(v []*string) *DescribeCostCheckResultsRequest {
	s.CheckIds = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetCheckPlanId(v int64) *DescribeCostCheckResultsRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetGroupBy(v string) *DescribeCostCheckResultsRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetLanguage(v string) *DescribeCostCheckResultsRequest {
	s.Language = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetProduct(v string) *DescribeCostCheckResultsRequest {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetRegionIds(v []*string) *DescribeCostCheckResultsRequest {
	s.RegionIds = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetResourceGroupIdList(v []*string) *DescribeCostCheckResultsRequest {
	s.ResourceGroupIdList = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetResourceId(v string) *DescribeCostCheckResultsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetResourceIds(v []*string) *DescribeCostCheckResultsRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetResourceName(v string) *DescribeCostCheckResultsRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetSeverity(v int32) *DescribeCostCheckResultsRequest {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetTagKeys(v []*string) *DescribeCostCheckResultsRequest {
	s.TagKeys = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetTagList(v []*DescribeCostCheckResultsRequestTagList) *DescribeCostCheckResultsRequest {
	s.TagList = v
	return s
}

func (s *DescribeCostCheckResultsRequest) SetTagValues(v []*string) *DescribeCostCheckResultsRequest {
	s.TagValues = v
	return s
}

func (s *DescribeCostCheckResultsRequest) Validate() error {
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostCheckResultsRequestTagList struct {
	// example:
	//
	// ERP
	TagKey   *string   `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckResultsRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsRequestTagList) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsRequestTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCostCheckResultsRequestTagList) GetTagValue() []*string {
	return s.TagValue
}

func (s *DescribeCostCheckResultsRequestTagList) SetTagKey(v string) *DescribeCostCheckResultsRequestTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeCostCheckResultsRequestTagList) SetTagValue(v []*string) *DescribeCostCheckResultsRequestTagList {
	s.TagValue = v
	return s
}

func (s *DescribeCostCheckResultsRequestTagList) Validate() error {
	return dara.Validate(s)
}
