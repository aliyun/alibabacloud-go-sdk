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
}

type ListTablesShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *ListTablesShrinkRequest) SetListQueryShrink(v string) *ListTablesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListTablesShrinkRequest) SetOpTenantId(v int64) *ListTablesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
