// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateToolData) *UpdateToolOutput
	GetData() *UpdateToolData
	SetMessage(v string) *UpdateToolOutput
	GetMessage() *string
	SetSuccess(v bool) *UpdateToolOutput
	GetSuccess() *bool
}

type UpdateToolOutput struct {
	Data    *UpdateToolData `json:"data,omitempty" xml:"data,omitempty"`
	Message *string         `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateToolOutput) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolOutput) GoString() string {
	return s.String()
}

func (s *UpdateToolOutput) GetData() *UpdateToolData {
	return s.Data
}

func (s *UpdateToolOutput) GetMessage() *string {
	return s.Message
}

func (s *UpdateToolOutput) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateToolOutput) SetData(v *UpdateToolData) *UpdateToolOutput {
	s.Data = v
	return s
}

func (s *UpdateToolOutput) SetMessage(v string) *UpdateToolOutput {
	s.Message = &v
	return s
}

func (s *UpdateToolOutput) SetSuccess(v bool) *UpdateToolOutput {
	s.Success = &v
	return s
}

func (s *UpdateToolOutput) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
