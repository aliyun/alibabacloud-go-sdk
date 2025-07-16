// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TranslateResponseBody
	GetCode() *int32
	SetData(v *TranslateResponseBodyData) *TranslateResponseBody
	GetData() *TranslateResponseBodyData
	SetMessage(v string) *TranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranslateResponseBody
	GetRequestId() *string
}

type TranslateResponseBody struct {
	// example:
	//
	// 200
	Code *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TranslateResponseBody) GetData() *TranslateResponseBodyData {
	return s.Data
}

func (s *TranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateResponseBody) SetCode(v int32) *TranslateResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateResponseBody) SetData(v *TranslateResponseBodyData) *TranslateResponseBody {
	s.Data = v
	return s
}

func (s *TranslateResponseBody) SetMessage(v string) *TranslateResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateResponseBody) SetRequestId(v string) *TranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateResponseBody) Validate() error {
	return dara.Validate(s)
}

type TranslateResponseBodyData struct {
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	// example:
	//
	// Hello
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateResponseBodyData) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *TranslateResponseBodyData) GetTranslated() *string {
	return s.Translated
}

func (s *TranslateResponseBodyData) GetWordCount() *string {
	return s.WordCount
}

func (s *TranslateResponseBodyData) SetDetectedLanguage(v string) *TranslateResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TranslateResponseBodyData) SetTranslated(v string) *TranslateResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateResponseBodyData) SetWordCount(v string) *TranslateResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *TranslateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
