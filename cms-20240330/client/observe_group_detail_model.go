// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *ObserveGroupDetail
	GetAliUid() *string
	SetCreateTime(v string) *ObserveGroupDetail
	GetCreateTime() *string
	SetDescription(v string) *ObserveGroupDetail
	GetDescription() *string
	SetDiscoverRules(v []*ObserveGroupDiscoverRule) *ObserveGroupDetail
	GetDiscoverRules() []*ObserveGroupDiscoverRule
	SetEntitySummaries(v []*ObserveGroupDetailEntitySummaries) *ObserveGroupDetail
	GetEntitySummaries() []*ObserveGroupDetailEntitySummaries
	SetExtraInfo(v string) *ObserveGroupDetail
	GetExtraInfo() *string
	SetFavorited(v bool) *ObserveGroupDetail
	GetFavorited() *bool
	SetGroupId(v string) *ObserveGroupDetail
	GetGroupId() *string
	SetGroupName(v string) *ObserveGroupDetail
	GetGroupName() *string
	SetGroupType(v string) *ObserveGroupDetail
	GetGroupType() *string
	SetModifyTime(v string) *ObserveGroupDetail
	GetModifyTime() *string
	SetOgEntityInfoEnabled(v bool) *ObserveGroupDetail
	GetOgEntityInfoEnabled() *bool
	SetOgEntityInfoPromInstances(v []*ObserveGroupPromInstance) *ObserveGroupDetail
	GetOgEntityInfoPromInstances() []*ObserveGroupPromInstance
	SetOriginGroupId(v string) *ObserveGroupDetail
	GetOriginGroupId() *string
	SetRegionId(v string) *ObserveGroupDetail
	GetRegionId() *string
	SetResourceGroupId(v string) *ObserveGroupDetail
	GetResourceGroupId() *string
	SetSourceOrigin(v string) *ObserveGroupDetail
	GetSourceOrigin() *string
	SetTags(v []*ObserveGroupDetailTags) *ObserveGroupDetail
	GetTags() []*ObserveGroupDetailTags
	SetWorkspaceId(v string) *ObserveGroupDetail
	GetWorkspaceId() *string
}

