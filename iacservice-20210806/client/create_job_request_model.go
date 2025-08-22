// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateJobRequest
	GetClientToken() *string
	SetDescription(v string) *CreateJobRequest
	GetDescription() *string
	SetSubCommand(v string) *CreateJobRequest
	GetSubCommand() *string
	SetTaskType(v string) *CreateJobRequest
	GetTaskType() *string
}

type CreateJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daf4227f747cbf11a5501f18cc5e004
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// null
	SubCommand *string `json:"subCommand,omitempty" xml:"subCommand,omitempty"`
	TaskType   *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateJobRequest) GetSubCommand() *string {
	return s.SubCommand
}

func (s *CreateJobRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateJobRequest) SetClientToken(v string) *CreateJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
	return s
}

func (s *CreateJobRequest) SetSubCommand(v string) *CreateJobRequest {
	s.SubCommand = &v
	return s
}

func (s *CreateJobRequest) SetTaskType(v string) *CreateJobRequest {
	s.TaskType = &v
	return s
}

func (s *CreateJobRequest) Validate() error {
	return dara.Validate(s)
}
