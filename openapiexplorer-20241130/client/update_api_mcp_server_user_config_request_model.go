// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiMcpServerUserConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePublicAccess(v bool) *UpdateApiMcpServerUserConfigRequest
	GetEnablePublicAccess() *bool
	SetVpcWhitelists(v []*string) *UpdateApiMcpServerUserConfigRequest
	GetVpcWhitelists() []*string
}

type UpdateApiMcpServerUserConfigRequest struct {
	// Specifies whether to enable public network access for all API MCP Servers in your account, including the system MCP Server. By default, public network access is enabled. If you disable it, you can only access the servers through their VPC domain names.
	//
	// example:
	//
	// true
	EnablePublicAccess *bool `json:"enablePublicAccess,omitempty" xml:"enablePublicAccess,omitempty"`
	// The VPC whitelist. After disabling public network access, use this parameter to specify allowed source VPCs. If you do not set this parameter or leave it empty, all source VPCs are allowed.
	VpcWhitelists []*string `json:"vpcWhitelists,omitempty" xml:"vpcWhitelists,omitempty" type:"Repeated"`
}

func (s UpdateApiMcpServerUserConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerUserConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerUserConfigRequest) GetEnablePublicAccess() *bool {
	return s.EnablePublicAccess
}

func (s *UpdateApiMcpServerUserConfigRequest) GetVpcWhitelists() []*string {
	return s.VpcWhitelists
}

func (s *UpdateApiMcpServerUserConfigRequest) SetEnablePublicAccess(v bool) *UpdateApiMcpServerUserConfigRequest {
	s.EnablePublicAccess = &v
	return s
}

func (s *UpdateApiMcpServerUserConfigRequest) SetVpcWhitelists(v []*string) *UpdateApiMcpServerUserConfigRequest {
	s.VpcWhitelists = v
	return s
}

func (s *UpdateApiMcpServerUserConfigRequest) Validate() error {
	return dara.Validate(s)
}
