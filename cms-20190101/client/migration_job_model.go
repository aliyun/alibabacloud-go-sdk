// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrationJob interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *MigrationJob
	GetCreateTime() *string
	SetDetail(v string) *MigrationJob
	GetDetail() *string
	SetJobStatus(v string) *MigrationJob
	GetJobStatus() *string
	SetPlan(v *MigrationJobPlan) *MigrationJob
	GetPlan() *MigrationJobPlan
	SetRuleNames(v []*string) *MigrationJob
	GetRuleNames() []*string
	SetSource(v []*MigrationJobSource) *MigrationJob
	GetSource() []*MigrationJobSource
	SetUpdateTime(v string) *MigrationJob
	GetUpdateTime() *string
	SetUuid(v string) *MigrationJob
	GetUuid() *string
}

type MigrationJob struct {
	CreateTime *string               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Detail     *string               `json:"Detail,omitempty" xml:"Detail,omitempty"`
	JobStatus  *string               `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	Plan       *MigrationJobPlan     `json:"Plan,omitempty" xml:"Plan,omitempty" type:"Struct"`
	RuleNames  []*string             `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
	Source     []*MigrationJobSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
	UpdateTime *string               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Uuid       *string               `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s MigrationJob) String() string {
	return dara.Prettify(s)
}

func (s MigrationJob) GoString() string {
	return s.String()
}

func (s *MigrationJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MigrationJob) GetDetail() *string {
	return s.Detail
}

func (s *MigrationJob) GetJobStatus() *string {
	return s.JobStatus
}

func (s *MigrationJob) GetPlan() *MigrationJobPlan {
	return s.Plan
}

func (s *MigrationJob) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *MigrationJob) GetSource() []*MigrationJobSource {
	return s.Source
}

func (s *MigrationJob) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *MigrationJob) GetUuid() *string {
	return s.Uuid
}

func (s *MigrationJob) SetCreateTime(v string) *MigrationJob {
	s.CreateTime = &v
	return s
}

func (s *MigrationJob) SetDetail(v string) *MigrationJob {
	s.Detail = &v
	return s
}

func (s *MigrationJob) SetJobStatus(v string) *MigrationJob {
	s.JobStatus = &v
	return s
}

func (s *MigrationJob) SetPlan(v *MigrationJobPlan) *MigrationJob {
	s.Plan = v
	return s
}

func (s *MigrationJob) SetRuleNames(v []*string) *MigrationJob {
	s.RuleNames = v
	return s
}

func (s *MigrationJob) SetSource(v []*MigrationJobSource) *MigrationJob {
	s.Source = v
	return s
}

func (s *MigrationJob) SetUpdateTime(v string) *MigrationJob {
	s.UpdateTime = &v
	return s
}

func (s *MigrationJob) SetUuid(v string) *MigrationJob {
	s.Uuid = &v
	return s
}

func (s *MigrationJob) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlan struct {
	Contacts      []*MigrationJobPlanContacts      `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	Escalations   []*MigrationJobPlanEscalations   `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Repeated"`
	Groups        []*MigrationJobPlanGroups        `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	RuleNames     []*string                        `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
	Strategies    []*MigrationJobPlanStrategies    `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
	Subscriptions []*MigrationJobPlanSubscriptions `json:"Subscriptions,omitempty" xml:"Subscriptions,omitempty" type:"Repeated"`
	Targets       []*MigrationJobPlanTargets       `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s MigrationJobPlan) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlan) GoString() string {
	return s.String()
}

func (s *MigrationJobPlan) GetContacts() []*MigrationJobPlanContacts {
	return s.Contacts
}

func (s *MigrationJobPlan) GetEscalations() []*MigrationJobPlanEscalations {
	return s.Escalations
}

func (s *MigrationJobPlan) GetGroups() []*MigrationJobPlanGroups {
	return s.Groups
}

func (s *MigrationJobPlan) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *MigrationJobPlan) GetStrategies() []*MigrationJobPlanStrategies {
	return s.Strategies
}

func (s *MigrationJobPlan) GetSubscriptions() []*MigrationJobPlanSubscriptions {
	return s.Subscriptions
}

func (s *MigrationJobPlan) GetTargets() []*MigrationJobPlanTargets {
	return s.Targets
}

func (s *MigrationJobPlan) SetContacts(v []*MigrationJobPlanContacts) *MigrationJobPlan {
	s.Contacts = v
	return s
}

