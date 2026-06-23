// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetDialogLogResponseBody
	GetCost() *int64
	SetData(v *GetDialogLogResponseBodyData) *GetDialogLogResponseBody
	GetData() *GetDialogLogResponseBodyData
	SetDataType(v string) *GetDialogLogResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetDialogLogResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetDialogLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDialogLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDialogLogResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetDialogLogResponseBody
	GetTime() *string
}

type GetDialogLogResponseBody struct {
	// The processing time.
	//
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// The response data.
	Data *GetDialogLogResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The data type.
	//
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 051EEB18-049A-17FF-A5E0-14A5B127C798
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDialogLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDialogLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetDialogLogResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetDialogLogResponseBody) GetData() *GetDialogLogResponseBodyData {
	return s.Data
}

func (s *GetDialogLogResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetDialogLogResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDialogLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDialogLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDialogLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDialogLogResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetDialogLogResponseBody) SetCost(v int64) *GetDialogLogResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDialogLogResponseBody) SetData(v *GetDialogLogResponseBodyData) *GetDialogLogResponseBody {
	s.Data = v
	return s
}

func (s *GetDialogLogResponseBody) SetDataType(v string) *GetDialogLogResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDialogLogResponseBody) SetErrCode(v string) *GetDialogLogResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDialogLogResponseBody) SetMessage(v string) *GetDialogLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetDialogLogResponseBody) SetRequestId(v string) *GetDialogLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDialogLogResponseBody) SetSuccess(v bool) *GetDialogLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetDialogLogResponseBody) SetTime(v string) *GetDialogLogResponseBody {
	s.Time = &v
	return s
}

