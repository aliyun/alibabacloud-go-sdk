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
	// 分组归属的阿里云主账号 UID。
	AliUid *string `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	// 创建时间，UTC 格式 yyyy-MM-ddTHH:mm:ssZ。
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 观测分组描述（业务用途说明）。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 实体发现规则列表，定义分组自动匹配哪些实体。
	DiscoverRules []*ObserveGroupDiscoverRule `json:"discoverRules,omitempty" xml:"discoverRules,omitempty" type:"Repeated"`
	// 分组内实体按类型的统计。
	EntitySummaries []*ObserveGroupDetailEntitySummaries `json:"entitySummaries,omitempty" xml:"entitySummaries,omitempty" type:"Repeated"`
	// 扩展信息（JSON 字符串，承载告警模板/联系人组/暂停策略等）。
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// 当前用户是否已关注该分组。
	Favorited *bool   `json:"favorited,omitempty" xml:"favorited,omitempty"`
	GroupId   *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 观测分组名称（同 workspace 下唯一）。
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 观测分组类型。
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
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
	Tags []*ObserveGroupDetailTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 分组归属的 workspace ID（workspace 级归属，创建后不可变）。
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
	// 实体分类。
	EntityCategory *string `json:"entityCategory,omitempty" xml:"entityCategory,omitempty"`
	// 实体数量。
	EntityCount *int32 `json:"entityCount,omitempty" xml:"entityCount,omitempty"`
	// 实体域。
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// 实体类型。
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
	// 标签键。
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// 标签值。
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
