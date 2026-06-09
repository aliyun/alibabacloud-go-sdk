// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCutQuestionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v string) *CutQuestionsRequest
	GetImage() *string
	SetParameters(v *CutQuestionsRequestParameters) *CutQuestionsRequest
	GetParameters() *CutQuestionsRequestParameters
	SetWorkspaceId(v string) *CutQuestionsRequest
	GetWorkspaceId() *string
}

type CutQuestionsRequest struct {
	// This parameter is required.
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// This parameter is required.
	Parameters *CutQuestionsRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// llm-1ijrzuv3v0ivvls7
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CutQuestionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CutQuestionsRequest) GoString() string {
	return s.String()
}

func (s *CutQuestionsRequest) GetImage() *string {
	return s.Image
}

func (s *CutQuestionsRequest) GetParameters() *CutQuestionsRequestParameters {
	return s.Parameters
}

func (s *CutQuestionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CutQuestionsRequest) SetImage(v string) *CutQuestionsRequest {
	s.Image = &v
	return s
}

func (s *CutQuestionsRequest) SetParameters(v *CutQuestionsRequestParameters) *CutQuestionsRequest {
	s.Parameters = v
	return s
}

func (s *CutQuestionsRequest) SetWorkspaceId(v string) *CutQuestionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CutQuestionsRequest) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CutQuestionsRequestParameters struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	ExtractImages *bool `json:"extract_images,omitempty" xml:"extract_images,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Struct *bool `json:"struct,omitempty" xml:"struct,omitempty"`
}

func (s CutQuestionsRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CutQuestionsRequestParameters) GoString() string {
	return s.String()
}

func (s *CutQuestionsRequestParameters) GetExtractImages() *bool {
	return s.ExtractImages
}

func (s *CutQuestionsRequestParameters) GetStruct() *bool {
	return s.Struct
}

func (s *CutQuestionsRequestParameters) SetExtractImages(v bool) *CutQuestionsRequestParameters {
	s.ExtractImages = &v
	return s
}

func (s *CutQuestionsRequestParameters) SetStruct(v bool) *CutQuestionsRequestParameters {
	s.Struct = &v
	return s
}

func (s *CutQuestionsRequestParameters) Validate() error {
	return dara.Validate(s)
}
