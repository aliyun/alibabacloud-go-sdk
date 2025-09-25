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
	// example:
	//
	// null
	Cost *int64                        `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDialogLogResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 051EEB18-049A-17FF-A5E0-14A5B127C798
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
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
	return dara.Validate(s)
}

type GetDialogLogResponseBodyData struct {
	AnalysisProcess  *string                                         `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	ConversationList *string                                         `json:"conversationList,omitempty" xml:"conversationList,omitempty"`
	HitIntentionList []*GetDialogLogResponseBodyDataHitIntentionList `json:"hitIntentionList,omitempty" xml:"hitIntentionList,omitempty" type:"Repeated"`
	IntentionList    []*GetDialogLogResponseBodyDataIntentionList    `json:"intentionList,omitempty" xml:"intentionList,omitempty" type:"Repeated"`
	// example:
	//
	// 1382
	ModelCostTime *int64  `json:"modelCostTime,omitempty" xml:"modelCostTime,omitempty"`
	RecallList    *string `json:"recallList,omitempty" xml:"recallList,omitempty"`
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
	return dara.Validate(s)
}

type GetDialogLogResponseBodyDataHitIntentionList struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	IntentionName   *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
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
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	IntentionName   *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
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
