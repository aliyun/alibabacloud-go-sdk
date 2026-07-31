// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepaidTextTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PrepaidTextTranslateResponseBody
	GetCode() *string
	SetData(v *PrepaidTextTranslateResponseBodyData) *PrepaidTextTranslateResponseBody
	GetData() *PrepaidTextTranslateResponseBodyData
	SetMessage(v string) *PrepaidTextTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *PrepaidTextTranslateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PrepaidTextTranslateResponseBody
	GetSuccess() *bool
}

type PrepaidTextTranslateResponseBody struct {
	// The response code. Returns "success" for normal calls.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The translation result data, including the translation list and usage information.
	Data *PrepaidTextTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. Returns "Success" for normal calls. Returns specific error information for exceptions, such as "The parameters contain sensitive information. Try other input."
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, used to identify a unique request call.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. true indicates success. false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PrepaidTextTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrepaidTextTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *PrepaidTextTranslateResponseBody) GetCode() *string {
	return s.Code
}

func (s *PrepaidTextTranslateResponseBody) GetData() *PrepaidTextTranslateResponseBodyData {
	return s.Data
}

func (s *PrepaidTextTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PrepaidTextTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrepaidTextTranslateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PrepaidTextTranslateResponseBody) SetCode(v string) *PrepaidTextTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *PrepaidTextTranslateResponseBody) SetData(v *PrepaidTextTranslateResponseBodyData) *PrepaidTextTranslateResponseBody {
	s.Data = v
	return s
}

func (s *PrepaidTextTranslateResponseBody) SetMessage(v string) *PrepaidTextTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *PrepaidTextTranslateResponseBody) SetRequestId(v string) *PrepaidTextTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrepaidTextTranslateResponseBody) SetSuccess(v bool) *PrepaidTextTranslateResponseBody {
	s.Success = &v
	return s
}

func (s *PrepaidTextTranslateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PrepaidTextTranslateResponseBodyData struct {
	// The translation result list. Each element corresponds to a translation result for an entry in the input text list.
	Translations []*PrepaidTextTranslateResponseBodyDataTranslations `json:"Translations,omitempty" xml:"Translations,omitempty" type:"Repeated"`
	// The usage information, including the input character count.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s PrepaidTextTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PrepaidTextTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *PrepaidTextTranslateResponseBodyData) GetTranslations() []*PrepaidTextTranslateResponseBodyDataTranslations {
	return s.Translations
}

func (s *PrepaidTextTranslateResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *PrepaidTextTranslateResponseBodyData) SetTranslations(v []*PrepaidTextTranslateResponseBodyDataTranslations) *PrepaidTextTranslateResponseBodyData {
	s.Translations = v
	return s
}

func (s *PrepaidTextTranslateResponseBodyData) SetUsageMap(v map[string]*int64) *PrepaidTextTranslateResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *PrepaidTextTranslateResponseBodyData) Validate() error {
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

type PrepaidTextTranslateResponseBodyDataTranslations struct {
	// The character count of the source text.
	//
	// example:
	//
	// 11
	Characters *int64 `json:"Characters,omitempty" xml:"Characters,omitempty"`
	// The automatically detected source language.
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

func (s PrepaidTextTranslateResponseBodyDataTranslations) String() string {
	return dara.Prettify(s)
}

func (s PrepaidTextTranslateResponseBodyDataTranslations) GoString() string {
	return s.String()
}

func (s *PrepaidTextTranslateResponseBodyDataTranslations) GetCharacters() *int64 {
	return s.Characters
}

func (s *PrepaidTextTranslateResponseBodyDataTranslations) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *PrepaidTextTranslateResponseBodyDataTranslations) GetTranslatedText() *string {
	return s.TranslatedText
}

func (s *PrepaidTextTranslateResponseBodyDataTranslations) SetCharacters(v int64) *PrepaidTextTranslateResponseBodyDataTranslations {
	s.Characters = &v
	return s
}

func (s *PrepaidTextTranslateResponseBodyDataTranslations) SetDetectedLanguage(v string) *PrepaidTextTranslateResponseBodyDataTranslations {
	s.DetectedLanguage = &v
	return s
}

func (s *PrepaidTextTranslateResponseBodyDataTranslations) SetTranslatedText(v string) *PrepaidTextTranslateResponseBodyDataTranslations {
	s.TranslatedText = &v
	return s
}

func (s *PrepaidTextTranslateResponseBodyDataTranslations) Validate() error {
	return dara.Validate(s)
}
