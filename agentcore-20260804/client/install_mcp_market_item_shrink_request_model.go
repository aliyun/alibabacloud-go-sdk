// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMcpMarketItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *InstallMcpMarketItemShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *InstallMcpMarketItemShrinkRequest
	GetClientToken() *string
	SetTemplateVersion(v string) *InstallMcpMarketItemShrinkRequest
	GetTemplateVersion() *string
}

type InstallMcpMarketItemShrinkRequest struct {
	// The MCP configuration submitted during template installation. The configuration must conform to the input schema of the template.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The template version to install. You can call GetMcpMarketItem to query available versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s InstallMcpMarketItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *InstallMcpMarketItemShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InstallMcpMarketItemShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *InstallMcpMarketItemShrinkRequest) SetBodyShrink(v string) *InstallMcpMarketItemShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *InstallMcpMarketItemShrinkRequest) SetClientToken(v string) *InstallMcpMarketItemShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallMcpMarketItemShrinkRequest) SetTemplateVersion(v string) *InstallMcpMarketItemShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *InstallMcpMarketItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
