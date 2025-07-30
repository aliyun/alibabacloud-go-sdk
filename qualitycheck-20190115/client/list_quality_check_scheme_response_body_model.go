// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityCheckSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQualityCheckSchemeResponseBody
	GetCode() *string
	SetCount(v int32) *ListQualityCheckSchemeResponseBody
	GetCount() *int32
	SetData(v []*ListQualityCheckSchemeResponseBodyData) *ListQualityCheckSchemeResponseBody
	GetData() []*ListQualityCheckSchemeResponseBodyData
	SetMessage(v string) *ListQualityCheckSchemeResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListQualityCheckSchemeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQualityCheckSchemeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListQualityCheckSchemeResponseBody
	GetRequestId() *string
	SetResultCountId(v string) *ListQualityCheckSchemeResponseBody
	GetResultCountId() *string
	SetSuccess(v bool) *ListQualityCheckSchemeResponseBody
	GetSuccess() *bool
}

type ListQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *int32                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*ListQualityCheckSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// XXX
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityCheckSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQualityCheckSchemeResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListQualityCheckSchemeResponseBody) GetData() []*ListQualityCheckSchemeResponseBodyData {
	return s.Data
}

func (s *ListQualityCheckSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQualityCheckSchemeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQualityCheckSchemeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityCheckSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityCheckSchemeResponseBody) GetResultCountId() *string {
	return s.ResultCountId
}

func (s *ListQualityCheckSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityCheckSchemeResponseBody) SetCode(v string) *ListQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetCount(v int32) *ListQualityCheckSchemeResponseBody {
	s.Count = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetData(v []*ListQualityCheckSchemeResponseBodyData) *ListQualityCheckSchemeResponseBody {
	s.Data = v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetMessage(v string) *ListQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetPageNumber(v int32) *ListQualityCheckSchemeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetPageSize(v int32) *ListQualityCheckSchemeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetRequestId(v string) *ListQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetResultCountId(v string) *ListQualityCheckSchemeResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) SetSuccess(v bool) *ListQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListQualityCheckSchemeResponseBodyData struct {
	// example:
	//
	// 2022-05-10T09:34Z
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 1
	DataType    *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleList            []*ListQualityCheckSchemeResponseBodyDataRuleList            `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	SchemeCheckTypeList []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList `json:"SchemeCheckTypeList,omitempty" xml:"SchemeCheckTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 112**
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2022-05-10T10:34Z
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQualityCheckSchemeResponseBodyData) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListQualityCheckSchemeResponseBodyData) GetDataType() *int32 {
	return s.DataType
}

func (s *ListQualityCheckSchemeResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListQualityCheckSchemeResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListQualityCheckSchemeResponseBodyData) GetRuleList() []*ListQualityCheckSchemeResponseBodyDataRuleList {
	return s.RuleList
}

func (s *ListQualityCheckSchemeResponseBodyData) GetSchemeCheckTypeList() []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	return s.SchemeCheckTypeList
}

func (s *ListQualityCheckSchemeResponseBodyData) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *ListQualityCheckSchemeResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListQualityCheckSchemeResponseBodyData) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *ListQualityCheckSchemeResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *ListQualityCheckSchemeResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListQualityCheckSchemeResponseBodyData) GetUpdateUserName() *string {
	return s.UpdateUserName
}

func (s *ListQualityCheckSchemeResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *ListQualityCheckSchemeResponseBodyData) SetCreateTime(v string) *ListQualityCheckSchemeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetCreateUserName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetDataType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.DataType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetDescription(v string) *ListQualityCheckSchemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetRuleList(v []*ListQualityCheckSchemeResponseBodyDataRuleList) *ListQualityCheckSchemeResponseBodyData {
	s.RuleList = v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetSchemeCheckTypeList(v []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) *ListQualityCheckSchemeResponseBodyData {
	s.SchemeCheckTypeList = v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetSchemeId(v int64) *ListQualityCheckSchemeResponseBodyData {
	s.SchemeId = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetStatus(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetTemplateType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.TemplateType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetType(v int32) *ListQualityCheckSchemeResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetUpdateTime(v string) *ListQualityCheckSchemeResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetUpdateUserName(v string) *ListQualityCheckSchemeResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) SetVersion(v int64) *ListQualityCheckSchemeResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListQualityCheckSchemeResponseBodyDataRuleList struct {
	Rules []*ListQualityCheckSchemeResponseBodyDataRuleListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListQualityCheckSchemeResponseBodyDataRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleList) GetRules() []*ListQualityCheckSchemeResponseBodyDataRuleListRules {
	return s.Rules
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleList) SetRules(v []*ListQualityCheckSchemeResponseBodyDataRuleListRules) *ListQualityCheckSchemeResponseBodyDataRuleList {
	s.Rules = v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleList) Validate() error {
	return dara.Validate(s)
}

type ListQualityCheckSchemeResponseBodyDataRuleListRules struct {
	// example:
	//
	// 1
	CheckType *int32  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// example:
	//
	// 2
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// example:
	//
	// 0
	ScoreNumType *int32 `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	// example:
	//
	// 1
	ScoreType *int32 `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	// example:
	//
	// 10
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyDataRuleListRules) String() string {
	return dara.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataRuleListRules) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetCheckType() *int32 {
	return s.CheckType
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetName() *string {
	return s.Name
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetRid() *int64 {
	return s.Rid
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetRuleScoreType() *int32 {
	return s.RuleScoreType
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetScoreNum() *int32 {
	return s.ScoreNum
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetScoreNumType() *int32 {
	return s.ScoreNumType
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetScoreType() *int32 {
	return s.ScoreType
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetCheckType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.CheckType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetName(v string) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.Name = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetRid(v int64) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.Rid = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetRuleScoreType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.RuleScoreType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreNum(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreNum = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreNumType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreNumType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetScoreType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.ScoreType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) SetTargetType(v int32) *ListQualityCheckSchemeResponseBodyDataRuleListRules {
	s.TargetType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataRuleListRules) Validate() error {
	return dara.Validate(s)
}

type ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList struct {
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 20
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 10
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GoString() string {
	return s.String()
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetCheckName() *string {
	return s.CheckName
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetCheckType() *int32 {
	return s.CheckType
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetEnable() *int32 {
	return s.Enable
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetScore() *int32 {
	return s.Score
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckName(v string) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckName = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckType(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetEnable(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Enable = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetScore(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Score = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetTargetType(v int32) *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.TargetType = &v
	return s
}

func (s *ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) Validate() error {
	return dara.Validate(s)
}
