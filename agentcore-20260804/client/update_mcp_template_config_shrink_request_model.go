// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpTemplateConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateMcpTemplateConfigShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateMcpTemplateConfigShrinkRequest
	GetClientToken() *string
	SetTemplateVersion(v string) *UpdateMcpTemplateConfigShrinkRequest
	GetTemplateVersion() *string
}

type UpdateMcpTemplateConfigShrinkRequest struct {
	// The MCP configuration to update by the specified template version. The configuration must conform to the input schema of the template.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The template version used for this update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s UpdateMcpTemplateConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpTemplateConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpTemplateConfigShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateMcpTemplateConfigShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMcpTemplateConfigShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateMcpTemplateConfigShrinkRequest) SetBodyShrink(v string) *UpdateMcpTemplateConfigShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateMcpTemplateConfigShrinkRequest) SetClientToken(v string) *UpdateMcpTemplateConfigShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMcpTemplateConfigShrinkRequest) SetTemplateVersion(v string) *UpdateMcpTemplateConfigShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateMcpTemplateConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
