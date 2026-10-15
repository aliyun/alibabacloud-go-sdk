// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPetHealthAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v []*string) *PetHealthAnalysisRequest
	GetImageUrl() []*string
	SetSystemPrompt(v string) *PetHealthAnalysisRequest
	GetSystemPrompt() *string
	SetUserPrompt(v string) *PetHealthAnalysisRequest
	GetUserPrompt() *string
}

type PetHealthAnalysisRequest struct {
	// This parameter is required.
	ImageUrl []*string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" type:"Repeated"`
	// example:
	//
	// 你是一名专业的宠物医生
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// example:
	//
	// 请分析这只宠物的健康状况
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s PetHealthAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s PetHealthAnalysisRequest) GoString() string {
	return s.String()
}

func (s *PetHealthAnalysisRequest) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *PetHealthAnalysisRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *PetHealthAnalysisRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *PetHealthAnalysisRequest) SetImageUrl(v []*string) *PetHealthAnalysisRequest {
	s.ImageUrl = v
	return s
}

func (s *PetHealthAnalysisRequest) SetSystemPrompt(v string) *PetHealthAnalysisRequest {
	s.SystemPrompt = &v
	return s
}

func (s *PetHealthAnalysisRequest) SetUserPrompt(v string) *PetHealthAnalysisRequest {
	s.UserPrompt = &v
	return s
}

func (s *PetHealthAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
