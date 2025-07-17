// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListAlertRulesResponseBodyPagingInfo) *ListAlertRulesResponseBody
	GetPagingInfo() *ListAlertRulesResponseBodyPagingInfo
	SetRequestId(v string) *ListAlertRulesResponseBody
	GetRequestId() *string
}

type ListAlertRulesResponseBody struct {
	// The pagination information.
	PagingInfo *ListAlertRulesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A6C6B486-E3A2-5D52-9E76-D9380485D946
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBody) GetPagingInfo() *ListAlertRulesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertRulesResponseBody) SetPagingInfo(v *ListAlertRulesResponseBodyPagingInfo) *ListAlertRulesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListAlertRulesResponseBody) SetRequestId(v string) *ListAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfo struct {
	// The rules.
	AlertRules []*ListAlertRulesResponseBodyPagingInfoAlertRules `json:"AlertRules,omitempty" xml:"AlertRules,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfo) GetAlertRules() []*ListAlertRulesResponseBodyPagingInfoAlertRules {
	return s.AlertRules
}

func (s *ListAlertRulesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlertRulesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertRulesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertRulesResponseBodyPagingInfo) SetAlertRules(v []*ListAlertRulesResponseBodyPagingInfoAlertRules) *ListAlertRulesResponseBodyPagingInfo {
	s.AlertRules = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfo) SetPageNumber(v int32) *ListAlertRulesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfo) SetPageSize(v int32) *ListAlertRulesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfo) SetTotalCount(v int32) *ListAlertRulesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRules struct {
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 22125
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// error_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account used by the owner of the rule.
	//
	// example:
	//
	// 1933790683****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The alert triggering condition.
	TriggerCondition *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition `json:"TriggerCondition,omitempty" xml:"TriggerCondition,omitempty" type:"Struct"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRules) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRules) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) GetId() *int64 {
	return s.Id
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) GetName() *string {
	return s.Name
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) GetOwner() *string {
	return s.Owner
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) GetTriggerCondition() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition {
	return s.TriggerCondition
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) SetEnabled(v bool) *ListAlertRulesResponseBodyPagingInfoAlertRules {
	s.Enabled = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) SetId(v int64) *ListAlertRulesResponseBodyPagingInfoAlertRules {
	s.Id = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) SetName(v string) *ListAlertRulesResponseBodyPagingInfoAlertRules {
	s.Name = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) SetOwner(v string) *ListAlertRulesResponseBodyPagingInfoAlertRules {
	s.Owner = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) SetTriggerCondition(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) *ListAlertRulesResponseBodyPagingInfoAlertRules {
	s.TriggerCondition = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRules) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition struct {
	// The extended information about the rule. This parameter is required for specific types of alerts.
	Extension *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The monitored objects.
	Target *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The alert type. Valid values:
	//
	// 	- Finished: An instance is successfully run.
	//
	// 	- UnFinished: An instance does not finish running before a specified point in time.
	//
	// 	- Error: An error occurs on an instance.
	//
	// 	- CycleUnfinished: An instance does not finish running as expected within a specific cycle.
	//
	// 	- Timeout: An instance times out.
	//
	// 	- InstanceTransferComplete: An instance is generated by the auto triggered node.
	//
	// 	- InstanceTransferFluctuate: The number of generated instances fluctuates.
	//
	// 	- ExhaustedError: An error persists after an instance is automatically rerun.
	//
	// 	- InstanceKeyword: An instance with errors contains specified keywords.
	//
	// 	- InstanceErrorCount: The number of instances on which an error occurs reaches a specified threshold.
	//
	// 	- InstanceErrorPercentage: The proportion of instances on which an error occurs in the workspace to the total number of instances reaches a specified threshold.
	//
	// 	- ResourceGroupPercentage: The usage rate of the resource group reaches a specified threshold.
	//
	// 	- ResourceGroupWaitCount: The number of instances that are waiting for resources in the resource group reaches a specified threshold.
	//
	// example:
	//
	// Error
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) GetExtension() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	return s.Extension
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) GetTarget() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget {
	return s.Target
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) GetType() *string {
	return s.Type
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) SetExtension(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition {
	s.Extension = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) SetTarget(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition {
	s.Target = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) SetType(v string) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition {
	s.Type = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerCondition) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension struct {
	// The configuration for an alert of the CycleUnfinished type.
	CycleUnfinished *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished `json:"CycleUnfinished,omitempty" xml:"CycleUnfinished,omitempty" type:"Struct"`
	// The configuration for an alert of the Error type.
	Error *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorCount type.
	InstanceErrorCount *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount `json:"InstanceErrorCount,omitempty" xml:"InstanceErrorCount,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorPercentage type.
	InstanceErrorPercentage *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage `json:"InstanceErrorPercentage,omitempty" xml:"InstanceErrorPercentage,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceTransferFluctuate type.
	InstanceTransferFluctuate *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate `json:"InstanceTransferFluctuate,omitempty" xml:"InstanceTransferFluctuate,omitempty" type:"Struct"`
	// The configuration for an alert of the Timeout type.
	Timeout *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
	// The configuration for an alert of the UnFinished type.
	UnFinished *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished `json:"UnFinished,omitempty" xml:"UnFinished,omitempty" type:"Struct"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GetCycleUnfinished() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished {
	return s.CycleUnfinished
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GetError() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError {
	return s.Error
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GetInstanceErrorCount() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount {
	return s.InstanceErrorCount
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GetInstanceErrorPercentage() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage {
	return s.InstanceErrorPercentage
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GetInstanceTransferFluctuate() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate {
	return s.InstanceTransferFluctuate
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GetTimeout() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout {
	return s.Timeout
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) GetUnFinished() *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished {
	return s.UnFinished
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) SetCycleUnfinished(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	s.CycleUnfinished = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) SetError(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	s.Error = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) SetInstanceErrorCount(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	s.InstanceErrorCount = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) SetInstanceErrorPercentage(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	s.InstanceErrorPercentage = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) SetInstanceTransferFluctuate(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	s.InstanceTransferFluctuate = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) SetTimeout(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	s.Timeout = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) SetUnFinished(v *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension {
	s.UnFinished = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtension) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished struct {
	// The configurations of the scheduling cycle and timeout period of the instance.
	CycleAndTime []*ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime `json:"CycleAndTime,omitempty" xml:"CycleAndTime,omitempty" type:"Repeated"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished) GetCycleAndTime() []*ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	return s.CycleAndTime
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished) SetCycleAndTime(v []*ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished {
	s.CycleAndTime = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinished) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime struct {
	// The ID of the scheduling cycle of the instance. Valid values: [1,288].
	//
	// example:
	//
	// 1
	CycleId *int32 `json:"CycleId,omitempty" xml:"CycleId,omitempty"`
	// The latest completion time of the instance within the scheduling cycle. The time is in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 01:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetCycleId() *int32 {
	return s.CycleId
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetTime() *string {
	return s.Time
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetCycleId(v int32) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.CycleId = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetTime(v string) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.Time = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionCycleUnfinishedCycleAndTime) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError struct {
	// Indicates whether an alert is triggered if a batch synchronization task is automatically rerun upon a failure.
	//
	// example:
	//
	// false
	AutoRerunAlertEnabled *bool `json:"AutoRerunAlertEnabled,omitempty" xml:"AutoRerunAlertEnabled,omitempty"`
	// The IDs of the real-time computing tasks. This parameter is required when you monitor real-time computing tasks.
	StreamTaskIds []*int64 `json:"StreamTaskIds,omitempty" xml:"StreamTaskIds,omitempty" type:"Repeated"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) GetAutoRerunAlertEnabled() *bool {
	return s.AutoRerunAlertEnabled
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) GetStreamTaskIds() []*int64 {
	return s.StreamTaskIds
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) SetAutoRerunAlertEnabled(v bool) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError {
	s.AutoRerunAlertEnabled = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) SetStreamTaskIds(v []*int64) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError {
	s.StreamTaskIds = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionError) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount struct {
	// The maximum number of instances on which an error occurs. Valid values: [1,10000].
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount) GetCount() *int32 {
	return s.Count
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount) SetCount(v int32) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount {
	s.Count = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorCount) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage struct {
	// The maximum percentage of instances on which an error occurs in the workspace to the total number of instances. Valid values: [1-100].
	//
	// example:
	//
	// 10
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage) GetPercentage() *int32 {
	return s.Percentage
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage) SetPercentage(v int32) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage {
	s.Percentage = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceErrorPercentage) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate struct {
	// The maximum percentage of fluctuation in the number of auto triggered node instances that are generated in your workspace. Valid values: [1-100].
	//
	// example:
	//
	// 10
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The way in which the number of auto triggered node instances that are generated in your workspace fluctuates. Valid values:
	//
	// 	- abs: the absolute value. The number of instances increases or decreases.
	//
	// 	- increase: The number of instances increases.
	//
	// 	- decrease: The number of instances decreases.
	//
	// example:
	//
	// abs
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) GetPercentage() *int32 {
	return s.Percentage
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) GetTrend() *string {
	return s.Trend
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) SetPercentage(v int32) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate {
	s.Percentage = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) SetTrend(v string) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate {
	s.Trend = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionInstanceTransferFluctuate) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout struct {
	// The timeout period. Unit: minutes.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout) SetTimeoutInMinutes(v int32) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout {
	s.TimeoutInMinutes = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionTimeout) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished struct {
	// The latest completion time of the instance. The period is in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 12:00
	UnFinishedTime *string `json:"UnFinishedTime,omitempty" xml:"UnFinishedTime,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished) GetUnFinishedTime() *string {
	return s.UnFinishedTime
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished) SetUnFinishedTime(v string) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished {
	s.UnFinishedTime = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionExtensionUnFinished) Validate() error {
	return dara.Validate(s)
}

type ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget struct {
	// The nodes that are not to be monitored.
	AllowTasks []*int64 `json:"AllowTasks,omitempty" xml:"AllowTasks,omitempty" type:"Repeated"`
	// The IDs of monitored objects.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The type of the monitored objects. Valid values:
	//
	// 	- Task: node
	//
	// 	- Baseline: baseline
	//
	// 	- Project: workspace
	//
	// 	- BizProcess: workflow
	//
	// example:
	//
	// Task
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) String() string {
	return dara.Prettify(s)
}

func (s ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) GoString() string {
	return s.String()
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) GetAllowTasks() []*int64 {
	return s.AllowTasks
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) GetType() *string {
	return s.Type
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) SetAllowTasks(v []*int64) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget {
	s.AllowTasks = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) SetIds(v []*int64) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget {
	s.Ids = v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) SetType(v string) *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget {
	s.Type = &v
	return s
}

func (s *ListAlertRulesResponseBodyPagingInfoAlertRulesTriggerConditionTarget) Validate() error {
	return dara.Validate(s)
}
