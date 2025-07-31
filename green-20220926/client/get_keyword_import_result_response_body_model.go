// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeywordImportResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetKeywordImportResultResponseBody
	GetCode() *int32
	SetData(v *GetKeywordImportResultResponseBodyData) *GetKeywordImportResultResponseBody
	GetData() *GetKeywordImportResultResponseBodyData
	SetMsg(v string) *GetKeywordImportResultResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetKeywordImportResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKeywordImportResultResponseBody
	GetSuccess() *bool
}

type GetKeywordImportResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetKeywordImportResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetKeywordImportResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordImportResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetKeywordImportResultResponseBody) GetData() *GetKeywordImportResultResponseBodyData {
	return s.Data
}

func (s *GetKeywordImportResultResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetKeywordImportResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKeywordImportResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKeywordImportResultResponseBody) SetCode(v int32) *GetKeywordImportResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetData(v *GetKeywordImportResultResponseBodyData) *GetKeywordImportResultResponseBody {
	s.Data = v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetMsg(v string) *GetKeywordImportResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetRequestId(v string) *GetKeywordImportResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetSuccess(v bool) *GetKeywordImportResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetKeywordImportResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetKeywordImportResultResponseBodyData struct {
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
	// customxx_xxx
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
	// xxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetKeywordImportResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordImportResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultResponseBodyData) GetI18nKey() *string {
	return s.I18nKey
}

func (s *GetKeywordImportResultResponseBodyData) GetIllegalLengthKeywords() []*string {
	return s.IllegalLengthKeywords
}

func (s *GetKeywordImportResultResponseBodyData) GetInvalidCount() *int32 {
	return s.InvalidCount
}

func (s *GetKeywordImportResultResponseBodyData) GetInvalidKeywords() []*string {
	return s.InvalidKeywords
}

func (s *GetKeywordImportResultResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *GetKeywordImportResultResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *GetKeywordImportResultResponseBodyData) GetRepeatCount() *int32 {
	return s.RepeatCount
}

func (s *GetKeywordImportResultResponseBodyData) GetRepeatKeywords() []*string {
	return s.RepeatKeywords
}

func (s *GetKeywordImportResultResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *GetKeywordImportResultResponseBodyData) GetTips() *string {
	return s.Tips
}

func (s *GetKeywordImportResultResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetKeywordImportResultResponseBodyData) SetI18nKey(v string) *GetKeywordImportResultResponseBodyData {
	s.I18nKey = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetIllegalLengthKeywords(v []*string) *GetKeywordImportResultResponseBodyData {
	s.IllegalLengthKeywords = v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetInvalidCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.InvalidCount = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetInvalidKeywords(v []*string) *GetKeywordImportResultResponseBodyData {
	s.InvalidKeywords = v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetLibId(v string) *GetKeywordImportResultResponseBodyData {
	s.LibId = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetProgress(v int32) *GetKeywordImportResultResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetRepeatCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.RepeatCount = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetRepeatKeywords(v []*string) *GetKeywordImportResultResponseBodyData {
	s.RepeatKeywords = v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetSuccessCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetTips(v string) *GetKeywordImportResultResponseBodyData {
	s.Tips = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetTotalCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
