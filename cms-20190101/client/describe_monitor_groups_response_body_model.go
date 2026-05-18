// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeMonitorGroupsResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMonitorGroupsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeMonitorGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitorGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMonitorGroupsResponseBody
	GetRequestId() *string
	SetResources(v *DescribeMonitorGroupsResponseBodyResources) *DescribeMonitorGroupsResponseBody
	GetResources() *DescribeMonitorGroupsResponseBodyResources
	SetSuccess(v bool) *DescribeMonitorGroupsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeMonitorGroupsResponseBody
	GetTotal() *int32
}

type DescribeMonitorGroupsResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F02B299A-D374-4595-9F55-7534D604F132
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources *DescribeMonitorGroupsResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMonitorGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMonitorGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitorGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitorGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorGroupsResponseBody) GetResources() *DescribeMonitorGroupsResponseBodyResources {
	return s.Resources
}

func (s *DescribeMonitorGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitorGroupsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMonitorGroupsResponseBody) SetCode(v int32) *DescribeMonitorGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetMessage(v string) *DescribeMonitorGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetPageNumber(v int32) *DescribeMonitorGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetPageSize(v int32) *DescribeMonitorGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetRequestId(v string) *DescribeMonitorGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetResources(v *DescribeMonitorGroupsResponseBodyResources) *DescribeMonitorGroupsResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetSuccess(v bool) *DescribeMonitorGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetTotal(v int32) *DescribeMonitorGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorGroupsResponseBodyResources struct {
	Resource []*DescribeMonitorGroupsResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResources) GetResource() []*DescribeMonitorGroupsResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeMonitorGroupsResponseBodyResources) SetResource(v []*DescribeMonitorGroupsResponseBodyResourcesResource) *DescribeMonitorGroupsResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResources) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorGroupsResponseBodyResourcesResource struct {
	BindUrl              *string                                                          `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	ContactGroups        *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	DynamicTagRuleId     *string                                                          `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	GmtCreate            *int64                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *int64                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupFounderTagKey   *string                                                          `json:"GroupFounderTagKey,omitempty" xml:"GroupFounderTagKey,omitempty"`
	GroupFounderTagValue *string                                                          `json:"GroupFounderTagValue,omitempty" xml:"GroupFounderTagValue,omitempty"`
	GroupId              *int64                                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName            *string                                                          `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ResourceGroupId      *string                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceId            *string                                                          `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Tags                 *DescribeMonitorGroupsResponseBodyResourcesResourceTags          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TemplateIds          *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds   `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Struct"`
	TemplateInfos        *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos `json:"TemplateInfos,omitempty" xml:"TemplateInfos,omitempty" type:"Struct"`
	Type                 *string                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetBindUrl() *string {
	return s.BindUrl
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetContactGroups() *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups {
	return s.ContactGroups
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetDynamicTagRuleId() *string {
	return s.DynamicTagRuleId
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetGroupFounderTagKey() *string {
	return s.GroupFounderTagKey
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetGroupFounderTagValue() *string {
	return s.GroupFounderTagValue
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetTags() *DescribeMonitorGroupsResponseBodyResourcesResourceTags {
	return s.Tags
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetTemplateIds() *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds {
	return s.TemplateIds
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetTemplateInfos() *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos {
	return s.TemplateInfos
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) GetType() *string {
	return s.Type
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetBindUrl(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.BindUrl = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetContactGroups(v *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.ContactGroups = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetDynamicTagRuleId(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGmtCreate(v int64) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGmtModified(v int64) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GmtModified = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupFounderTagKey(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupFounderTagKey = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupFounderTagValue(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupFounderTagValue = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupId(v int64) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupName(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupName = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetResourceGroupId(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetServiceId(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.ServiceId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetTags(v *DescribeMonitorGroupsResponseBodyResourcesResourceTags) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.Tags = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetTemplateIds(v *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.TemplateIds = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetTemplateInfos(v *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.TemplateInfos = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetType(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.Type = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) Validate() error {
	if s.ContactGroups != nil {
		if err := s.ContactGroups.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateIds != nil {
		if err := s.TemplateIds.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateInfos != nil {
		if err := s.TemplateInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups struct {
	ContactGroup []*DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) GetContactGroup() []*DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup {
	return s.ContactGroup
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) SetContactGroup(v []*DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups {
	s.ContactGroup = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) Validate() error {
	if s.ContactGroup != nil {
		for _, item := range s.ContactGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) GetName() *string {
	return s.Name
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) SetName(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup {
	s.Name = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTags struct {
	Tag []*DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTags) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTags) GetTag() []*DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag {
	return s.Tag
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTags) SetTag(v []*DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) *DescribeMonitorGroupsResponseBodyResourcesResourceTags {
	s.Tag = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) SetKey(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) SetValue(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds struct {
	TemplateId []*string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) GetTemplateId() []*string {
	return s.TemplateId
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) SetTemplateId(v []*string) *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds {
	s.TemplateId = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos struct {
	TemplateInfo []*DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos) GetTemplateInfo() []*DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo {
	return s.TemplateInfo
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos) SetTemplateInfo(v []*DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos {
	s.TemplateInfo = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfos) Validate() error {
	if s.TemplateInfo != nil {
		for _, item := range s.TemplateInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo struct {
	EffectTime *int64  `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Ver        *string `json:"Ver,omitempty" xml:"Ver,omitempty"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) GetEffectTime() *int64 {
	return s.EffectTime
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) GetVer() *string {
	return s.Ver
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) SetEffectTime(v int64) *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo {
	s.EffectTime = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) SetTemplateId(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) SetVer(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo {
	s.Ver = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateInfosTemplateInfo) Validate() error {
	return dara.Validate(s)
}
