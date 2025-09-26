// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateToolData) *CreateToolOutput
	GetData() *CreateToolData
	SetMessage(v string) *CreateToolOutput
	GetMessage() *string
	SetSuccess(v bool) *CreateToolOutput
	GetSuccess() *bool
}

type CreateToolOutput struct {
	Data    *CreateToolData `json:"data,omitempty" xml:"data,omitempty"`
	Message *string         `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateToolOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateToolOutput) GoString() string {
	return s.String()
}

func (s *CreateToolOutput) GetData() *CreateToolData {
	return s.Data
}

func (s *CreateToolOutput) GetMessage() *string {
	return s.Message
}

func (s *CreateToolOutput) GetSuccess() *bool {
	return s.Success
}

func (s *CreateToolOutput) SetData(v *CreateToolData) *CreateToolOutput {
	s.Data = v
	return s
}

func (s *CreateToolOutput) SetMessage(v string) *CreateToolOutput {
	s.Message = &v
	return s
}

func (s *CreateToolOutput) SetSuccess(v bool) *CreateToolOutput {
	s.Success = &v
	return s
}

func (s *CreateToolOutput) Validate() error {
	return dara.Validate(s)
}
