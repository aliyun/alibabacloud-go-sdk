// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiStageVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteApiStageVariableRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DeleteApiStageVariableRequest
	GetSecurityToken() *string
	SetStageId(v string) *DeleteApiStageVariableRequest
	GetStageId() *string
	SetVariableName(v string) *DeleteApiStageVariableRequest
	GetVariableName() *string
}

type DeleteApiStageVariableRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 523e8dc7bbe04613b5b1d726c2a7889d
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The name of the variable to be deleted. This parameter is case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// serverName
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
}

func (s DeleteApiStageVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiStageVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiStageVariableRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteApiStageVariableRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteApiStageVariableRequest) GetStageId() *string {
	return s.StageId
}

func (s *DeleteApiStageVariableRequest) GetVariableName() *string {
	return s.VariableName
}

func (s *DeleteApiStageVariableRequest) SetGroupId(v string) *DeleteApiStageVariableRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteApiStageVariableRequest) SetSecurityToken(v string) *DeleteApiStageVariableRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteApiStageVariableRequest) SetStageId(v string) *DeleteApiStageVariableRequest {
	s.StageId = &v
	return s
}

func (s *DeleteApiStageVariableRequest) SetVariableName(v string) *DeleteApiStageVariableRequest {
	s.VariableName = &v
	return s
}

func (s *DeleteApiStageVariableRequest) Validate() error {
	return dara.Validate(s)
}
