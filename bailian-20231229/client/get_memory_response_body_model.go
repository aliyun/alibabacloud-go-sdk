// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetMemoryResponseBody
	GetDescription() *string
	SetMemoryId(v string) *GetMemoryResponseBody
	GetMemoryId() *string
	SetRequestId(v string) *GetMemoryResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *GetMemoryResponseBody
	GetWorkspaceId() *string
}

type GetMemoryResponseBody struct {
	// The description of the long-term memory.
	//
	// example:
	//
	// 我的大模型应用$APP_ID关于A用户的长期记忆体
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the long-term memory.
	//
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbdxxxx
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-11402910xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the workspace to which the long-term memory belongs.
	//
	// example:
	//
	// llm-3z7uw7fwz0vexxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMemoryResponseBody) GetMemoryId() *string {
	return s.MemoryId
}

func (s *GetMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMemoryResponseBody) SetDescription(v string) *GetMemoryResponseBody {
	s.Description = &v
	return s
}

func (s *GetMemoryResponseBody) SetMemoryId(v string) *GetMemoryResponseBody {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryResponseBody) SetRequestId(v string) *GetMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryResponseBody) SetWorkspaceId(v string) *GetMemoryResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
