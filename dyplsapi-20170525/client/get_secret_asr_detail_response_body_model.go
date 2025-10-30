// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretAsrDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecretAsrDetailResponseBody
	GetCode() *string
	SetData(v *GetSecretAsrDetailResponseBodyData) *GetSecretAsrDetailResponseBody
	GetData() *GetSecretAsrDetailResponseBodyData
	SetMessage(v string) *GetSecretAsrDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecretAsrDetailResponseBody
	GetRequestId() *string
}

type GetSecretAsrDetailResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ASR details.
	Data *GetSecretAsrDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretAsrDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretAsrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecretAsrDetailResponseBody) GetData() *GetSecretAsrDetailResponseBodyData {
	return s.Data
}

func (s *GetSecretAsrDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecretAsrDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretAsrDetailResponseBody) SetCode(v string) *GetSecretAsrDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetData(v *GetSecretAsrDetailResponseBodyData) *GetSecretAsrDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetMessage(v string) *GetSecretAsrDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecretAsrDetailResponseBody) SetRequestId(v string) *GetSecretAsrDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecretAsrDetailResponseBodyData struct {
	// The total duration of the audio file that was recognized. Unit: milliseconds.
	//
	// example:
	//
	// 10944
	BizDuration *int64 `json:"BizDuration,omitempty" xml:"BizDuration,omitempty"`
	// The ID of the business process.
	//
	// example:
	//
	// 435ee78c7a019650@!FC100000074672458@!2020061522****
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The business keyword.
	//
	// example:
	//
	// JCGTncltuNao****
	BusinessKey *string `json:"BusinessKey,omitempty" xml:"BusinessKey,omitempty"`
	// The status code. The status code 21050000 indicates that the request was successful.
	//
	// example:
	//
	// 21050000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description.
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8d2329d407a83447a83be441681f4872ac74nE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ASR result.
	Sentences []*GetSecretAsrDetailResponseBodyDataSentences `json:"Sentences,omitempty" xml:"Sentences,omitempty" type:"Repeated"`
	// The type.
	//
	// example:
	//
	// asrResult
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSecretAsrDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSecretAsrDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBodyData) GetBizDuration() *int64 {
	return s.BizDuration
}

func (s *GetSecretAsrDetailResponseBodyData) GetBusinessId() *string {
	return s.BusinessId
}

func (s *GetSecretAsrDetailResponseBodyData) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *GetSecretAsrDetailResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetSecretAsrDetailResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *GetSecretAsrDetailResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretAsrDetailResponseBodyData) GetSentences() []*GetSecretAsrDetailResponseBodyDataSentences {
	return s.Sentences
}

func (s *GetSecretAsrDetailResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetSecretAsrDetailResponseBodyData) SetBizDuration(v int64) *GetSecretAsrDetailResponseBodyData {
	s.BizDuration = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetBusinessId(v string) *GetSecretAsrDetailResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetBusinessKey(v string) *GetSecretAsrDetailResponseBodyData {
	s.BusinessKey = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetCode(v string) *GetSecretAsrDetailResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetMsg(v string) *GetSecretAsrDetailResponseBodyData {
	s.Msg = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetRequestId(v string) *GetSecretAsrDetailResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetSentences(v []*GetSecretAsrDetailResponseBodyDataSentences) *GetSecretAsrDetailResponseBodyData {
	s.Sentences = v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) SetType(v string) *GetSecretAsrDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyData) Validate() error {
	if s.Sentences != nil {
		for _, item := range s.Sentences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSecretAsrDetailResponseBodyDataSentences struct {
	// The start time offset of the sentence. Unit: milliseconds.
	//
	// example:
	//
	// 1020
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The ID of the audio track to which the sentence belongs.
	//
	// example:
	//
	// 0
	ChannelId *int32 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The emotion value. Value range: 1 to 10. The higher the value, the stronger the emotion.
	//
	// example:
	//
	// 5.7
	EmotionValue *string `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// The end time offset of the sentence. Unit: milliseconds.
	//
	// example:
	//
	// 1770
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The silence duration between the current sentence and the previous sentence. Unit: seconds.
	//
	// example:
	//
	// 0
	SilenceDuration *int64 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// The average speech rate of the sentence. Unit: number of words per minute.
	//
	// example:
	//
	// 80
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The recognition result of the sentence.
	//
	// example:
	//
	// Hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetSecretAsrDetailResponseBodyDataSentences) String() string {
	return dara.Prettify(s)
}

func (s GetSecretAsrDetailResponseBodyDataSentences) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) GetChannelId() *int32 {
	return s.ChannelId
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) GetEmotionValue() *string {
	return s.EmotionValue
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) GetSilenceDuration() *int64 {
	return s.SilenceDuration
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) GetText() *string {
	return s.Text
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetBeginTime(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.BeginTime = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetChannelId(v int32) *GetSecretAsrDetailResponseBodyDataSentences {
	s.ChannelId = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetEmotionValue(v string) *GetSecretAsrDetailResponseBodyDataSentences {
	s.EmotionValue = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetEndTime(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.EndTime = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetSilenceDuration(v int64) *GetSecretAsrDetailResponseBodyDataSentences {
	s.SilenceDuration = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetSpeechRate(v int32) *GetSecretAsrDetailResponseBodyDataSentences {
	s.SpeechRate = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) SetText(v string) *GetSecretAsrDetailResponseBodyDataSentences {
	s.Text = &v
	return s
}

func (s *GetSecretAsrDetailResponseBodyDataSentences) Validate() error {
	return dara.Validate(s)
}
