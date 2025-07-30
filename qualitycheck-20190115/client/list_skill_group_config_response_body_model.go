// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSkillGroupConfigResponseBody
	GetCode() *string
	SetData(v *ListSkillGroupConfigResponseBodyData) *ListSkillGroupConfigResponseBody
	GetData() *ListSkillGroupConfigResponseBodyData
	SetMessage(v string) *ListSkillGroupConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSkillGroupConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSkillGroupConfigResponseBody
	GetSuccess() *bool
}

type ListSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListSkillGroupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSkillGroupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSkillGroupConfigResponseBody) GetData() *ListSkillGroupConfigResponseBodyData {
	return s.Data
}

func (s *ListSkillGroupConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSkillGroupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillGroupConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSkillGroupConfigResponseBody) SetCode(v string) *ListSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetData(v *ListSkillGroupConfigResponseBodyData) *ListSkillGroupConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetMessage(v string) *ListSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetRequestId(v string) *ListSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) SetSuccess(v bool) *ListSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListSkillGroupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyData struct {
	SkillGroupConfig []*ListSkillGroupConfigResponseBodyDataSkillGroupConfig `json:"SkillGroupConfig,omitempty" xml:"SkillGroupConfig,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyData) GetSkillGroupConfig() []*ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	return s.SkillGroupConfig
}

func (s *ListSkillGroupConfigResponseBodyData) SetSkillGroupConfig(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfig) *ListSkillGroupConfigResponseBodyData {
	s.SkillGroupConfig = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfig struct {
	// example:
	//
	// 1
	AllContentQualityCheck *int32 `json:"AllContentQualityCheck,omitempty" xml:"AllContentQualityCheck,omitempty"`
	// example:
	//
	// 223
	AllRids     *string                                                          `json:"AllRids,omitempty" xml:"AllRids,omitempty"`
	AllRuleList *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList `json:"AllRuleList,omitempty" xml:"AllRuleList,omitempty" type:"Struct"`
	// example:
	//
	// 2020-12-01T15:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 221
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 211
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// xxx
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	QualityCheckType *int32 `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	// example:
	//
	// 2221
	Rid      *string                                                       `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleList *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// example:
	//
	// true
	ScreenSwitch *bool `json:"ScreenSwitch,omitempty" xml:"ScreenSwitch,omitempty"`
	// example:
	//
	// 0
	SkillGroupFrom *int32 `json:"SkillGroupFrom,omitempty" xml:"SkillGroupFrom,omitempty"`
	// example:
	//
	// 123
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// xxx
	SkillGroupName    *string                                                                `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	SkillGroupScreens *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens `json:"SkillGroupScreens,omitempty" xml:"SkillGroupScreens,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2020-12-01T19:28Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 323
	VocabId *int64 `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	// example:
	//
	// xxx
	VocabName *string `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetAllContentQualityCheck() *int32 {
	return s.AllContentQualityCheck
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetAllRids() *string {
	return s.AllRids
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetAllRuleList() *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList {
	return s.AllRuleList
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetId() *int64 {
	return s.Id
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetModelId() *int64 {
	return s.ModelId
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetModelName() *string {
	return s.ModelName
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetName() *string {
	return s.Name
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetQualityCheckType() *int32 {
	return s.QualityCheckType
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetRid() *string {
	return s.Rid
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetRuleList() *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList {
	return s.RuleList
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetScreenSwitch() *bool {
	return s.ScreenSwitch
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetSkillGroupFrom() *int32 {
	return s.SkillGroupFrom
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetSkillGroupScreens() *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens {
	return s.SkillGroupScreens
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetStatus() *int32 {
	return s.Status
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetType() *int32 {
	return s.Type
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetVocabId() *int64 {
	return s.VocabId
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) GetVocabName() *string {
	return s.VocabName
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllContentQualityCheck(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllContentQualityCheck = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllRids(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllRids = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetAllRuleList(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.AllRuleList = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetCreateTime(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.CreateTime = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Id = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetInstanceId(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetModelId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ModelId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetModelName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ModelName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Name = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetQualityCheckType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.QualityCheckType = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetRid(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetRuleList(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.RuleList = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetScreenSwitch(v bool) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.ScreenSwitch = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupFrom(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupFrom = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupId(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetSkillGroupScreens(v *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.SkillGroupScreens = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetStatus(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Status = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.Type = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetUpdateTime(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.UpdateTime = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetVocabId(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.VocabId = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) SetVocabName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfig {
	s.VocabName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfig) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList struct {
	RuleNameInfo []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) GetRuleNameInfo() []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo {
	return s.RuleNameInfo
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) SetRuleNameInfo(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList {
	s.RuleNameInfo = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleList) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo struct {
	// example:
	//
	// 221
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) GetRid() *int64 {
	return s.Rid
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) SetRid(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) SetRuleName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigAllRuleListRuleNameInfo) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList struct {
	RuleNameInfo []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) GetRuleNameInfo() []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo {
	return s.RuleNameInfo
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) SetRuleNameInfo(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList {
	s.RuleNameInfo = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleList) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo struct {
	// example:
	//
	// 2221
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// x\\"x\\"x
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) GetRid() *int64 {
	return s.Rid
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) SetRid(v int64) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) SetRuleName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigRuleListRuleNameInfo) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens struct {
	SkillGroupScreen []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen `json:"SkillGroupScreen,omitempty" xml:"SkillGroupScreen,omitempty" type:"Repeated"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) GetSkillGroupScreen() []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	return s.SkillGroupScreen
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) SetSkillGroupScreen(v []*ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens {
	s.SkillGroupScreen = v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreens) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen struct {
	// example:
	//
	// 0
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// customerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Symbol *int32  `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) GetDataType() *int32 {
	return s.DataType
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) GetName() *string {
	return s.Name
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) GetSymbol() *int32 {
	return s.Symbol
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) GetValue() *string {
	return s.Value
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetDataType(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.DataType = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetName(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Name = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetSymbol(v int32) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Symbol = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) SetValue(v string) *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen {
	s.Value = &v
	return s
}

func (s *ListSkillGroupConfigResponseBodyDataSkillGroupConfigSkillGroupScreensSkillGroupScreen) Validate() error {
	return dara.Validate(s)
}
