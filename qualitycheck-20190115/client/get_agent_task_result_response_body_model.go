// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentTaskResultResponseBody
	GetCode() *string
	SetData(v *GetAgentTaskResultResponseBodyData) *GetAgentTaskResultResponseBody
	GetData() *GetAgentTaskResultResponseBodyData
	SetMessage(v string) *GetAgentTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAgentTaskResultResponseBody
	GetSuccess() *string
}

type GetAgentTaskResultResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *GetAgentTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. You can use this field to determine whether the request succeeded:
	//
	// - **true**: The request was successful.
	//
	// - **false/null**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentTaskResultResponseBody) GetData() *GetAgentTaskResultResponseBodyData {
	return s.Data
}

func (s *GetAgentTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentTaskResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAgentTaskResultResponseBody) SetCode(v string) *GetAgentTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentTaskResultResponseBody) SetData(v *GetAgentTaskResultResponseBodyData) *GetAgentTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentTaskResultResponseBody) SetMessage(v string) *GetAgentTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentTaskResultResponseBody) SetRequestId(v string) *GetAgentTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentTaskResultResponseBody) SetSuccess(v string) *GetAgentTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentTaskResultResponseBodyData struct {
	Dialogues    []*GetAgentTaskResultResponseBodyDataDialogues `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Repeated"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The number of input tokens.
	//
	// example:
	//
	// 100
	InputTokens *string `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The request ID returned by the large language model service.
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	LlmRequestId *string `json:"LlmRequestId,omitempty" xml:"LlmRequestId,omitempty"`
	// The number of output tokens.
	//
	// example:
	//
	// 200
	OutputTokens *string `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The result of the computation task.
	Response *GetAgentTaskResultResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// - 1: pending
	//
	// - 2: running
	//
	// - 3: succeeded
	//
	// - 4: failed
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// A6BEC8D-9A5B-4BE5-8432-4F635E***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of tokens.
	//
	// example:
	//
	// 300
	TotalTokens *string `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// The number of times the plus model is used.
	//
	// example:
	//
	// 1
	TyxmPlusCount *string `json:"TyxmPlusCount,omitempty" xml:"TyxmPlusCount,omitempty"`
	// The number of times the turbo model is used.
	//
	// example:
	//
	// 1
	TyxmTurboCount *string `json:"TyxmTurboCount,omitempty" xml:"TyxmTurboCount,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 6fa76916-3ce6-45d8-ac64-01b7f31***
	Vid *string `json:"Vid,omitempty" xml:"Vid,omitempty"`
}

func (s GetAgentTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyData) GetDialogues() []*GetAgentTaskResultResponseBodyDataDialogues {
	return s.Dialogues
}

func (s *GetAgentTaskResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgentTaskResultResponseBodyData) GetInputTokens() *string {
	return s.InputTokens
}

func (s *GetAgentTaskResultResponseBodyData) GetLlmRequestId() *string {
	return s.LlmRequestId
}

func (s *GetAgentTaskResultResponseBodyData) GetOutputTokens() *string {
	return s.OutputTokens
}

func (s *GetAgentTaskResultResponseBodyData) GetResponse() *GetAgentTaskResultResponseBodyDataResponse {
	return s.Response
}

func (s *GetAgentTaskResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAgentTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAgentTaskResultResponseBodyData) GetTotalTokens() *string {
	return s.TotalTokens
}

func (s *GetAgentTaskResultResponseBodyData) GetTyxmPlusCount() *string {
	return s.TyxmPlusCount
}

func (s *GetAgentTaskResultResponseBodyData) GetTyxmTurboCount() *string {
	return s.TyxmTurboCount
}

func (s *GetAgentTaskResultResponseBodyData) GetVid() *string {
	return s.Vid
}

func (s *GetAgentTaskResultResponseBodyData) SetDialogues(v []*GetAgentTaskResultResponseBodyDataDialogues) *GetAgentTaskResultResponseBodyData {
	s.Dialogues = v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetErrorMessage(v string) *GetAgentTaskResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetInputTokens(v string) *GetAgentTaskResultResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetLlmRequestId(v string) *GetAgentTaskResultResponseBodyData {
	s.LlmRequestId = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetOutputTokens(v string) *GetAgentTaskResultResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetResponse(v *GetAgentTaskResultResponseBodyDataResponse) *GetAgentTaskResultResponseBodyData {
	s.Response = v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetStatus(v string) *GetAgentTaskResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetTaskId(v string) *GetAgentTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetTotalTokens(v string) *GetAgentTaskResultResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetTyxmPlusCount(v string) *GetAgentTaskResultResponseBodyData {
	s.TyxmPlusCount = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetTyxmTurboCount(v string) *GetAgentTaskResultResponseBodyData {
	s.TyxmTurboCount = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) SetVid(v string) *GetAgentTaskResultResponseBodyData {
	s.Vid = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyData) Validate() error {
	if s.Dialogues != nil {
		for _, item := range s.Dialogues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Response != nil {
		if err := s.Response.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentTaskResultResponseBodyDataDialogues struct {
	Begin        *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	EmotionValue *int32  `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	End          *int64  `json:"End,omitempty" xml:"End,omitempty"`
	HourMinSec   *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SpeechRate   *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words        *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetAgentTaskResultResponseBodyDataDialogues) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataDialogues) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) GetBegin() *int64 {
	return s.Begin
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) GetEnd() *int64 {
	return s.End
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) GetHourMinSec() *string {
	return s.HourMinSec
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) GetRole() *string {
	return s.Role
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) GetWords() *string {
	return s.Words
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) SetBegin(v int64) *GetAgentTaskResultResponseBodyDataDialogues {
	s.Begin = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) SetEmotionValue(v int32) *GetAgentTaskResultResponseBodyDataDialogues {
	s.EmotionValue = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) SetEnd(v int64) *GetAgentTaskResultResponseBodyDataDialogues {
	s.End = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) SetHourMinSec(v string) *GetAgentTaskResultResponseBodyDataDialogues {
	s.HourMinSec = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) SetRole(v string) *GetAgentTaskResultResponseBodyDataDialogues {
	s.Role = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) SetSpeechRate(v int32) *GetAgentTaskResultResponseBodyDataDialogues {
	s.SpeechRate = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) SetWords(v string) *GetAgentTaskResultResponseBodyDataDialogues {
	s.Words = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataDialogues) Validate() error {
	return dara.Validate(s)
}

