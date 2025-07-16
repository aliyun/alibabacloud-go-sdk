// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectLanguageVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceText(v string) *GetDetectLanguageVpcRequest
	GetSourceText() *string
}

type GetDetectLanguageVpcRequest struct {
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
}

func (s GetDetectLanguageVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDetectLanguageVpcRequest) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageVpcRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *GetDetectLanguageVpcRequest) SetSourceText(v string) *GetDetectLanguageVpcRequest {
	s.SourceText = &v
	return s
}

func (s *GetDetectLanguageVpcRequest) Validate() error {
	return dara.Validate(s)
}