func (s *MigrationJobPlan) SetEscalations(v []*MigrationJobPlanEscalations) *MigrationJobPlan {
	s.Escalations = v
	return s
}

func (s *MigrationJobPlan) SetGroups(v []*MigrationJobPlanGroups) *MigrationJobPlan {
	s.Groups = v
	return s
}

func (s *MigrationJobPlan) SetRuleNames(v []*string) *MigrationJobPlan {
	s.RuleNames = v
	return s
}

func (s *MigrationJobPlan) SetStrategies(v []*MigrationJobPlanStrategies) *MigrationJobPlan {
	s.Strategies = v
	return s
}

func (s *MigrationJobPlan) SetSubscriptions(v []*MigrationJobPlanSubscriptions) *MigrationJobPlan {
	s.Subscriptions = v
	return s
}

func (s *MigrationJobPlan) SetTargets(v []*MigrationJobPlanTargets) *MigrationJobPlan {
	s.Targets = v
	return s
}

func (s *MigrationJobPlan) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanContacts struct {
	Channels []*MigrationJobPlanContactsChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	Name     *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MigrationJobPlanContacts) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanContacts) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanContacts) GetChannels() []*MigrationJobPlanContactsChannels {
	return s.Channels
}

func (s *MigrationJobPlanContacts) GetName() *string {
	return s.Name
}

func (s *MigrationJobPlanContacts) SetChannels(v []*MigrationJobPlanContactsChannels) *MigrationJobPlanContacts {
	s.Channels = v
	return s
}

func (s *MigrationJobPlanContacts) SetName(v string) *MigrationJobPlanContacts {
	s.Name = &v
	return s
}

func (s *MigrationJobPlanContacts) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanContactsChannels struct {
	Level *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MigrationJobPlanContactsChannels) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanContactsChannels) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanContactsChannels) GetLevel() *int64 {
	return s.Level
}

func (s *MigrationJobPlanContactsChannels) GetType() *string {
	return s.Type
}

func (s *MigrationJobPlanContactsChannels) GetValue() *string {
	return s.Value
}

func (s *MigrationJobPlanContactsChannels) SetLevel(v int64) *MigrationJobPlanContactsChannels {
	s.Level = &v
	return s
}

func (s *MigrationJobPlanContactsChannels) SetType(v string) *MigrationJobPlanContactsChannels {
	s.Type = &v
	return s
}

func (s *MigrationJobPlanContactsChannels) SetValue(v string) *MigrationJobPlanContactsChannels {
	s.Value = &v
	return s
}

func (s *MigrationJobPlanContactsChannels) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanEscalations struct {
	Escalations []*MigrationJobPlanEscalationsEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Repeated"`
	Name        *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Uuid        *string                                   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s MigrationJobPlanEscalations) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanEscalations) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanEscalations) GetEscalations() []*MigrationJobPlanEscalationsEscalations {
	return s.Escalations
}

func (s *MigrationJobPlanEscalations) GetName() *string {
	return s.Name
}

func (s *MigrationJobPlanEscalations) GetUuid() *string {
	return s.Uuid
}

func (s *MigrationJobPlanEscalations) SetEscalations(v []*MigrationJobPlanEscalationsEscalations) *MigrationJobPlanEscalations {
	s.Escalations = v
	return s
}

func (s *MigrationJobPlanEscalations) SetName(v string) *MigrationJobPlanEscalations {
	s.Name = &v
	return s
}

func (s *MigrationJobPlanEscalations) SetUuid(v string) *MigrationJobPlanEscalations {
	s.Uuid = &v
	return s
}

