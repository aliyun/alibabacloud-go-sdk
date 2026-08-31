// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetTopicsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListAssetTopicsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListAssetTopicsShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListAssetTopicsShrinkRequest
	GetOpUserId() *string
}

type ListAssetTopicsShrinkRequest struct {
	// The query parameters.
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListAssetTopicsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListAssetTopicsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAssetTopicsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListAssetTopicsShrinkRequest) SetListQueryShrink(v string) *ListAssetTopicsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListAssetTopicsShrinkRequest) SetOpTenantId(v int64) *ListAssetTopicsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAssetTopicsShrinkRequest) SetOpUserId(v string) *ListAssetTopicsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListAssetTopicsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
