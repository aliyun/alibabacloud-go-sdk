// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindBestMatchSecurityStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlModule(v string) *FindBestMatchSecurityStrategyRequest
	GetControlModule() *string
	SetControlSubModule(v string) *FindBestMatchSecurityStrategyRequest
	GetControlSubModule() *string
	SetWorkspaceId(v int64) *FindBestMatchSecurityStrategyRequest
	GetWorkspaceId() *int64
}

type FindBestMatchSecurityStrategyRequest struct {
	// example:
	//
	// DataQuery
	ControlModule *string `json:"ControlModule,omitempty" xml:"ControlModule,omitempty"`
	// example:
	//
	// MyCatalog
	ControlSubModule *string `json:"ControlSubModule,omitempty" xml:"ControlSubModule,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s FindBestMatchSecurityStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s FindBestMatchSecurityStrategyRequest) GoString() string {
	return s.String()
}

func (s *FindBestMatchSecurityStrategyRequest) GetControlModule() *string {
	return s.ControlModule
}

func (s *FindBestMatchSecurityStrategyRequest) GetControlSubModule() *string {
	return s.ControlSubModule
}

func (s *FindBestMatchSecurityStrategyRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *FindBestMatchSecurityStrategyRequest) SetControlModule(v string) *FindBestMatchSecurityStrategyRequest {
	s.ControlModule = &v
	return s
}

func (s *FindBestMatchSecurityStrategyRequest) SetControlSubModule(v string) *FindBestMatchSecurityStrategyRequest {
	s.ControlSubModule = &v
	return s
}

func (s *FindBestMatchSecurityStrategyRequest) SetWorkspaceId(v int64) *FindBestMatchSecurityStrategyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *FindBestMatchSecurityStrategyRequest) Validate() error {
	return dara.Validate(s)
}
