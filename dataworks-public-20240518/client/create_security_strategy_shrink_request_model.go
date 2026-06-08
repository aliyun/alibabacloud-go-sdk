// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSecurityStrategyShrinkRequest
	GetClientToken() *string
	SetContentShrink(v string) *CreateSecurityStrategyShrinkRequest
	GetContentShrink() *string
	SetControlDwScope(v string) *CreateSecurityStrategyShrinkRequest
	GetControlDwScope() *string
	SetControlModule(v string) *CreateSecurityStrategyShrinkRequest
	GetControlModule() *string
	SetControlSubModule(v string) *CreateSecurityStrategyShrinkRequest
	GetControlSubModule() *string
	SetDescription(v string) *CreateSecurityStrategyShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateSecurityStrategyShrinkRequest
	GetName() *string
	SetSchemaName(v string) *CreateSecurityStrategyShrinkRequest
	GetSchemaName() *string
	SetWorkspacesShrink(v string) *CreateSecurityStrategyShrinkRequest
	GetWorkspacesShrink() *string
}

type CreateSecurityStrategyShrinkRequest struct {
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Tenant
	ControlDwScope *string `json:"ControlDwScope,omitempty" xml:"ControlDwScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DataQuerySecurityStrategySchema
	SchemaName       *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	WorkspacesShrink *string `json:"Workspaces,omitempty" xml:"Workspaces,omitempty"`
}

func (s CreateSecurityStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityStrategyShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSecurityStrategyShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *CreateSecurityStrategyShrinkRequest) GetControlDwScope() *string {
	return s.ControlDwScope
}

func (s *CreateSecurityStrategyShrinkRequest) GetControlModule() *string {
	return s.ControlModule
}

func (s *CreateSecurityStrategyShrinkRequest) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *CreateSecurityStrategyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityStrategyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateSecurityStrategyShrinkRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *CreateSecurityStrategyShrinkRequest) GetWorkspacesShrink() *string {
	return s.WorkspacesShrink
}

func (s *CreateSecurityStrategyShrinkRequest) SetClientToken(v string) *CreateSecurityStrategyShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetContentShrink(v string) *CreateSecurityStrategyShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetControlDwScope(v string) *CreateSecurityStrategyShrinkRequest {
	s.ControlDwScope = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetControlModule(v string) *CreateSecurityStrategyShrinkRequest {
	s.ControlModule = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetControlSubModule(v string) *CreateSecurityStrategyShrinkRequest {
	s.ControlSubModule = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetDescription(v string) *CreateSecurityStrategyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetName(v string) *CreateSecurityStrategyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetSchemaName(v string) *CreateSecurityStrategyShrinkRequest {
	s.SchemaName = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) SetWorkspacesShrink(v string) *CreateSecurityStrategyShrinkRequest {
	s.WorkspacesShrink = &v
	return s
}

func (s *CreateSecurityStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
