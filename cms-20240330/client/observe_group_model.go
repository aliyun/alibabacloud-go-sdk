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
	// 分组归属的阿里云主账号 UID。
	AliUid *string `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	// 创建时间，UTC 格式 yyyy-MM-ddTHH:mm:ssZ。
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 观测分组描述（业务用途说明）。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 实体发现规则列表，定义分组自动匹配哪些实体。
	DiscoverRules []*ObserveGroupDiscoverRule `json:"discoverRules,omitempty" xml:"discoverRules,omitempty" type:"Repeated"`
	// 各分类实体数量。key=分类 domain（acs=云产品 / apm / rum，可扩展），value=该分类下属于本分组的实体数。仅 withEntityCount=true 时返回。
	EntityCounts map[string]*int32 `json:"entityCounts,omitempty" xml:"entityCounts,omitempty"`
	// 扩展信息（JSON 字符串，承载告警模板/联系人组/暂停策略等）。
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// 当前用户是否已关注该分组（「我的关注」筛选依据）。
	Favorited *bool `json:"favorited,omitempty" xml:"favorited,omitempty"`
	// 观测分组全局唯一 ID，格式 og-<16位hash>，指标/告警/控制台统一使用。
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 观测分组名称（同 workspace 下唯一）。
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 观测分组类型。
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// 分组健康状态：-1 未知（占位）/ 1 健康 / 0 不健康。
	Health *int32 `json:"health,omitempty" xml:"health,omitempty"`
	// 最后修改时间，UTC 格式 yyyy-MM-ddTHH:mm:ssZ。资源任意属性变更时自动更新。
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 是否开启 og_entity_info 指标输出。开启后数据面会把本分组归属写入目标 Prometheus 实例。
	OgEntityInfoEnabled *bool `json:"ogEntityInfoEnabled,omitempty" xml:"ogEntityInfoEnabled,omitempty"`
	// 写入 og_entity_info 的 Prometheus 实例集合，包含系统自动识别(system)和用户自定义(custom)两类来源。
	OgEntityInfoPromInstances []*ObserveGroupPromInstance `json:"ogEntityInfoPromInstances,omitempty" xml:"ogEntityInfoPromInstances,omitempty" type:"Repeated"`
	// 仅 synced_from_1_0 时填值，存 1.0 应用分组的 product_group.id。
	OriginGroupId *string `json:"originGroupId,omitempty" xml:"originGroupId,omitempty"`
	// 分组所属地域 ID。
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 阿里云资源组 ID。
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 数据来源：native_2_0（2.0 原生）/ synced_from_1_0（从 1.0 应用分组同步）。
	SourceOrigin *string `json:"sourceOrigin,omitempty" xml:"sourceOrigin,omitempty"`
	// 资源标签（阿里云标准标签），键值对数组。
	Tags []*ObserveGroupTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 分组归属的 workspace ID（workspace 级归属，创建后不可变）。
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
	// 标签键。
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// 标签值。
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
