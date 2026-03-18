// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITeacherExpansionDialogueSuggestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAITeacherExpansionDialogueSuggestionResponseBodyData) *GetAITeacherExpansionDialogueSuggestionResponseBody
	GetData() *GetAITeacherExpansionDialogueSuggestionResponseBodyData
	SetErrCode(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAITeacherExpansionDialogueSuggestionResponseBody
	GetSuccess() *bool
}

type GetAITeacherExpansionDialogueSuggestionResponseBody struct {
	// example:
	//
	// []
	Data *GetAITeacherExpansionDialogueSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) GetData() *GetAITeacherExpansionDialogueSuggestionResponseBodyData {
	return s.Data
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetData(v *GetAITeacherExpansionDialogueSuggestionResponseBodyData) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.Data = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetErrCode(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetErrMessage(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetHttpStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetRequestId(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetSuccess(v bool) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.Success = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAITeacherExpansionDialogueSuggestionResponseBodyData struct {
	// example:
	//
	// 谢谢莉莉.你喜欢吃肉吗，莉莉？
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) GetChineseResult() *string {
	return s.ChineseResult
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) GetEnglishResult() *string {
	return s.EnglishResult
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) SetChineseResult(v string) *GetAITeacherExpansionDialogueSuggestionResponseBodyData {
	s.ChineseResult = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) SetEnglishResult(v string) *GetAITeacherExpansionDialogueSuggestionResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
