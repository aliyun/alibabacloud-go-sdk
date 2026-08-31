// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListFilesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListFilesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListFilesShrinkRequest
	GetOpUserId() *string
}

type ListFilesShrinkRequest struct {
	// Query conditions
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFilesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListFilesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListFilesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListFilesShrinkRequest) SetListQueryShrink(v string) *ListFilesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListFilesShrinkRequest) SetOpTenantId(v int64) *ListFilesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListFilesShrinkRequest) SetOpUserId(v string) *ListFilesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
