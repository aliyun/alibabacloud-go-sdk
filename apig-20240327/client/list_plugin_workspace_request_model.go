// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *ListPluginWorkspaceRequest
	GetGatewayType() *string
}

type ListPluginWorkspaceRequest struct {
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
}

func (s ListPluginWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *ListPluginWorkspaceRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListPluginWorkspaceRequest) SetGatewayType(v string) *ListPluginWorkspaceRequest {
	s.GatewayType = &v
	return s
}

func (s *ListPluginWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
