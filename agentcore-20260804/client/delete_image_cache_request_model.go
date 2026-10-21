// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteImageCacheRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DeleteImageCacheRequest
	GetWorkspaceId() *string
}

type DeleteImageCacheRequest struct {
	// The product entry region, which is determined by the selected regional endpoint.
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

func (s DeleteImageCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteImageCacheRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteImageCacheRequest) SetRegionId(v string) *DeleteImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetWorkspaceId(v string) *DeleteImageCacheRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteImageCacheRequest) Validate() error {
	return dara.Validate(s)
}
