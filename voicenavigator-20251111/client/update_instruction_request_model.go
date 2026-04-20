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
	// example:
	//
	// Transfer00
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
