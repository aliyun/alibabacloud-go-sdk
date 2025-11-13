// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractRuleGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunContractRuleGenerationRequest
	GetAppId() *string
	SetAssistant(v *RunContractRuleGenerationRequestAssistant) *RunContractRuleGenerationRequest
	GetAssistant() *RunContractRuleGenerationRequestAssistant
	SetStream(v bool) *RunContractRuleGenerationRequest
	GetStream() *bool
}

type RunContractRuleGenerationRequest struct {
	// example:
	//
	// farui
	AppId     *string                                    `json:"appId,omitempty" xml:"appId,omitempty"`
	Assistant *RunContractRuleGenerationRequestAssistant `json:"assistant,omitempty" xml:"assistant,omitempty" type:"Struct"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s RunContractRuleGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunContractRuleGenerationRequest) GetAssistant() *RunContractRuleGenerationRequestAssistant {
	return s.Assistant
}

func (s *RunContractRuleGenerationRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunContractRuleGenerationRequest) SetAppId(v string) *RunContractRuleGenerationRequest {
	s.AppId = &v
	return s
}

func (s *RunContractRuleGenerationRequest) SetAssistant(v *RunContractRuleGenerationRequestAssistant) *RunContractRuleGenerationRequest {
	s.Assistant = v
	return s
}

func (s *RunContractRuleGenerationRequest) SetStream(v bool) *RunContractRuleGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunContractRuleGenerationRequest) Validate() error {
	if s.Assistant != nil {
		if err := s.Assistant.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunContractRuleGenerationRequestAssistant struct {
	MetaData *RunContractRuleGenerationRequestAssistantMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// example:
	//
	// contract_examime
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RunContractRuleGenerationRequestAssistant) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationRequestAssistant) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationRequestAssistant) GetMetaData() *RunContractRuleGenerationRequestAssistantMetaData {
	return s.MetaData
}

func (s *RunContractRuleGenerationRequestAssistant) GetType() *string {
	return s.Type
}

func (s *RunContractRuleGenerationRequestAssistant) GetVersion() *string {
	return s.Version
}

func (s *RunContractRuleGenerationRequestAssistant) SetMetaData(v *RunContractRuleGenerationRequestAssistantMetaData) *RunContractRuleGenerationRequestAssistant {
	s.MetaData = v
	return s
}

func (s *RunContractRuleGenerationRequestAssistant) SetType(v string) *RunContractRuleGenerationRequestAssistant {
	s.Type = &v
	return s
}

func (s *RunContractRuleGenerationRequestAssistant) SetVersion(v string) *RunContractRuleGenerationRequestAssistant {
	s.Version = &v
	return s
}

func (s *RunContractRuleGenerationRequestAssistant) Validate() error {
	if s.MetaData != nil {
		if err := s.MetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunContractRuleGenerationRequestAssistantMetaData struct {
	// example:
	//
	// 9a6b1ba60d9944249363ec3cc1529b7b
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 1
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
}

func (s RunContractRuleGenerationRequestAssistantMetaData) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationRequestAssistantMetaData) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationRequestAssistantMetaData) GetFileId() *string {
	return s.FileId
}

func (s *RunContractRuleGenerationRequestAssistantMetaData) GetPosition() *string {
	return s.Position
}

func (s *RunContractRuleGenerationRequestAssistantMetaData) SetFileId(v string) *RunContractRuleGenerationRequestAssistantMetaData {
	s.FileId = &v
	return s
}

func (s *RunContractRuleGenerationRequestAssistantMetaData) SetPosition(v string) *RunContractRuleGenerationRequestAssistantMetaData {
	s.Position = &v
	return s
}

func (s *RunContractRuleGenerationRequestAssistantMetaData) Validate() error {
	return dara.Validate(s)
}
