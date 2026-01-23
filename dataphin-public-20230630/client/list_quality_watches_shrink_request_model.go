// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityWatchesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListQualityWatchesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListQualityWatchesShrinkRequest
	GetOpTenantId() *int64
}

type ListQualityWatchesShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityWatchesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityWatchesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQualityWatchesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListQualityWatchesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityWatchesShrinkRequest) SetListQueryShrink(v string) *ListQualityWatchesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListQualityWatchesShrinkRequest) SetOpTenantId(v int64) *ListQualityWatchesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityWatchesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
