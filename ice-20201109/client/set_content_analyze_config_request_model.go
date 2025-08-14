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
	// example:
	//
	// true
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// example:
	//
	// TEXT,FACE
	SaveType *string `json:"SaveType,omitempty" xml:"SaveType,omitempty"`
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
