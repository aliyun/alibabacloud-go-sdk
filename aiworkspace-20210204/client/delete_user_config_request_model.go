// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *DeleteUserConfigRequest
	GetConfigKey() *string
	SetScope(v string) *DeleteUserConfigRequest
	GetScope() *string
}

type DeleteUserConfigRequest struct {
	// The configuration item keys. Currently, only customizePAIAssumedRole.
	//
	// example:
	//
	// tempStoragePath
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The scope. Valid values: subUser and owner.
	//
	// example:
	//
	// subUser
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DeleteUserConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DeleteUserConfigRequest) GetScope() *string {
	return s.Scope
}

func (s *DeleteUserConfigRequest) SetConfigKey(v string) *DeleteUserConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *DeleteUserConfigRequest) SetScope(v string) *DeleteUserConfigRequest {
	s.Scope = &v
	return s
}

func (s *DeleteUserConfigRequest) Validate() error {
	return dara.Validate(s)
}
