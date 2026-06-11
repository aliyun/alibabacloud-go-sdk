// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMemoryShrinkRequest
	GetAppId() *string
	SetContent(v string) *UpdateMemoryShrinkRequest
	GetContent() *string
	SetMemoryNodeId(v string) *UpdateMemoryShrinkRequest
	GetMemoryNodeId() *string
	SetMetaDataShrink(v string) *UpdateMemoryShrinkRequest
	GetMetaDataShrink() *string
	SetSource(v string) *UpdateMemoryShrinkRequest
	GetSource() *string
	SetUserDefinedId(v string) *UpdateMemoryShrinkRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *UpdateMemoryShrinkRequest
	GetWorkspaceId() *string
}

type UpdateMemoryShrinkRequest struct {
	// This parameter is required.
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	MemoryNodeId   *string `json:"MemoryNodeId,omitempty" xml:"MemoryNodeId,omitempty"`
	MetaDataShrink *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMemoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMemoryShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateMemoryShrinkRequest) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *UpdateMemoryShrinkRequest) GetMetaDataShrink() *string {
	return s.MetaDataShrink
}

func (s *UpdateMemoryShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateMemoryShrinkRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *UpdateMemoryShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMemoryShrinkRequest) SetAppId(v string) *UpdateMemoryShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) SetContent(v string) *UpdateMemoryShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) SetMemoryNodeId(v string) *UpdateMemoryShrinkRequest {
	s.MemoryNodeId = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) SetMetaDataShrink(v string) *UpdateMemoryShrinkRequest {
	s.MetaDataShrink = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) SetSource(v string) *UpdateMemoryShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) SetUserDefinedId(v string) *UpdateMemoryShrinkRequest {
	s.UserDefinedId = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) SetWorkspaceId(v string) *UpdateMemoryShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMemoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
