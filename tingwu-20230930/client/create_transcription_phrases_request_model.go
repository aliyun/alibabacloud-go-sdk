// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTranscriptionPhrasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTranscriptionPhrasesRequest
	GetDescription() *string
	SetName(v string) *CreateTranscriptionPhrasesRequest
	GetName() *string
	SetWordWeights(v string) *CreateTranscriptionPhrasesRequest
	GetWordWeights() *string
}

type CreateTranscriptionPhrasesRequest struct {
	// example:
	//
	// custom fruit phrases list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fruit_phrase
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"苹果":3,"西瓜":3}
	WordWeights *string `json:"WordWeights,omitempty" xml:"WordWeights,omitempty"`
}

func (s CreateTranscriptionPhrasesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTranscriptionPhrasesRequest) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTranscriptionPhrasesRequest) GetName() *string {
	return s.Name
}

func (s *CreateTranscriptionPhrasesRequest) GetWordWeights() *string {
	return s.WordWeights
}

func (s *CreateTranscriptionPhrasesRequest) SetDescription(v string) *CreateTranscriptionPhrasesRequest {
	s.Description = &v
	return s
}

func (s *CreateTranscriptionPhrasesRequest) SetName(v string) *CreateTranscriptionPhrasesRequest {
	s.Name = &v
	return s
}

func (s *CreateTranscriptionPhrasesRequest) SetWordWeights(v string) *CreateTranscriptionPhrasesRequest {
	s.WordWeights = &v
	return s
}

func (s *CreateTranscriptionPhrasesRequest) Validate() error {
	return dara.Validate(s)
}
