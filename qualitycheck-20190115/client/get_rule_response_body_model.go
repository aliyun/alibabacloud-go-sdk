// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRuleResponseBody
	GetCode() *string
	SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody
	GetData() *GetRuleResponseBodyData
	SetMessage(v string) *GetRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRuleResponseBody
	GetSuccess() *bool
}

type GetRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRuleResponseBody) GetData() *GetRuleResponseBodyData {
	return s.Data
}

func (s *GetRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRuleResponseBody) SetCode(v string) *GetRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleResponseBody) SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleResponseBody) SetMessage(v string) *GetRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleResponseBody) SetSuccess(v bool) *GetRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyData struct {
	Rules *GetRuleResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s GetRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyData) GetRules() *GetRuleResponseBodyDataRules {
	return s.Rules
}

func (s *GetRuleResponseBodyData) SetRules(v *GetRuleResponseBodyDataRules) *GetRuleResponseBodyData {
	s.Rules = v
	return s
}

func (s *GetRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyDataRules struct {
	RuleInfo []*GetRuleResponseBodyDataRulesRuleInfo `json:"RuleInfo,omitempty" xml:"RuleInfo,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRules) GetRuleInfo() []*GetRuleResponseBodyDataRulesRuleInfo {
	return s.RuleInfo
}

func (s *GetRuleResponseBodyDataRules) SetRuleInfo(v []*GetRuleResponseBodyDataRulesRuleInfo) *GetRuleResponseBodyDataRules {
	s.RuleInfo = v
	return s
}

func (s *GetRuleResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyDataRulesRuleInfo struct {
	// example:
	//
	// 1
	AutoReview               *int32                                                        `json:"AutoReview,omitempty" xml:"AutoReview,omitempty"`
	BusinessCategoryNameList *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Struct"`
	Comments                 *string                                                       `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 123
	CreateEmpid *string `json:"CreateEmpid,omitempty" xml:"CreateEmpid,omitempty"`
	// example:
	//
	// 2016-08-05 10:37:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2016-08-05 10:37:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// example:
	//
	// 1
	IsOnline *int32 `json:"IsOnline,omitempty" xml:"IsOnline,omitempty"`
	// example:
	//
	// 123
	LastUpdateEmpid *string `json:"LastUpdateEmpid,omitempty" xml:"LastUpdateEmpid,omitempty"`
	// example:
	//
	// 2019-10-28 14:23:28
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// a && b
	RuleLambda *string `json:"RuleLambda,omitempty" xml:"RuleLambda,omitempty"`
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// example:
	//
	// 123
	ScoreId   *int32  `json:"ScoreId,omitempty" xml:"ScoreId,omitempty"`
	ScoreName *string `json:"ScoreName,omitempty" xml:"ScoreName,omitempty"`
	// example:
	//
	// 22
	ScoreSubId   *int32  `json:"ScoreSubId,omitempty" xml:"ScoreSubId,omitempty"`
	ScoreSubName *string `json:"ScoreSubName,omitempty" xml:"ScoreSubName,omitempty"`
	// example:
	//
	// 2016-08-05 10:37:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// 1
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetRuleResponseBodyDataRulesRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyDataRulesRuleInfo) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetAutoReview() *int32 {
	return s.AutoReview
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetBusinessCategoryNameList() *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList {
	return s.BusinessCategoryNameList
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetComments() *string {
	return s.Comments
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetCreateEmpid() *string {
	return s.CreateEmpid
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetIsDelete() *int32 {
	return s.IsDelete
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetIsOnline() *int32 {
	return s.IsOnline
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetLastUpdateEmpid() *string {
	return s.LastUpdateEmpid
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetName() *string {
	return s.Name
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetRid() *string {
	return s.Rid
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetRuleLambda() *string {
	return s.RuleLambda
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetRuleScoreType() *int32 {
	return s.RuleScoreType
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetScoreId() *int32 {
	return s.ScoreId
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetScoreName() *string {
	return s.ScoreName
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetScoreSubId() *int32 {
	return s.ScoreSubId
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetScoreSubName() *string {
	return s.ScoreSubName
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetStatus() *int32 {
	return s.Status
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetType() *int32 {
	return s.Type
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) GetWeight() *string {
	return s.Weight
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetAutoReview(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.AutoReview = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetBusinessCategoryNameList(v *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) *GetRuleResponseBodyDataRulesRuleInfo {
	s.BusinessCategoryNameList = v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetComments(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Comments = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetCreateEmpid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.CreateEmpid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetCreateTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.CreateTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetEndTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.EndTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetIsDelete(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.IsDelete = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetIsOnline(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.IsOnline = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetLastUpdateEmpid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.LastUpdateEmpid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetLastUpdateTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.LastUpdateTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Name = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRid(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Rid = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRuleLambda(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.RuleLambda = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetRuleScoreType(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.RuleScoreType = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreId(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreId = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreName = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreSubId(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreSubId = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetScoreSubName(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.ScoreSubName = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetStartTime(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.StartTime = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetStatus(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Status = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetType(v int32) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Type = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) SetWeight(v string) *GetRuleResponseBodyDataRulesRuleInfo {
	s.Weight = &v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfo) Validate() error {
	return dara.Validate(s)
}

type GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList struct {
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
}

func (s GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) GetBusinessCategoryNameList() []*string {
	return s.BusinessCategoryNameList
}

func (s *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) SetBusinessCategoryNameList(v []*string) *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList {
	s.BusinessCategoryNameList = v
	return s
}

func (s *GetRuleResponseBodyDataRulesRuleInfoBusinessCategoryNameList) Validate() error {
	return dara.Validate(s)
}