func (s *GetDialogLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDialogLogResponseBodyData struct {
	// The analysis process. This field has a value if the analysis process is enabled during the real-time conversation.
	//
	// example:
	//
	// 客户回答的内容与提供的意图列表描述均不匹配，没有表达出对账单、还款、天气或其他服务的具体需求或问题。
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// The `conversationList` field records the conversation content.
	//
	// example:
	//
	// ##客服##:您好，请问是张三先生是吧？\\n ##客户##:人工客服\\n ##客服##:您好，我是2804，很高兴为您服务！\\n ##客服##:您好，请问有什么可以帮到您？\\n ##客户##:好的 谢谢\\n
	ConversationList *string `json:"conversationList,omitempty" xml:"conversationList,omitempty"`
	// The list of hit intents.
	HitIntentionList []*GetDialogLogResponseBodyDataHitIntentionList `json:"hitIntentionList,omitempty" xml:"hitIntentionList,omitempty" type:"Repeated"`
	// The intent list.
	IntentionList []*GetDialogLogResponseBodyDataIntentionList `json:"intentionList,omitempty" xml:"intentionList,omitempty" type:"Repeated"`
	// The model processing time, in milliseconds.
	//
	// example:
	//
	// 1382
	ModelCostTime *int64 `json:"modelCostTime,omitempty" xml:"modelCostTime,omitempty"`
	// The recall list.
	//
	// example:
	//
	// ## Example:\\n- 对话内容为：\\"##客服##:您好，请问有什么可以帮到您？\\n ##客户##:暂时没有了。谢谢。\\"时，用户意图为：\\"客户想要挂断电话\\"\\n- 对话内容为：\\"##客服##:您好，请问有什么可以帮到您？\\n ##客户##:哎你好。\\"时，用户意图为：\\"客户询问来电目的\\"\\n- 对话内容为：\\"##客服##:您好，请问有什么可以帮到您？\\n ##客户##:我现在财务状况很好，谢谢关心。\\"时，用户意图为：\\"客户拒绝贷款\\"\\n- 对话内容为：\\"##客服##:您好，请问有什么可以帮到您？\\n ##客户##:不用了，谢谢，不要再打电话了，谢谢。\\"时，用户意图为：\\"投诉/退订/不要打电话/骂人\\"\\n- 对话内容为：\\"##客服##:您好，请问有什么可以帮到您？\\n ##客户##:你好。\\"时，用户意图为：\\"客户询问来电目的\\"
	RecallList *string `json:"recallList,omitempty" xml:"recallList,omitempty"`
}

func (s GetDialogLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDialogLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDialogLogResponseBodyData) GetAnalysisProcess() *string {
	return s.AnalysisProcess
}

func (s *GetDialogLogResponseBodyData) GetConversationList() *string {
	return s.ConversationList
}

func (s *GetDialogLogResponseBodyData) GetHitIntentionList() []*GetDialogLogResponseBodyDataHitIntentionList {
	return s.HitIntentionList
}

func (s *GetDialogLogResponseBodyData) GetIntentionList() []*GetDialogLogResponseBodyDataIntentionList {
	return s.IntentionList
}

func (s *GetDialogLogResponseBodyData) GetModelCostTime() *int64 {
	return s.ModelCostTime
}

func (s *GetDialogLogResponseBodyData) GetRecallList() *string {
	return s.RecallList
}

func (s *GetDialogLogResponseBodyData) SetAnalysisProcess(v string) *GetDialogLogResponseBodyData {
	s.AnalysisProcess = &v
	return s
}

func (s *GetDialogLogResponseBodyData) SetConversationList(v string) *GetDialogLogResponseBodyData {
	s.ConversationList = &v
	return s
}

func (s *GetDialogLogResponseBodyData) SetHitIntentionList(v []*GetDialogLogResponseBodyDataHitIntentionList) *GetDialogLogResponseBodyData {
	s.HitIntentionList = v
	return s
}

func (s *GetDialogLogResponseBodyData) SetIntentionList(v []*GetDialogLogResponseBodyDataIntentionList) *GetDialogLogResponseBodyData {
	s.IntentionList = v
	return s
}

func (s *GetDialogLogResponseBodyData) SetModelCostTime(v int64) *GetDialogLogResponseBodyData {
	s.ModelCostTime = &v
	return s
}

func (s *GetDialogLogResponseBodyData) SetRecallList(v string) *GetDialogLogResponseBodyData {
	s.RecallList = &v
	return s
}

func (s *GetDialogLogResponseBodyData) Validate() error {
	if s.HitIntentionList != nil {
		for _, item := range s.HitIntentionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IntentionList != nil {
		for _, item := range s.IntentionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDialogLogResponseBodyDataHitIntentionList struct {
	// A description of the customer\\"s intent.
	//
	// example:
	//
	// 客户希望与真人接触，不想和AI客服继续对话。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The intent name.
	//
	// example:
	//
	// 客户要求转人工
	IntentionName *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	// The scripted reply based on the customer\\"s intent.
	//
	// example:
	//
	// 很抱歉，我这里无法直接为您转接，您可以拨打我司客服热线进行咨询。
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s GetDialogLogResponseBodyDataHitIntentionList) String() string {
	return dara.Prettify(s)
}

func (s GetDialogLogResponseBodyDataHitIntentionList) GoString() string {
	return s.String()
}

func (s *GetDialogLogResponseBodyDataHitIntentionList) GetDescription() *string {
	return s.Description
}

func (s *GetDialogLogResponseBodyDataHitIntentionList) GetIntentionName() *string {
	return s.IntentionName
}

func (s *GetDialogLogResponseBodyDataHitIntentionList) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *GetDialogLogResponseBodyDataHitIntentionList) SetDescription(v string) *GetDialogLogResponseBodyDataHitIntentionList {
	s.Description = &v
	return s
}

func (s *GetDialogLogResponseBodyDataHitIntentionList) SetIntentionName(v string) *GetDialogLogResponseBodyDataHitIntentionList {
	s.IntentionName = &v
	return s
}

func (s *GetDialogLogResponseBodyDataHitIntentionList) SetIntentionScript(v string) *GetDialogLogResponseBodyDataHitIntentionList {
	s.IntentionScript = &v
	return s
}

func (s *GetDialogLogResponseBodyDataHitIntentionList) Validate() error {
	return dara.Validate(s)
}

type GetDialogLogResponseBodyDataIntentionList struct {
	// The `description` field provides a detailed description of the user\\"s intent.
	//
	// example:
	//
	// 客户明确表示投诉/退订/不要打电话/骂人等拒绝营销
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The intent name.
	//
	// example:
	//
	// 客户明确表示拒绝营销
	IntentionName *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	// The `intentionScript` field contains the service agent\\"s reply script for the user\\"s intent.
	//
	// example:
	//
	// 非常抱歉，给您带来了不好的体验。如您无需再接受我们的官方来电，请回复“我要退订”四个字！
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s GetDialogLogResponseBodyDataIntentionList) String() string {
	return dara.Prettify(s)
}

func (s GetDialogLogResponseBodyDataIntentionList) GoString() string {
	return s.String()
}

func (s *GetDialogLogResponseBodyDataIntentionList) GetDescription() *string {
	return s.Description
}

func (s *GetDialogLogResponseBodyDataIntentionList) GetIntentionName() *string {
	return s.IntentionName
}

func (s *GetDialogLogResponseBodyDataIntentionList) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *GetDialogLogResponseBodyDataIntentionList) SetDescription(v string) *GetDialogLogResponseBodyDataIntentionList {
	s.Description = &v
	return s
}

func (s *GetDialogLogResponseBodyDataIntentionList) SetIntentionName(v string) *GetDialogLogResponseBodyDataIntentionList {
	s.IntentionName = &v
	return s
}

func (s *GetDialogLogResponseBodyDataIntentionList) SetIntentionScript(v string) *GetDialogLogResponseBodyDataIntentionList {
	s.IntentionScript = &v
	return s
}

func (s *GetDialogLogResponseBodyDataIntentionList) Validate() error {
	return dara.Validate(s)
}
