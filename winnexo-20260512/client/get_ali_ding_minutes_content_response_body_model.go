// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliDingMinutesContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAliDingMinutesContentResponseBody
	GetCode() *string
	SetMessage(v string) *GetAliDingMinutesContentResponseBody
	GetMessage() *string
	SetMinutesId(v string) *GetAliDingMinutesContentResponseBody
	GetMinutesId() *string
	SetRequestId(v string) *GetAliDingMinutesContentResponseBody
	GetRequestId() *string
	SetSummary(v string) *GetAliDingMinutesContentResponseBody
	GetSummary() *string
	SetTitle(v string) *GetAliDingMinutesContentResponseBody
	GetTitle() *string
	SetTodoContent(v string) *GetAliDingMinutesContentResponseBody
	GetTodoContent() *string
	SetTranscription(v []*GetAliDingMinutesContentResponseBodyTranscription) *GetAliDingMinutesContentResponseBody
	GetTranscription() []*GetAliDingMinutesContentResponseBodyTranscription
}

type GetAliDingMinutesContentResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The DingTalk minutes ID.
	//
	// example:
	//
	// 76327569643231383535353939365f3436383537393431335f32
	MinutesId *string `json:"minutesId,omitempty" xml:"minutesId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The intelligent meeting summary content.
	//
	// example:
	//
	// # Meeting Summary
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The new session title.
	//
	// example:
	//
	// Weekly Project Meeting
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The to-do item details.
	//
	// example:
	//
	// {"dingtalkTodoList":[]}
	TodoContent *string `json:"todoContent,omitempty" xml:"todoContent,omitempty"`
	// The speech-type execution parameters.
	Transcription []*GetAliDingMinutesContentResponseBodyTranscription `json:"transcription,omitempty" xml:"transcription,omitempty" type:"Repeated"`
}

func (s GetAliDingMinutesContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAliDingMinutesContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAliDingMinutesContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAliDingMinutesContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAliDingMinutesContentResponseBody) GetMinutesId() *string {
	return s.MinutesId
}

func (s *GetAliDingMinutesContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAliDingMinutesContentResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *GetAliDingMinutesContentResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetAliDingMinutesContentResponseBody) GetTodoContent() *string {
	return s.TodoContent
}

func (s *GetAliDingMinutesContentResponseBody) GetTranscription() []*GetAliDingMinutesContentResponseBodyTranscription {
	return s.Transcription
}

func (s *GetAliDingMinutesContentResponseBody) SetCode(v string) *GetAliDingMinutesContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) SetMessage(v string) *GetAliDingMinutesContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) SetMinutesId(v string) *GetAliDingMinutesContentResponseBody {
	s.MinutesId = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) SetRequestId(v string) *GetAliDingMinutesContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) SetSummary(v string) *GetAliDingMinutesContentResponseBody {
	s.Summary = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) SetTitle(v string) *GetAliDingMinutesContentResponseBody {
	s.Title = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) SetTodoContent(v string) *GetAliDingMinutesContentResponseBody {
	s.TodoContent = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) SetTranscription(v []*GetAliDingMinutesContentResponseBodyTranscription) *GetAliDingMinutesContentResponseBody {
	s.Transcription = v
	return s
}

func (s *GetAliDingMinutesContentResponseBody) Validate() error {
	if s.Transcription != nil {
		for _, item := range s.Transcription {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAliDingMinutesContentResponseBodyTranscription struct {
	// The returned content.
	//
	// example:
	//
	// Meeting started
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The speaker.
	//
	// example:
	//
	// John
	Speaker *string `json:"speaker,omitempty" xml:"speaker,omitempty"`
	// The avatar of the speaker. An empty string is returned if no avatar is available.
	//
	// example:
	//
	// https://robject.oss-cn-shanghai.aliyuncs.com/robject-daily1/
	SpeakerAvatar *string `json:"speakerAvatar,omitempty" xml:"speakerAvatar,omitempty"`
	// The end time of the segment.
	//
	// example:
	//
	// 1200
	TimeEnd *int64 `json:"timeEnd,omitempty" xml:"timeEnd,omitempty"`
	// The start time of the segment.
	//
	// example:
	//
	// 0
	TimeStart *int64 `json:"timeStart,omitempty" xml:"timeStart,omitempty"`
}

func (s GetAliDingMinutesContentResponseBodyTranscription) String() string {
	return dara.Prettify(s)
}

func (s GetAliDingMinutesContentResponseBodyTranscription) GoString() string {
	return s.String()
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) GetContent() *string {
	return s.Content
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) GetSpeaker() *string {
	return s.Speaker
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) GetSpeakerAvatar() *string {
	return s.SpeakerAvatar
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) GetTimeEnd() *int64 {
	return s.TimeEnd
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) GetTimeStart() *int64 {
	return s.TimeStart
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) SetContent(v string) *GetAliDingMinutesContentResponseBodyTranscription {
	s.Content = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) SetSpeaker(v string) *GetAliDingMinutesContentResponseBodyTranscription {
	s.Speaker = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) SetSpeakerAvatar(v string) *GetAliDingMinutesContentResponseBodyTranscription {
	s.SpeakerAvatar = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) SetTimeEnd(v int64) *GetAliDingMinutesContentResponseBodyTranscription {
	s.TimeEnd = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) SetTimeStart(v int64) *GetAliDingMinutesContentResponseBodyTranscription {
	s.TimeStart = &v
	return s
}

func (s *GetAliDingMinutesContentResponseBodyTranscription) Validate() error {
	return dara.Validate(s)
}
