// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnswerImportProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetI18nKey(v string) *GetAnswerImportProgressResponseBody
	GetI18nKey() *string
	SetIllegalLengthSamples(v []*string) *GetAnswerImportProgressResponseBody
	GetIllegalLengthSamples() []*string
	SetInvalidCount(v int32) *GetAnswerImportProgressResponseBody
	GetInvalidCount() *int32
	SetLibId(v string) *GetAnswerImportProgressResponseBody
	GetLibId() *string
	SetProgress(v int32) *GetAnswerImportProgressResponseBody
	GetProgress() *int32
	SetRepeatCount(v int32) *GetAnswerImportProgressResponseBody
	GetRepeatCount() *int32
	SetRepeatSamples(v []*string) *GetAnswerImportProgressResponseBody
	GetRepeatSamples() []*string
	SetRequestId(v string) *GetAnswerImportProgressResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *GetAnswerImportProgressResponseBody
	GetSuccessCount() *int32
	SetTaskId(v string) *GetAnswerImportProgressResponseBody
	GetTaskId() *string
	SetTips(v string) *GetAnswerImportProgressResponseBody
	GetTips() *string
	SetTotalCount(v int32) *GetAnswerImportProgressResponseBody
	GetTotalCount() *int32
}

type GetAnswerImportProgressResponseBody struct {
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
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// xxxxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAnswerImportProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAnswerImportProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetAnswerImportProgressResponseBody) GetI18nKey() *string {
	return s.I18nKey
}

func (s *GetAnswerImportProgressResponseBody) GetIllegalLengthSamples() []*string {
	return s.IllegalLengthSamples
}

func (s *GetAnswerImportProgressResponseBody) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *GetAnswerImportProgressResponseBody) GetLibId() *string {
	return s.LibId
}

func (s *GetAnswerImportProgressResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *GetAnswerImportProgressResponseBody) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *GetAnswerImportProgressResponseBody) GetRepeatSamples() []*string {
	return s.RepeatSamples
}

func (s *GetAnswerImportProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAnswerImportProgressResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *GetAnswerImportProgressResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAnswerImportProgressResponseBody) GetTips() *string {
	return s.Tips
}

func (s *GetAnswerImportProgressResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAnswerImportProgressResponseBody) SetI18nKey(v string) *GetAnswerImportProgressResponseBody {
	s.I18nKey = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetIllegalLengthSamples(v []*string) *GetAnswerImportProgressResponseBody {
	s.IllegalLengthSamples = v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetInvalidCount(v int32) *GetAnswerImportProgressResponseBody {
	s.InvalidCount = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetLibId(v string) *GetAnswerImportProgressResponseBody {
	s.LibId = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetProgress(v int32) *GetAnswerImportProgressResponseBody {
	s.Progress = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetRepeatCount(v int32) *GetAnswerImportProgressResponseBody {
	s.RepeatCount = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetRepeatSamples(v []*string) *GetAnswerImportProgressResponseBody {
	s.RepeatSamples = v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetRequestId(v string) *GetAnswerImportProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetSuccessCount(v int32) *GetAnswerImportProgressResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetTaskId(v string) *GetAnswerImportProgressResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetTips(v string) *GetAnswerImportProgressResponseBody {
	s.Tips = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetTotalCount(v int32) *GetAnswerImportProgressResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) Validate() error {
	return dara.Validate(s)
}
