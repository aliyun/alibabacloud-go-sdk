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
}

type ListProjectsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *ListProjectsShrinkRequest) SetListQueryShrink(v string) *ListProjectsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetOpTenantId(v int64) *ListProjectsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
