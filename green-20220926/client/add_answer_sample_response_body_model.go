// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAnswerSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *AddAnswerSampleResponseBody
	GetLibId() *string
	SetRequestId(v string) *AddAnswerSampleResponseBody
	GetRequestId() *string
	SetResult(v *AddAnswerSampleResponseBodyResult) *AddAnswerSampleResponseBody
	GetResult() *AddAnswerSampleResponseBodyResult
	SetTaskId(v string) *AddAnswerSampleResponseBody
	GetTaskId() *string
}

type AddAnswerSampleResponseBody struct {
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddAnswerSampleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddAnswerSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAnswerSampleResponseBody) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleResponseBody) GetLibId() *string {
	return s.LibId
}

func (s *AddAnswerSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAnswerSampleResponseBody) GetResult() *AddAnswerSampleResponseBodyResult {
	return s.Result
}

func (s *AddAnswerSampleResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *AddAnswerSampleResponseBody) SetLibId(v string) *AddAnswerSampleResponseBody {
	s.LibId = &v
	return s
}

func (s *AddAnswerSampleResponseBody) SetRequestId(v string) *AddAnswerSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAnswerSampleResponseBody) SetResult(v *AddAnswerSampleResponseBodyResult) *AddAnswerSampleResponseBody {
	s.Result = v
	return s
}

func (s *AddAnswerSampleResponseBody) SetTaskId(v string) *AddAnswerSampleResponseBody {
	s.TaskId = &v
	return s
}

func (s *AddAnswerSampleResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddAnswerSampleResponseBodyResult struct {
	// example:
	//
	// xxx
	I18nKey              *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthSamples []*string `json:"IllegalLengthSamples,omitempty" xml:"IllegalLengthSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 118
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
	// 98
	RepeatCount   *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatSamples []*string `json:"RepeatSamples,omitempty" xml:"RepeatSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 318
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 534
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddAnswerSampleResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddAnswerSampleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleResponseBodyResult) GetI18nKey() *string {
	return s.I18nKey
}

func (s *AddAnswerSampleResponseBodyResult) GetIllegalLengthSamples() []*string {
	return s.IllegalLengthSamples
}

func (s *AddAnswerSampleResponseBodyResult) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *AddAnswerSampleResponseBodyResult) GetLibId() *string {
	return s.LibId
}

func (s *AddAnswerSampleResponseBodyResult) GetProgress() *int32 {
	return s.Progress
}

func (s *AddAnswerSampleResponseBodyResult) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *AddAnswerSampleResponseBodyResult) GetRepeatSamples() []*string {
	return s.RepeatSamples
}

func (s *AddAnswerSampleResponseBodyResult) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *AddAnswerSampleResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *AddAnswerSampleResponseBodyResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AddAnswerSampleResponseBodyResult) SetI18nKey(v string) *AddAnswerSampleResponseBodyResult {
	s.I18nKey = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetIllegalLengthSamples(v []*string) *AddAnswerSampleResponseBodyResult {
	s.IllegalLengthSamples = v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetInvalidCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.InvalidCount = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetLibId(v string) *AddAnswerSampleResponseBodyResult {
	s.LibId = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetProgress(v int32) *AddAnswerSampleResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetRepeatCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.RepeatCount = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetRepeatSamples(v []*string) *AddAnswerSampleResponseBodyResult {
	s.RepeatSamples = v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetSuccessCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.SuccessCount = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetTaskId(v string) *AddAnswerSampleResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetTotalCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
