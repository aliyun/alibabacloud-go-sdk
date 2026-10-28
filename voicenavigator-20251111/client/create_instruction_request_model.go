// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstructionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstructionRequest
	GetCode() *string
	SetConfig(v string) *CreateInstructionRequest
	GetConfig() *string
	SetInstanceId(v string) *CreateInstructionRequest
	GetInstanceId() *string
	SetName(v string) *CreateInstructionRequest
	GetName() *string
	SetType(v string) *CreateInstructionRequest
	GetType() *string
}

type CreateInstructionRequest struct {
	// The instruction code.
	//
	// 	Notice: Globally unique.</notice>
	//
	// example:
	//
	// Transfer00
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instruction configuration.
	//
	// > - providerId: the ID of the human agent service provider. You can obtain the ID by calling ListCallCenterProviders.
	//
	// > - transferMethod: the transfer method. Valid values: BYE, BRIDGE, and REFER.
	//
	// > - transferType: the transfer target type. Valid values: External (external number), Agent (agent), and SkillGroup (skill group).
	//
	// > - transferTarget: the transfer target.
	//
	// > - transferor: the caller number. If this parameter is left empty, the mobile phone number of the calling customer is used by default.
	//
	// > - timeoutSeconds: the timeout period for transferring to a human agent, in seconds. Default value: 30.
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
	// The instance ID.
	//
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instruction name.
	//
	// example:
	//
	// Transfer to human agent instruction 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instruction type.
	//
	// example:
	//
	// TRANSFER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateInstructionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstructionRequest) GoString() string {
	return s.String()
}

func (s *CreateInstructionRequest) GetCode() *string {
	return s.Code
}

func (s *CreateInstructionRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateInstructionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstructionRequest) GetName() *string {
	return s.Name
}

func (s *CreateInstructionRequest) GetType() *string {
	return s.Type
}

func (s *CreateInstructionRequest) SetCode(v string) *CreateInstructionRequest {
	s.Code = &v
	return s
}

func (s *CreateInstructionRequest) SetConfig(v string) *CreateInstructionRequest {
	s.Config = &v
	return s
}

func (s *CreateInstructionRequest) SetInstanceId(v string) *CreateInstructionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstructionRequest) SetName(v string) *CreateInstructionRequest {
	s.Name = &v
	return s
}

func (s *CreateInstructionRequest) SetType(v string) *CreateInstructionRequest {
	s.Type = &v
	return s
}

func (s *CreateInstructionRequest) Validate() error {
	return dara.Validate(s)
}
