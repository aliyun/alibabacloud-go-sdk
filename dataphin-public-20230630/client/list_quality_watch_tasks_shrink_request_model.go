// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListQualityWatchTasksShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListQualityWatchTasksShrinkRequest
	GetOpTenantId() *int64
}

type ListQualityWatchTasksShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityWatchTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQualityWatchTasksShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListQualityWatchTasksShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityWatchTasksShrinkRequest) SetListQueryShrink(v string) *ListQualityWatchTasksShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListQualityWatchTasksShrinkRequest) SetOpTenantId(v int64) *ListQualityWatchTasksShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityWatchTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
