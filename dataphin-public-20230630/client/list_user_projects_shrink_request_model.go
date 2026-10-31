// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListUserProjectsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListUserProjectsShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListUserProjectsShrinkRequest
	GetOpUserId() *string
}

type ListUserProjectsShrinkRequest struct {
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
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListUserProjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserProjectsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListUserProjectsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUserProjectsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListUserProjectsShrinkRequest) SetListQueryShrink(v string) *ListUserProjectsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListUserProjectsShrinkRequest) SetOpTenantId(v int64) *ListUserProjectsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUserProjectsShrinkRequest) SetOpUserId(v string) *ListUserProjectsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListUserProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
