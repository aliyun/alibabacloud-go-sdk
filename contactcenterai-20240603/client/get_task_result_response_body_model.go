// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTaskResultResponseBodyData) *GetTaskResultResponseBody
	GetData() *GetTaskResultResponseBodyData
	SetRequestId(v string) *GetTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetTaskResultResponseBody
	GetSuccess() *string
}

type GetTaskResultResponseBody struct {
	Data *GetTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-C552DED7E8BF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) GetData() *GetTaskResultResponseBodyData {
	return s.Data
}

func (s *GetTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetTaskResultResponseBody) SetData(v *GetTaskResultResponseBodyData) *GetTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetSuccess(v string) *GetTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTaskResultResponseBodyData struct {
	AsrResult        []*GetTaskResultResponseBodyDataAsrResult `json:"asrResult,omitempty" xml:"asrResult,omitempty" type:"Repeated"`
	Extra            *string                                   `json:"extra,omitempty" xml:"extra,omitempty"`
	TaskErrorMessage *string                                   `json:"taskErrorMessage,omitempty" xml:"taskErrorMessage,omitempty"`
	// example:
	//
	// 20240905-********-93E9-5D45-B4EF-045743A34071
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// FINISH
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyData) GetAsrResult() []*GetTaskResultResponseBodyDataAsrResult {
	return s.AsrResult
}

func (s *GetTaskResultResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetTaskResultResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *GetTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResultResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetTaskResultResponseBodyData) GetText() *string {
	return s.Text
}

func (s *GetTaskResultResponseBodyData) SetAsrResult(v []*GetTaskResultResponseBodyDataAsrResult) *GetTaskResultResponseBodyData {
	s.AsrResult = v
	return s
}

func (s *GetTaskResultResponseBodyData) SetExtra(v string) *GetTaskResultResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskErrorMessage(v string) *GetTaskResultResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskId(v string) *GetTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskStatus(v string) *GetTaskResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetText(v string) *GetTaskResultResponseBodyData {
	s.Text = &v
	return s
}

func (s *GetTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetTaskResultResponseBodyDataAsrResult struct {
	Begin        *int64  `json:"begin,omitempty" xml:"begin,omitempty"`
	EmotionValue *int32  `json:"emotionValue,omitempty" xml:"emotionValue,omitempty"`
	End          *int64  `json:"end,omitempty" xml:"end,omitempty"`
	Role         *string `json:"role,omitempty" xml:"role,omitempty"`
	RoleName     *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	SpeechRate   *int32  `json:"speechRate,omitempty" xml:"speechRate,omitempty"`
	Words        *string `json:"words,omitempty" xml:"words,omitempty"`
}

func (s GetTaskResultResponseBodyDataAsrResult) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyDataAsrResult) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetBegin() *int64 {
	return s.Begin
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetEnd() *int64 {
	return s.End
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetRole() *string {
	return s.Role
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetRoleName() *string {
	return s.RoleName
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetWords() *string {
	return s.Words
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetBegin(v int64) *GetTaskResultResponseBodyDataAsrResult {
	s.Begin = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetEmotionValue(v int32) *GetTaskResultResponseBodyDataAsrResult {
	s.EmotionValue = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetEnd(v int64) *GetTaskResultResponseBodyDataAsrResult {
	s.End = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetRole(v string) *GetTaskResultResponseBodyDataAsrResult {
	s.Role = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetRoleName(v string) *GetTaskResultResponseBodyDataAsrResult {
	s.RoleName = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetSpeechRate(v int32) *GetTaskResultResponseBodyDataAsrResult {
	s.SpeechRate = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetWords(v string) *GetTaskResultResponseBodyDataAsrResult {
	s.Words = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) Validate() error {
	return dara.Validate(s)
}
