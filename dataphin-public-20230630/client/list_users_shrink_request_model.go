// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListUsersShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListUsersShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListUsersShrinkRequest
	GetOpUserId() *string
}

type ListUsersShrinkRequest struct {
	// The paged query for paging.
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
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUsersShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListUsersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUsersShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListUsersShrinkRequest) SetListQueryShrink(v string) *ListUsersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListUsersShrinkRequest) SetOpTenantId(v int64) *ListUsersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUsersShrinkRequest) SetOpUserId(v string) *ListUsersShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
