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
	// The configuration information of the connection. The information is configured in key-value pairs. The keys in the configuration vary based on the connection type. For more information, see the supplementary description of request parameters for the CreateConnection operation.
	Configs map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The connection ID. For more information about how to obtain the connection ID, see [ListConnections](~~ListConnections~~).
	//
	// example:
	//
	// conn-x2wz5zvvbyf2420fi9
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The type of the connection. You can specify a connection type that supports public network access. Valid values:
	//
	// - DashScopeConnection: The Model Studio connection.
	//
	// - DeepSeekConnection: The DeepSeek connection.
	//
	// - FunctionAIMCPConnection: The Function AI MCP connection.
	//
	// - CustomMCPConnection: The custom MCP connection.
	//
	// example:
	//
	// DashScopeConnection
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The key-value configurations that require encryption, such as the database logon password and the key for the model connection.
	Secrets map[string]*string `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
	// The validation type. Valid values:
	//
	// - Connectivity: The connectivity test.
	//
	// example:
	//
	// Connectivity
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
	// The workspace ID. You can obtain the workspace ID by calling the [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) operation.
	//
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
