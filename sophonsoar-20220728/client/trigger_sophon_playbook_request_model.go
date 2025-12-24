// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerSophonPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandName(v string) *TriggerSophonPlaybookRequest
	GetCommandName() *string
	SetInputParams(v string) *TriggerSophonPlaybookRequest
	GetInputParams() *string
	SetSophonTaskId(v string) *TriggerSophonPlaybookRequest
	GetSophonTaskId() *string
	SetTriggerType(v string) *TriggerSophonPlaybookRequest
	GetTriggerType() *string
	SetUuid(v string) *TriggerSophonPlaybookRequest
	GetUuid() *string
}

type TriggerSophonPlaybookRequest struct {
	// The name of the command that you want to trigger.
	//
	// >  You can call the [DescribeSophonCommands](~~DescribeSophonCommands~~) operation to query the command name.
	//
	// example:
	//
	// waf_process_command
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The input parameters of the command or playbook that you want to trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "param1": "xx.xx.xx.xx",
	//
	//     "param2": "7d"
	//
	// }
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The custom ID. If you do not specify this parameter when the playbook is triggered, a random ID is generated for fault locating and troubleshooting.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxxxxxxx
	SophonTaskId *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	// The task type. Valid values:
	//
	// 	- **command**
	//
	// 	- **playbook**
	//
	// example:
	//
	// playbook
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s TriggerSophonPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerSophonPlaybookRequest) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookRequest) GetCommandName() *string {
	return s.CommandName
}

func (s *TriggerSophonPlaybookRequest) GetInputParams() *string {
	return s.InputParams
}

func (s *TriggerSophonPlaybookRequest) GetSophonTaskId() *string {
	return s.SophonTaskId
}

func (s *TriggerSophonPlaybookRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *TriggerSophonPlaybookRequest) GetUuid() *string {
	return s.Uuid
}

func (s *TriggerSophonPlaybookRequest) SetCommandName(v string) *TriggerSophonPlaybookRequest {
	s.CommandName = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetInputParams(v string) *TriggerSophonPlaybookRequest {
	s.InputParams = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetSophonTaskId(v string) *TriggerSophonPlaybookRequest {
	s.SophonTaskId = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetTriggerType(v string) *TriggerSophonPlaybookRequest {
	s.TriggerType = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) SetUuid(v string) *TriggerSophonPlaybookRequest {
	s.Uuid = &v
	return s
}

func (s *TriggerSophonPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
