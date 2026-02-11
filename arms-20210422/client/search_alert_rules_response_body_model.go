// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchAlertRulesResponseBodyPageBean) *SearchAlertRulesResponseBody
	GetPageBean() *SearchAlertRulesResponseBodyPageBean
	SetRequestId(v string) *SearchAlertRulesResponseBody
	GetRequestId() *string
}

type SearchAlertRulesResponseBody struct {
	PageBean  *SearchAlertRulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBody) GetPageBean() *SearchAlertRulesResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAlertRulesResponseBody) SetPageBean(v *SearchAlertRulesResponseBodyPageBean) *SearchAlertRulesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertRulesResponseBody) SetRequestId(v string) *SearchAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAlertRulesResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchAlertRulesResponseBodyPageBean struct {
	AlertRules []*SearchAlertRulesResponseBodyPageBeanAlertRules `json:"AlertRules,omitempty" xml:"AlertRules,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBean) GetAlertRules() []*SearchAlertRulesResponseBodyPageBeanAlertRules {
	return s.AlertRules
}

func (s *SearchAlertRulesResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAlertRulesResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAlertRulesResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchAlertRulesResponseBodyPageBean) SetAlertRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRules) *SearchAlertRulesResponseBodyPageBean {
	s.AlertRules = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageSize(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) Validate() error {
	if s.AlertRules != nil {
		for _, item := range s.AlertRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAlertRulesResponseBodyPageBeanAlertRules struct {
	AlarmContext       *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext `json:"AlarmContext,omitempty" xml:"AlarmContext,omitempty" type:"Struct"`
	AlertLevel         *string                                                     `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	AlertRule          *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule    `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	AlertTitle         *string                                                     `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	AlertType          *int32                                                      `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertVersion       *int32                                                      `json:"AlertVersion,omitempty" xml:"AlertVersion,omitempty"`
	AlertWay           []*string                                                   `json:"AlertWay,omitempty" xml:"AlertWay,omitempty" type:"Repeated"`
	AlertWays          []*string                                                   `json:"AlertWays,omitempty" xml:"AlertWays,omitempty" type:"Repeated"`
	Config             *string                                                     `json:"Config,omitempty" xml:"Config,omitempty"`
	ContactGroupIdList *string                                                     `json:"ContactGroupIdList,omitempty" xml:"ContactGroupIdList,omitempty"`
	ContactGroupIds    *string                                                     `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	CreateTime         *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id                 *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	MetricParam        *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam  `json:"MetricParam,omitempty" xml:"MetricParam,omitempty" type:"Struct"`
	Notice             *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice       `json:"Notice,omitempty" xml:"Notice,omitempty" type:"Struct"`
	RegionId           *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId             *int64                                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus         *string                                                     `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Title              *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime         *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId             *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlarmContext() *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	return s.AlarmContext
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertRule() *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	return s.AlertRule
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertType() *int32 {
	return s.AlertType
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertVersion() *int32 {
	return s.AlertVersion
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertWay() []*string {
	return s.AlertWay
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertWays() []*string {
	return s.AlertWays
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetConfig() *string {
	return s.Config
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetContactGroupIdList() *string {
	return s.ContactGroupIdList
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetId() *int64 {
	return s.Id
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetMetricParam() *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	return s.MetricParam
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetNotice() *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	return s.Notice
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetStatus() *string {
	return s.Status
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetTitle() *string {
	return s.Title
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetUserId() *string {
	return s.UserId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlarmContext(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlarmContext = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertLevel(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertLevel = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertRule(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertRule = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertType(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertType = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertVersion(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertVersion = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWay(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWay = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWays(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWays = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetConfig(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Config = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIdList(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIdList = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIds(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetCreateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Id = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetMetricParam(v *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.MetricParam = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetNotice(v *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Notice = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetRegionId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Status = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskStatus = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUpdateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUserId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UserId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) Validate() error {
	if s.AlarmContext != nil {
		if err := s.AlarmContext.Validate(); err != nil {
			return err
		}
	}
	if s.AlertRule != nil {
		if err := s.AlertRule.Validate(); err != nil {
			return err
		}
	}
	if s.MetricParam != nil {
		if err := s.MetricParam.Validate(); err != nil {
			return err
		}
	}
	if s.Notice != nil {
		if err := s.Notice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext struct {
	AlarmContentSubTitle *string `json:"AlarmContentSubTitle,omitempty" xml:"AlarmContentSubTitle,omitempty"`
	AlarmContentTemplate *string `json:"AlarmContentTemplate,omitempty" xml:"AlarmContentTemplate,omitempty"`
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SubTitle             *string `json:"SubTitle,omitempty" xml:"SubTitle,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetAlarmContentSubTitle() *string {
	return s.AlarmContentSubTitle
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetAlarmContentTemplate() *string {
	return s.AlarmContentTemplate
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetContent() *string {
	return s.Content
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetSubTitle() *string {
	return s.SubTitle
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentSubTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentTemplate(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentTemplate = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetContent(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.Content = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.SubTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule struct {
	Operator *string                                                         `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Rules    []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GetOperator() *string {
	return s.Operator
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GetRules() []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	return s.Rules
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Rules = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules struct {
	Aggregates *string  `json:"Aggregates,omitempty" xml:"Aggregates,omitempty"`
	Alias      *string  `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Measure    *string  `json:"Measure,omitempty" xml:"Measure,omitempty"`
	NValue     *int32   `json:"NValue,omitempty" xml:"NValue,omitempty"`
	Operator   *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value      *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetAggregates() *string {
	return s.Aggregates
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetAlias() *string {
	return s.Alias
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetMeasure() *string {
	return s.Measure
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetNValue() *int32 {
	return s.NValue
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetOperator() *string {
	return s.Operator
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetValue() *float32 {
	return s.Value
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAggregates(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Aggregates = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAlias(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Alias = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetMeasure(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Measure = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetNValue(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.NValue = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetValue(v float32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Value = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam struct {
	AppGroupId *string                                                                `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	AppId      *string                                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Dimensions []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Pid        *string                                                                `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Type       *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetAppGroupId() *string {
	return s.AppGroupId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetAppId() *string {
	return s.AppId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetDimensions() []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	return s.Dimensions
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetPid() *string {
	return s.Pid
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetType() *string {
	return s.Type
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppGroupId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppGroupId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetDimensions(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Dimensions = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetPid(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GetKey() *string {
	return s.Key
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GetType() *string {
	return s.Type
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GetValue() *string {
	return s.Value
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetKey(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Key = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetValue(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Value = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesNotice struct {
	EndTime         *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NoticeEndTime   *int64 `json:"NoticeEndTime,omitempty" xml:"NoticeEndTime,omitempty"`
	NoticeStartTime *int64 `json:"NoticeStartTime,omitempty" xml:"NoticeStartTime,omitempty"`
	StartTime       *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetNoticeEndTime() *int64 {
	return s.NoticeEndTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetNoticeStartTime() *int64 {
	return s.NoticeStartTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.EndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeEndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeStartTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.StartTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) Validate() error {
	return dara.Validate(s)
}
