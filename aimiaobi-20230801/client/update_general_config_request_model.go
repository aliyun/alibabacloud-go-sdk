// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGeneralConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *UpdateGeneralConfigRequest
	GetConfigKey() *string
	SetConfigValue(v string) *UpdateGeneralConfigRequest
	GetConfigValue() *string
	SetWorkspaceId(v string) *UpdateGeneralConfigRequest
	GetWorkspaceId() *string
}

type UpdateGeneralConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xx
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateGeneralConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGeneralConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneralConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *UpdateGeneralConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *UpdateGeneralConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateGeneralConfigRequest) SetConfigKey(v string) *UpdateGeneralConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *UpdateGeneralConfigRequest) SetConfigValue(v string) *UpdateGeneralConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *UpdateGeneralConfigRequest) SetWorkspaceId(v string) *UpdateGeneralConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateGeneralConfigRequest) Validate() error {
	return dara.Validate(s)
}
