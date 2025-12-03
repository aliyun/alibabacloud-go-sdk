// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateECommerceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TranslateECommerceResponseBody
	GetCode() *int32
	SetData(v *TranslateECommerceResponseBodyData) *TranslateECommerceResponseBody
	GetData() *TranslateECommerceResponseBodyData
	SetMessage(v string) *TranslateECommerceResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranslateECommerceResponseBody
	GetRequestId() *string
}

type TranslateECommerceResponseBody struct {
	// example:
	//
	// 200
	Code *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateECommerceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// An error occurred while translating.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CC93BB5C-EAB5-579B-AA44-F61528F48FF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateECommerceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateECommerceResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TranslateECommerceResponseBody) GetData() *TranslateECommerceResponseBodyData {
	return s.Data
}

func (s *TranslateECommerceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranslateECommerceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateECommerceResponseBody) SetCode(v int32) *TranslateECommerceResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetData(v *TranslateECommerceResponseBodyData) *TranslateECommerceResponseBody {
	s.Data = v
	return s
}

func (s *TranslateECommerceResponseBody) SetMessage(v string) *TranslateECommerceResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateECommerceResponseBody) SetRequestId(v string) *TranslateECommerceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateECommerceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TranslateECommerceResponseBodyData struct {
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	Translated       *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s TranslateECommerceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateECommerceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponseBodyData) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *TranslateECommerceResponseBodyData) GetTranslated() *string {
	return s.Translated
}

func (s *TranslateECommerceResponseBodyData) GetWordCount() *string {
	return s.WordCount
}

func (s *TranslateECommerceResponseBodyData) SetDetectedLanguage(v string) *TranslateECommerceResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TranslateECommerceResponseBodyData) SetTranslated(v string) *TranslateECommerceResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateECommerceResponseBodyData) SetWordCount(v string) *TranslateECommerceResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *TranslateECommerceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
