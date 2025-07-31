// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordsToLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddKeywordsToLibResponseBodyData) *AddKeywordsToLibResponseBody
	GetData() *AddKeywordsToLibResponseBodyData
	SetRequestId(v string) *AddKeywordsToLibResponseBody
	GetRequestId() *string
}

type AddKeywordsToLibResponseBody struct {
	Data *AddKeywordsToLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddKeywordsToLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsToLibResponseBody) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponseBody) GetData() *AddKeywordsToLibResponseBodyData {
	return s.Data
}

func (s *AddKeywordsToLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKeywordsToLibResponseBody) SetData(v *AddKeywordsToLibResponseBodyData) *AddKeywordsToLibResponseBody {
	s.Data = v
	return s
}

func (s *AddKeywordsToLibResponseBody) SetRequestId(v string) *AddKeywordsToLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKeywordsToLibResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddKeywordsToLibResponseBodyData struct {
	KeywordsResult *AddKeywordsToLibResponseBodyDataKeywordsResult `json:"KeywordsResult,omitempty" xml:"KeywordsResult,omitempty" type:"Struct"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddKeywordsToLibResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsToLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponseBodyData) GetKeywordsResult() *AddKeywordsToLibResponseBodyDataKeywordsResult {
	return s.KeywordsResult
}

func (s *AddKeywordsToLibResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordsToLibResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AddKeywordsToLibResponseBodyData) SetKeywordsResult(v *AddKeywordsToLibResponseBodyDataKeywordsResult) *AddKeywordsToLibResponseBodyData {
	s.KeywordsResult = v
	return s
}

func (s *AddKeywordsToLibResponseBodyData) SetLibId(v string) *AddKeywordsToLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyData) SetTaskId(v string) *AddKeywordsToLibResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type AddKeywordsToLibResponseBodyDataKeywordsResult struct {
	// example:
	//
	// xxx
	I18nKey               *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthKeywords []*string `json:"IllegalLengthKeywords,omitempty" xml:"IllegalLengthKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvalidCount    *int32    `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	InvalidKeywords []*string `json:"InvalidKeywords,omitempty" xml:"InvalidKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	RepeatCount    *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatKeywords []*string `json:"RepeatKeywords,omitempty" xml:"RepeatKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddKeywordsToLibResponseBodyDataKeywordsResult) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsToLibResponseBodyDataKeywordsResult) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetI18nKey() *string {
	return s.I18nKey
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetIllegalLengthKeywords() []*string {
	return s.IllegalLengthKeywords
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetInvalidKeywords() []*string {
	return s.InvalidKeywords
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetProgress() *int32 {
	return s.Progress
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetRepeatKeywords() []*string {
	return s.RepeatKeywords
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetI18nKey(v string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.I18nKey = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetIllegalLengthKeywords(v []*string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.IllegalLengthKeywords = v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetInvalidCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.InvalidCount = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetInvalidKeywords(v []*string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.InvalidKeywords = v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetLibId(v string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.LibId = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetProgress(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.Progress = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetRepeatCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.RepeatCount = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetRepeatKeywords(v []*string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.RepeatKeywords = v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetSuccessCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.SuccessCount = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetTotalCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.TotalCount = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) Validate() error {
	return dara.Validate(s)
}
