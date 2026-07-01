// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetContentAnalyzeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuto(v bool) *SetContentAnalyzeConfigRequest
	GetAuto() *bool
	SetSaveType(v string) *SetContentAnalyzeConfigRequest
	GetSaveType() *string
	SetTemplateId(v string) *SetContentAnalyzeConfigRequest
	GetTemplateId() *string
}

type SetContentAnalyzeConfigRequest struct {
	// Specifies whether to automatically start Intelligent Content Analysis after a Media Asset is registered.
	//
	// Valid values:
	//
	// - true: Enable
	//
	// - false: Disable (Default)
	//
	// example:
	//
	// true
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// The storage type for analysis results. This parameter applies only when Auto is set to true. Default: Empty.
	//
	// - TEXT: Label
	//
	// - FACE: Face
	//
	// - DNA: Similar Image
	//
	// You can specify multiple values separated by commas. If this parameter is empty, Intelligent Content Analysis results are not saved to any search library, and you cannot perform content searches.
	//
	// example:
	//
	// TEXT,FACE
	SaveType *string `json:"SaveType,omitempty" xml:"SaveType,omitempty"`
	// The ID of the Intelligent Content Analysis Template. Each Template includes the following AI analysis features:
	//
	// - S00000101-100040: Text Recognition
	//
	// - S00000101-100060: Video Categorization and Face Recognition
	//
	// - S00000101-100070: Text Recognition, Video Categorization, and Face Recognition
	//
	// example:
	//
	// S00000101-100070
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SetContentAnalyzeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetContentAnalyzeConfigRequest) GoString() string {
	return s.String()
}

func (s *SetContentAnalyzeConfigRequest) GetAuto() *bool {
	return s.Auto
}

func (s *SetContentAnalyzeConfigRequest) GetSaveType() *string {
	return s.SaveType
}

func (s *SetContentAnalyzeConfigRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetContentAnalyzeConfigRequest) SetAuto(v bool) *SetContentAnalyzeConfigRequest {
	s.Auto = &v
	return s
}

func (s *SetContentAnalyzeConfigRequest) SetSaveType(v string) *SetContentAnalyzeConfigRequest {
	s.SaveType = &v
	return s
}

func (s *SetContentAnalyzeConfigRequest) SetTemplateId(v string) *SetContentAnalyzeConfigRequest {
	s.TemplateId = &v
	return s
}

func (s *SetContentAnalyzeConfigRequest) Validate() error {
	return dara.Validate(s)
}
