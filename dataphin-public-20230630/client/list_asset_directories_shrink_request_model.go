// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetDirectoriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListAssetDirectoriesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListAssetDirectoriesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListAssetDirectoriesShrinkRequest
	GetOpUserId() *string
}

type ListAssetDirectoriesShrinkRequest struct {
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
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListAssetDirectoriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListAssetDirectoriesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAssetDirectoriesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListAssetDirectoriesShrinkRequest) SetListQueryShrink(v string) *ListAssetDirectoriesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListAssetDirectoriesShrinkRequest) SetOpTenantId(v int64) *ListAssetDirectoriesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAssetDirectoriesShrinkRequest) SetOpUserId(v string) *ListAssetDirectoriesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListAssetDirectoriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
