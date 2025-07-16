// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectLanguageVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectedLanguage(v string) *GetDetectLanguageVpcResponseBody
	GetDetectedLanguage() *string
	SetLanguageProbabilities(v string) *GetDetectLanguageVpcResponseBody
	GetLanguageProbabilities() *string
	SetRequestId(v string) *GetDetectLanguageVpcResponseBody
	GetRequestId() *string
}

type GetDetectLanguageVpcResponseBody struct {
	DetectedLanguage      *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	LanguageProbabilities *string `json:"LanguageProbabilities,omitempty" xml:"LanguageProbabilities,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectLanguageVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDetectLanguageVpcResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageVpcResponseBody) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *GetDetectLanguageVpcResponseBody) GetLanguageProbabilities() *string {
	return s.LanguageProbabilities
}

func (s *GetDetectLanguageVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDetectLanguageVpcResponseBody) SetDetectedLanguage(v string) *GetDetectLanguageVpcResponseBody {
	s.DetectedLanguage = &v
	return s
}

func (s *GetDetectLanguageVpcResponseBody) SetLanguageProbabilities(v string) *GetDetectLanguageVpcResponseBody {
	s.LanguageProbabilities = &v
	return s
}

func (s *GetDetectLanguageVpcResponseBody) SetRequestId(v string) *GetDetectLanguageVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectLanguageVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
