// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTranslateByVPCRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiType(v string) *GetBatchTranslateByVPCRequest
	GetApiType() *string
	SetFormatType(v string) *GetBatchTranslateByVPCRequest
	GetFormatType() *string
	SetScene(v string) *GetBatchTranslateByVPCRequest
	GetScene() *string
	SetSourceLanguage(v string) *GetBatchTranslateByVPCRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *GetBatchTranslateByVPCRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *GetBatchTranslateByVPCRequest
	GetTargetLanguage() *string
}

type GetBatchTranslateByVPCRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// translate_standard
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s GetBatchTranslateByVPCRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTranslateByVPCRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateByVPCRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetBatchTranslateByVPCRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *GetBatchTranslateByVPCRequest) GetScene() *string {
	return s.Scene
}

func (s *GetBatchTranslateByVPCRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *GetBatchTranslateByVPCRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *GetBatchTranslateByVPCRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *GetBatchTranslateByVPCRequest) SetApiType(v string) *GetBatchTranslateByVPCRequest {
	s.ApiType = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetFormatType(v string) *GetBatchTranslateByVPCRequest {
	s.FormatType = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetScene(v string) *GetBatchTranslateByVPCRequest {
	s.Scene = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetSourceLanguage(v string) *GetBatchTranslateByVPCRequest {
	s.SourceLanguage = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetSourceText(v string) *GetBatchTranslateByVPCRequest {
	s.SourceText = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) SetTargetLanguage(v string) *GetBatchTranslateByVPCRequest {
	s.TargetLanguage = &v
	return s
}

func (s *GetBatchTranslateByVPCRequest) Validate() error {
	return dara.Validate(s)
}
