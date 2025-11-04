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
	// The input file.
	Input *SetAIAgentVoiceprintRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// A unique identifier for the voiceprint. Generate this ID based on your own business rules. Requirement: 1 to 127 characters in length.
	//
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
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetAIAgentVoiceprintRequestInput struct {
	// The media access link.
	//
	// example:
	//
	// https://my-bucket.oss-cn-hangzhou.aliyuncs.com/audio/sample.wav
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The audio file format. Only WAV is supported.
	//
	// example:
	//
	// wav
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Specifies the access type for the audio file. The system will verify file accessibility via HEAD or GET requests. Valid values:
	//
	// 	- url: An HTTP(S) link to the audio file.
	//
	// 	- oss: An OSS object. Supports the following formats:
	//
	//     1.  OSS URI: oss://bucket-name/object-key
	//
	//         Example: oss://my-bucket/audio/sample.wav
	//
	//     2.  OSS public URL: http(s)://${bucket}.oss-${region}.aliyuncs.com/${object}
	//
	//         Example: https://my-bucket.oss-cn-hangzhou.aliyuncs.com/audio/sample.wav
	//
	// >  The OSS bucket must be in the same region as the service. Otherwise, the access fails.
	//
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
