// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubmitRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListSubmitRecordsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListSubmitRecordsShrinkRequest
	GetOpTenantId() *int64
}

type ListSubmitRecordsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSubmitRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListSubmitRecordsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSubmitRecordsShrinkRequest) SetListQueryShrink(v string) *ListSubmitRecordsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListSubmitRecordsShrinkRequest) SetOpTenantId(v int64) *ListSubmitRecordsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSubmitRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
