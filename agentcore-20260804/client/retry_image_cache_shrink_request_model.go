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
	// The request body is a fixed empty object {}. WorkspaceId is passed as a query parameter. Do not pass the Pool, role, Digest, or network parameters.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// This parameter is required and cannot be empty. It is reserved to keep the API compatible. This field does not currently deduplicate requests or recover the original result. If a timeout occurs or the result is not confirmed, query the query cache status first before deciding whether to commit again.
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
	// The current routed workspace. The server verifies the account, region, and resource ownership. You cannot omit this parameter or use it to access other workspaces.
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
