// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRuntimeVersionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PublishRuntimeVersionInput
	GetDescription() *string
	SetPublisher(v string) *PublishRuntimeVersionInput
	GetPublisher() *string
}

type PublishRuntimeVersionInput struct {
	// Description of this version
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Identity of the user or system that publishes this version
	//
	// example:
	//
	// user123
	Publisher *string `json:"publisher,omitempty" xml:"publisher,omitempty"`
}

func (s PublishRuntimeVersionInput) String() string {
	return dara.Prettify(s)
}

func (s PublishRuntimeVersionInput) GoString() string {
	return s.String()
}

func (s *PublishRuntimeVersionInput) GetDescription() *string {
	return s.Description
}

func (s *PublishRuntimeVersionInput) GetPublisher() *string {
	return s.Publisher
}

func (s *PublishRuntimeVersionInput) SetDescription(v string) *PublishRuntimeVersionInput {
	s.Description = &v
	return s
}

func (s *PublishRuntimeVersionInput) SetPublisher(v string) *PublishRuntimeVersionInput {
	s.Publisher = &v
	return s
}

func (s *PublishRuntimeVersionInput) Validate() error {
	return dara.Validate(s)
}
