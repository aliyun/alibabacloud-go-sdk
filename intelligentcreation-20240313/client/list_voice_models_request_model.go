// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListVoiceModelsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoiceModelsRequest
	GetPageSize() *int32
	SetResSpecType(v string) *ListVoiceModelsRequest
	GetResSpecType() *string
	SetUseScene(v string) *ListVoiceModelsRequest
	GetUseScene() *string
	SetVoiceLanguage(v string) *ListVoiceModelsRequest
	GetVoiceLanguage() *string
	SetVoiceType(v string) *ListVoiceModelsRequest
	GetVoiceType() *string
}

type ListVoiceModelsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResSpecType *string `json:"resSpecType,omitempty" xml:"resSpecType,omitempty"`
	// example:
	//
	// offlineSynthesis
	UseScene      *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	// example:
	//
	// PRIVATE_VOICE
	VoiceType *string `json:"voiceType,omitempty" xml:"voiceType,omitempty"`
}

func (s ListVoiceModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceModelsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceModelsRequest) GetResSpecType() *string {
	return s.ResSpecType
}

func (s *ListVoiceModelsRequest) GetUseScene() *string {
	return s.UseScene
}

func (s *ListVoiceModelsRequest) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *ListVoiceModelsRequest) GetVoiceType() *string {
	return s.VoiceType
}

func (s *ListVoiceModelsRequest) SetPageNumber(v int32) *ListVoiceModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceModelsRequest) SetPageSize(v int32) *ListVoiceModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoiceModelsRequest) SetResSpecType(v string) *ListVoiceModelsRequest {
	s.ResSpecType = &v
	return s
}

func (s *ListVoiceModelsRequest) SetUseScene(v string) *ListVoiceModelsRequest {
	s.UseScene = &v
	return s
}

func (s *ListVoiceModelsRequest) SetVoiceLanguage(v string) *ListVoiceModelsRequest {
	s.VoiceLanguage = &v
	return s
}

func (s *ListVoiceModelsRequest) SetVoiceType(v string) *ListVoiceModelsRequest {
	s.VoiceType = &v
	return s
}

func (s *ListVoiceModelsRequest) Validate() error {
	return dara.Validate(s)
}
