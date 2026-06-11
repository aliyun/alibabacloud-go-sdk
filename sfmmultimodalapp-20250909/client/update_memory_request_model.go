// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMemoryRequest
	GetAppId() *string
	SetContent(v string) *UpdateMemoryRequest
	GetContent() *string
	SetMemoryNodeId(v string) *UpdateMemoryRequest
	GetMemoryNodeId() *string
	SetMetaData(v map[string]*string) *UpdateMemoryRequest
	GetMetaData() map[string]*string
	SetSource(v string) *UpdateMemoryRequest
	GetSource() *string
	SetUserDefinedId(v string) *UpdateMemoryRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *UpdateMemoryRequest
	GetWorkspaceId() *string
}

type UpdateMemoryRequest struct {
	// This parameter is required.
	AppId        *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content      *string            `json:"Content,omitempty" xml:"Content,omitempty"`
	MemoryNodeId *string            `json:"MemoryNodeId,omitempty" xml:"MemoryNodeId,omitempty"`
	MetaData     map[string]*string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	Source       *string            `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateMemoryRequest) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *UpdateMemoryRequest) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *UpdateMemoryRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateMemoryRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *UpdateMemoryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMemoryRequest) SetAppId(v string) *UpdateMemoryRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMemoryRequest) SetContent(v string) *UpdateMemoryRequest {
	s.Content = &v
	return s
}

func (s *UpdateMemoryRequest) SetMemoryNodeId(v string) *UpdateMemoryRequest {
	s.MemoryNodeId = &v
	return s
}

func (s *UpdateMemoryRequest) SetMetaData(v map[string]*string) *UpdateMemoryRequest {
	s.MetaData = v
	return s
}

func (s *UpdateMemoryRequest) SetSource(v string) *UpdateMemoryRequest {
	s.Source = &v
	return s
}

func (s *UpdateMemoryRequest) SetUserDefinedId(v string) *UpdateMemoryRequest {
	s.UserDefinedId = &v
	return s
}

func (s *UpdateMemoryRequest) SetWorkspaceId(v string) *UpdateMemoryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
