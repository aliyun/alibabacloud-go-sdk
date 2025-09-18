// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateCommandShrinkRequest
	GetAppId() *string
	SetDomainCode(v string) *CreateCommandShrinkRequest
	GetDomainCode() *string
	SetDomainName(v string) *CreateCommandShrinkRequest
	GetDomainName() *string
	SetToolDescription(v string) *CreateCommandShrinkRequest
	GetToolDescription() *string
	SetToolExamplesShrink(v string) *CreateCommandShrinkRequest
	GetToolExamplesShrink() *string
	SetToolName(v string) *CreateCommandShrinkRequest
	GetToolName() *string
	SetToolParamsShrink(v string) *CreateCommandShrinkRequest
	GetToolParamsShrink() *string
	SetWorkspaceId(v string) *CreateCommandShrinkRequest
	GetWorkspaceId() *string
}

type CreateCommandShrinkRequest struct {
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
	// This parameter is required.
	//
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
	// open_bx
	ToolName         *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParamsShrink *string `json:"ToolParams,omitempty" xml:"ToolParams,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCommandShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateCommandShrinkRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *CreateCommandShrinkRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCommandShrinkRequest) GetToolDescription() *string {
	return s.ToolDescription
}

func (s *CreateCommandShrinkRequest) GetToolExamplesShrink() *string {
	return s.ToolExamplesShrink
}

func (s *CreateCommandShrinkRequest) GetToolName() *string {
	return s.ToolName
}

func (s *CreateCommandShrinkRequest) GetToolParamsShrink() *string {
	return s.ToolParamsShrink
}

func (s *CreateCommandShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCommandShrinkRequest) SetAppId(v string) *CreateCommandShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateCommandShrinkRequest) SetDomainCode(v string) *CreateCommandShrinkRequest {
	s.DomainCode = &v
	return s
}

func (s *CreateCommandShrinkRequest) SetDomainName(v string) *CreateCommandShrinkRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCommandShrinkRequest) SetToolDescription(v string) *CreateCommandShrinkRequest {
	s.ToolDescription = &v
	return s
}

func (s *CreateCommandShrinkRequest) SetToolExamplesShrink(v string) *CreateCommandShrinkRequest {
	s.ToolExamplesShrink = &v
	return s
}

func (s *CreateCommandShrinkRequest) SetToolName(v string) *CreateCommandShrinkRequest {
	s.ToolName = &v
	return s
}

func (s *CreateCommandShrinkRequest) SetToolParamsShrink(v string) *CreateCommandShrinkRequest {
	s.ToolParamsShrink = &v
	return s
}

func (s *CreateCommandShrinkRequest) SetWorkspaceId(v string) *CreateCommandShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
