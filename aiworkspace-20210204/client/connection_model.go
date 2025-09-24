// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnection interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Connection
	GetAccessibility() *string
	SetConfigs(v map[string]*string) *Connection
	GetConfigs() map[string]*string
	SetConnectionId(v string) *Connection
	GetConnectionId() *string
	SetConnectionName(v string) *Connection
	GetConnectionName() *string
	SetConnectionType(v string) *Connection
	GetConnectionType() *string
	SetCreator(v string) *Connection
	GetCreator() *string
	SetDescription(v string) *Connection
	GetDescription() *string
	SetGmtCreateTime(v string) *Connection
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Connection
	GetGmtModifiedTime() *string
	SetModels(v []*ConnectionModels) *Connection
	GetModels() []*ConnectionModels
	SetResourceMeta(v *ConnectionResourceMeta) *Connection
	GetResourceMeta() *ConnectionResourceMeta
	SetSecrets(v map[string]*string) *Connection
	GetSecrets() map[string]*string
	SetWorkspaceId(v string) *Connection
	GetWorkspaceId() *string
}

type Connection struct {
	Accessibility   *string                 `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Configs         map[string]*string      `json:"Configs,omitempty" xml:"Configs,omitempty"`
	ConnectionId    *string                 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	ConnectionName  *string                 `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	ConnectionType  *string                 `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	Creator         *string                 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string                 `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                 `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                 `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Models          []*ConnectionModels     `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	ResourceMeta    *ConnectionResourceMeta `json:"ResourceMeta,omitempty" xml:"ResourceMeta,omitempty" type:"Struct"`
	Secrets         map[string]*string      `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
	WorkspaceId     *string                 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Connection) String() string {
	return dara.Prettify(s)
}

func (s Connection) GoString() string {
	return s.String()
}

func (s *Connection) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Connection) GetConfigs() map[string]*string {
	return s.Configs
}

func (s *Connection) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *Connection) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *Connection) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *Connection) GetCreator() *string {
	return s.Creator
}

func (s *Connection) GetDescription() *string {
	return s.Description
}

func (s *Connection) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Connection) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Connection) GetModels() []*ConnectionModels {
	return s.Models
}

func (s *Connection) GetResourceMeta() *ConnectionResourceMeta {
	return s.ResourceMeta
}

func (s *Connection) GetSecrets() map[string]*string {
	return s.Secrets
}

func (s *Connection) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Connection) SetAccessibility(v string) *Connection {
	s.Accessibility = &v
	return s
}

func (s *Connection) SetConfigs(v map[string]*string) *Connection {
	s.Configs = v
	return s
}

func (s *Connection) SetConnectionId(v string) *Connection {
	s.ConnectionId = &v
	return s
}

func (s *Connection) SetConnectionName(v string) *Connection {
	s.ConnectionName = &v
	return s
}

func (s *Connection) SetConnectionType(v string) *Connection {
	s.ConnectionType = &v
	return s
}

func (s *Connection) SetCreator(v string) *Connection {
	s.Creator = &v
	return s
}

func (s *Connection) SetDescription(v string) *Connection {
	s.Description = &v
	return s
}

func (s *Connection) SetGmtCreateTime(v string) *Connection {
	s.GmtCreateTime = &v
	return s
}

func (s *Connection) SetGmtModifiedTime(v string) *Connection {
	s.GmtModifiedTime = &v
	return s
}

func (s *Connection) SetModels(v []*ConnectionModels) *Connection {
	s.Models = v
	return s
}

func (s *Connection) SetResourceMeta(v *ConnectionResourceMeta) *Connection {
	s.ResourceMeta = v
	return s
}

func (s *Connection) SetSecrets(v map[string]*string) *Connection {
	s.Secrets = v
	return s
}

func (s *Connection) SetWorkspaceId(v string) *Connection {
	s.WorkspaceId = &v
	return s
}

func (s *Connection) Validate() error {
	return dara.Validate(s)
}

type ConnectionModels struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ModelType   *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ToolCall    *bool   `json:"ToolCall,omitempty" xml:"ToolCall,omitempty"`
}

func (s ConnectionModels) String() string {
	return dara.Prettify(s)
}

func (s ConnectionModels) GoString() string {
	return s.String()
}

func (s *ConnectionModels) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ConnectionModels) GetModel() *string {
	return s.Model
}

func (s *ConnectionModels) GetModelType() *string {
	return s.ModelType
}

func (s *ConnectionModels) GetToolCall() *bool {
	return s.ToolCall
}

func (s *ConnectionModels) SetDisplayName(v string) *ConnectionModels {
	s.DisplayName = &v
	return s
}

func (s *ConnectionModels) SetModel(v string) *ConnectionModels {
	s.Model = &v
	return s
}

func (s *ConnectionModels) SetModelType(v string) *ConnectionModels {
	s.ModelType = &v
	return s
}

func (s *ConnectionModels) SetToolCall(v bool) *ConnectionModels {
	s.ToolCall = &v
	return s
}

func (s *ConnectionModels) Validate() error {
	return dara.Validate(s)
}

type ConnectionResourceMeta struct {
	Extra        *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ConnectionResourceMeta) String() string {
	return dara.Prettify(s)
}

func (s ConnectionResourceMeta) GoString() string {
	return s.String()
}

func (s *ConnectionResourceMeta) GetExtra() *string {
	return s.Extra
}

func (s *ConnectionResourceMeta) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConnectionResourceMeta) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ConnectionResourceMeta) SetExtra(v string) *ConnectionResourceMeta {
	s.Extra = &v
	return s
}

func (s *ConnectionResourceMeta) SetInstanceId(v string) *ConnectionResourceMeta {
	s.InstanceId = &v
	return s
}

func (s *ConnectionResourceMeta) SetInstanceName(v string) *ConnectionResourceMeta {
	s.InstanceName = &v
	return s
}

func (s *ConnectionResourceMeta) Validate() error {
	return dara.Validate(s)
}
