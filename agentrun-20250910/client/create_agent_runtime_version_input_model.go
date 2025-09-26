// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRuntimeVersionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAgentRuntimeVersionInput
	GetDescription() *string
}

type CreateAgentRuntimeVersionInput struct {
	// 版本描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateAgentRuntimeVersionInput) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRuntimeVersionInput) GoString() string {
	return s.String()
}

func (s *CreateAgentRuntimeVersionInput) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRuntimeVersionInput) SetDescription(v string) *CreateAgentRuntimeVersionInput {
	s.Description = &v
	return s
}

func (s *CreateAgentRuntimeVersionInput) Validate() error {
	return dara.Validate(s)
}
