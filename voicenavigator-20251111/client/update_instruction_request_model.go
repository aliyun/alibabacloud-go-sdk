// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstructionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstructionRequest
	GetCode() *string
	SetConfig(v string) *UpdateInstructionRequest
	GetConfig() *string
	SetInstanceId(v string) *UpdateInstructionRequest
	GetInstanceId() *string
	SetName(v string) *UpdateInstructionRequest
	GetName() *string
	SetType(v string) *UpdateInstructionRequest
	GetType() *string
}

type UpdateInstructionRequest struct {
	// The unique code for the instruction.
	//
	// example:
	//
	// Transfer00
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration for the instruction, specified as a JSON string.
	//
	// > - `providerId`: The ID of the live agent service provider. Call the `ListCallCenterProviders` operation to obtain this ID.
	//
	// >
	//
	// > - `transferMethod`: The method used for the transfer. Valid values: `BYE`, `BRIDGE`, and `REFER`.
	//
	// >
	//
	// > - `transferType`: The transfer destination type. Valid values: `External` (external number), `Agent` (agent), and `SkillGroup` (skill group).
	//
	// >
	//
	// > - `transferTarget`: The destination of the transfer, such as a phone number, an agent ID, or a skill group ID. The value of this parameter depends on the `transferType`.
	//
	// >
	//
	// > - `transferor`: The caller ID to display for the transfer. If this parameter is not specified, the caller\\"s original phone number is used.
	//
	// >
	//
	// > - `timeoutSeconds`: The timeout in seconds for transferring a call to an agent. Default: 30.
	//
	// example:
	//
	// {
	//
	// 	"providerId": "xxxxxxxxx",
	//
	// 	"transferMethod": "BYE",
	//
	// 	"transferType": "External",
	//
	// 	"transferTarget": "152******",
	//
	// 	"transferor": "01*****",
	//
	// 	"timeoutSeconds": "10"
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instruction.
	//
	// example:
	//
	// 转人工指令一
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the instruction.
	//
	// example:
	//
	// TRANSFER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateInstructionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstructionRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstructionRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateInstructionRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateInstructionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstructionRequest) GetName() *string {
	return s.Name
}

func (s *UpdateInstructionRequest) GetType() *string {
	return s.Type
}

func (s *UpdateInstructionRequest) SetCode(v string) *UpdateInstructionRequest {
	s.Code = &v
	return s
}

func (s *UpdateInstructionRequest) SetConfig(v string) *UpdateInstructionRequest {
	s.Config = &v
	return s
}

func (s *UpdateInstructionRequest) SetInstanceId(v string) *UpdateInstructionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstructionRequest) SetName(v string) *UpdateInstructionRequest {
	s.Name = &v
	return s
}

func (s *UpdateInstructionRequest) SetType(v string) *UpdateInstructionRequest {
	s.Type = &v
	return s
}

func (s *UpdateInstructionRequest) Validate() error {
	return dara.Validate(s)
}
