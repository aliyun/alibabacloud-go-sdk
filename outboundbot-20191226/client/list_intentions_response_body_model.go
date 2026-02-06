// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIntentionsResponseBody
	GetCode() *string
	SetData(v *ListIntentionsResponseBodyData) *ListIntentionsResponseBody
	GetData() *ListIntentionsResponseBodyData
	SetHttpStatusCode(v int32) *ListIntentionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIntentionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIntentionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIntentionsResponseBody
	GetSuccess() *bool
}

type ListIntentionsResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListIntentionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIntentionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIntentionsResponseBody) GetData() *ListIntentionsResponseBodyData {
	return s.Data
}

func (s *ListIntentionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIntentionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIntentionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntentionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIntentionsResponseBody) SetCode(v string) *ListIntentionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntentionsResponseBody) SetData(v *ListIntentionsResponseBodyData) *ListIntentionsResponseBody {
	s.Data = v
	return s
}

func (s *ListIntentionsResponseBody) SetHttpStatusCode(v int32) *ListIntentionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntentionsResponseBody) SetMessage(v string) *ListIntentionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntentionsResponseBody) SetRequestId(v string) *ListIntentionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntentionsResponseBody) SetSuccess(v bool) *ListIntentionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListIntentionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIntentionsResponseBodyData struct {
	// example:
	//
	// chatbot-cn-n7QmzrUnNe
	BotId      *string                                     `json:"BotId,omitempty" xml:"BotId,omitempty"`
	IntentList []*ListIntentionsResponseBodyDataIntentList `json:"IntentList,omitempty" xml:"IntentList,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIntentionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyData) GetBotId() *string {
	return s.BotId
}

func (s *ListIntentionsResponseBodyData) GetIntentList() []*ListIntentionsResponseBodyDataIntentList {
	return s.IntentList
}

func (s *ListIntentionsResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListIntentionsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ListIntentionsResponseBodyData) SetBotId(v string) *ListIntentionsResponseBodyData {
	s.BotId = &v
	return s
}

func (s *ListIntentionsResponseBodyData) SetIntentList(v []*ListIntentionsResponseBodyDataIntentList) *ListIntentionsResponseBodyData {
	s.IntentList = v
	return s
}

func (s *ListIntentionsResponseBodyData) SetMessage(v string) *ListIntentionsResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListIntentionsResponseBodyData) SetSuccess(v bool) *ListIntentionsResponseBodyData {
	s.Success = &v
	return s
}

func (s *ListIntentionsResponseBodyData) Validate() error {
	if s.IntentList != nil {
		for _, item := range s.IntentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntentionsResponseBodyDataIntentList struct {
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// example:
	//
	// chatbot-cn-n7QmzrUnNe
	BotId *int64 `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// 111
	BotName *string `json:"BotName,omitempty" xml:"BotName,omitempty"`
	// example:
	//
	// 50099
	DialogId *string `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// example:
	//
	// 10717802
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// zh-cn
	Language  *string                                              `json:"Language,omitempty" xml:"Language,omitempty"`
	Name      *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleCheck []*ListIntentionsResponseBodyDataIntentListRuleCheck `json:"RuleCheck,omitempty" xml:"RuleCheck,omitempty" type:"Repeated"`
	Slot      []*ListIntentionsResponseBodyDataIntentListSlot      `json:"Slot,omitempty" xml:"Slot,omitempty" type:"Repeated"`
	// example:
	//
	// 43258
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// example:
	//
	// 0
	Type    *int32                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	UserSay []*ListIntentionsResponseBodyDataIntentListUserSay `json:"UserSay,omitempty" xml:"UserSay,omitempty" type:"Repeated"`
}

func (s ListIntentionsResponseBodyDataIntentList) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentList) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentList) GetAlias() []*string {
	return s.Alias
}

func (s *ListIntentionsResponseBodyDataIntentList) GetBotId() *int64 {
	return s.BotId
}

func (s *ListIntentionsResponseBodyDataIntentList) GetBotName() *string {
	return s.BotName
}

func (s *ListIntentionsResponseBodyDataIntentList) GetDialogId() *string {
	return s.DialogId
}

func (s *ListIntentionsResponseBodyDataIntentList) GetId() *int64 {
	return s.Id
}

func (s *ListIntentionsResponseBodyDataIntentList) GetLanguage() *string {
	return s.Language
}

func (s *ListIntentionsResponseBodyDataIntentList) GetName() *string {
	return s.Name
}

func (s *ListIntentionsResponseBodyDataIntentList) GetRuleCheck() []*ListIntentionsResponseBodyDataIntentListRuleCheck {
	return s.RuleCheck
}

func (s *ListIntentionsResponseBodyDataIntentList) GetSlot() []*ListIntentionsResponseBodyDataIntentListSlot {
	return s.Slot
}

func (s *ListIntentionsResponseBodyDataIntentList) GetTableId() *int64 {
	return s.TableId
}

func (s *ListIntentionsResponseBodyDataIntentList) GetType() *int32 {
	return s.Type
}

func (s *ListIntentionsResponseBodyDataIntentList) GetUserSay() []*ListIntentionsResponseBodyDataIntentListUserSay {
	return s.UserSay
}

func (s *ListIntentionsResponseBodyDataIntentList) SetAlias(v []*string) *ListIntentionsResponseBodyDataIntentList {
	s.Alias = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetBotId(v int64) *ListIntentionsResponseBodyDataIntentList {
	s.BotId = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetBotName(v string) *ListIntentionsResponseBodyDataIntentList {
	s.BotName = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetDialogId(v string) *ListIntentionsResponseBodyDataIntentList {
	s.DialogId = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetId(v int64) *ListIntentionsResponseBodyDataIntentList {
	s.Id = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetLanguage(v string) *ListIntentionsResponseBodyDataIntentList {
	s.Language = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetName(v string) *ListIntentionsResponseBodyDataIntentList {
	s.Name = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetRuleCheck(v []*ListIntentionsResponseBodyDataIntentListRuleCheck) *ListIntentionsResponseBodyDataIntentList {
	s.RuleCheck = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetSlot(v []*ListIntentionsResponseBodyDataIntentListSlot) *ListIntentionsResponseBodyDataIntentList {
	s.Slot = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetTableId(v int64) *ListIntentionsResponseBodyDataIntentList {
	s.TableId = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetType(v int32) *ListIntentionsResponseBodyDataIntentList {
	s.Type = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) SetUserSay(v []*ListIntentionsResponseBodyDataIntentListUserSay) *ListIntentionsResponseBodyDataIntentList {
	s.UserSay = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentList) Validate() error {
	if s.RuleCheck != nil {
		for _, item := range s.RuleCheck {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Slot != nil {
		for _, item := range s.Slot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserSay != nil {
		for _, item := range s.UserSay {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntentionsResponseBodyDataIntentListRuleCheck struct {
	Error []*string `json:"Error,omitempty" xml:"Error,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Strict  *bool     `json:"Strict,omitempty" xml:"Strict,omitempty"`
	Text    *string   `json:"Text,omitempty" xml:"Text,omitempty"`
	Warning []*string `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s ListIntentionsResponseBodyDataIntentListRuleCheck) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentListRuleCheck) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) GetError() []*string {
	return s.Error
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) GetStrict() *bool {
	return s.Strict
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) GetText() *string {
	return s.Text
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) GetWarning() []*string {
	return s.Warning
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) SetError(v []*string) *ListIntentionsResponseBodyDataIntentListRuleCheck {
	s.Error = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) SetStrict(v bool) *ListIntentionsResponseBodyDataIntentListRuleCheck {
	s.Strict = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) SetText(v string) *ListIntentionsResponseBodyDataIntentListRuleCheck {
	s.Text = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) SetWarning(v []*string) *ListIntentionsResponseBodyDataIntentListRuleCheck {
	s.Warning = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListRuleCheck) Validate() error {
	return dara.Validate(s)
}

type ListIntentionsResponseBodyDataIntentListSlot struct {
	FeedbackFunctions []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions `json:"FeedbackFunctions,omitempty" xml:"FeedbackFunctions,omitempty" type:"Repeated"`
	// example:
	//
	// test
	FeedbackType *string `json:"FeedbackType,omitempty" xml:"FeedbackType,omitempty"`
	// example:
	//
	// 9ec31b50-32b8-11eb-9478-19d2d885afdb
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsArray *bool `json:"IsArray,omitempty" xml:"IsArray,omitempty"`
	// example:
	//
	// false
	IsEncrypt *bool `json:"IsEncrypt,omitempty" xml:"IsEncrypt,omitempty"`
	// example:
	//
	// false
	IsInteractive *bool `json:"IsInteractive,omitempty" xml:"IsInteractive,omitempty"`
	// example:
	//
	// true
	IsNecessary *bool `json:"IsNecessary,omitempty" xml:"IsNecessary,omitempty"`
	// example:
	//
	// 0
	LifeSpan *int32                                              `json:"LifeSpan,omitempty" xml:"LifeSpan,omitempty"`
	Name     *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Question []*string                                           `json:"Question,omitempty" xml:"Question,omitempty" type:"Repeated"`
	Tags     []*ListIntentionsResponseBodyDataIntentListSlotTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Value    *string                                             `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIntentionsResponseBodyDataIntentListSlot) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentListSlot) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetFeedbackFunctions() []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	return s.FeedbackFunctions
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetFeedbackType() *string {
	return s.FeedbackType
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetId() *string {
	return s.Id
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetIsArray() *bool {
	return s.IsArray
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetIsEncrypt() *bool {
	return s.IsEncrypt
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetIsInteractive() *bool {
	return s.IsInteractive
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetIsNecessary() *bool {
	return s.IsNecessary
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetLifeSpan() *int32 {
	return s.LifeSpan
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetName() *string {
	return s.Name
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetQuestion() []*string {
	return s.Question
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetTags() []*ListIntentionsResponseBodyDataIntentListSlotTags {
	return s.Tags
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) GetValue() *string {
	return s.Value
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetFeedbackFunctions(v []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) *ListIntentionsResponseBodyDataIntentListSlot {
	s.FeedbackFunctions = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetFeedbackType(v string) *ListIntentionsResponseBodyDataIntentListSlot {
	s.FeedbackType = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetId(v string) *ListIntentionsResponseBodyDataIntentListSlot {
	s.Id = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetIsArray(v bool) *ListIntentionsResponseBodyDataIntentListSlot {
	s.IsArray = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetIsEncrypt(v bool) *ListIntentionsResponseBodyDataIntentListSlot {
	s.IsEncrypt = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetIsInteractive(v bool) *ListIntentionsResponseBodyDataIntentListSlot {
	s.IsInteractive = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetIsNecessary(v bool) *ListIntentionsResponseBodyDataIntentListSlot {
	s.IsNecessary = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetLifeSpan(v int32) *ListIntentionsResponseBodyDataIntentListSlot {
	s.LifeSpan = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetName(v string) *ListIntentionsResponseBodyDataIntentListSlot {
	s.Name = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetQuestion(v []*string) *ListIntentionsResponseBodyDataIntentListSlot {
	s.Question = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetTags(v []*ListIntentionsResponseBodyDataIntentListSlotTags) *ListIntentionsResponseBodyDataIntentListSlot {
	s.Tags = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) SetValue(v string) *ListIntentionsResponseBodyDataIntentListSlot {
	s.Value = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlot) Validate() error {
	if s.FeedbackFunctions != nil {
		for _, item := range s.FeedbackFunctions {
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

type ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions struct {
	// example:
	//
	// test
	AliyunFunction *string `json:"AliyunFunction,omitempty" xml:"AliyunFunction,omitempty"`
	// example:
	//
	// test
	AliyunService *string `json:"AliyunService,omitempty" xml:"AliyunService,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// GA setup for HPC cn4-HPC-EndUserServer-GlobalAccelerator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cn-hangzhou.log.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// example:
	//
	// count
	Function *string                                                                `json:"Function,omitempty" xml:"Function,omitempty"`
	Name     *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Params   map[string]interface{}                                                 `json:"Params,omitempty" xml:"Params,omitempty"`
	Switch   []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch `json:"Switch,omitempty" xml:"Switch,omitempty" type:"Repeated"`
	// example:
	//
	// cluster
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetAliyunFunction() *string {
	return s.AliyunFunction
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetAliyunService() *string {
	return s.AliyunService
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetCode() *string {
	return s.Code
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetDescription() *string {
	return s.Description
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetEndPoint() *string {
	return s.EndPoint
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetFunction() *string {
	return s.Function
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetName() *string {
	return s.Name
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetParams() map[string]interface{} {
	return s.Params
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetSwitch() []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch {
	return s.Switch
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) GetType() *string {
	return s.Type
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetAliyunFunction(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.AliyunFunction = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetAliyunService(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.AliyunService = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetCode(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.Code = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetDescription(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.Description = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetEndPoint(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.EndPoint = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetFunction(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.Function = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetName(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.Name = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetParams(v map[string]interface{}) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.Params = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetSwitch(v []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.Switch = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) SetType(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions {
	s.Type = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions) Validate() error {
	if s.Switch != nil {
		for _, item := range s.Switch {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch struct {
	// example:
	//
	// b9932604-08ae-4525-bbe5-c8cce3066070
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// SQL_SUB_QUERY
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PASSKEY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// BASE_VALIDATE_FILTER_SWITCH
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) GetId() *string {
	return s.Id
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) GetLabel() *string {
	return s.Label
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) GetName() *string {
	return s.Name
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) GetType() *string {
	return s.Type
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) GetValue() *string {
	return s.Value
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) SetId(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch {
	s.Id = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) SetLabel(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch {
	s.Label = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) SetName(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch {
	s.Name = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) SetType(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch {
	s.Type = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) SetValue(v string) *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch {
	s.Value = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch) Validate() error {
	return dara.Validate(s)
}

type ListIntentionsResponseBodyDataIntentListSlotTags struct {
	// example:
	//
	// 17448458
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIntentionsResponseBodyDataIntentListSlotTags) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentListSlotTags) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentListSlotTags) GetUserSayId() *string {
	return s.UserSayId
}

func (s *ListIntentionsResponseBodyDataIntentListSlotTags) GetValue() *string {
	return s.Value
}

func (s *ListIntentionsResponseBodyDataIntentListSlotTags) SetUserSayId(v string) *ListIntentionsResponseBodyDataIntentListSlotTags {
	s.UserSayId = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotTags) SetValue(v string) *ListIntentionsResponseBodyDataIntentListSlotTags {
	s.Value = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListSlotTags) Validate() error {
	return dara.Validate(s)
}

type ListIntentionsResponseBodyDataIntentListUserSay struct {
	// example:
	//
	// 1234567
	FromId *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
	// example:
	//
	// 17448458
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	Strict      *bool                                                         `json:"Strict,omitempty" xml:"Strict,omitempty"`
	UserSayData []*ListIntentionsResponseBodyDataIntentListUserSayUserSayData `json:"UserSayData,omitempty" xml:"UserSayData,omitempty" type:"Repeated"`
}

func (s ListIntentionsResponseBodyDataIntentListUserSay) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentListUserSay) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) GetFromId() *string {
	return s.FromId
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) GetId() *string {
	return s.Id
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) GetStrict() *bool {
	return s.Strict
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) GetUserSayData() []*ListIntentionsResponseBodyDataIntentListUserSayUserSayData {
	return s.UserSayData
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) SetFromId(v string) *ListIntentionsResponseBodyDataIntentListUserSay {
	s.FromId = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) SetId(v string) *ListIntentionsResponseBodyDataIntentListUserSay {
	s.Id = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) SetStrict(v bool) *ListIntentionsResponseBodyDataIntentListUserSay {
	s.Strict = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) SetUserSayData(v []*ListIntentionsResponseBodyDataIntentListUserSayUserSayData) *ListIntentionsResponseBodyDataIntentListUserSay {
	s.UserSayData = v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListUserSay) Validate() error {
	if s.UserSayData != nil {
		for _, item := range s.UserSayData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntentionsResponseBodyDataIntentListUserSayUserSayData struct {
	// example:
	//
	// 9ec31b50-32b8-11eb-9478-19d2d885afdb
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Text   *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListIntentionsResponseBodyDataIntentListUserSayUserSayData) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsResponseBodyDataIntentListUserSayUserSayData) GoString() string {
	return s.String()
}

func (s *ListIntentionsResponseBodyDataIntentListUserSayUserSayData) GetSlotId() *string {
	return s.SlotId
}

func (s *ListIntentionsResponseBodyDataIntentListUserSayUserSayData) GetText() *string {
	return s.Text
}

func (s *ListIntentionsResponseBodyDataIntentListUserSayUserSayData) SetSlotId(v string) *ListIntentionsResponseBodyDataIntentListUserSayUserSayData {
	s.SlotId = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListUserSayUserSayData) SetText(v string) *ListIntentionsResponseBodyDataIntentListUserSayUserSayData {
	s.Text = &v
	return s
}

func (s *ListIntentionsResponseBodyDataIntentListUserSayUserSayData) Validate() error {
	return dara.Validate(s)
}
