// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TextModerationResponseBody
	GetCode() *int32
	SetData(v *TextModerationResponseBodyData) *TextModerationResponseBody
	GetData() *TextModerationResponseBodyData
	SetMessage(v string) *TextModerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *TextModerationResponseBody
	GetRequestId() *string
}

type TextModerationResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation result data.
	Data *TextModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message for the request.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TextModerationResponseBody) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TextModerationResponseBody) GetData() *TextModerationResponseBodyData {
	return s.Data
}

func (s *TextModerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TextModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TextModerationResponseBody) SetCode(v int32) *TextModerationResponseBody {
	s.Code = &v
	return s
}

func (s *TextModerationResponseBody) SetData(v *TextModerationResponseBodyData) *TextModerationResponseBody {
	s.Data = v
	return s
}

func (s *TextModerationResponseBody) SetMessage(v string) *TextModerationResponseBody {
	s.Message = &v
	return s
}

func (s *TextModerationResponseBody) SetRequestId(v string) *TextModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *TextModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextModerationResponseBodyData struct {
	// The \\`accountId\\` specified in the request.
	//
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The data ID of the moderated object.
	//
	// > If you specify the dataId parameter in the request, its value is returned in this parameter.
	//
	// example:
	//
	// text1234
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The description of the label.
	//
	// example:
	//
	// 疑似广告内容
	Descriptions *string `json:"descriptions,omitempty" xml:"descriptions,omitempty"`
	// The \\`deviceId\\` specified in the request.
	//
	// example:
	//
	// xxxxxx
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// Auxiliary reference information for the text.
	Ext *TextModerationResponseBodyDataExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	// The moderation labels. If multiple labels are returned, they are separated by commas (,). Valid values: ad: advertisement profanity: profanity contraband: contraband sexual_content: sexual content violence: violent and terrorist content nonsense: meaningless content spam: spam negative_content: undesirable content cyberbullying: cyberbullying C_customized: A match in a custom library
	//
	// example:
	//
	// porn
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The ID of the manual review task.
	//
	// example:
	//
	// xxxxx-xxxxx
	ManualTaskId *string `json:"manualTaskId,omitempty" xml:"manualTaskId,omitempty"`
	// A JSON string that contains the reason for the moderation result. The string includes the following fields:
	//
	// 1. riskTips: The sub-labels.
	//
	// 2. riskWords: The detected risk words.
	//
	// 3. adNums: The detected ad-related numbers.
	//
	// 4. customizedWords: The detected custom words.
	//
	// 5. customizedLibs: The names of the custom libraries that contain a match.
	//
	// 6. riskLevel: The risk level, which is recommended by the system. Valid values:
	//
	// - high: high risk
	//
	// - medium: medium risk
	//
	// - low: low risk
	//
	// example:
	//
	// {\\"riskLevel\\":\\"high\\",\\"riskTips\\":\\"色情_低俗词\\",\\"riskWords\\":\\"色情服务\\"}
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s TextModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TextModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *TextModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *TextModerationResponseBodyData) GetDescriptions() *string {
	return s.Descriptions
}

func (s *TextModerationResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *TextModerationResponseBodyData) GetExt() *TextModerationResponseBodyDataExt {
	return s.Ext
}

func (s *TextModerationResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *TextModerationResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *TextModerationResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *TextModerationResponseBodyData) SetAccountId(v string) *TextModerationResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetDataId(v string) *TextModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetDescriptions(v string) *TextModerationResponseBodyData {
	s.Descriptions = &v
	return s
}

func (s *TextModerationResponseBodyData) SetDeviceId(v string) *TextModerationResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetExt(v *TextModerationResponseBodyDataExt) *TextModerationResponseBodyData {
	s.Ext = v
	return s
}

func (s *TextModerationResponseBodyData) SetLabels(v string) *TextModerationResponseBodyData {
	s.Labels = &v
	return s
}

func (s *TextModerationResponseBodyData) SetManualTaskId(v string) *TextModerationResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetReason(v string) *TextModerationResponseBodyData {
	s.Reason = &v
	return s
}

func (s *TextModerationResponseBodyData) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextModerationResponseBodyDataExt struct {
	// The output from the Large Language Model (LLM).
	LlmContent *TextModerationResponseBodyDataExtLlmContent `json:"llmContent,omitempty" xml:"llmContent,omitempty" type:"Struct"`
}

func (s TextModerationResponseBodyDataExt) String() string {
	return dara.Prettify(s)
}

func (s TextModerationResponseBodyDataExt) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBodyDataExt) GetLlmContent() *TextModerationResponseBodyDataExtLlmContent {
	return s.LlmContent
}

func (s *TextModerationResponseBodyDataExt) SetLlmContent(v *TextModerationResponseBodyDataExtLlmContent) *TextModerationResponseBodyDataExt {
	s.LlmContent = v
	return s
}

func (s *TextModerationResponseBodyDataExt) Validate() error {
	if s.LlmContent != nil {
		if err := s.LlmContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextModerationResponseBodyDataExtLlmContent struct {
	// The output content.
	//
	// example:
	//
	// 正常。文本中无风险内容。
	OutputText *string `json:"outputText,omitempty" xml:"outputText,omitempty"`
}

func (s TextModerationResponseBodyDataExtLlmContent) String() string {
	return dara.Prettify(s)
}

func (s TextModerationResponseBodyDataExtLlmContent) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBodyDataExtLlmContent) GetOutputText() *string {
	return s.OutputText
}

func (s *TextModerationResponseBodyDataExtLlmContent) SetOutputText(v string) *TextModerationResponseBodyDataExtLlmContent {
	s.OutputText = &v
	return s
}

func (s *TextModerationResponseBodyDataExtLlmContent) Validate() error {
	return dara.Validate(s)
}
