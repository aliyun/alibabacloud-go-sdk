// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGeneralConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *GetGeneralConfigRequest
	GetConfigKey() *string
	SetWorkspaceId(v string) *GetGeneralConfigRequest
	GetWorkspaceId() *string
}

type GetGeneralConfigRequest struct {
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

func (s GetGeneralConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGeneralConfigRequest) GoString() string {
	return s.String()
}

func (s *GetGeneralConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *GetGeneralConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetGeneralConfigRequest) SetConfigKey(v string) *GetGeneralConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *GetGeneralConfigRequest) SetWorkspaceId(v string) *GetGeneralConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetGeneralConfigRequest) Validate() error {
	return dara.Validate(s)
}
