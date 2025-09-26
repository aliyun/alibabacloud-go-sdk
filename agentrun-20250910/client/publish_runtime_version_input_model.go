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
}

type PublishRuntimeVersionInput struct {
	// 此版本的描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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

func (s *PublishRuntimeVersionInput) SetDescription(v string) *PublishRuntimeVersionInput {
	s.Description = &v
	return s
}

func (s *PublishRuntimeVersionInput) Validate() error {
	return dara.Validate(s)
}
