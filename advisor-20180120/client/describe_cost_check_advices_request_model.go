// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckAdvicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunIdList(v []*int64) *DescribeCostCheckAdvicesRequest
	GetAssumeAliyunIdList() []*int64
	SetCheckId(v string) *DescribeCostCheckAdvicesRequest
	GetCheckId() *string
	SetCheckPlanId(v int64) *DescribeCostCheckAdvicesRequest
	GetCheckPlanId() *int64
	SetLanguage(v string) *DescribeCostCheckAdvicesRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeCostCheckAdvicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCostCheckAdvicesRequest
	GetPageSize() *int32
	SetRegionIds(v []*string) *DescribeCostCheckAdvicesRequest
	GetRegionIds() []*string
	SetResourceGroupIdList(v []*string) *DescribeCostCheckAdvicesRequest
	GetResourceGroupIdList() []*string
	SetResourceId(v string) *DescribeCostCheckAdvicesRequest
	GetResourceId() *string
	SetResourceIds(v []*string) *DescribeCostCheckAdvicesRequest
	GetResourceIds() []*string
	SetResourceName(v string) *DescribeCostCheckAdvicesRequest
	GetResourceName() *string
	SetSeverity(v string) *DescribeCostCheckAdvicesRequest
	GetSeverity() *string
	SetTagKeys(v []*string) *DescribeCostCheckAdvicesRequest
	GetTagKeys() []*string
	SetTagList(v []*DescribeCostCheckAdvicesRequestTagList) *DescribeCostCheckAdvicesRequest
	GetTagList() []*DescribeCostCheckAdvicesRequestTagList
	SetTagValues(v []*string) *DescribeCostCheckAdvicesRequest
	GetTagValues() []*string
}

type DescribeCostCheckAdvicesRequest struct {
	AssumeAliyunIdList []*int64 `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty" type:"Repeated"`
	// example:
	//
	// EcsCostLowUtilizationCheck
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 6
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize            *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds           []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	ResourceGroupIdList []*string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty" type:"Repeated"`
	ResourceId          *string   `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIds         []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity  *string                                   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeys   []*string                                 `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	TagList   []*DescribeCostCheckAdvicesRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	TagValues []*string                                 `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckAdvicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesRequest) GetAssumeAliyunIdList() []*int64 {
	return s.AssumeAliyunIdList
}

func (s *DescribeCostCheckAdvicesRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeCostCheckAdvicesRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeCostCheckAdvicesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeCostCheckAdvicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCostCheckAdvicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCostCheckAdvicesRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *DescribeCostCheckAdvicesRequest) GetResourceGroupIdList() []*string {
	return s.ResourceGroupIdList
}

func (s *DescribeCostCheckAdvicesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCostCheckAdvicesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeCostCheckAdvicesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCostCheckAdvicesRequest) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeCostCheckAdvicesRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *DescribeCostCheckAdvicesRequest) GetTagList() []*DescribeCostCheckAdvicesRequestTagList {
	return s.TagList
}

func (s *DescribeCostCheckAdvicesRequest) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeCostCheckAdvicesRequest) SetAssumeAliyunIdList(v []*int64) *DescribeCostCheckAdvicesRequest {
	s.AssumeAliyunIdList = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetCheckId(v string) *DescribeCostCheckAdvicesRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetCheckPlanId(v int64) *DescribeCostCheckAdvicesRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetLanguage(v string) *DescribeCostCheckAdvicesRequest {
	s.Language = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetPageNumber(v int32) *DescribeCostCheckAdvicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetPageSize(v int32) *DescribeCostCheckAdvicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetRegionIds(v []*string) *DescribeCostCheckAdvicesRequest {
	s.RegionIds = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetResourceGroupIdList(v []*string) *DescribeCostCheckAdvicesRequest {
	s.ResourceGroupIdList = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetResourceId(v string) *DescribeCostCheckAdvicesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetResourceIds(v []*string) *DescribeCostCheckAdvicesRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetResourceName(v string) *DescribeCostCheckAdvicesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetSeverity(v string) *DescribeCostCheckAdvicesRequest {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetTagKeys(v []*string) *DescribeCostCheckAdvicesRequest {
	s.TagKeys = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetTagList(v []*DescribeCostCheckAdvicesRequestTagList) *DescribeCostCheckAdvicesRequest {
	s.TagList = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) SetTagValues(v []*string) *DescribeCostCheckAdvicesRequest {
	s.TagValues = v
	return s
}

func (s *DescribeCostCheckAdvicesRequest) Validate() error {
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

type DescribeCostCheckAdvicesRequestTagList struct {
	// example:
	//
	// ecs_***_shanghai
	TagKey   *string   `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s DescribeCostCheckAdvicesRequestTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesRequestTagList) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesRequestTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCostCheckAdvicesRequestTagList) GetTagValue() []*string {
	return s.TagValue
}

func (s *DescribeCostCheckAdvicesRequestTagList) SetTagKey(v string) *DescribeCostCheckAdvicesRequestTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeCostCheckAdvicesRequestTagList) SetTagValue(v []*string) *DescribeCostCheckAdvicesRequestTagList {
	s.TagValue = v
	return s
}

func (s *DescribeCostCheckAdvicesRequestTagList) Validate() error {
	return dara.Validate(s)
}