func (s *MigrationJobPlanEscalations) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanEscalationsEscalations struct {
	Groups      []*string                                          `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	LevelGroups *MigrationJobPlanEscalationsEscalationsLevelGroups `json:"LevelGroups,omitempty" xml:"LevelGroups,omitempty" type:"Struct"`
}

func (s MigrationJobPlanEscalationsEscalations) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanEscalationsEscalations) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanEscalationsEscalations) GetGroups() []*string {
	return s.Groups
}

func (s *MigrationJobPlanEscalationsEscalations) GetLevelGroups() *MigrationJobPlanEscalationsEscalationsLevelGroups {
	return s.LevelGroups
}

func (s *MigrationJobPlanEscalationsEscalations) SetGroups(v []*string) *MigrationJobPlanEscalationsEscalations {
	s.Groups = v
	return s
}

func (s *MigrationJobPlanEscalationsEscalations) SetLevelGroups(v *MigrationJobPlanEscalationsEscalationsLevelGroups) *MigrationJobPlanEscalationsEscalations {
	s.LevelGroups = v
	return s
}

func (s *MigrationJobPlanEscalationsEscalations) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanEscalationsEscalationsLevelGroups struct {
	Critical []*string `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Repeated"`
	Info     []*string `json:"Info,omitempty" xml:"Info,omitempty" type:"Repeated"`
	Resolved []*string `json:"Resolved,omitempty" xml:"Resolved,omitempty" type:"Repeated"`
	Warning  []*string `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s MigrationJobPlanEscalationsEscalationsLevelGroups) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanEscalationsEscalationsLevelGroups) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) GetCritical() []*string {
	return s.Critical
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) GetInfo() []*string {
	return s.Info
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) GetResolved() []*string {
	return s.Resolved
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) GetWarning() []*string {
	return s.Warning
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) SetCritical(v []*string) *MigrationJobPlanEscalationsEscalationsLevelGroups {
	s.Critical = v
	return s
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) SetInfo(v []*string) *MigrationJobPlanEscalationsEscalationsLevelGroups {
	s.Info = v
	return s
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) SetResolved(v []*string) *MigrationJobPlanEscalationsEscalationsLevelGroups {
	s.Resolved = v
	return s
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) SetWarning(v []*string) *MigrationJobPlanEscalationsEscalationsLevelGroups {
	s.Warning = v
	return s
}

func (s *MigrationJobPlanEscalationsEscalationsLevelGroups) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanGroups struct {
	Contacts []*string `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	Name     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MigrationJobPlanGroups) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanGroups) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanGroups) GetContacts() []*string {
	return s.Contacts
}

func (s *MigrationJobPlanGroups) GetName() *string {
	return s.Name
}

func (s *MigrationJobPlanGroups) SetContacts(v []*string) *MigrationJobPlanGroups {
	s.Contacts = v
	return s
}

func (s *MigrationJobPlanGroups) SetName(v string) *MigrationJobPlanGroups {
	s.Name = &v
	return s
}

func (s *MigrationJobPlanGroups) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanStrategies struct {
	EscalationSetting *MigrationJobPlanStrategiesEscalationSetting `json:"EscalationSetting,omitempty" xml:"EscalationSetting,omitempty" type:"Struct"`
	Name              *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	PushingSetting    *MigrationJobPlanStrategiesPushingSetting    `json:"PushingSetting,omitempty" xml:"PushingSetting,omitempty" type:"Struct"`
}

func (s MigrationJobPlanStrategies) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanStrategies) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanStrategies) GetEscalationSetting() *MigrationJobPlanStrategiesEscalationSetting {
	return s.EscalationSetting
}

func (s *MigrationJobPlanStrategies) GetName() *string {
	return s.Name
}

func (s *MigrationJobPlanStrategies) GetPushingSetting() *MigrationJobPlanStrategiesPushingSetting {
	return s.PushingSetting
}

func (s *MigrationJobPlanStrategies) SetEscalationSetting(v *MigrationJobPlanStrategiesEscalationSetting) *MigrationJobPlanStrategies {
	s.EscalationSetting = v
	return s
}

func (s *MigrationJobPlanStrategies) SetName(v string) *MigrationJobPlanStrategies {
	s.Name = &v
	return s
}

func (s *MigrationJobPlanStrategies) SetPushingSetting(v *MigrationJobPlanStrategiesPushingSetting) *MigrationJobPlanStrategies {
	s.PushingSetting = v
	return s
}

func (s *MigrationJobPlanStrategies) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanStrategiesEscalationSetting struct {
	EscalationUuid *string `json:"escalationUuid,omitempty" xml:"escalationUuid,omitempty"`
}

func (s MigrationJobPlanStrategiesEscalationSetting) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanStrategiesEscalationSetting) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanStrategiesEscalationSetting) GetEscalationUuid() *string {
	return s.EscalationUuid
}

func (s *MigrationJobPlanStrategiesEscalationSetting) SetEscalationUuid(v string) *MigrationJobPlanStrategiesEscalationSetting {
	s.EscalationUuid = &v
	return s
}

func (s *MigrationJobPlanStrategiesEscalationSetting) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanStrategiesPushingSetting struct {
	TargetUuids []*string `json:"TargetUuids,omitempty" xml:"TargetUuids,omitempty" type:"Repeated"`
}

func (s MigrationJobPlanStrategiesPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanStrategiesPushingSetting) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanStrategiesPushingSetting) GetTargetUuids() []*string {
	return s.TargetUuids
}

func (s *MigrationJobPlanStrategiesPushingSetting) SetTargetUuids(v []*string) *MigrationJobPlanStrategiesPushingSetting {
	s.TargetUuids = v
	return s
}

func (s *MigrationJobPlanStrategiesPushingSetting) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanSubscriptions struct {
	Conditions   []*MigrationJobPlanSubscriptionsConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Name         *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	StrategyUuid *string                                    `json:"StrategyUuid,omitempty" xml:"StrategyUuid,omitempty"`
}

func (s MigrationJobPlanSubscriptions) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanSubscriptions) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanSubscriptions) GetConditions() []*MigrationJobPlanSubscriptionsConditions {
	return s.Conditions
}

func (s *MigrationJobPlanSubscriptions) GetName() *string {
	return s.Name
}

func (s *MigrationJobPlanSubscriptions) GetStrategyUuid() *string {
	return s.StrategyUuid
}

func (s *MigrationJobPlanSubscriptions) SetConditions(v []*MigrationJobPlanSubscriptionsConditions) *MigrationJobPlanSubscriptions {
	s.Conditions = v
	return s
}

func (s *MigrationJobPlanSubscriptions) SetName(v string) *MigrationJobPlanSubscriptions {
	s.Name = &v
	return s
}

func (s *MigrationJobPlanSubscriptions) SetStrategyUuid(v string) *MigrationJobPlanSubscriptions {
	s.StrategyUuid = &v
	return s
}

func (s *MigrationJobPlanSubscriptions) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanSubscriptionsConditions struct {
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Op    *string `json:"Op,omitempty" xml:"Op,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MigrationJobPlanSubscriptionsConditions) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanSubscriptionsConditions) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanSubscriptionsConditions) GetField() *string {
	return s.Field
}

func (s *MigrationJobPlanSubscriptionsConditions) GetOp() *string {
	return s.Op
}

func (s *MigrationJobPlanSubscriptionsConditions) GetValue() *string {
	return s.Value
}

func (s *MigrationJobPlanSubscriptionsConditions) SetField(v string) *MigrationJobPlanSubscriptionsConditions {
	s.Field = &v
	return s
}

func (s *MigrationJobPlanSubscriptionsConditions) SetOp(v string) *MigrationJobPlanSubscriptionsConditions {
	s.Op = &v
	return s
}

func (s *MigrationJobPlanSubscriptionsConditions) SetValue(v string) *MigrationJobPlanSubscriptionsConditions {
	s.Value = &v
	return s
}

func (s *MigrationJobPlanSubscriptionsConditions) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanTargets struct {
	Arn               *string                                   `json:"Arn,omitempty" xml:"Arn,omitempty"`
	HttpRequestTarget *MigrationJobPlanTargetsHttpRequestTarget `json:"HttpRequestTarget,omitempty" xml:"HttpRequestTarget,omitempty" type:"Struct"`
	Name              *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Type              *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuid              *string                                   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s MigrationJobPlanTargets) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanTargets) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanTargets) GetArn() *string {
	return s.Arn
}

func (s *MigrationJobPlanTargets) GetHttpRequestTarget() *MigrationJobPlanTargetsHttpRequestTarget {
	return s.HttpRequestTarget
}

func (s *MigrationJobPlanTargets) GetName() *string {
	return s.Name
}

func (s *MigrationJobPlanTargets) GetType() *string {
	return s.Type
}

func (s *MigrationJobPlanTargets) GetUuid() *string {
	return s.Uuid
}

func (s *MigrationJobPlanTargets) SetArn(v string) *MigrationJobPlanTargets {
	s.Arn = &v
	return s
}

func (s *MigrationJobPlanTargets) SetHttpRequestTarget(v *MigrationJobPlanTargetsHttpRequestTarget) *MigrationJobPlanTargets {
	s.HttpRequestTarget = v
	return s
}

func (s *MigrationJobPlanTargets) SetName(v string) *MigrationJobPlanTargets {
	s.Name = &v
	return s
}

func (s *MigrationJobPlanTargets) SetType(v string) *MigrationJobPlanTargets {
	s.Type = &v
	return s
}

func (s *MigrationJobPlanTargets) SetUuid(v string) *MigrationJobPlanTargets {
	s.Uuid = &v
	return s
}

func (s *MigrationJobPlanTargets) Validate() error {
	return dara.Validate(s)
}

type MigrationJobPlanTargetsHttpRequestTarget struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MigrationJobPlanTargetsHttpRequestTarget) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobPlanTargetsHttpRequestTarget) GoString() string {
	return s.String()
}

func (s *MigrationJobPlanTargetsHttpRequestTarget) GetContentType() *string {
	return s.ContentType
}

func (s *MigrationJobPlanTargetsHttpRequestTarget) GetMethod() *string {
	return s.Method
}

func (s *MigrationJobPlanTargetsHttpRequestTarget) GetUrl() *string {
	return s.Url
}

func (s *MigrationJobPlanTargetsHttpRequestTarget) SetContentType(v string) *MigrationJobPlanTargetsHttpRequestTarget {
	s.ContentType = &v
	return s
}

func (s *MigrationJobPlanTargetsHttpRequestTarget) SetMethod(v string) *MigrationJobPlanTargetsHttpRequestTarget {
	s.Method = &v
	return s
}

func (s *MigrationJobPlanTargetsHttpRequestTarget) SetUrl(v string) *MigrationJobPlanTargetsHttpRequestTarget {
	s.Url = &v
	return s
}

func (s *MigrationJobPlanTargetsHttpRequestTarget) Validate() error {
	return dara.Validate(s)
}

type MigrationJobSource struct {
	Rule    *MigrationJobSourceRule      `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	Targets []*MigrationJobSourceTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s MigrationJobSource) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobSource) GoString() string {
	return s.String()
}

