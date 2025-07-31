// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnswerLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *CreateAnswerLibResponseBody
	GetLibId() *string
	SetRequestId(v string) *CreateAnswerLibResponseBody
	GetRequestId() *string
	SetResult(v *CreateAnswerLibResponseBodyResult) *CreateAnswerLibResponseBody
	GetResult() *CreateAnswerLibResponseBodyResult
	SetTaskId(v string) *CreateAnswerLibResponseBody
	GetTaskId() *string
}

type CreateAnswerLibResponseBody struct {
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAnswerLibResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateAnswerLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibResponseBody) GetLibId() *string {
	return s.LibId
}

func (s *CreateAnswerLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAnswerLibResponseBody) GetResult() *CreateAnswerLibResponseBodyResult {
	return s.Result
}

func (s *CreateAnswerLibResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAnswerLibResponseBody) SetLibId(v string) *CreateAnswerLibResponseBody {
	s.LibId = &v
	return s
}

func (s *CreateAnswerLibResponseBody) SetRequestId(v string) *CreateAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnswerLibResponseBody) SetResult(v *CreateAnswerLibResponseBodyResult) *CreateAnswerLibResponseBody {
	s.Result = v
	return s
}

func (s *CreateAnswerLibResponseBody) SetTaskId(v string) *CreateAnswerLibResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateAnswerLibResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAnswerLibResponseBodyResult struct {
	// example:
	//
	// xxx
	I18nKey              *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthSamples []*string `json:"IllegalLengthSamples,omitempty" xml:"IllegalLengthSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvalidCount *int32 `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	RepeatCount   *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatSamples []*string `json:"RepeatSamples,omitempty" xml:"RepeatSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// alAaaaxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateAnswerLibResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAnswerLibResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibResponseBodyResult) GetI18nKey() *string {
	return s.I18nKey
}

func (s *CreateAnswerLibResponseBodyResult) GetIllegalLengthSamples() []*string {
	return s.IllegalLengthSamples
}

func (s *CreateAnswerLibResponseBodyResult) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *CreateAnswerLibResponseBodyResult) GetLibId() *string {
	return s.LibId
}

func (s *CreateAnswerLibResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *CreateAnswerLibResponseBodyResult) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *CreateAnswerLibResponseBodyResult) GetRepeatSamples() []*string {
	return s.RepeatSamples
}

func (s *CreateAnswerLibResponseBodyResult) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateAnswerLibResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAnswerLibResponseBodyResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CreateAnswerLibResponseBodyResult) SetI18nKey(v string) *CreateAnswerLibResponseBodyResult {
	s.I18nKey = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetIllegalLengthSamples(v []*string) *CreateAnswerLibResponseBodyResult {
	s.IllegalLengthSamples = v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetInvalidCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.InvalidCount = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetLibId(v string) *CreateAnswerLibResponseBodyResult {
	s.LibId = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetProgress(v int32) *CreateAnswerLibResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetRepeatCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.RepeatCount = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetRepeatSamples(v []*string) *CreateAnswerLibResponseBodyResult {
	s.RepeatSamples = v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetSuccessCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.SuccessCount = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetTaskId(v string) *CreateAnswerLibResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetTotalCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
