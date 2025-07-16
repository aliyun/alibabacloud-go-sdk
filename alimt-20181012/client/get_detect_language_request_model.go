// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectLanguageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceText(v string) *GetDetectLanguageRequest
	GetSourceText() *string
}

type GetDetectLanguageRequest struct {
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
}

func (s GetDetectLanguageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDetectLanguageRequest) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *GetDetectLanguageRequest) SetSourceText(v string) *GetDetectLanguageRequest {
	s.SourceText = &v
	return s
}

func (s *GetDetectLanguageRequest) Validate() error {
	return dara.Validate(s)
}
