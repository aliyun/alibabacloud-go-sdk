// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryImageCacheShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *RetryImageCacheShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *RetryImageCacheShrinkRequest
	GetClientToken() *string
	SetRegionId(v string) *RetryImageCacheShrinkRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *RetryImageCacheShrinkRequest
	GetWorkspaceId() *string
}

type RetryImageCacheShrinkRequest struct {
	// The request body for retrying image cache creation.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The non-empty client token passed as a query parameter. The current implementation does not guarantee deduplication by this token or restoration of the original acceptance result.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The region ID. This parameter is passed as a query parameter and must match the selected region endpoint and the workspace region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s RetryImageCacheShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryImageCacheShrinkRequest) GoString() string {
	return s.String()
}

func (s *RetryImageCacheShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *RetryImageCacheShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RetryImageCacheShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RetryImageCacheShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RetryImageCacheShrinkRequest) SetBodyShrink(v string) *RetryImageCacheShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *RetryImageCacheShrinkRequest) SetClientToken(v string) *RetryImageCacheShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RetryImageCacheShrinkRequest) SetRegionId(v string) *RetryImageCacheShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RetryImageCacheShrinkRequest) SetWorkspaceId(v string) *RetryImageCacheShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RetryImageCacheShrinkRequest) Validate() error {
	return dara.Validate(s)
}