type GetAgentTaskResultResponseBodyDataResponse struct {
	// The result of the custom prompt.
	CustomerPromptResponse *GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse `json:"CustomerPromptResponse,omitempty" xml:"CustomerPromptResponse,omitempty" type:"Struct"`
	// The field extraction result.
	FieldResponse *GetAgentTaskResultResponseBodyDataResponseFieldResponse `json:"FieldResponse,omitempty" xml:"FieldResponse,omitempty" type:"Struct"`
	// The service quality inspection result.
	ServiceInspectionResponse *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse `json:"ServiceInspectionResponse,omitempty" xml:"ServiceInspectionResponse,omitempty" type:"Struct"`
	// The tag categorization result.
	TagCategoryResponse *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse `json:"TagCategoryResponse,omitempty" xml:"TagCategoryResponse,omitempty" type:"Struct"`
}

func (s GetAgentTaskResultResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponse) GetCustomerPromptResponse() *GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse {
	return s.CustomerPromptResponse
}

func (s *GetAgentTaskResultResponseBodyDataResponse) GetFieldResponse() *GetAgentTaskResultResponseBodyDataResponseFieldResponse {
	return s.FieldResponse
}

func (s *GetAgentTaskResultResponseBodyDataResponse) GetServiceInspectionResponse() *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse {
	return s.ServiceInspectionResponse
}

func (s *GetAgentTaskResultResponseBodyDataResponse) GetTagCategoryResponse() *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse {
	return s.TagCategoryResponse
}

