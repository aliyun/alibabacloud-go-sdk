// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAsyncTranslateResponseBody
	GetCode() *int32
	SetData(v *GetAsyncTranslateResponseBodyData) *GetAsyncTranslateResponseBody
	GetData() *GetAsyncTranslateResponseBodyData
	SetMessage(v string) *GetAsyncTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAsyncTranslateResponseBody
	GetRequestId() *string
}

type GetAsyncTranslateResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAsyncTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAsyncTranslateResponseBody) GetData() *GetAsyncTranslateResponseBodyData {
	return s.Data
}

func (s *GetAsyncTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAsyncTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsyncTranslateResponseBody) SetCode(v int32) *GetAsyncTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncTranslateResponseBody) SetData(v *GetAsyncTranslateResponseBodyData) *GetAsyncTranslateResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncTranslateResponseBody) SetMessage(v string) *GetAsyncTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncTranslateResponseBody) SetRequestId(v string) *GetAsyncTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncTranslateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAsyncTranslateResponseBodyData struct {
	// example:
	//
	// zh
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	// example:
	//
	// ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// hello
	TranslatedText *string `json:"TranslatedText,omitempty" xml:"TranslatedText,omitempty"`
	// example:
	//
	// 2
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s GetAsyncTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateResponseBodyData) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *GetAsyncTranslateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAsyncTranslateResponseBodyData) GetTranslatedText() *string {
	return s.TranslatedText
}

func (s *GetAsyncTranslateResponseBodyData) GetWordCount() *string {
	return s.WordCount
}

func (s *GetAsyncTranslateResponseBodyData) SetDetectedLanguage(v string) *GetAsyncTranslateResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *GetAsyncTranslateResponseBodyData) SetStatus(v string) *GetAsyncTranslateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAsyncTranslateResponseBodyData) SetTranslatedText(v string) *GetAsyncTranslateResponseBodyData {
	s.TranslatedText = &v
	return s
}

func (s *GetAsyncTranslateResponseBodyData) SetWordCount(v string) *GetAsyncTranslateResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *GetAsyncTranslateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
