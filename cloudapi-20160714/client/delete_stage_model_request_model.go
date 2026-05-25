// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStageModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DeleteStageModelRequest
	GetSecurityToken() *string
	SetStageModelId(v string) *DeleteStageModelRequest
	GetStageModelId() *string
}

type DeleteStageModelRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3294sdufo3nenexxx
	StageModelId *string `json:"StageModelId,omitempty" xml:"StageModelId,omitempty"`
}

func (s DeleteStageModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStageModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteStageModelRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteStageModelRequest) GetStageModelId() *string {
	return s.StageModelId
}

func (s *DeleteStageModelRequest) SetSecurityToken(v string) *DeleteStageModelRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteStageModelRequest) SetStageModelId(v string) *DeleteStageModelRequest {
	s.StageModelId = &v
	return s
}

func (s *DeleteStageModelRequest) Validate() error {
	return dara.Validate(s)
}
