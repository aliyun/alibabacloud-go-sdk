// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListTablesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListTablesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListTablesShrinkRequest
	GetOpUserId() *string
}

type ListTablesShrinkRequest struct {
	// The paged query conditions.
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

func (s ListTablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTablesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListTablesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListTablesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListTablesShrinkRequest) SetListQueryShrink(v string) *ListTablesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListTablesShrinkRequest) SetOpTenantId(v int64) *ListTablesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTablesShrinkRequest) SetOpUserId(v string) *ListTablesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListTablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
