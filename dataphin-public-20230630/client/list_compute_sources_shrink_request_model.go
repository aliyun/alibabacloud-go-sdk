// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeSourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListComputeSourcesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListComputeSourcesShrinkRequest
	GetOpTenantId() *int64
}

type ListComputeSourcesShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListComputeSourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListComputeSourcesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListComputeSourcesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListComputeSourcesShrinkRequest) SetListQueryShrink(v string) *ListComputeSourcesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListComputeSourcesShrinkRequest) SetOpTenantId(v int64) *ListComputeSourcesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListComputeSourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
