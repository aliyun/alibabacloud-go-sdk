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
	// The status code. A value of **200*	- indicates success. Other values indicate a failure. Use this code to identify the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 22
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The data of the quality check scheme list.
	Data []*ListQualityCheckSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message. If the request is successful, \\`successful\\` is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An internal parameter. You can ignore this parameter.
	//
	// example:
	//
	// XXX
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates success. A value of **false*	- or **null*	- indicates failure.
	//
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
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityCheckSchemeResponseBodyData struct {
	// The time when the scheme was created.
	//
	// example:
	//
	// 2022-05-10T09:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// 张三
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// The data type of the quality check scheme. Valid values: 0 (text) and 1 (audio).
	//
	// example:
	//
	// 1
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The description.
	//
	// example:
	//
	// 售前使用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the quality check scheme.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of check items.
	RuleList []*ListQualityCheckSchemeResponseBodyDataRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// The list of quality check dimensions.
	SchemeCheckTypeList []*ListQualityCheckSchemeResponseBodyDataSchemeCheckTypeList `json:"SchemeCheckTypeList,omitempty" xml:"SchemeCheckTypeList,omitempty" type:"Repeated"`
	// The ID of the quality check scheme.
	//
	// example:
	//
	// 112**
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// The status of the quality check scheme. Valid values: 0 (deleted), 1 (published), 2 (unpublished), and 3 (updated but not published).
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the quality check scheme template. Valid values: 1 (built-in) and 2 (user-defined).
	//
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The type of the quality check scheme. Valid values: 0 (built-in) and 1 (user-defined).
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the scheme was last updated.
	//
	// example:
	//
	// 2022-05-10T10:34Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The name of the user who last updated the scheme.
	//
	// example:
	//
	// 李四
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// The version of the quality check scheme.
	//
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
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SchemeCheckTypeList != nil {
		for _, item := range s.SchemeCheckTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityCheckSchemeResponseBodyDataRuleList struct {
	// The rule information. This parameter is reserved for future use. Currently, only one rule is returned.
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

type ListQualityCheckSchemeResponseBodyDataRuleListRules struct {
	// The quality check dimension to which the item belongs.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// 测试规则
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// 12
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Indicates whether to calculate a score. Valid values: 1 (no score) and 3 (score).
	//
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// The score.
	//
	// example:
	//
	// 2
	ScoreNum *int32 `json:"ScoreNum,omitempty" xml:"ScoreNum,omitempty"`
	// The scoring type. Valid values: 0 (Points are added or deducted each time the rule is triggered) and 1 (A one-time score is given when the rule is triggered).
	//
	// example:
	//
	// 0
	ScoreNumType *int32 `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	// The scoring method. Valid values: 1 (add points) and 3 (deduct points).
	//
	// example:
	//
	// 1
	ScoreType *int32 `json:"ScoreType,omitempty" xml:"ScoreType,omitempty"`
	// The application scenario of the check item. Valid values: 10 (common check item) and 11 (SOP flow check item).
	//
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
	// The name of the quality check dimension.
	//
	// example:
	//
	// 服务规范性检测
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The type of the quality check dimension.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The enabled status. Valid values: 0 (disabled) and 1 (enabled).
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The score of the quality check dimension.
	//
	// example:
	//
	// 20
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// An internal parameter. You can ignore this parameter.
	//
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
