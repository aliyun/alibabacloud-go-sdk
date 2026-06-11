// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMemoryRequest
	GetAppId() *string
	SetMemoryNodeId(v string) *DeleteMemoryRequest
	GetMemoryNodeId() *string
	SetUserDefinedId(v string) *DeleteMemoryRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *DeleteMemoryRequest
	GetWorkspaceId() *string
}

type DeleteMemoryRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MemoryNodeId *string `json:"MemoryNodeId,omitempty" xml:"MemoryNodeId,omitempty"`
	// This parameter is required.
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemoryRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMemoryRequest) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *DeleteMemoryRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *DeleteMemoryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMemoryRequest) SetAppId(v string) *DeleteMemoryRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMemoryRequest) SetMemoryNodeId(v string) *DeleteMemoryRequest {
	s.MemoryNodeId = &v
	return s
}

func (s *DeleteMemoryRequest) SetUserDefinedId(v string) *DeleteMemoryRequest {
	s.UserDefinedId = &v
	return s
}

func (s *DeleteMemoryRequest) SetWorkspaceId(v string) *DeleteMemoryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMemoryRequest) Validate() error {
	return dara.Validate(s)
}
