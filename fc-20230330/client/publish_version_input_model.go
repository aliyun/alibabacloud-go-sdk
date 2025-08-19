// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVersionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PublishVersionInput
	GetDescription() *string
}

type PublishVersionInput struct {
	// example:
	//
	// my version
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s PublishVersionInput) String() string {
	return dara.Prettify(s)
}

func (s PublishVersionInput) GoString() string {
	return s.String()
}

func (s *PublishVersionInput) GetDescription() *string {
	return s.Description
}

func (s *PublishVersionInput) SetDescription(v string) *PublishVersionInput {
	s.Description = &v
	return s
}

func (s *PublishVersionInput) Validate() error {
	return dara.Validate(s)
}
