// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListPublishRecordsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListPublishRecordsShrinkRequest
	GetOpTenantId() *int64
}

type ListPublishRecordsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListPublishRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListPublishRecordsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListPublishRecordsShrinkRequest) SetListQueryShrink(v string) *ListPublishRecordsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListPublishRecordsShrinkRequest) SetOpTenantId(v int64) *ListPublishRecordsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListPublishRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
