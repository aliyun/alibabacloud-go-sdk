// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishApplicationVersionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PublishApplicationVersionInput
	GetDescription() *string
}

type PublishApplicationVersionInput struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s PublishApplicationVersionInput) String() string {
	return dara.Prettify(s)
}

func (s PublishApplicationVersionInput) GoString() string {
	return s.String()
}

func (s *PublishApplicationVersionInput) GetDescription() *string {
	return s.Description
}

func (s *PublishApplicationVersionInput) SetDescription(v string) *PublishApplicationVersionInput {
	s.Description = &v
	return s
}

func (s *PublishApplicationVersionInput) Validate() error {
	return dara.Validate(s)
}
