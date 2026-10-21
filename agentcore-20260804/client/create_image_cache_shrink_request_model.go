// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageCacheShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateImageCacheShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateImageCacheShrinkRequest
	GetClientToken() *string
	SetRegionId(v string) *CreateImageCacheShrinkRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CreateImageCacheShrinkRequest
	GetWorkspaceId() *string
}

type CreateImageCacheShrinkRequest struct {
	// The business parameters for creating an image cache.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// This parameter is required and cannot be empty. It is reserved to be compatible with the API. Currently, it does not provide the ability to deduplicate requests or recover the original result based on this field. If a timeout occurs or the result is unconfirmed, query the query cache status first before deciding whether to recommit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
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

func (s CreateImageCacheShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageCacheShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateImageCacheShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageCacheShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageCacheShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateImageCacheShrinkRequest) SetBodyShrink(v string) *CreateImageCacheShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetClientToken(v string) *CreateImageCacheShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetRegionId(v string) *CreateImageCacheShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) SetWorkspaceId(v string) *CreateImageCacheShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateImageCacheShrinkRequest) Validate() error {
	return dara.Validate(s)
}
