// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowVersionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PublishFlowVersionInput
	GetDescription() *string
}

type PublishFlowVersionInput struct {
	// 工作流版本的描述信息
	//
	// example:
	//
	// Version 1.0 - Initial release
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s PublishFlowVersionInput) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowVersionInput) GoString() string {
	return s.String()
}

func (s *PublishFlowVersionInput) GetDescription() *string {
	return s.Description
}

func (s *PublishFlowVersionInput) SetDescription(v string) *PublishFlowVersionInput {
	s.Description = &v
	return s
}

func (s *PublishFlowVersionInput) Validate() error {
	return dara.Validate(s)
}