type ObserveGroupDetail struct {
	// The UID of the Alibaba Cloud account to which the group belongs.
	AliUid *string `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	// The time when the group was created, in UTC format (yyyy-MM-ddTHH:mm:ssZ).
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the observability group, which explains its business purpose.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of entity discovery rules that define which entities the group automatically matches.
	DiscoverRules []*ObserveGroupDiscoverRule `json:"discoverRules,omitempty" xml:"discoverRules,omitempty" type:"Repeated"`
	// The statistics of entities in the group, categorized by entity type.
	EntitySummaries []*ObserveGroupDetailEntitySummaries `json:"entitySummaries,omitempty" xml:"entitySummaries,omitempty" type:"Repeated"`
	// The extended information in JSON string format, which carries alert templates, alert contact groups, suspension policies, and other configurations.
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// Indicates whether the current user has followed the group.
	Favorited *bool   `json:"favorited,omitempty" xml:"favorited,omitempty"`
	GroupId   *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The name of the observability group. The name must be unique within the same workspace.
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The type of the observability group.
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The time when the group was last modified, in UTC format (yyyy-MM-ddTHH:mm:ssZ). This value is automatically updated when any property of the resource changes.
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// Specifies whether to enable the og_entity_info metric output. When enabled, the data plane writes the group ownership information to the target Prometheus instance.
	OgEntityInfoEnabled *bool `json:"ogEntityInfoEnabled,omitempty" xml:"ogEntityInfoEnabled,omitempty"`
	// The set of Prometheus instances to which og_entity_info is written. This includes two source types: system (automatically identified by the system) and custom (user-defined).
	OgEntityInfoPromInstances []*ObserveGroupPromInstance `json:"ogEntityInfoPromInstances,omitempty" xml:"ogEntityInfoPromInstances,omitempty" type:"Repeated"`
	// The product_group.id of the version 1.0 application group. This parameter is valid only when sourceOrigin is set to synced_from_1_0.
	OriginGroupId *string `json:"originGroupId,omitempty" xml:"originGroupId,omitempty"`
	// The region ID of the group.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The Alibaba Cloud resource group ID.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The data source. Valid values:
	//
	// - native_2_0: created natively in version 2.0.
	//
	// - synced_from_1_0: synchronized from a version 1.0 application group.
	SourceOrigin *string `json:"sourceOrigin,omitempty" xml:"sourceOrigin,omitempty"`
	// The resource tags (Alibaba Cloud standard tags), represented as an array of key-value pairs.
	Tags []*ObserveGroupDetailTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace ID to which the group belongs. This value is set at the workspace level and cannot be changed after the group is created.
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ObserveGroupDetail) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDetail) GoString() string {
	return s.String()
}

func (s *ObserveGroupDetail) GetAliUid() *string {
	return s.AliUid
}

func (s *ObserveGroupDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ObserveGroupDetail) GetDescription() *string {
	return s.Description
}

func (s *ObserveGroupDetail) GetDiscoverRules() []*ObserveGroupDiscoverRule {
	return s.DiscoverRules
}

func (s *ObserveGroupDetail) GetEntitySummaries() []*ObserveGroupDetailEntitySummaries {
	return s.EntitySummaries
}

func (s *ObserveGroupDetail) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *ObserveGroupDetail) GetFavorited() *bool {
	return s.Favorited
}

func (s *ObserveGroupDetail) GetGroupId() *string {
	return s.GroupId
}

func (s *ObserveGroupDetail) GetGroupName() *string {
	return s.GroupName
}

func (s *ObserveGroupDetail) GetGroupType() *string {
	return s.GroupType
}

func (s *ObserveGroupDetail) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ObserveGroupDetail) GetOgEntityInfoEnabled() *bool {
	return s.OgEntityInfoEnabled
}

func (s *ObserveGroupDetail) GetOgEntityInfoPromInstances() []*ObserveGroupPromInstance {
	return s.OgEntityInfoPromInstances
}

func (s *ObserveGroupDetail) GetOriginGroupId() *string {
	return s.OriginGroupId
}

func (s *ObserveGroupDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *ObserveGroupDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ObserveGroupDetail) GetSourceOrigin() *string {
	return s.SourceOrigin
}

func (s *ObserveGroupDetail) GetTags() []*ObserveGroupDetailTags {
	return s.Tags
}

func (s *ObserveGroupDetail) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ObserveGroupDetail) SetAliUid(v string) *ObserveGroupDetail {
	s.AliUid = &v
	return s
}

func (s *ObserveGroupDetail) SetCreateTime(v string) *ObserveGroupDetail {
	s.CreateTime = &v
	return s
}

func (s *ObserveGroupDetail) SetDescription(v string) *ObserveGroupDetail {
	s.Description = &v
	return s
}

func (s *ObserveGroupDetail) SetDiscoverRules(v []*ObserveGroupDiscoverRule) *ObserveGroupDetail {
	s.DiscoverRules = v
	return s
}

func (s *ObserveGroupDetail) SetEntitySummaries(v []*ObserveGroupDetailEntitySummaries) *ObserveGroupDetail {
	s.EntitySummaries = v
	return s
}

func (s *ObserveGroupDetail) SetExtraInfo(v string) *ObserveGroupDetail {
	s.ExtraInfo = &v
	return s
}

func (s *ObserveGroupDetail) SetFavorited(v bool) *ObserveGroupDetail {
	s.Favorited = &v
	return s
}

func (s *ObserveGroupDetail) SetGroupId(v string) *ObserveGroupDetail {
	s.GroupId = &v
	return s
}

func (s *ObserveGroupDetail) SetGroupName(v string) *ObserveGroupDetail {
	s.GroupName = &v
	return s
}

func (s *ObserveGroupDetail) SetGroupType(v string) *ObserveGroupDetail {
	s.GroupType = &v
	return s
}

func (s *ObserveGroupDetail) SetModifyTime(v string) *ObserveGroupDetail {
	s.ModifyTime = &v
	return s
}

func (s *ObserveGroupDetail) SetOgEntityInfoEnabled(v bool) *ObserveGroupDetail {
	s.OgEntityInfoEnabled = &v
	return s
}

func (s *ObserveGroupDetail) SetOgEntityInfoPromInstances(v []*ObserveGroupPromInstance) *ObserveGroupDetail {
	s.OgEntityInfoPromInstances = v
	return s
}

func (s *ObserveGroupDetail) SetOriginGroupId(v string) *ObserveGroupDetail {
	s.OriginGroupId = &v
	return s
}

func (s *ObserveGroupDetail) SetRegionId(v string) *ObserveGroupDetail {
	s.RegionId = &v
	return s
}

func (s *ObserveGroupDetail) SetResourceGroupId(v string) *ObserveGroupDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *ObserveGroupDetail) SetSourceOrigin(v string) *ObserveGroupDetail {
	s.SourceOrigin = &v
	return s
}

func (s *ObserveGroupDetail) SetTags(v []*ObserveGroupDetailTags) *ObserveGroupDetail {
	s.Tags = v
	return s
}

func (s *ObserveGroupDetail) SetWorkspaceId(v string) *ObserveGroupDetail {
	s.WorkspaceId = &v
	return s
}

func (s *ObserveGroupDetail) Validate() error {
	if s.DiscoverRules != nil {
		for _, item := range s.DiscoverRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EntitySummaries != nil {
		for _, item := range s.EntitySummaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OgEntityInfoPromInstances != nil {
		for _, item := range s.OgEntityInfoPromInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ObserveGroupDetailEntitySummaries struct {
	// The entity category.
	EntityCategory *string `json:"entityCategory,omitempty" xml:"entityCategory,omitempty"`
	// The number of entities.
	EntityCount *int32 `json:"entityCount,omitempty" xml:"entityCount,omitempty"`
	// The entity domain.
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// The entity type.
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
}

func (s ObserveGroupDetailEntitySummaries) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDetailEntitySummaries) GoString() string {
	return s.String()
}

func (s *ObserveGroupDetailEntitySummaries) GetEntityCategory() *string {
	return s.EntityCategory
}

func (s *ObserveGroupDetailEntitySummaries) GetEntityCount() *int32 {
	return s.EntityCount
}

func (s *ObserveGroupDetailEntitySummaries) GetEntityDomain() *string {
	return s.EntityDomain
}

func (s *ObserveGroupDetailEntitySummaries) GetEntityType() *string {
	return s.EntityType
}

func (s *ObserveGroupDetailEntitySummaries) SetEntityCategory(v string) *ObserveGroupDetailEntitySummaries {
	s.EntityCategory = &v
	return s
}

func (s *ObserveGroupDetailEntitySummaries) SetEntityCount(v int32) *ObserveGroupDetailEntitySummaries {
	s.EntityCount = &v
	return s
}

func (s *ObserveGroupDetailEntitySummaries) SetEntityDomain(v string) *ObserveGroupDetailEntitySummaries {
	s.EntityDomain = &v
	return s
}

func (s *ObserveGroupDetailEntitySummaries) SetEntityType(v string) *ObserveGroupDetailEntitySummaries {
	s.EntityType = &v
	return s
}

func (s *ObserveGroupDetailEntitySummaries) Validate() error {
	return dara.Validate(s)
}

type ObserveGroupDetailTags struct {
	// The tag key.
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ObserveGroupDetailTags) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDetailTags) GoString() string {
	return s.String()
}

func (s *ObserveGroupDetailTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ObserveGroupDetailTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ObserveGroupDetailTags) SetTagKey(v string) *ObserveGroupDetailTags {
	s.TagKey = &v
	return s
}

func (s *ObserveGroupDetailTags) SetTagValue(v string) *ObserveGroupDetailTags {
	s.TagValue = &v
	return s
}

func (s *ObserveGroupDetailTags) Validate() error {
	return dara.Validate(s)
}
