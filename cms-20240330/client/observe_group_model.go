// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroup interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *ObserveGroup
	GetAliUid() *string
	SetCreateTime(v string) *ObserveGroup
	GetCreateTime() *string
	SetDescription(v string) *ObserveGroup
	GetDescription() *string
	SetDiscoverRules(v string) *ObserveGroup
	GetDiscoverRules() *string
	SetEntityCounts(v map[string]*int32) *ObserveGroup
	GetEntityCounts() map[string]*int32
	SetExtraInfo(v string) *ObserveGroup
	GetExtraInfo() *string
	SetFavorited(v bool) *ObserveGroup
	GetFavorited() *bool
	SetGroupId(v string) *ObserveGroup
	GetGroupId() *string
	SetGroupName(v string) *ObserveGroup
	GetGroupName() *string
	SetGroupType(v string) *ObserveGroup
	GetGroupType() *string
	SetHealth(v int32) *ObserveGroup
	GetHealth() *int32
	SetModifyTime(v string) *ObserveGroup
	GetModifyTime() *string
	SetOriginGroupId(v string) *ObserveGroup
	GetOriginGroupId() *string
	SetRegionId(v string) *ObserveGroup
	GetRegionId() *string
	SetResourceGroupId(v string) *ObserveGroup
	GetResourceGroupId() *string
	SetSourceOrigin(v string) *ObserveGroup
	GetSourceOrigin() *string
	SetTags(v []*ObserveGroupTags) *ObserveGroup
	GetTags() []*ObserveGroupTags
	SetWorkspaceId(v string) *ObserveGroup
	GetWorkspaceId() *string
}

type ObserveGroup struct {
	// The UID of the Alibaba Cloud account to which the group belongs.
	AliUid *string `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	// The time when the group was created, in UTC format (yyyy-MM-ddTHH:mm:ssZ).
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the observability group, which describes the business purpose.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of entity discovery rules that define which entities the group automatically matches.
	DiscoverRules *string `json:"discoverRules,omitempty" xml:"discoverRules,omitempty"`
	// The number of entities in each category. The key is the category domain (such as acs for Alibaba Cloud services, apm, or rum, which is extensible). The value is the number of entities in that category that belong to this group. This parameter is returned only when withEntityCount is set to true.
	EntityCounts map[string]*int32 `json:"entityCounts,omitempty" xml:"entityCounts,omitempty"`
	// The extended information, which is a JSON string that contains alert templates, alert contact groups, pause policies, and other settings.
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// Indicates whether the current user has favorited this group. This value is used as the filter criterion for the My Favorites feature.
	Favorited *bool `json:"favorited,omitempty" xml:"favorited,omitempty"`
	// The globally unique ID of the observability group, in the format of og-<16-character hash>. This ID is used across metrics, alerts, and the console.
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The name of the observability group. The name must be unique within the workspace.
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The type of the observability group.
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The health status of the group. Valid values:
	//
	// - -1: unknown (placeholder).
	//
	// - 1: healthy.
	//
	// - 0: unhealthy.
	Health *int32 `json:"health,omitempty" xml:"health,omitempty"`
	// The time when the group was last modified, in UTC format (yyyy-MM-ddTHH:mm:ssZ). This value is automatically updated when any property of the resource changes.
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The product_group.id of the version 1.0 application group. This parameter is returned only when sourceOrigin is set to synced_from_1_0.
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
	// The resource tags (Alibaba Cloud standard tags), which are an array of key-value pairs.
	Tags []*ObserveGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace ID to which the group belongs. This value is set at the workspace level and cannot be changed after the group is created.
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ObserveGroup) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroup) GoString() string {
	return s.String()
}

func (s *ObserveGroup) GetAliUid() *string {
	return s.AliUid
}

func (s *ObserveGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ObserveGroup) GetDescription() *string {
	return s.Description
}

func (s *ObserveGroup) GetDiscoverRules() *string {
	return s.DiscoverRules
}

func (s *ObserveGroup) GetEntityCounts() map[string]*int32 {
	return s.EntityCounts
}

func (s *ObserveGroup) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *ObserveGroup) GetFavorited() *bool {
	return s.Favorited
}

func (s *ObserveGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *ObserveGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ObserveGroup) GetGroupType() *string {
	return s.GroupType
}

func (s *ObserveGroup) GetHealth() *int32 {
	return s.Health
}

func (s *ObserveGroup) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ObserveGroup) GetOriginGroupId() *string {
	return s.OriginGroupId
}

func (s *ObserveGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *ObserveGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ObserveGroup) GetSourceOrigin() *string {
	return s.SourceOrigin
}

func (s *ObserveGroup) GetTags() []*ObserveGroupTags {
	return s.Tags
}

func (s *ObserveGroup) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ObserveGroup) SetAliUid(v string) *ObserveGroup {
	s.AliUid = &v
	return s
}

func (s *ObserveGroup) SetCreateTime(v string) *ObserveGroup {
	s.CreateTime = &v
	return s
}

func (s *ObserveGroup) SetDescription(v string) *ObserveGroup {
	s.Description = &v
	return s
}

func (s *ObserveGroup) SetDiscoverRules(v string) *ObserveGroup {
	s.DiscoverRules = &v
	return s
}

func (s *ObserveGroup) SetEntityCounts(v map[string]*int32) *ObserveGroup {
	s.EntityCounts = v
	return s
}

func (s *ObserveGroup) SetExtraInfo(v string) *ObserveGroup {
	s.ExtraInfo = &v
	return s
}

func (s *ObserveGroup) SetFavorited(v bool) *ObserveGroup {
	s.Favorited = &v
	return s
}

func (s *ObserveGroup) SetGroupId(v string) *ObserveGroup {
	s.GroupId = &v
	return s
}

func (s *ObserveGroup) SetGroupName(v string) *ObserveGroup {
	s.GroupName = &v
	return s
}

func (s *ObserveGroup) SetGroupType(v string) *ObserveGroup {
	s.GroupType = &v
	return s
}

func (s *ObserveGroup) SetHealth(v int32) *ObserveGroup {
	s.Health = &v
	return s
}

func (s *ObserveGroup) SetModifyTime(v string) *ObserveGroup {
	s.ModifyTime = &v
	return s
}

func (s *ObserveGroup) SetOriginGroupId(v string) *ObserveGroup {
	s.OriginGroupId = &v
	return s
}

func (s *ObserveGroup) SetRegionId(v string) *ObserveGroup {
	s.RegionId = &v
	return s
}

func (s *ObserveGroup) SetResourceGroupId(v string) *ObserveGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *ObserveGroup) SetSourceOrigin(v string) *ObserveGroup {
	s.SourceOrigin = &v
	return s
}

func (s *ObserveGroup) SetTags(v []*ObserveGroupTags) *ObserveGroup {
	s.Tags = v
	return s
}

func (s *ObserveGroup) SetWorkspaceId(v string) *ObserveGroup {
	s.WorkspaceId = &v
	return s
}

func (s *ObserveGroup) Validate() error {
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

type ObserveGroupTags struct {
	// The tag key.
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ObserveGroupTags) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupTags) GoString() string {
	return s.String()
}

func (s *ObserveGroupTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ObserveGroupTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ObserveGroupTags) SetTagKey(v string) *ObserveGroupTags {
	s.TagKey = &v
	return s
}

func (s *ObserveGroupTags) SetTagValue(v string) *ObserveGroupTags {
	s.TagValue = &v
	return s
}

func (s *ObserveGroupTags) Validate() error {
	return dara.Validate(s)
}
