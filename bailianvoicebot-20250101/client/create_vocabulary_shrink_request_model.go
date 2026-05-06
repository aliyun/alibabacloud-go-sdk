// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabularyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *CreateVocabularyShrinkRequest
	GetBusinessUnitId() *string
	SetDescription(v string) *CreateVocabularyShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateVocabularyShrinkRequest
	GetName() *string
	SetWordsShrink(v string) *CreateVocabularyShrinkRequest
	GetWordsShrink() *string
}

type CreateVocabularyShrinkRequest struct {
	// example:
	//
	// llm-baployoyopf22m2r
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WordsShrink    *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s CreateVocabularyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabularyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVocabularyShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateVocabularyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVocabularyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateVocabularyShrinkRequest) GetWordsShrink() *string {
	return s.WordsShrink
}

func (s *CreateVocabularyShrinkRequest) SetBusinessUnitId(v string) *CreateVocabularyShrinkRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *CreateVocabularyShrinkRequest) SetDescription(v string) *CreateVocabularyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateVocabularyShrinkRequest) SetName(v string) *CreateVocabularyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateVocabularyShrinkRequest) SetWordsShrink(v string) *CreateVocabularyShrinkRequest {
	s.WordsShrink = &v
	return s
}

func (s *CreateVocabularyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
