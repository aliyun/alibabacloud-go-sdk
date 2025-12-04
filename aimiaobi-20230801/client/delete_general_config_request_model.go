// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGeneralConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *DeleteGeneralConfigRequest
	GetConfigKey() *string
	SetWorkspaceId(v string) *DeleteGeneralConfigRequest
	GetWorkspaceId() *string
}

type DeleteGeneralConfigRequest struct {
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
	// llm-
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteGeneralConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGeneralConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteGeneralConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DeleteGeneralConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteGeneralConfigRequest) SetConfigKey(v string) *DeleteGeneralConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *DeleteGeneralConfigRequest) SetWorkspaceId(v string) *DeleteGeneralConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteGeneralConfigRequest) Validate() error {
	return dara.Validate(s)
}
