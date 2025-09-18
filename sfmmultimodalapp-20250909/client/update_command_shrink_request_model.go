// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateCommandShrinkRequest
	GetAppId() *string
	SetDomainCode(v string) *UpdateCommandShrinkRequest
	GetDomainCode() *string
	SetDomainName(v string) *UpdateCommandShrinkRequest
	GetDomainName() *string
	SetToolDescription(v string) *UpdateCommandShrinkRequest
	GetToolDescription() *string
	SetToolExamplesShrink(v string) *UpdateCommandShrinkRequest
	GetToolExamplesShrink() *string
	SetToolId(v string) *UpdateCommandShrinkRequest
	GetToolId() *string
	SetToolName(v string) *UpdateCommandShrinkRequest
	GetToolName() *string
	SetToolParamsShrink(v string) *UpdateCommandShrinkRequest
	GetToolParamsShrink() *string
	SetWorkspaceId(v string) *UpdateCommandShrinkRequest
	GetWorkspaceId() *string
}

type UpdateCommandShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_axaxaaa
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 659864545
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// example:
	//
	// shopping_t
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	ToolDescription    *string `json:"ToolDescription,omitempty" xml:"ToolDescription,omitempty"`
	ToolExamplesShrink *string `json:"ToolExamples,omitempty" xml:"ToolExamples,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8293382932xxx
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open_bx
	ToolName         *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParamsShrink *string `json:"ToolParams,omitempty" xml:"ToolParams,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommandShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCommandShrinkRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *UpdateCommandShrinkRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateCommandShrinkRequest) GetToolDescription() *string {
	return s.ToolDescription
}

func (s *UpdateCommandShrinkRequest) GetToolExamplesShrink() *string {
	return s.ToolExamplesShrink
}

func (s *UpdateCommandShrinkRequest) GetToolId() *string {
	return s.ToolId
}

func (s *UpdateCommandShrinkRequest) GetToolName() *string {
	return s.ToolName
}

func (s *UpdateCommandShrinkRequest) GetToolParamsShrink() *string {
	return s.ToolParamsShrink
}

func (s *UpdateCommandShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCommandShrinkRequest) SetAppId(v string) *UpdateCommandShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetDomainCode(v string) *UpdateCommandShrinkRequest {
	s.DomainCode = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetDomainName(v string) *UpdateCommandShrinkRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetToolDescription(v string) *UpdateCommandShrinkRequest {
	s.ToolDescription = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetToolExamplesShrink(v string) *UpdateCommandShrinkRequest {
	s.ToolExamplesShrink = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetToolId(v string) *UpdateCommandShrinkRequest {
	s.ToolId = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetToolName(v string) *UpdateCommandShrinkRequest {
	s.ToolName = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetToolParamsShrink(v string) *UpdateCommandShrinkRequest {
	s.ToolParamsShrink = &v
	return s
}

func (s *UpdateCommandShrinkRequest) SetWorkspaceId(v string) *UpdateCommandShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
