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
	// The response code. A value of **200*	- indicates a successful response.
	//
	// > Other values indicate a failed response. You can use this field to identify the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the skill group rule configuration.
	Data *GetSkillGroupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. If the request is successful, a value of **successful*	- is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// The caller can use this field to determine whether the request was successful:
	//
	// - **true**: The request was successful.
	//
	// - false or **null**: The request failed.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillGroupConfigResponseBodyData struct {
	// Indicates whether to perform a full-text quality check after the real-time quality check is complete. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
	//
	// example:
	//
	// 1
	AllContentQualityCheck *int32 `json:"AllContentQualityCheck,omitempty" xml:"AllContentQualityCheck,omitempty"`
	// The ID of the rule used for the full-text quality check.
	//
	// example:
	//
	// 223
	AllRids     *string                                         `json:"AllRids,omitempty" xml:"AllRids,omitempty"`
	AllRuleList *GetSkillGroupConfigResponseBodyDataAllRuleList `json:"AllRuleList,omitempty" xml:"AllRuleList,omitempty" type:"Struct"`
	// The time when the configuration was created.
	//
	// example:
	//
	// 2020-12-01T15:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 1212
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This is an internal parameter. You can ignore it.
	//
	// example:
	//
	// xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language model ID.
	//
	// example:
	//
	// 1321
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The language model name.
	//
	// example:
	//
	// xxx
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The configuration name.
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The quality check type. Valid values:
	//
	// - 0: offline
	//
	// - 1: real-time
	//
	// example:
	//
	// 0
	QualityCheckType *int32 `json:"QualityCheckType,omitempty" xml:"QualityCheckType,omitempty"`
	// The quality check rule ID.
	//
	// example:
	//
	// 2332
	Rid      *string                                      `json:"Rid,omitempty" xml:"Rid,omitempty"`
	RuleList *GetSkillGroupConfigResponseBodyDataRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// The source of the skill group. The value is fixed at 0.
	//
	// example:
	//
	// 0
	SkillGroupFrom *int32 `json:"SkillGroupFrom,omitempty" xml:"SkillGroupFrom,omitempty"`
	// The skill group ID.
	//
	// example:
	//
	// 111
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The skill group name.
	//
	// example:
	//
	// xxx
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// The status of the configuration. Valid values: 0 (disabled) and 1 (enabled).
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The configuration type. Valid values: 1 (custom configuration) and 0 (built-in configuration).
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the configuration was last updated.
	//
	// example:
	//
	// 2020-12-01T19:28Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The hotword ID.
	//
	// example:
	//
	// 123
	VocabId *int64 `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	// The hotword name.
	//
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
	if s.AllRuleList != nil {
		if err := s.AllRuleList.Validate(); err != nil {
			return err
		}
	}
	if s.RuleList != nil {
		if err := s.RuleList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.RuleNameInfo != nil {
		for _, item := range s.RuleNameInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSkillGroupConfigResponseBodyDataAllRuleListRuleNameInfo struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
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
	if s.RuleNameInfo != nil {
		for _, item := range s.RuleNameInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSkillGroupConfigResponseBodyDataRuleListRuleNameInfo struct {
	Rid      *int64  `json:"Rid,omitempty" xml:"Rid,omitempty"`
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
