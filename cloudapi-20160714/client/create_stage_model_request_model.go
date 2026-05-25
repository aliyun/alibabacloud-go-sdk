// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStageModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateStageModelRequest
	GetDescription() *string
	SetSecurityToken(v string) *CreateStageModelRequest
	GetSecurityToken() *string
	SetStageAlias(v string) *CreateStageModelRequest
	GetStageAlias() *string
	SetStageName(v string) *CreateStageModelRequest
	GetStageName() *string
}

type CreateStageModelRequest struct {
	// example:
	//
	// Model Description
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	StageAlias *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s CreateStageModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStageModelRequest) GoString() string {
	return s.String()
}

func (s *CreateStageModelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStageModelRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateStageModelRequest) GetStageAlias() *string {
	return s.StageAlias
}

func (s *CreateStageModelRequest) GetStageName() *string {
	return s.StageName
}

func (s *CreateStageModelRequest) SetDescription(v string) *CreateStageModelRequest {
	s.Description = &v
	return s
}

func (s *CreateStageModelRequest) SetSecurityToken(v string) *CreateStageModelRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateStageModelRequest) SetStageAlias(v string) *CreateStageModelRequest {
	s.StageAlias = &v
	return s
}

func (s *CreateStageModelRequest) SetStageName(v string) *CreateStageModelRequest {
	s.StageName = &v
	return s
}

func (s *CreateStageModelRequest) Validate() error {
	return dara.Validate(s)
}
