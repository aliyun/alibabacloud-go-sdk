// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiGroupAttributes(v *DescribeApiGroupsResponseBodyApiGroupAttributes) *DescribeApiGroupsResponseBody
	GetApiGroupAttributes() *DescribeApiGroupsResponseBodyApiGroupAttributes
	SetPageNumber(v int32) *DescribeApiGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeApiGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeApiGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeApiGroupsResponseBody struct {
	ApiGroupAttributes *DescribeApiGroupsResponseBodyApiGroupAttributes `json:"ApiGroupAttributes,omitempty" xml:"ApiGroupAttributes,omitempty" type:"Struct"`
	// The number of pages to return the results on.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBody) GetApiGroupAttributes() *DescribeApiGroupsResponseBodyApiGroupAttributes {
	return s.ApiGroupAttributes
}

func (s *DescribeApiGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeApiGroupsResponseBody) SetApiGroupAttributes(v *DescribeApiGroupsResponseBodyApiGroupAttributes) *DescribeApiGroupsResponseBody {
	s.ApiGroupAttributes = v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetPageNumber(v int32) *DescribeApiGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetPageSize(v int32) *DescribeApiGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetRequestId(v string) *DescribeApiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetTotalCount(v int32) *DescribeApiGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) Validate() error {
	if s.ApiGroupAttributes != nil {
		if err := s.ApiGroupAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiGroupsResponseBodyApiGroupAttributes struct {
	ApiGroupAttribute []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute `json:"ApiGroupAttribute,omitempty" xml:"ApiGroupAttribute,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributes) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributes) GetApiGroupAttribute() []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	return s.ApiGroupAttribute
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributes) SetApiGroupAttribute(v []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) *DescribeApiGroupsResponseBodyApiGroupAttributes {
	s.ApiGroupAttribute = v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributes) Validate() error {
	if s.ApiGroupAttribute != nil {
		for _, item := range s.ApiGroupAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute struct {
	BasePath      *string                                                               `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	BillingStatus *string                                                               `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	CreatedTime   *string                                                               `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string                                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                                               `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpsPolicy   *string                                                               `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	IllegalStatus *string                                                               `json:"IllegalStatus,omitempty" xml:"IllegalStatus,omitempty"`
	InstanceId    *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType  *string                                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	ModifiedTime  *string                                                               `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId      *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubDomain     *string                                                               `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	Tags          *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TrafficLimit  *int32                                                                `json:"TrafficLimit,omitempty" xml:"TrafficLimit,omitempty"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetBasePath() *string {
	return s.BasePath
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetBillingStatus() *string {
	return s.BillingStatus
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetHttpsPolicy() *string {
	return s.HttpsPolicy
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetIllegalStatus() *string {
	return s.IllegalStatus
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetTags() *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags {
	return s.Tags
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GetTrafficLimit() *int32 {
	return s.TrafficLimit
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetBasePath(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.BasePath = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetBillingStatus(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.BillingStatus = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetCreatedTime(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetDescription(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetGroupId(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetGroupName(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetHttpsPolicy(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.HttpsPolicy = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetIllegalStatus(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.IllegalStatus = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetInstanceId(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetInstanceType(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.InstanceType = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetModifiedTime(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetRegionId(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetSubDomain(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.SubDomain = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetTags(v *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.Tags = v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetTrafficLimit(v int32) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.TrafficLimit = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags struct {
	TagInfo []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) GetTagInfo() []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo {
	return s.TagInfo
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) SetTagInfo(v []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags {
	s.TagInfo = v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) Validate() error {
	if s.TagInfo != nil {
		for _, item := range s.TagInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) SetKey(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) SetValue(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo {
	s.Value = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) Validate() error {
	return dara.Validate(s)
}
