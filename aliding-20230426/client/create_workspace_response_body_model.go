// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkspaceResponseBody
	GetDescription() *string
	SetName(v string) *CreateWorkspaceResponseBody
	GetName() *string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetUrl(v string) *CreateWorkspaceResponseBody
	GetUrl() *string
	SetWorkspaceId(v string) *CreateWorkspaceResponseBody
	GetWorkspaceId() *string
}

type CreateWorkspaceResponseBody struct {
	// example:
	//
	// 知识库描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 知识库
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// https://xxx/workspaceId
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// YRBGvyxxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CreateWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResponseBody) SetDescription(v string) *CreateWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetName(v string) *CreateWorkspaceResponseBody {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetUrl(v string) *CreateWorkspaceResponseBody {
	s.Url = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v string) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
