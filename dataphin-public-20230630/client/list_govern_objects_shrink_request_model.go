// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGovernObjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListGovernObjectsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListGovernObjectsShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListGovernObjectsShrinkRequest
	GetOpUserId() *string
}

type ListGovernObjectsShrinkRequest struct {
	// The paged query conditions.
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

func (s ListGovernObjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListGovernObjectsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListGovernObjectsShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListGovernObjectsShrinkRequest) SetListQueryShrink(v string) *ListGovernObjectsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListGovernObjectsShrinkRequest) SetOpTenantId(v int64) *ListGovernObjectsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListGovernObjectsShrinkRequest) SetOpUserId(v string) *ListGovernObjectsShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListGovernObjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
