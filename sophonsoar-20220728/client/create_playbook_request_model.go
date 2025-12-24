// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePlaybookRequest
	GetDescription() *string
	SetDisplayName(v string) *CreatePlaybookRequest
	GetDisplayName() *string
	SetInputParams(v string) *CreatePlaybookRequest
	GetInputParams() *string
	SetLang(v string) *CreatePlaybookRequest
	GetLang() *string
	SetOutputParams(v string) *CreatePlaybookRequest
	GetOutputParams() *string
	SetTaskflowType(v string) *CreatePlaybookRequest
	GetTaskflowType() *string
}

type CreatePlaybookRequest struct {
	// Description of the playbook.
	//
	// example:
	//
	// This is a new version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Name of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// test09
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// {\"key1\": \"value1\", \"key2\": \"value2\"}
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// Language type for receiving messages. Values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// {\"result\": \"success\"}
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// Playbook TaskFlow type.
	//
	// - **x6*	- : x6
	//
	// - **bpmn**: bpmn
	//
	// example:
	//
	// x6
	TaskflowType *string `json:"TaskflowType,omitempty" xml:"TaskflowType,omitempty"`
}

func (s CreatePlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlaybookRequest) GoString() string {
	return s.String()
}

func (s *CreatePlaybookRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePlaybookRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreatePlaybookRequest) GetInputParams() *string {
	return s.InputParams
}

func (s *CreatePlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *CreatePlaybookRequest) GetOutputParams() *string {
	return s.OutputParams
}

func (s *CreatePlaybookRequest) GetTaskflowType() *string {
	return s.TaskflowType
}

func (s *CreatePlaybookRequest) SetDescription(v string) *CreatePlaybookRequest {
	s.Description = &v
	return s
}

func (s *CreatePlaybookRequest) SetDisplayName(v string) *CreatePlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *CreatePlaybookRequest) SetInputParams(v string) *CreatePlaybookRequest {
	s.InputParams = &v
	return s
}

func (s *CreatePlaybookRequest) SetLang(v string) *CreatePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *CreatePlaybookRequest) SetOutputParams(v string) *CreatePlaybookRequest {
	s.OutputParams = &v
	return s
}

func (s *CreatePlaybookRequest) SetTaskflowType(v string) *CreatePlaybookRequest {
	s.TaskflowType = &v
	return s
}

func (s *CreatePlaybookRequest) Validate() error {
	return dara.Validate(s)
}
