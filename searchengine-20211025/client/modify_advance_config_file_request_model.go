// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvanceConfigFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ModifyAdvanceConfigFileRequest
	GetContent() *string
	SetVariables(v map[string]*VariablesValue) *ModifyAdvanceConfigFileRequest
	GetVariables() map[string]*VariablesValue
	SetFileName(v string) *ModifyAdvanceConfigFileRequest
	GetFileName() *string
}

type ModifyAdvanceConfigFileRequest struct {
	// The file content.
	//
	// example:
	//
	// "ha3"
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The variables.
	Variables map[string]*VariablesValue `json:"variables,omitempty" xml:"variables,omitempty"`
	// The name of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// /qrs.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ModifyAdvanceConfigFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvanceConfigFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileRequest) GetContent() *string {
	return s.Content
}

func (s *ModifyAdvanceConfigFileRequest) GetVariables() map[string]*VariablesValue {
	return s.Variables
}

func (s *ModifyAdvanceConfigFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *ModifyAdvanceConfigFileRequest) SetContent(v string) *ModifyAdvanceConfigFileRequest {
	s.Content = &v
	return s
}

func (s *ModifyAdvanceConfigFileRequest) SetVariables(v map[string]*VariablesValue) *ModifyAdvanceConfigFileRequest {
	s.Variables = v
	return s
}

func (s *ModifyAdvanceConfigFileRequest) SetFileName(v string) *ModifyAdvanceConfigFileRequest {
	s.FileName = &v
	return s
}

func (s *ModifyAdvanceConfigFileRequest) Validate() error {
	return dara.Validate(s)
}
