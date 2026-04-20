// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingMcpShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MmAppBindingMcpShrinkRequest
	GetAppId() *string
	SetMcpsShrink(v string) *MmAppBindingMcpShrinkRequest
	GetMcpsShrink() *string
	SetWorkspaceId(v string) *MmAppBindingMcpShrinkRequest
	GetWorkspaceId() *string
}

type MmAppBindingMcpShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_a2eb4e04b48041108edb1f6de815
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	McpsShrink *string `json:"Mcps,omitempty" xml:"Mcps,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-o8ixktz41iyd2b6p
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MmAppBindingMcpShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingMcpShrinkRequest) GoString() string {
	return s.String()
}

func (s *MmAppBindingMcpShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *MmAppBindingMcpShrinkRequest) GetMcpsShrink() *string {
	return s.McpsShrink
}

func (s *MmAppBindingMcpShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *MmAppBindingMcpShrinkRequest) SetAppId(v string) *MmAppBindingMcpShrinkRequest {
	s.AppId = &v
	return s
}

func (s *MmAppBindingMcpShrinkRequest) SetMcpsShrink(v string) *MmAppBindingMcpShrinkRequest {
	s.McpsShrink = &v
	return s
}

func (s *MmAppBindingMcpShrinkRequest) SetWorkspaceId(v string) *MmAppBindingMcpShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *MmAppBindingMcpShrinkRequest) Validate() error {
	return dara.Validate(s)
}
