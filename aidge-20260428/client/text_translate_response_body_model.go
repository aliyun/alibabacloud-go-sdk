// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TextTranslateResponseBody
	GetCode() *string
	SetData(v *TextTranslateResponseBodyData) *TextTranslateResponseBody
	GetData() *TextTranslateResponseBodyData
	SetMessage(v string) *TextTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *TextTranslateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TextTranslateResponseBody
	GetSuccess() *bool
}

type TextTranslateResponseBody struct {
	// The response code. The value "success" is returned for a successful call.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The translation result data, which contains the translation list and usage information.
	Data *TextTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. The value "Success" is returned for a successful call. For a failed call, a specific error message is returned, such as "The parameters contain sensitive information. Try other input.".
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// 922E43BB-EE0E-1A29-B143-BB91BB3EA6AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TextTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *TextTranslateResponseBody) GetCode() *string {
	return s.Code
}

func (s *TextTranslateResponseBody) GetData() *TextTranslateResponseBodyData {
	return s.Data
}

func (s *TextTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TextTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TextTranslateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TextTranslateResponseBody) SetCode(v string) *TextTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *TextTranslateResponseBody) SetData(v *TextTranslateResponseBodyData) *TextTranslateResponseBody {
	s.Data = v
	return s
}

func (s *TextTranslateResponseBody) SetMessage(v string) *TextTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *TextTranslateResponseBody) SetRequestId(v string) *TextTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *TextTranslateResponseBody) SetSuccess(v bool) *TextTranslateResponseBody {
	s.Success = &v
	return s
}

func (s *TextTranslateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextTranslateResponseBodyData struct {
	// The list of translation results. Each element corresponds to a translation result for an entry in the input text list.
	Translations []*TextTranslateResponseBodyDataTranslations `json:"Translations,omitempty" xml:"Translations,omitempty" type:"Repeated"`
	// The usage information, including the number of input characters.
	//
	// example:
	//
	// {"InputCharacterCount":5}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s TextTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextTranslateResponseBodyData) GetTranslations() []*TextTranslateResponseBodyDataTranslations {
	return s.Translations
}

func (s *TextTranslateResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *TextTranslateResponseBodyData) SetTranslations(v []*TextTranslateResponseBodyDataTranslations) *TextTranslateResponseBodyData {
	s.Translations = v
	return s
}

func (s *TextTranslateResponseBodyData) SetUsageMap(v map[string]*int64) *TextTranslateResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *TextTranslateResponseBodyData) Validate() error {
	if s.Translations != nil {
		for _, item := range s.Translations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TextTranslateResponseBodyDataTranslations struct {
	// The number of characters in the source text.
	//
	// example:
	//
	// 11
	Characters *int64 `json:"Characters,omitempty" xml:"Characters,omitempty"`
	// The automatically detected source language code.
	//
	// example:
	//
	// en
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	// The translated text.
	//
	// example:
	//
	// 你好世界
	TranslatedText *string `json:"TranslatedText,omitempty" xml:"TranslatedText,omitempty"`
}

func (s TextTranslateResponseBodyDataTranslations) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateResponseBodyDataTranslations) GoString() string {
	return s.String()
}

func (s *TextTranslateResponseBodyDataTranslations) GetCharacters() *int64 {
	return s.Characters
}

func (s *TextTranslateResponseBodyDataTranslations) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *TextTranslateResponseBodyDataTranslations) GetTranslatedText() *string {
	return s.TranslatedText
}

func (s *TextTranslateResponseBodyDataTranslations) SetCharacters(v int64) *TextTranslateResponseBodyDataTranslations {
	s.Characters = &v
	return s
}

func (s *TextTranslateResponseBodyDataTranslations) SetDetectedLanguage(v string) *TextTranslateResponseBodyDataTranslations {
	s.DetectedLanguage = &v
	return s
}

func (s *TextTranslateResponseBodyDataTranslations) SetTranslatedText(v string) *TextTranslateResponseBodyDataTranslations {
	s.TranslatedText = &v
	return s
}

func (s *TextTranslateResponseBodyDataTranslations) Validate() error {
	return dara.Validate(s)
}
