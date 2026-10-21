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
	// The request body is a fixed empty object {}. WorkspaceId is passed as a query parameter. Do not pass the Pool, role, Digest, or network parameters.
	Body *RetryImageCacheRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
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
