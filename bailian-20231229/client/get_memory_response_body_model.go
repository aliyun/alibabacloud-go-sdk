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
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbd5fc3
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// llm-us9hjmt32nysdm5v
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
