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
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	Data *ListIntentionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
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
	// Bot ID
	//
	// example:
	//
	// chatbot-cn-n7QmzrUnNe
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// Intent list
	IntentList []*ListIntentionsResponseBodyDataIntentList `json:"IntentList,omitempty" xml:"IntentList,omitempty" type:"Repeated"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation succeeded.
	//
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
	// Intent alias
	Alias []*string `json:"Alias,omitempty" xml:"Alias,omitempty" type:"Repeated"`
	// Bot ID [Deprecated]
	//
	// example:
	//
	// chatbot-cn-n7QmzrUnNe
	BotId *int64 `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// bot Name
	//
	// example:
	//
	// 111
	BotName *string `json:"BotName,omitempty" xml:"BotName,omitempty"`
	// Conversation flow ID
	//
	// example:
	//
	// 50099
	DialogId *string `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 10717802
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Language
	//
	// - English en-us
	//
	// - Chinese zh-cn
	//
	// example:
	//
	// zh-cn
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Intent name
	//
	// example:
	//
	// 知道了
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// LGF intent expression
	RuleCheck []*ListIntentionsResponseBodyDataIntentListRuleCheck `json:"RuleCheck,omitempty" xml:"RuleCheck,omitempty" type:"Repeated"`
	// Intent slot information
	Slot []*ListIntentionsResponseBodyDataIntentListSlot `json:"Slot,omitempty" xml:"Slot,omitempty" type:"Repeated"`
	// Table ID
	//
	// example:
	//
	// 43258
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// Intent type:
	//
	// - 0: Normal intent;
	//
	// - 1: UNKNOWN;
	//
	// - 2: TableQA intent;
	//
	// - 3: Generated from standard intent
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// User utterance list
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
	// error message
	Error []*string `json:"Error,omitempty" xml:"Error,omitempty" type:"Repeated"`
	// is strict match
	//
	// example:
	//
	// true
	Strict *bool `json:"Strict,omitempty" xml:"Strict,omitempty"`
	// expression value
	//
	// example:
	//
	// 知道
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// warning message
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
	// Follow-up question function
	//
	// 	Notice: Invalid content
	FeedbackFunctions []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctions `json:"FeedbackFunctions,omitempty" xml:"FeedbackFunctions,omitempty" type:"Repeated"`
	// Feedback type
	//
	// 	Notice: Invalid content
	//
	// example:
	//
	// test
	FeedbackType *string `json:"FeedbackType,omitempty" xml:"FeedbackType,omitempty"`
	// Unique slot identity
	//
	// example:
	//
	// 9ec31b50-32b8-11eb-9478-19d2d885afdb
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Is array:
	//
	// - true: Yes
	//
	// - false: No
	//
	// example:
	//
	// false
	IsArray *bool `json:"IsArray,omitempty" xml:"IsArray,omitempty"`
	// Whether encrypted
	//
	// 	Notice: Invalid content
	//
	// example:
	//
	// false
	IsEncrypt *bool `json:"IsEncrypt,omitempty" xml:"IsEncrypt,omitempty"`
	// Is interactive
	//
	// 	Notice: Invalid content
	//
	// example:
	//
	// false
	IsInteractive *bool `json:"IsInteractive,omitempty" xml:"IsInteractive,omitempty"`
	// Is the slot Required:
	//
	// - true: Yes
	//
	// - false: No
	//
	// example:
	//
	// true
	IsNecessary *bool `json:"IsNecessary,omitempty" xml:"IsNecessary,omitempty"`
	// Slot lifecycle
	//
	// example:
	//
	// 0
	LifeSpan *int32 `json:"LifeSpan,omitempty" xml:"LifeSpan,omitempty"`
	// Slot Name
	//
	// example:
	//
	// 知道
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Slot follow-up question script
	Question []*string `json:"Question,omitempty" xml:"Question,omitempty" type:"Repeated"`
	// Slot tagging labels
	Tags []*ListIntentionsResponseBodyDataIntentListSlotTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Slot tagging Result
	//
	// example:
	//
	// @知道
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// Alibaba Cloud function name
	//
	// example:
	//
	// test
	AliyunFunction *string `json:"AliyunFunction,omitempty" xml:"AliyunFunction,omitempty"`
	// Alibaba Cloud service
	//
	// example:
	//
	// test
	AliyunService *string `json:"AliyunService,omitempty" xml:"AliyunService,omitempty"`
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Follow-up question description
	//
	// example:
	//
	// GA setup for HPC cn4-HPC-EndUserServer-GlobalAccelerator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Access edge zone
	//
	// example:
	//
	// cn-hangzhou.log.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// Follow-up question function
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// Function Name
	//
	// example:
	//
	// 方欣云呼系统每日拨测_2024年11月
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Parameter
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// Feature toggle. Valid values:
	//
	// - **on**: Enabled
	//
	// - **off**: shutdown
	Switch []*ListIntentionsResponseBodyDataIntentListSlotFeedbackFunctionsSwitch `json:"Switch,omitempty" xml:"Switch,omitempty" type:"Repeated"`
	// Follow-up question type
	//
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
	// Condition ID
	//
	// example:
	//
	// b9932604-08ae-4525-bbe5-c8cce3066070
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Follow-up question label
	//
	// example:
	//
	// SQL_SUB_QUERY
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Condition name
	//
	// example:
	//
	// 测试0609_20241021_101018_复制_复制_复制
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Type
	//
	// example:
	//
	// PASSKEY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Conditional value
	//
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
	// Unique identity of the intent expression corresponding to the label
	//
	// example:
	//
	// 17448458
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
	// Tagging result of the intent expression corresponding to the label
	//
	// example:
	//
	// 你知道xxxx？
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// Source ID
	//
	// 	Notice: Invalid content
	//
	// example:
	//
	// 1234567
	FromId *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
	// Utterance ID
	//
	// example:
	//
	// 17448458
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Whether to use strict matching
	//
	// example:
	//
	// true
	Strict *bool `json:"Strict,omitempty" xml:"Strict,omitempty"`
	// Utterance list
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
	// Slot identity
	//
	// example:
	//
	// 9ec31b50-32b8-11eb-9478-19d2d885afdb
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// Expression
	//
	// example:
	//
	// 知道
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
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
