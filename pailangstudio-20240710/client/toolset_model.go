// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolset interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Toolset
	GetAccessibility() *string
	SetCreator(v string) *Toolset
	GetCreator() *string
	SetDescription(v string) *Toolset
	GetDescription() *string
	SetGmtCreateTime(v string) *Toolset
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Toolset
	GetGmtModifiedTime() *string
	SetTools(v string) *Toolset
	GetTools() *string
	SetToolsetConfig(v string) *Toolset
	GetToolsetConfig() *string
	SetToolsetId(v string) *Toolset
	GetToolsetId() *string
	SetToolsetName(v string) *Toolset
	GetToolsetName() *string
	SetToolsetType(v string) *Toolset
	GetToolsetType() *string
	SetWorkspaceId(v string) *Toolset
	GetWorkspaceId() *string
}

type Toolset struct {
	Accessibility   *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Creator         *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Tools           *string `json:"Tools,omitempty" xml:"Tools,omitempty"`
	ToolsetConfig   *string `json:"ToolsetConfig,omitempty" xml:"ToolsetConfig,omitempty"`
	ToolsetId       *string `json:"ToolsetId,omitempty" xml:"ToolsetId,omitempty"`
	ToolsetName     *string `json:"ToolsetName,omitempty" xml:"ToolsetName,omitempty"`
	ToolsetType     *string `json:"ToolsetType,omitempty" xml:"ToolsetType,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Toolset) String() string {
	return dara.Prettify(s)
}

func (s Toolset) GoString() string {
	return s.String()
}

func (s *Toolset) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Toolset) GetCreator() *string {
	return s.Creator
}

func (s *Toolset) GetDescription() *string {
	return s.Description
}

func (s *Toolset) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Toolset) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Toolset) GetTools() *string {
	return s.Tools
}

func (s *Toolset) GetToolsetConfig() *string {
	return s.ToolsetConfig
}

func (s *Toolset) GetToolsetId() *string {
	return s.ToolsetId
}

func (s *Toolset) GetToolsetName() *string {
	return s.ToolsetName
}

func (s *Toolset) GetToolsetType() *string {
	return s.ToolsetType
}

func (s *Toolset) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Toolset) SetAccessibility(v string) *Toolset {
	s.Accessibility = &v
	return s
}

func (s *Toolset) SetCreator(v string) *Toolset {
	s.Creator = &v
	return s
}

func (s *Toolset) SetDescription(v string) *Toolset {
	s.Description = &v
	return s
}

func (s *Toolset) SetGmtCreateTime(v string) *Toolset {
	s.GmtCreateTime = &v
	return s
}

func (s *Toolset) SetGmtModifiedTime(v string) *Toolset {
	s.GmtModifiedTime = &v
	return s
}

func (s *Toolset) SetTools(v string) *Toolset {
	s.Tools = &v
	return s
}

func (s *Toolset) SetToolsetConfig(v string) *Toolset {
	s.ToolsetConfig = &v
	return s
}

func (s *Toolset) SetToolsetId(v string) *Toolset {
	s.ToolsetId = &v
	return s
}

func (s *Toolset) SetToolsetName(v string) *Toolset {
	s.ToolsetName = &v
	return s
}

func (s *Toolset) SetToolsetType(v string) *Toolset {
	s.ToolsetType = &v
	return s
}

func (s *Toolset) SetWorkspaceId(v string) *Toolset {
	s.WorkspaceId = &v
	return s
}

func (s *Toolset) Validate() error {
	return dara.Validate(s)
}
