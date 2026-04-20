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
