// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioModelCode(v string) *CreateVocabRequest
	GetAudioModelCode() *string
	SetDescription(v string) *CreateVocabRequest
	GetDescription() *string
	SetName(v string) *CreateVocabRequest
	GetName() *string
	SetWordWeightList(v []*CreateVocabRequestWordWeightList) *CreateVocabRequest
	GetWordWeightList() []*CreateVocabRequestWordWeightList
	SetWorkspaceId(v string) *CreateVocabRequest
	GetWorkspaceId() *string
}

type CreateVocabRequest struct {
	// example:
	//
	// nls
	AudioModelCode *string `json:"audioModelCode,omitempty" xml:"audioModelCode,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	WordWeightList []*CreateVocabRequestWordWeightList `json:"wordWeightList,omitempty" xml:"wordWeightList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-9****me1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabRequest) GoString() string {
	return s.String()
}

func (s *CreateVocabRequest) GetAudioModelCode() *string {
	return s.AudioModelCode
}

func (s *CreateVocabRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVocabRequest) GetName() *string {
	return s.Name
}

func (s *CreateVocabRequest) GetWordWeightList() []*CreateVocabRequestWordWeightList {
	return s.WordWeightList
}

func (s *CreateVocabRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateVocabRequest) SetAudioModelCode(v string) *CreateVocabRequest {
	s.AudioModelCode = &v
	return s
}

func (s *CreateVocabRequest) SetDescription(v string) *CreateVocabRequest {
	s.Description = &v
	return s
}

func (s *CreateVocabRequest) SetName(v string) *CreateVocabRequest {
	s.Name = &v
	return s
}

func (s *CreateVocabRequest) SetWordWeightList(v []*CreateVocabRequestWordWeightList) *CreateVocabRequest {
	s.WordWeightList = v
	return s
}

func (s *CreateVocabRequest) SetWorkspaceId(v string) *CreateVocabRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateVocabRequest) Validate() error {
	return dara.Validate(s)
}

type CreateVocabRequestWordWeightList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
	// This parameter is required.
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s CreateVocabRequestWordWeightList) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabRequestWordWeightList) GoString() string {
	return s.String()
}

func (s *CreateVocabRequestWordWeightList) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVocabRequestWordWeightList) GetWord() *string {
	return s.Word
}

func (s *CreateVocabRequestWordWeightList) SetWeight(v int32) *CreateVocabRequestWordWeightList {
	s.Weight = &v
	return s
}

func (s *CreateVocabRequestWordWeightList) SetWord(v string) *CreateVocabRequestWordWeightList {
	s.Word = &v
	return s
}

func (s *CreateVocabRequestWordWeightList) Validate() error {
	return dara.Validate(s)
}
