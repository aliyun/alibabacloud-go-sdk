// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetMemoryNodeResponseBody
	GetContent() *string
	SetMemoryId(v string) *GetMemoryNodeResponseBody
	GetMemoryId() *string
	SetMemoryNodeId(v string) *GetMemoryNodeResponseBody
	GetMemoryNodeId() *string
	SetRequestId(v string) *GetMemoryNodeResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *GetMemoryNodeResponseBody
	GetWorkspaceId() *string
}

type GetMemoryNodeResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbd5fc3
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// example:
	//
	// 68de06c95368463a8be4a84efc872cc5
	MemoryNodeId *string `json:"memoryNodeId,omitempty" xml:"memoryNodeId,omitempty"`
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// llm-us9hjmt32nysdm5v
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetMemoryNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryNodeResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetMemoryNodeResponseBody) GetMemoryId() *string {
	return s.MemoryId
}

func (s *GetMemoryNodeResponseBody) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *GetMemoryNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryNodeResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMemoryNodeResponseBody) SetContent(v string) *GetMemoryNodeResponseBody {
	s.Content = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetMemoryId(v string) *GetMemoryNodeResponseBody {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetMemoryNodeId(v string) *GetMemoryNodeResponseBody {
	s.MemoryNodeId = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetRequestId(v string) *GetMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetWorkspaceId(v string) *GetMemoryNodeResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetMemoryNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
