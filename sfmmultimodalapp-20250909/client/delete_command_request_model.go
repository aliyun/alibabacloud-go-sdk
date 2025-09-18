// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteCommandRequest
	GetAppId() *string
	SetDomainCode(v string) *DeleteCommandRequest
	GetDomainCode() *string
	SetToolId(v string) *DeleteCommandRequest
	GetToolId() *string
	SetWorkspaceId(v string) *DeleteCommandRequest
	GetWorkspaceId() *string
}

type DeleteCommandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7533545
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 564646456
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// example:
	//
	// llm-xxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommandRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommandRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteCommandRequest) GetDomainCode() *string {
	return s.DomainCode
}

func (s *DeleteCommandRequest) GetToolId() *string {
	return s.ToolId
}

func (s *DeleteCommandRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteCommandRequest) SetAppId(v string) *DeleteCommandRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCommandRequest) SetDomainCode(v string) *DeleteCommandRequest {
	s.DomainCode = &v
	return s
}

func (s *DeleteCommandRequest) SetToolId(v string) *DeleteCommandRequest {
	s.ToolId = &v
	return s
}

func (s *DeleteCommandRequest) SetWorkspaceId(v string) *DeleteCommandRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteCommandRequest) Validate() error {
	return dara.Validate(s)
}
