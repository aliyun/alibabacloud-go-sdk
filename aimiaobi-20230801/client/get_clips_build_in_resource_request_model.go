// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClipsBuildInResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v int32) *GetClipsBuildInResourceRequest
	GetResourceType() *int32
	SetWorkspaceId(v string) *GetClipsBuildInResourceRequest
	GetWorkspaceId() *string
}

type GetClipsBuildInResourceRequest struct {
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// llm-az2gglkjauwnnhpq
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetClipsBuildInResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClipsBuildInResourceRequest) GoString() string {
	return s.String()
}

func (s *GetClipsBuildInResourceRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GetClipsBuildInResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetClipsBuildInResourceRequest) SetResourceType(v int32) *GetClipsBuildInResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *GetClipsBuildInResourceRequest) SetWorkspaceId(v string) *GetClipsBuildInResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetClipsBuildInResourceRequest) Validate() error {
	return dara.Validate(s)
}
