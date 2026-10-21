// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetImageCacheRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *GetImageCacheRequest
	GetWorkspaceId() *string
}

type GetImageCacheRequest struct {
	// The product entry region, which is determined by the selected region endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The current routed workspace. The server verifies the account, region, and resource ownership. This parameter cannot be omitted or used to access other workspaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetImageCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheRequest) GoString() string {
	return s.String()
}

func (s *GetImageCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageCacheRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetImageCacheRequest) SetRegionId(v string) *GetImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *GetImageCacheRequest) SetWorkspaceId(v string) *GetImageCacheRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetImageCacheRequest) Validate() error {
	return dara.Validate(s)
}
