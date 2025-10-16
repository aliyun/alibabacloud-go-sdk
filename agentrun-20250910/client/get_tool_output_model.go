// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetToolOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ToolInfo) *GetToolOutput
	GetData() *ToolInfo
	SetSuccess(v bool) *GetToolOutput
	GetSuccess() *bool
}

type GetToolOutput struct {
	Data    *ToolInfo `json:"data,omitempty" xml:"data,omitempty"`
	Success *bool     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetToolOutput) String() string {
	return dara.Prettify(s)
}

func (s GetToolOutput) GoString() string {
	return s.String()
}

func (s *GetToolOutput) GetData() *ToolInfo {
	return s.Data
}

func (s *GetToolOutput) GetSuccess() *bool {
	return s.Success
}

func (s *GetToolOutput) SetData(v *ToolInfo) *GetToolOutput {
	s.Data = v
	return s
}

func (s *GetToolOutput) SetSuccess(v bool) *GetToolOutput {
	s.Success = &v
	return s
}

func (s *GetToolOutput) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
