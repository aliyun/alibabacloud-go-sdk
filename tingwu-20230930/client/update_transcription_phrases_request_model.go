// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTranscriptionPhrasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTranscriptionPhrasesRequest
	GetDescription() *string
	SetName(v string) *UpdateTranscriptionPhrasesRequest
	GetName() *string
	SetWordWeights(v string) *UpdateTranscriptionPhrasesRequest
	GetWordWeights() *string
}

type UpdateTranscriptionPhrasesRequest struct {
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

func (s UpdateTranscriptionPhrasesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTranscriptionPhrasesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTranscriptionPhrasesRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTranscriptionPhrasesRequest) GetWordWeights() *string {
	return s.WordWeights
}

func (s *UpdateTranscriptionPhrasesRequest) SetDescription(v string) *UpdateTranscriptionPhrasesRequest {
	s.Description = &v
	return s
}

func (s *UpdateTranscriptionPhrasesRequest) SetName(v string) *UpdateTranscriptionPhrasesRequest {
	s.Name = &v
	return s
}

func (s *UpdateTranscriptionPhrasesRequest) SetWordWeights(v string) *UpdateTranscriptionPhrasesRequest {
	s.WordWeights = &v
	return s
}

func (s *UpdateTranscriptionPhrasesRequest) Validate() error {
	return dara.Validate(s)
}
