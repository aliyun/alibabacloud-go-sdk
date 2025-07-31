// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddKeywordLibResponseBody
	GetCode() *int32
	SetData(v *AddKeywordLibResponseBodyData) *AddKeywordLibResponseBody
	GetData() *AddKeywordLibResponseBodyData
	SetMsg(v string) *AddKeywordLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *AddKeywordLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddKeywordLibResponseBody
	GetSuccess() *bool
}

type AddKeywordLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddKeywordLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s AddKeywordLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddKeywordLibResponseBody) GetData() *AddKeywordLibResponseBodyData {
	return s.Data
}

func (s *AddKeywordLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *AddKeywordLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKeywordLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddKeywordLibResponseBody) SetCode(v int32) *AddKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *AddKeywordLibResponseBody) SetData(v *AddKeywordLibResponseBodyData) *AddKeywordLibResponseBody {
	s.Data = v
	return s
}

func (s *AddKeywordLibResponseBody) SetMsg(v string) *AddKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *AddKeywordLibResponseBody) SetRequestId(v string) *AddKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKeywordLibResponseBody) SetSuccess(v bool) *AddKeywordLibResponseBody {
	s.Success = &v
	return s
}

func (s *AddKeywordLibResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddKeywordLibResponseBodyData struct {
	KeywordsResult *AddKeywordLibResponseBodyDataKeywordsResult `json:"KeywordsResult,omitempty" xml:"KeywordsResult,omitempty" type:"Struct"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddKeywordLibResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponseBodyData) GetKeywordsResult() *AddKeywordLibResponseBodyDataKeywordsResult {
	return s.KeywordsResult
}

func (s *AddKeywordLibResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordLibResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AddKeywordLibResponseBodyData) SetKeywordsResult(v *AddKeywordLibResponseBodyDataKeywordsResult) *AddKeywordLibResponseBodyData {
	s.KeywordsResult = v
	return s
}

func (s *AddKeywordLibResponseBodyData) SetLibId(v string) *AddKeywordLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *AddKeywordLibResponseBodyData) SetTaskId(v string) *AddKeywordLibResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AddKeywordLibResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type AddKeywordLibResponseBodyDataKeywordsResult struct {
	// example:
	//
	// xxx
	I18nKey               *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthKeywords []*string `json:"IllegalLengthKeywords,omitempty" xml:"IllegalLengthKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 133
	InvalidCount    *int32    `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	InvalidKeywords []*string `json:"InvalidKeywords,omitempty" xml:"InvalidKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 118
	RepeatCount    *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatKeywords []*string `json:"RepeatKeywords,omitempty" xml:"RepeatKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 278
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// xxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 529
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddKeywordLibResponseBodyDataKeywordsResult) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordLibResponseBodyDataKeywordsResult) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetI18nKey() *string {
	return s.I18nKey
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetIllegalLengthKeywords() []*string {
	return s.IllegalLengthKeywords
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetInvalidKeywords() []*string {
	return s.InvalidKeywords
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetLibId() *string {
	return s.LibId
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetRepeatKeywords() []*string {
	return s.RepeatKeywords
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetTips() *string {
	return s.Tips
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetI18nKey(v string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.I18nKey = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetIllegalLengthKeywords(v []*string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.IllegalLengthKeywords = v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetInvalidCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.InvalidCount = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetInvalidKeywords(v []*string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.InvalidKeywords = v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetLibId(v string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.LibId = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetRepeatCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.RepeatCount = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetRepeatKeywords(v []*string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.RepeatKeywords = v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetSuccessCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.SuccessCount = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetTips(v string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.Tips = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetTotalCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.TotalCount = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) Validate() error {
	return dara.Validate(s)
}
