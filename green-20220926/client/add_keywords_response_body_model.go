// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddKeywordsResponseBody
	GetCode() *int32
	SetData(v *AddKeywordsResponseBodyData) *AddKeywordsResponseBody
	GetData() *AddKeywordsResponseBodyData
	SetMsg(v string) *AddKeywordsResponseBody
	GetMsg() *string
	SetRequestId(v string) *AddKeywordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddKeywordsResponseBody
	GetSuccess() *bool
}

type AddKeywordsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddKeywordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddKeywordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddKeywordsResponseBody) GetData() *AddKeywordsResponseBodyData {
	return s.Data
}

func (s *AddKeywordsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *AddKeywordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKeywordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddKeywordsResponseBody) SetCode(v int32) *AddKeywordsResponseBody {
	s.Code = &v
	return s
}

func (s *AddKeywordsResponseBody) SetData(v *AddKeywordsResponseBodyData) *AddKeywordsResponseBody {
	s.Data = v
	return s
}

func (s *AddKeywordsResponseBody) SetMsg(v string) *AddKeywordsResponseBody {
	s.Msg = &v
	return s
}

func (s *AddKeywordsResponseBody) SetRequestId(v string) *AddKeywordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKeywordsResponseBody) SetSuccess(v bool) *AddKeywordsResponseBody {
	s.Success = &v
	return s
}

func (s *AddKeywordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddKeywordsResponseBodyData struct {
	KeywordsResult *AddKeywordsResponseBodyDataKeywordsResult `json:"KeywordsResult,omitempty" xml:"KeywordsResult,omitempty" type:"Struct"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddKeywordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponseBodyData) GetKeywordsResult() *AddKeywordsResponseBodyDataKeywordsResult {
	return s.KeywordsResult
}

func (s *AddKeywordsResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordsResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AddKeywordsResponseBodyData) SetKeywordsResult(v *AddKeywordsResponseBodyDataKeywordsResult) *AddKeywordsResponseBodyData {
	s.KeywordsResult = v
	return s
}

func (s *AddKeywordsResponseBodyData) SetLibId(v string) *AddKeywordsResponseBodyData {
	s.LibId = &v
	return s
}

func (s *AddKeywordsResponseBodyData) SetTaskId(v string) *AddKeywordsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AddKeywordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type AddKeywordsResponseBodyDataKeywordsResult struct {
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
	// 6
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// xxxxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddKeywordsResponseBodyDataKeywordsResult) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordsResponseBodyDataKeywordsResult) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetI18nKey() *string {
	return s.I18nKey
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetIllegalLengthKeywords() []*string {
	return s.IllegalLengthKeywords
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetInvalidKeywords() []*string {
	return s.InvalidKeywords
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetProgress() *int32 {
	return s.Progress
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetRepeatKeywords() []*string {
	return s.RepeatKeywords
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetTips() *string {
	return s.Tips
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetI18nKey(v string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.I18nKey = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetIllegalLengthKeywords(v []*string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.IllegalLengthKeywords = v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetInvalidCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.InvalidCount = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetInvalidKeywords(v []*string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.InvalidKeywords = v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetLibId(v string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.LibId = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetProgress(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.Progress = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetRepeatCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.RepeatCount = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetRepeatKeywords(v []*string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.RepeatKeywords = v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetSuccessCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.SuccessCount = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetTips(v string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.Tips = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetTotalCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.TotalCount = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) Validate() error {
	return dara.Validate(s)
}
