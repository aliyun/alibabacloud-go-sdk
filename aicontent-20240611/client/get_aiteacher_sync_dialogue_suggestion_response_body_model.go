// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITeacherSyncDialogueSuggestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAITeacherSyncDialogueSuggestionResponseBodyData) *GetAITeacherSyncDialogueSuggestionResponseBody
	GetData() *GetAITeacherSyncDialogueSuggestionResponseBodyData
	SetErrCode(v string) *GetAITeacherSyncDialogueSuggestionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetAITeacherSyncDialogueSuggestionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetAITeacherSyncDialogueSuggestionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAITeacherSyncDialogueSuggestionResponseBody
	GetSuccess() *bool
}

type GetAITeacherSyncDialogueSuggestionResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {"chineseResult":"谢谢莉莉.你喜欢吃肉吗，莉莉？","englishResult":"Thanks, Lily. Do you like meat, Lily?"}
	Data *GetAITeacherSyncDialogueSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault message encoding.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) GetData() *GetAITeacherSyncDialogueSuggestionResponseBodyData {
	return s.Data
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetData(v *GetAITeacherSyncDialogueSuggestionResponseBodyData) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.Data = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetErrCode(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetErrMessage(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetHttpStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetRequestId(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetSuccess(v bool) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.Success = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAITeacherSyncDialogueSuggestionResponseBodyData struct {
	// The Chinese prompt from the assistant.
	//
	// example:
	//
	// 谢谢莉莉.你喜欢吃肉吗，莉莉？
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// The English prompt from the assistant.
	//
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) GetChineseResult() *string {
	return s.ChineseResult
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) GetEnglishResult() *string {
	return s.EnglishResult
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) SetChineseResult(v string) *GetAITeacherSyncDialogueSuggestionResponseBodyData {
	s.ChineseResult = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) SetEnglishResult(v string) *GetAITeacherSyncDialogueSuggestionResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
