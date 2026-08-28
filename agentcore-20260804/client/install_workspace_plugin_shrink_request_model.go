// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallWorkspacePluginShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *InstallWorkspacePluginShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *InstallWorkspacePluginShrinkRequest
	GetClientToken() *string
}

type InstallWorkspacePluginShrinkRequest struct {
	// The request body for installing a plugin.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client idempotency token.
	//
	// example:
	//
	// workspace-plugin-install-20260810-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InstallWorkspacePluginShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *InstallWorkspacePluginShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InstallWorkspacePluginShrinkRequest) SetBodyShrink(v string) *InstallWorkspacePluginShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *InstallWorkspacePluginShrinkRequest) SetClientToken(v string) *InstallWorkspacePluginShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallWorkspacePluginShrinkRequest) Validate() error {
	return dara.Validate(s)
}
