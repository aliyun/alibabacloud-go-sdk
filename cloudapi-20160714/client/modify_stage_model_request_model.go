// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStageModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyStageModelRequest
	GetDescription() *string
	SetSecurityToken(v string) *ModifyStageModelRequest
	GetSecurityToken() *string
	SetStageAlias(v string) *ModifyStageModelRequest
	GetStageAlias() *string
	SetStageModelId(v string) *ModifyStageModelRequest
	GetStageModelId() *string
}

type ModifyStageModelRequest struct {
	// example:
	//
	// test
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageAlias    *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdljflsjflj324xxx
	StageModelId *string `json:"StageModelId,omitempty" xml:"StageModelId,omitempty"`
}

func (s ModifyStageModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStageModelRequest) GoString() string {
	return s.String()
}

func (s *ModifyStageModelRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyStageModelRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyStageModelRequest) GetStageAlias() *string {
	return s.StageAlias
}

func (s *ModifyStageModelRequest) GetStageModelId() *string {
	return s.StageModelId
}

func (s *ModifyStageModelRequest) SetDescription(v string) *ModifyStageModelRequest {
	s.Description = &v
	return s
}

func (s *ModifyStageModelRequest) SetSecurityToken(v string) *ModifyStageModelRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyStageModelRequest) SetStageAlias(v string) *ModifyStageModelRequest {
	s.StageAlias = &v
	return s
}

func (s *ModifyStageModelRequest) SetStageModelId(v string) *ModifyStageModelRequest {
	s.StageModelId = &v
	return s
}

func (s *ModifyStageModelRequest) Validate() error {
	return dara.Validate(s)
}
