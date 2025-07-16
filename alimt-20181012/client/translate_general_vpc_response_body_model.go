// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateGeneralVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TranslateGeneralVpcResponseBody
	GetCode() *int32
	SetData(v *TranslateGeneralVpcResponseBodyData) *TranslateGeneralVpcResponseBody
	GetData() *TranslateGeneralVpcResponseBodyData
	SetMessage(v string) *TranslateGeneralVpcResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranslateGeneralVpcResponseBody
	GetRequestId() *string
}

type TranslateGeneralVpcResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 200
	Code *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateGeneralVpcResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// translate from source to target not support
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateGeneralVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralVpcResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TranslateGeneralVpcResponseBody) GetData() *TranslateGeneralVpcResponseBodyData {
	return s.Data
}

func (s *TranslateGeneralVpcResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranslateGeneralVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateGeneralVpcResponseBody) SetCode(v int32) *TranslateGeneralVpcResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateGeneralVpcResponseBody) SetData(v *TranslateGeneralVpcResponseBodyData) *TranslateGeneralVpcResponseBody {
	s.Data = v
	return s
}

func (s *TranslateGeneralVpcResponseBody) SetMessage(v string) *TranslateGeneralVpcResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateGeneralVpcResponseBody) SetRequestId(v string) *TranslateGeneralVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateGeneralVpcResponseBody) Validate() error {
	return dara.Validate(s)
}

type TranslateGeneralVpcResponseBodyData struct {
	// example:
	//
	// zh
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

func (s TranslateGeneralVpcResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralVpcResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcResponseBodyData) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *TranslateGeneralVpcResponseBodyData) GetTranslated() *string {
	return s.Translated
}

func (s *TranslateGeneralVpcResponseBodyData) GetWordCount() *string {
	return s.WordCount
}

func (s *TranslateGeneralVpcResponseBodyData) SetDetectedLanguage(v string) *TranslateGeneralVpcResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TranslateGeneralVpcResponseBodyData) SetTranslated(v string) *TranslateGeneralVpcResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateGeneralVpcResponseBodyData) SetWordCount(v string) *TranslateGeneralVpcResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *TranslateGeneralVpcResponseBodyData) Validate() error {
	return dara.Validate(s)
}
