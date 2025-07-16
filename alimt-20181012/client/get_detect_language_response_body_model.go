// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectLanguageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectedLanguage(v string) *GetDetectLanguageResponseBody
	GetDetectedLanguage() *string
	SetLanguageProbabilities(v string) *GetDetectLanguageResponseBody
	GetLanguageProbabilities() *string
	SetRequestId(v string) *GetDetectLanguageResponseBody
	GetRequestId() *string
}

type GetDetectLanguageResponseBody struct {
	// example:
	//
	// zh
	DetectedLanguage      *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	LanguageProbabilities *string `json:"LanguageProbabilities,omitempty" xml:"LanguageProbabilities,omitempty"`
	// example:
	//
	// 0C5EC1EC-1A06-4D60-97E6-4D41350945E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectLanguageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDetectLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageResponseBody) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *GetDetectLanguageResponseBody) GetLanguageProbabilities() *string {
	return s.LanguageProbabilities
}

func (s *GetDetectLanguageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDetectLanguageResponseBody) SetDetectedLanguage(v string) *GetDetectLanguageResponseBody {
	s.DetectedLanguage = &v
	return s
}

func (s *GetDetectLanguageResponseBody) SetLanguageProbabilities(v string) *GetDetectLanguageResponseBody {
	s.LanguageProbabilities = &v
	return s
}

func (s *GetDetectLanguageResponseBody) SetRequestId(v string) *GetDetectLanguageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectLanguageResponseBody) Validate() error {
	return dara.Validate(s)
}
