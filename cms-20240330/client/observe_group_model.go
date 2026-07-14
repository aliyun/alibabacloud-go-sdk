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
	SetDiscoverRules(v []*ObserveGroupDiscoverRule) *ObserveGroup
	GetDiscoverRules() []*ObserveGroupDiscoverRule
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
	SetOgEntityInfoEnabled(v bool) *ObserveGroup
	GetOgEntityInfoEnabled() *bool
	SetOgEntityInfoPromInstances(v []*ObserveGroupPromInstance) *ObserveGroup
	GetOgEntityInfoPromInstances() []*ObserveGroupPromInstance
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
	// The creation time in UTC format: yyyy-MM-ddTHH:mm:ssZ.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the observability group, which explains its business purpose.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of entity discovery rules that define which entities the group automatically matches.
	DiscoverRules []*ObserveGroupDiscoverRule `json:"discoverRules,omitempty" xml:"discoverRules,omitempty" type:"Repeated"`
	// The number of entities in each category. The key is the category domain (acs for cloud services, apm, or rum, which is extensible). The value is the number of entities in that category that belong to this group. Returned only when withEntityCount is set to true.
	EntityCounts map[string]*int32 `json:"entityCounts,omitempty" xml:"entityCounts,omitempty"`
	// The extended information as a JSON string, which carries alert templates, alert contact groups, and suspension policies.
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// Indicates whether the current user has followed this group. This value is used as the filter criterion for the My Favorites feature.
	Favorited *bool `json:"favorited,omitempty" xml:"favorited,omitempty"`
	// The globally unique ID of the observability group. Format: og-<16-character hash>. Used uniformly across metrics, alerts, and consoles.
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The name of the observability group. Must be unique within the same workspace.
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
	// The last modification time in UTC format: yyyy-MM-ddTHH:mm:ssZ. Automatically updated when any resource attribute changes.
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// Specifies whether og_entity_info metric output is enabled. When enabled, the data plane writes the group membership information to the target Prometheus instance.
	OgEntityInfoEnabled *bool `json:"ogEntityInfoEnabled,omitempty" xml:"ogEntityInfoEnabled,omitempty"`
	// The set of Prometheus instances to which og_entity_info is written. Includes two source types: system (automatically identified by the system) and custom (user-defined).
	OgEntityInfoPromInstances []*ObserveGroupPromInstance `json:"ogEntityInfoPromInstances,omitempty" xml:"ogEntityInfoPromInstances,omitempty" type:"Repeated"`
	// The product_group.id of the version 1.0 application group. This parameter is valid only when sourceOrigin is set to synced_from_1_0.
	OriginGroupId *string `json:"originGroupId,omitempty" xml:"originGroupId,omitempty"`
	// The region ID to which the group belongs.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The Alibaba Cloud resource group ID.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The data source. Valid values:
	//
	// - native_2_0: created natively in version 2.0.
	//
	// - synced_from_1_0: synchronized from a version 1.0 application group.
	SourceOrigin *string `json:"sourceOrigin,omitempty" xml:"sourceOrigin,omitempty"`
	// The resource tags (Alibaba Cloud standard tags) as an array of key-value pairs.
	Tags []*ObserveGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace ID to which the group belongs. This value is set at the workspace level and cannot be changed after creation.
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

func (s *ObserveGroup) GetDiscoverRules() []*ObserveGroupDiscoverRule {
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

func (s *ObserveGroup) GetOgEntityInfoEnabled() *bool {
	return s.OgEntityInfoEnabled
}

func (s *ObserveGroup) GetOgEntityInfoPromInstances() []*ObserveGroupPromInstance {
	return s.OgEntityInfoPromInstances
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

func (s *ObserveGroup) SetDiscoverRules(v []*ObserveGroupDiscoverRule) *ObserveGroup {
	s.DiscoverRules = v
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

func (s *ObserveGroup) SetOgEntityInfoEnabled(v bool) *ObserveGroup {
	s.OgEntityInfoEnabled = &v
	return s
}

func (s *ObserveGroup) SetOgEntityInfoPromInstances(v []*ObserveGroupPromInstance) *ObserveGroup {
	s.OgEntityInfoPromInstances = v
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
	if s.DiscoverRules != nil {
		for _, item := range s.DiscoverRules {
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
