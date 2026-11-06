// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryImageCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RetryImageCacheRequestBody) *RetryImageCacheRequest
	GetBody() *RetryImageCacheRequestBody
	SetClientToken(v string) *RetryImageCacheRequest
	GetClientToken() *string
	SetRegionId(v string) *RetryImageCacheRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *RetryImageCacheRequest
	GetWorkspaceId() *string
}

type RetryImageCacheRequest struct {
	// The request body for retrying image cache creation.
	Body *RetryImageCacheRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The non-empty request token passed as a query parameter. The current implementation does not guarantee deduplication or recovery of the original acceptance result based on this token.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The product entry region. This parameter is passed as a query parameter and must be consistent with the selected region endpoint and the workspace region.
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

func (s RetryImageCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryImageCacheRequest) GoString() string {
	return s.String()
}

func (s *RetryImageCacheRequest) GetBody() *RetryImageCacheRequestBody {
	return s.Body
}

func (s *RetryImageCacheRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RetryImageCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RetryImageCacheRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RetryImageCacheRequest) SetBody(v *RetryImageCacheRequestBody) *RetryImageCacheRequest {
	s.Body = v
	return s
}

func (s *RetryImageCacheRequest) SetClientToken(v string) *RetryImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *RetryImageCacheRequest) SetRegionId(v string) *RetryImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *RetryImageCacheRequest) SetWorkspaceId(v string) *RetryImageCacheRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RetryImageCacheRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryImageCacheRequestBody struct {
}

func (s RetryImageCacheRequestBody) String() string {
	return dara.Prettify(s)
}

func (s RetryImageCacheRequestBody) GoString() string {
	return s.String()
}

func (s *RetryImageCacheRequestBody) Validate() error {
	return dara.Validate(s)
}
