// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabularyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVocabularyShrinkRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateVocabularyShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateVocabularyShrinkRequest
	GetName() *string
	SetWordsShrink(v string) *CreateVocabularyShrinkRequest
	GetWordsShrink() *string
}

type CreateVocabularyShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5a
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WordsShrink *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s CreateVocabularyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabularyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVocabularyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVocabularyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVocabularyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateVocabularyShrinkRequest) GetWordsShrink() *string {
	return s.WordsShrink
}

func (s *CreateVocabularyShrinkRequest) SetDescription(v string) *CreateVocabularyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateVocabularyShrinkRequest) SetInstanceId(v string) *CreateVocabularyShrinkRequest {
	s.InstanceId = &v
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