func (s *GetAgentTaskResultResponseBodyDataResponse) SetCustomerPromptResponse(v *GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse) *GetAgentTaskResultResponseBodyDataResponse {
	s.CustomerPromptResponse = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponse) SetFieldResponse(v *GetAgentTaskResultResponseBodyDataResponseFieldResponse) *GetAgentTaskResultResponseBodyDataResponse {
	s.FieldResponse = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponse) SetServiceInspectionResponse(v *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse) *GetAgentTaskResultResponseBodyDataResponse {
	s.ServiceInspectionResponse = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponse) SetTagCategoryResponse(v *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse) *GetAgentTaskResultResponseBodyDataResponse {
	s.TagCategoryResponse = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponse) Validate() error {
	if s.CustomerPromptResponse != nil {
		if err := s.CustomerPromptResponse.Validate(); err != nil {
			return err
		}
	}
	if s.FieldResponse != nil {
		if err := s.FieldResponse.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceInspectionResponse != nil {
		if err := s.ServiceInspectionResponse.Validate(); err != nil {
			return err
		}
	}
	if s.TagCategoryResponse != nil {
		if err := s.TagCategoryResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse struct {
	// The result returned by the large language model.
	//
	// example:
	//
	// 175/xl the fabric feels very comfortable, looks slim when worn, great clothes super good-looking, quality and feel are top-notch, very satisfied with this purchase.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse) GetText() *string {
	return s.Text
}

func (s *GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse) SetText(v string) *GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse {
	s.Text = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseCustomerPromptResponse) Validate() error {
	return dara.Validate(s)
}

type GetAgentTaskResultResponseBodyDataResponseFieldResponse struct {
	// The list of fields.
	FieldVoList []*GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList `json:"FieldVoList,omitempty" xml:"FieldVoList,omitempty" type:"Repeated"`
}

func (s GetAgentTaskResultResponseBodyDataResponseFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponseFieldResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponse) GetFieldVoList() []*GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList {
	return s.FieldVoList
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponse) SetFieldVoList(v []*GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) *GetAgentTaskResultResponseBodyDataResponseFieldResponse {
	s.FieldVoList = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponse) Validate() error {
	if s.FieldVoList != nil {
		for _, item := range s.FieldVoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList struct {
	// The field name.
	//
	// example:
	//
	// phone
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sentences referenced in the reasoning.
	OriginalUtterances []*int32 `json:"OriginalUtterances,omitempty" xml:"OriginalUtterances,omitempty" type:"Repeated"`
	// The reasoning for the judgment.
	//
	// example:
	//
	// Determined by the first sentence of the agent.
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The field value.
	//
	// example:
	//
	// 1234561
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) GetName() *string {
	return s.Name
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) GetOriginalUtterances() []*int32 {
	return s.OriginalUtterances
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) GetRemarks() *string {
	return s.Remarks
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) GetValue() *string {
	return s.Value
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) SetName(v string) *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList {
	s.Name = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) SetOriginalUtterances(v []*int32) *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList {
	s.OriginalUtterances = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) SetRemarks(v string) *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList {
	s.Remarks = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) SetValue(v string) *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList {
	s.Value = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseFieldResponseFieldVoList) Validate() error {
	return dara.Validate(s)
}

type GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse struct {
	// The list of inspection items.
	ServiceInspectionVoList []*GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList `json:"ServiceInspectionVoList,omitempty" xml:"ServiceInspectionVoList,omitempty" type:"Repeated"`
}

func (s GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse) GetServiceInspectionVoList() []*GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList {
	return s.ServiceInspectionVoList
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse) SetServiceInspectionVoList(v []*GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse {
	s.ServiceInspectionVoList = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponse) Validate() error {
	if s.ServiceInspectionVoList != nil {
		for _, item := range s.ServiceInspectionVoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList struct {
	// The inspection dimension.
	//
	// example:
	//
	// Service attitude.
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// Indicates whether the tag is matched.
	//
	// example:
	//
	// true
	IsMatch *bool `json:"IsMatch,omitempty" xml:"IsMatch,omitempty"`
	// The sentences referenced in the reasoning.
	OriginalUtterances []*string `json:"OriginalUtterances,omitempty" xml:"OriginalUtterances,omitempty" type:"Repeated"`
	// The reasoning for the judgment.
	//
	// example:
	//
	// Determined by the first sentence of the agent.
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
}

func (s GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) GetDimension() *string {
	return s.Dimension
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) GetIsMatch() *bool {
	return s.IsMatch
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) GetOriginalUtterances() []*string {
	return s.OriginalUtterances
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) GetRemarks() *string {
	return s.Remarks
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) SetDimension(v string) *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList {
	s.Dimension = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) SetIsMatch(v bool) *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList {
	s.IsMatch = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) SetOriginalUtterances(v []*string) *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList {
	s.OriginalUtterances = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) SetRemarks(v string) *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList {
	s.Remarks = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseServiceInspectionResponseServiceInspectionVoList) Validate() error {
	return dara.Validate(s)
}

type GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse struct {
	// The list of tags.
	TagCategoryVoList []*GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList `json:"TagCategoryVoList,omitempty" xml:"TagCategoryVoList,omitempty" type:"Repeated"`
}

func (s GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse) GetTagCategoryVoList() []*GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList {
	return s.TagCategoryVoList
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse) SetTagCategoryVoList(v []*GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse {
	s.TagCategoryVoList = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponse) Validate() error {
	if s.TagCategoryVoList != nil {
		for _, item := range s.TagCategoryVoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList struct {
	// The tag dimension.
	//
	// example:
	//
	// Customer intent.
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// Indicates whether the tag is matched.
	//
	// example:
	//
	// true
	IsMatch *bool `json:"IsMatch,omitempty" xml:"IsMatch,omitempty"`
	// The sentences referenced in the reasoning.
	OriginalUtterances []*string `json:"OriginalUtterances,omitempty" xml:"OriginalUtterances,omitempty" type:"Repeated"`
	// The reasoning for the judgment.
	//
	// example:
	//
	// Determined by the first sentence of the agent.
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The list of matched labels.
	ResultLabels []*string `json:"ResultLabels,omitempty" xml:"ResultLabels,omitempty" type:"Repeated"`
}

func (s GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) GetDimension() *string {
	return s.Dimension
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) GetIsMatch() *bool {
	return s.IsMatch
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) GetOriginalUtterances() []*string {
	return s.OriginalUtterances
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) GetRemarks() *string {
	return s.Remarks
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) GetResultLabels() []*string {
	return s.ResultLabels
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) SetDimension(v string) *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList {
	s.Dimension = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) SetIsMatch(v bool) *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList {
	s.IsMatch = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) SetOriginalUtterances(v []*string) *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList {
	s.OriginalUtterances = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) SetRemarks(v string) *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList {
	s.Remarks = &v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) SetResultLabels(v []*string) *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList {
	s.ResultLabels = v
	return s
}

func (s *GetAgentTaskResultResponseBodyDataResponseTagCategoryResponseTagCategoryVoList) Validate() error {
	return dara.Validate(s)
}
