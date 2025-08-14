// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAIAgentVoiceprintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SetAIAgentVoiceprintRequestInput) *SetAIAgentVoiceprintRequest
	GetInput() *SetAIAgentVoiceprintRequestInput
	SetVoiceprintId(v string) *SetAIAgentVoiceprintRequest
	GetVoiceprintId() *string
}

type SetAIAgentVoiceprintRequest struct {
	Input *SetAIAgentVoiceprintRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// example:
	//
	// vp_1699123456_8527
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s SetAIAgentVoiceprintRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAIAgentVoiceprintRequest) GoString() string {
	return s.String()
}

func (s *SetAIAgentVoiceprintRequest) GetInput() *SetAIAgentVoiceprintRequestInput {
	return s.Input
}

func (s *SetAIAgentVoiceprintRequest) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *SetAIAgentVoiceprintRequest) SetInput(v *SetAIAgentVoiceprintRequestInput) *SetAIAgentVoiceprintRequest {
	s.Input = v
	return s
}

func (s *SetAIAgentVoiceprintRequest) SetVoiceprintId(v string) *SetAIAgentVoiceprintRequest {
	s.VoiceprintId = &v
	return s
}

func (s *SetAIAgentVoiceprintRequest) Validate() error {
	return dara.Validate(s)
}

type SetAIAgentVoiceprintRequestInput struct {
	// example:
	//
	// https://my-bucket.oss-cn-hangzhou.aliyuncs.com/audio/sample.wav
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// wav
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// example:
	//
	// url
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SetAIAgentVoiceprintRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SetAIAgentVoiceprintRequestInput) GoString() string {
	return s.String()
}

func (s *SetAIAgentVoiceprintRequestInput) GetData() *string {
	return s.Data
}

func (s *SetAIAgentVoiceprintRequestInput) GetFormat() *string {
	return s.Format
}

func (s *SetAIAgentVoiceprintRequestInput) GetType() *string {
	return s.Type
}

func (s *SetAIAgentVoiceprintRequestInput) SetData(v string) *SetAIAgentVoiceprintRequestInput {
	s.Data = &v
	return s
}

func (s *SetAIAgentVoiceprintRequestInput) SetFormat(v string) *SetAIAgentVoiceprintRequestInput {
	s.Format = &v
	return s
}

func (s *SetAIAgentVoiceprintRequestInput) SetType(v string) *SetAIAgentVoiceprintRequestInput {
	s.Type = &v
	return s
}

func (s *SetAIAgentVoiceprintRequestInput) Validate() error {
	return dara.Validate(s)
}
