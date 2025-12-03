// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateGeneralResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TranslateGeneralResponseBody
	GetCode() *int32
	SetData(v *TranslateGeneralResponseBodyData) *TranslateGeneralResponseBody
	GetData() *TranslateGeneralResponseBodyData
	SetMessage(v string) *TranslateGeneralResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranslateGeneralResponseBody
	GetRequestId() *string
}

type TranslateGeneralResponseBody struct {
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateGeneralResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateGeneralResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TranslateGeneralResponseBody) GetData() *TranslateGeneralResponseBodyData {
	return s.Data
}

func (s *TranslateGeneralResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranslateGeneralResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateGeneralResponseBody) SetCode(v int32) *TranslateGeneralResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetData(v *TranslateGeneralResponseBodyData) *TranslateGeneralResponseBody {
	s.Data = v
	return s
}

func (s *TranslateGeneralResponseBody) SetMessage(v string) *TranslateGeneralResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateGeneralResponseBody) SetRequestId(v string) *TranslateGeneralResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateGeneralResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TranslateGeneralResponseBodyData struct {
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

func (s TranslateGeneralResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponseBodyData) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *TranslateGeneralResponseBodyData) GetTranslated() *string {
	return s.Translated
}

func (s *TranslateGeneralResponseBodyData) GetWordCount() *string {
	return s.WordCount
}

func (s *TranslateGeneralResponseBodyData) SetDetectedLanguage(v string) *TranslateGeneralResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TranslateGeneralResponseBodyData) SetTranslated(v string) *TranslateGeneralResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateGeneralResponseBodyData) SetWordCount(v string) *TranslateGeneralResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *TranslateGeneralResponseBodyData) Validate() error {
	return dara.Validate(s)
}