func (s *MigrationJobSource) GetRule() *MigrationJobSourceRule {
	return s.Rule
}

func (s *MigrationJobSource) GetTargets() []*MigrationJobSourceTargets {
	return s.Targets
}

func (s *MigrationJobSource) SetRule(v *MigrationJobSourceRule) *MigrationJobSource {
	s.Rule = v
	return s
}

func (s *MigrationJobSource) SetTargets(v []*MigrationJobSourceTargets) *MigrationJobSource {
	s.Targets = v
	return s
}

func (s *MigrationJobSource) Validate() error {
	return dara.Validate(s)
}

type MigrationJobSourceRule struct {
	KeywordFilter  *MigrationJobSourceRuleKeywordFilter    `json:"KeywordFilter,omitempty" xml:"KeywordFilter,omitempty" type:"Struct"`
	Name           *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PrimaryFilters []*MigrationJobSourceRulePrimaryFilters `json:"PrimaryFilters,omitempty" xml:"PrimaryFilters,omitempty" type:"Repeated"`
}

func (s MigrationJobSourceRule) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobSourceRule) GoString() string {
	return s.String()
}

func (s *MigrationJobSourceRule) GetKeywordFilter() *MigrationJobSourceRuleKeywordFilter {
	return s.KeywordFilter
}

func (s *MigrationJobSourceRule) GetName() *string {
	return s.Name
}

func (s *MigrationJobSourceRule) GetPrimaryFilters() []*MigrationJobSourceRulePrimaryFilters {
	return s.PrimaryFilters
}

func (s *MigrationJobSourceRule) SetKeywordFilter(v *MigrationJobSourceRuleKeywordFilter) *MigrationJobSourceRule {
	s.KeywordFilter = v
	return s
}

func (s *MigrationJobSourceRule) SetName(v string) *MigrationJobSourceRule {
	s.Name = &v
	return s
}

func (s *MigrationJobSourceRule) SetPrimaryFilters(v []*MigrationJobSourceRulePrimaryFilters) *MigrationJobSourceRule {
	s.PrimaryFilters = v
	return s
}

func (s *MigrationJobSourceRule) Validate() error {
	return dara.Validate(s)
}

