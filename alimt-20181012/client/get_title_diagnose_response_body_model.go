// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleDiagnoseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTitleDiagnoseResponseBody
	GetCode() *int32
	SetData(v *GetTitleDiagnoseResponseBodyData) *GetTitleDiagnoseResponseBody
	GetData() *GetTitleDiagnoseResponseBodyData
	SetMessage(v string) *GetTitleDiagnoseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTitleDiagnoseResponseBody
	GetRequestId() *string
}

type GetTitleDiagnoseResponseBody struct {
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTitleDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTitleDiagnoseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTitleDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTitleDiagnoseResponseBody) GetData() *GetTitleDiagnoseResponseBodyData {
	return s.Data
}

func (s *GetTitleDiagnoseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTitleDiagnoseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTitleDiagnoseResponseBody) SetCode(v int32) *GetTitleDiagnoseResponseBody {
	s.Code = &v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetData(v *GetTitleDiagnoseResponseBodyData) *GetTitleDiagnoseResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetMessage(v string) *GetTitleDiagnoseResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleDiagnoseResponseBody) SetRequestId(v string) *GetTitleDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTitleDiagnoseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTitleDiagnoseResponseBodyData struct {
	// example:
	//
	// Boy
	AllUppercaseWords *string `json:"AllUppercaseWords,omitempty" xml:"AllUppercaseWords,omitempty"`
	// example:
	//
	// true
	ContainCoreClasses *string `json:"ContainCoreClasses,omitempty" xml:"ContainCoreClasses,omitempty"`
	// example:
	//
	// baba
	DisableWords *string `json:"DisableWords,omitempty" xml:"DisableWords,omitempty"`
	// example:
	//
	// hi
	DuplicateWords *string `json:"DuplicateWords,omitempty" xml:"DuplicateWords,omitempty"`
	// example:
	//
	// 2
	LanguageQualityScore *string `json:"LanguageQualityScore,omitempty" xml:"LanguageQualityScore,omitempty"`
	// example:
	//
	// no
	NoFirstUppercaseList *string `json:"NoFirstUppercaseList,omitempty" xml:"NoFirstUppercaseList,omitempty"`
	// example:
	//
	// 100
	OverLengthLimit *string `json:"OverLengthLimit,omitempty" xml:"OverLengthLimit,omitempty"`
	// example:
	//
	// 3
	TotalScore *string `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
	// example:
	//
	// 10
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	// example:
	//
	// ahk
	WordSpelledCorrectError *string `json:"WordSpelledCorrectError,omitempty" xml:"WordSpelledCorrectError,omitempty"`
}

func (s GetTitleDiagnoseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTitleDiagnoseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseResponseBodyData) GetAllUppercaseWords() *string {
	return s.AllUppercaseWords
}

func (s *GetTitleDiagnoseResponseBodyData) GetContainCoreClasses() *string {
	return s.ContainCoreClasses
}

func (s *GetTitleDiagnoseResponseBodyData) GetDisableWords() *string {
	return s.DisableWords
}

func (s *GetTitleDiagnoseResponseBodyData) GetDuplicateWords() *string {
	return s.DuplicateWords
}

func (s *GetTitleDiagnoseResponseBodyData) GetLanguageQualityScore() *string {
	return s.LanguageQualityScore
}

func (s *GetTitleDiagnoseResponseBodyData) GetNoFirstUppercaseList() *string {
	return s.NoFirstUppercaseList
}

func (s *GetTitleDiagnoseResponseBodyData) GetOverLengthLimit() *string {
	return s.OverLengthLimit
}

func (s *GetTitleDiagnoseResponseBodyData) GetTotalScore() *string {
	return s.TotalScore
}

func (s *GetTitleDiagnoseResponseBodyData) GetWordCount() *string {
	return s.WordCount
}

func (s *GetTitleDiagnoseResponseBodyData) GetWordSpelledCorrectError() *string {
	return s.WordSpelledCorrectError
}

func (s *GetTitleDiagnoseResponseBodyData) SetAllUppercaseWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.AllUppercaseWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetContainCoreClasses(v string) *GetTitleDiagnoseResponseBodyData {
	s.ContainCoreClasses = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetDisableWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.DisableWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetDuplicateWords(v string) *GetTitleDiagnoseResponseBodyData {
	s.DuplicateWords = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetLanguageQualityScore(v string) *GetTitleDiagnoseResponseBodyData {
	s.LanguageQualityScore = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetNoFirstUppercaseList(v string) *GetTitleDiagnoseResponseBodyData {
	s.NoFirstUppercaseList = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetOverLengthLimit(v string) *GetTitleDiagnoseResponseBodyData {
	s.OverLengthLimit = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetTotalScore(v string) *GetTitleDiagnoseResponseBodyData {
	s.TotalScore = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetWordCount(v string) *GetTitleDiagnoseResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) SetWordSpelledCorrectError(v string) *GetTitleDiagnoseResponseBodyData {
	s.WordSpelledCorrectError = &v
	return s
}

func (s *GetTitleDiagnoseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
