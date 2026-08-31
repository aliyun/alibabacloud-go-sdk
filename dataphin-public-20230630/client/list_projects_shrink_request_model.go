// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListProjectsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListProjectsShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListProjectsShrinkRequest
	GetOpUserId() *string
}

type ListProjectsShrinkRequest struct {
	// The query conditions.
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
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListProjectsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListProjectsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListProjectsShrinkRequest) SetListQueryShrink(v string) *ListProjectsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetOpTenantId(v int64) *ListProjectsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetOpUserId(v string) *ListProjectsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