type MigrationJobSourceRuleKeywordFilter struct {
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Relation *string   `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s MigrationJobSourceRuleKeywordFilter) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobSourceRuleKeywordFilter) GoString() string {
	return s.String()
}

func (s *MigrationJobSourceRuleKeywordFilter) GetKeywords() []*string {
	return s.Keywords
}

func (s *MigrationJobSourceRuleKeywordFilter) GetRelation() *string {
	return s.Relation
}

func (s *MigrationJobSourceRuleKeywordFilter) SetKeywords(v []*string) *MigrationJobSourceRuleKeywordFilter {
	s.Keywords = v
	return s
}

func (s *MigrationJobSourceRuleKeywordFilter) SetRelation(v string) *MigrationJobSourceRuleKeywordFilter {
	s.Relation = &v
	return s
}

func (s *MigrationJobSourceRuleKeywordFilter) Validate() error {
	return dara.Validate(s)
}

type MigrationJobSourceRulePrimaryFilters struct {
	Field  *string `json:"Field,omitempty" xml:"Field,omitempty"`
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MigrationJobSourceRulePrimaryFilters) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobSourceRulePrimaryFilters) GoString() string {
	return s.String()
}

func (s *MigrationJobSourceRulePrimaryFilters) GetField() *string {
	return s.Field
}

func (s *MigrationJobSourceRulePrimaryFilters) GetOpType() *string {
	return s.OpType
}

func (s *MigrationJobSourceRulePrimaryFilters) GetValue() *string {
	return s.Value
}

func (s *MigrationJobSourceRulePrimaryFilters) SetField(v string) *MigrationJobSourceRulePrimaryFilters {
	s.Field = &v
	return s
}

func (s *MigrationJobSourceRulePrimaryFilters) SetOpType(v string) *MigrationJobSourceRulePrimaryFilters {
	s.OpType = &v
	return s
}

func (s *MigrationJobSourceRulePrimaryFilters) SetValue(v string) *MigrationJobSourceRulePrimaryFilters {
	s.Value = &v
	return s
}

func (s *MigrationJobSourceRulePrimaryFilters) Validate() error {
	return dara.Validate(s)
}

type MigrationJobSourceTargets struct {
	Content *MigrationJobSourceTargetsContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Type    *string                           `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MigrationJobSourceTargets) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobSourceTargets) GoString() string {
	return s.String()
}

func (s *MigrationJobSourceTargets) GetContent() *MigrationJobSourceTargetsContent {
	return s.Content
}

func (s *MigrationJobSourceTargets) GetType() *string {
	return s.Type
}

func (s *MigrationJobSourceTargets) SetContent(v *MigrationJobSourceTargetsContent) *MigrationJobSourceTargets {
	s.Content = v
	return s
}

func (s *MigrationJobSourceTargets) SetType(v string) *MigrationJobSourceTargets {
	s.Type = &v
	return s
}

func (s *MigrationJobSourceTargets) Validate() error {
	return dara.Validate(s)
}

type MigrationJobSourceTargetsContent struct {
	Group        *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Method       *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourcePath *string `json:"ResourcePath,omitempty" xml:"ResourcePath,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MigrationJobSourceTargetsContent) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobSourceTargetsContent) GoString() string {
	return s.String()
}

func (s *MigrationJobSourceTargetsContent) GetGroup() *string {
	return s.Group
}

func (s *MigrationJobSourceTargetsContent) GetLevel() *string {
	return s.Level
}

func (s *MigrationJobSourceTargetsContent) GetMethod() *string {
	return s.Method
}

func (s *MigrationJobSourceTargetsContent) GetRegion() *string {
	return s.Region
}

func (s *MigrationJobSourceTargetsContent) GetResourcePath() *string {
	return s.ResourcePath
}

func (s *MigrationJobSourceTargetsContent) GetUrl() *string {
	return s.Url
}

func (s *MigrationJobSourceTargetsContent) SetGroup(v string) *MigrationJobSourceTargetsContent {
	s.Group = &v
	return s
}

func (s *MigrationJobSourceTargetsContent) SetLevel(v string) *MigrationJobSourceTargetsContent {
	s.Level = &v
	return s
}

func (s *MigrationJobSourceTargetsContent) SetMethod(v string) *MigrationJobSourceTargetsContent {
	s.Method = &v
	return s
}

func (s *MigrationJobSourceTargetsContent) SetRegion(v string) *MigrationJobSourceTargetsContent {
	s.Region = &v
	return s
}

func (s *MigrationJobSourceTargetsContent) SetResourcePath(v string) *MigrationJobSourceTargetsContent {
	s.ResourcePath = &v
	return s
}

func (s *MigrationJobSourceTargetsContent) SetUrl(v string) *MigrationJobSourceTargetsContent {
	s.Url = &v
	return s
}

func (s *MigrationJobSourceTargetsContent) Validate() error {
	return dara.Validate(s)
}
