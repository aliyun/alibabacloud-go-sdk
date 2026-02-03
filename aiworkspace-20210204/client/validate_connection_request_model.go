// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v map[string]*string) *ValidateConnectionRequest
	GetConfigs() map[string]*string
	SetConnectionId(v string) *ValidateConnectionRequest
	GetConnectionId() *string
	SetConnectionType(v string) *ValidateConnectionRequest
	GetConnectionType() *string
	SetSecrets(v map[string]*string) *ValidateConnectionRequest
	GetSecrets() map[string]*string
	SetValidateType(v string) *ValidateConnectionRequest
	GetValidateType() *string
	SetWorkspaceId(v string) *ValidateConnectionRequest
	GetWorkspaceId() *string
}

type ValidateConnectionRequest struct {
	Configs map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// example:
	//
	// conn-x2wz5zvvbyf2420fi9
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// example:
	//
	// DashScopeConnection
	ConnectionType *string            `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	Secrets        map[string]*string `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
	// example:
	//
	// Connectivity
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ValidateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateConnectionRequest) GoString() string {
	return s.String()
}

func (s *ValidateConnectionRequest) GetConfigs() map[string]*string {
	return s.Configs
}

func (s *ValidateConnectionRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *ValidateConnectionRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *ValidateConnectionRequest) GetSecrets() map[string]*string {
	return s.Secrets
}

func (s *ValidateConnectionRequest) GetValidateType() *string {
	return s.ValidateType
}

func (s *ValidateConnectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ValidateConnectionRequest) SetConfigs(v map[string]*string) *ValidateConnectionRequest {
	s.Configs = v
	return s
}

func (s *ValidateConnectionRequest) SetConnectionId(v string) *ValidateConnectionRequest {
	s.ConnectionId = &v
	return s
}

func (s *ValidateConnectionRequest) SetConnectionType(v string) *ValidateConnectionRequest {
	s.ConnectionType = &v
	return s
}

func (s *ValidateConnectionRequest) SetSecrets(v map[string]*string) *ValidateConnectionRequest {
	s.Secrets = v
	return s
}

func (s *ValidateConnectionRequest) SetValidateType(v string) *ValidateConnectionRequest {
	s.ValidateType = &v
	return s
}

func (s *ValidateConnectionRequest) SetWorkspaceId(v string) *ValidateConnectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ValidateConnectionRequest) Validate() error {
	return dara.Validate(s)
}
