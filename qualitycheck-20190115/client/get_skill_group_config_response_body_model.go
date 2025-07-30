// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupConfigResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupConfigResponseBodyData) *GetSkillGroupConfigResponseBody
	GetData() *GetSkillGroupConfigResponseBodyData
	SetMessage(v string) *GetSkillGroupConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSkillGroupConfigResponseBody
	GetSuccess() *bool
}

type GetSkillGroupConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetSkillGroupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupConfigResponseBody) GetData() *GetSkillGroupConfigResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSkillGroupConfigResponseBody) SetCode(v string) *GetSkillGroupConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetData(v *GetSkillGroupConfigResponseBodyData) *GetSkillGroupConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetMessage(v string) *GetSkillGroupConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetRequestId(v string) *GetSkillGroupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) SetSuccess(v bool) *GetSkillGroupConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSkillGroupConfigResponseBodyData struct {
	// example:
	//
	// 1
	AllContentQualityCheck *int32 `json:"AllContentQualityCheck,omitempty" xml:"AllContentQualityCheck,omitempty"`
	// example:
	//
	// 223
	AllRids     *string                                         `json:"AllRids,omitempty" xml:"AllRids,omitempty"`
	AllRuleList *GetSkillGroupConfigResponseBodyDataAllRuleList `json:"AllRuleList,omitempty" xml:"AllRuleList,omitempty" type:"Struct"`
	// example:
	//
	// 2020-12-01T15:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1212
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1321
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
	// 2332
	Rid      *string                                      `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleList *GetSkillGroupConfigResponseBodyDataRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SkillGroupFrom *int32 `json:"SkillGroupFrom,omitempty" xml:"SkillGroupFrom,omitempty"`
	// example:
	//
	// 111
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// xxx
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// example:
	//
	// 0
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
	// 123
	VocabId *int64 `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	// example:
	//
	// test
	VocabName *string `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyData) GetAllContentQualityCheck() *int32 {
	return s.AllContentQualityCheck
}

func (s *GetSkillGroupConfigResponseBodyData) GetAllRids() *string {
	return s.AllRids
}

func (s *GetSkillGroupConfigResponseBodyData) GetAllRuleList() *GetSkillGroupConfigResponseBodyDataAllRuleList {
	return s.AllRuleList
}

func (s *GetSkillGroupConfigResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSkillGroupConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetSkillGroupConfigResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSkillGroupConfigResponseBodyData) GetModelId() *int64 {
	return s.ModelId
}

func (s *GetSkillGroupConfigResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *GetSkillGroupConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSkillGroupConfigResponseBodyData) GetQualityCheckType() *int32 {
	return s.QualityCheckType
}

func (s *GetSkillGroupConfigResponseBodyData) GetRid() *string {
	return s.Rid
}

func (s *GetSkillGroupConfigResponseBodyData) GetRuleList() *GetSkillGroupConfigResponseBodyDataRuleList {
	return s.RuleList
}

func (s *GetSkillGroupConfigResponseBodyData) GetSkillGroupFrom() *int32 {
	return s.SkillGroupFrom
}

func (s *GetSkillGroupConfigResponseBodyData) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *GetSkillGroupConfigResponseBodyData) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *GetSkillGroupConfigResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetSkillGroupConfigResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *GetSkillGroupConfigResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSkillGroupConfigResponseBodyData) GetVocabId() *int64 {
	return s.VocabId
}

func (s *GetSkillGroupConfigResponseBodyData) GetVocabName() *string {
	return s.VocabName
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllContentQualityCheck(v int32) *GetSkillGroupConfigResponseBodyData {
	s.AllContentQualityCheck = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllRids(v string) *GetSkillGroupConfigResponseBodyData {
	s.AllRids = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetAllRuleList(v *GetSkillGroupConfigResponseBodyDataAllRuleList) *GetSkillGroupConfigResponseBodyData {
	s.AllRuleList = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetCreateTime(v string) *GetSkillGroupConfigResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetInstanceId(v string) *GetSkillGroupConfigResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetModelId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetModelName(v string) *GetSkillGroupConfigResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetName(v string) *GetSkillGroupConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetQualityCheckType(v int32) *GetSkillGroupConfigResponseBodyData {
	s.QualityCheckType = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetRid(v string) *GetSkillGroupConfigResponseBodyData {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetRuleList(v *GetSkillGroupConfigResponseBodyDataRuleList) *GetSkillGroupConfigResponseBodyData {
	s.RuleList = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupFrom(v int32) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupFrom = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupId(v string) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetSkillGroupName(v string) *GetSkillGroupConfigResponseBodyData {
	s.SkillGroupName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetStatus(v int32) *GetSkillGroupConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetType(v int32) *GetSkillGroupConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetUpdateTime(v string) *GetSkillGroupConfigResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetVocabId(v int64) *GetSkillGroupConfigResponseBodyData {
	s.VocabId = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) SetVocabName(v string) *GetSkillGroupConfigResponseBodyData {
	s.VocabName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSkillGroupConfigResponseBodyDataAllRuleList struct {
	RuleNameInfo []*GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleList) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleList) GetRuleNameInfo() []*GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo {
	return s.RuleNameInfo
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleList) SetRuleNameInfo(v []*GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) *GetSkillGroupConfigResponseBodyDataAllRuleList {
	s.RuleNameInfo = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleList) Validate() error {
	return dara.Validate(s)
}

type GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo struct {
	// example:
	//
	// 12
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) GetRid() *int64 {
	return s.Rid
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) GetRuleName() *string {
	return s.RuleName
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) SetRid(v int64) *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) SetRuleName(v string) *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo) Validate() error {
	return dara.Validate(s)
}

type GetSkillGroupConfigResponseBodyDataRuleList struct {
	RuleNameInfo []*GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo `json:"RuleNameInfo,omitempty" xml:"RuleNameInfo,omitempty" type:"Repeated"`
}

func (s GetSkillGroupConfigResponseBodyDataRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataRuleList) GetRuleNameInfo() []*GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo {
	return s.RuleNameInfo
}

func (s *GetSkillGroupConfigResponseBodyDataRuleList) SetRuleNameInfo(v []*GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) *GetSkillGroupConfigResponseBodyDataRuleList {
	s.RuleNameInfo = v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataRuleList) Validate() error {
	return dara.Validate(s)
}

type GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo struct {
	// example:
	//
	// 222
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) GoString() string {
	return s.String()
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) GetRid() *int64 {
	return s.Rid
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) GetRuleName() *string {
	return s.RuleName
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) SetRid(v int64) *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo {
	s.Rid = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) SetRuleName(v string) *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo {
	s.RuleName = &v
	return s
}

func (s *GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo) Validate() error {
	return dara.Validate(s)
}
