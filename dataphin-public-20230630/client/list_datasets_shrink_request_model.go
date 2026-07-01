// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetQueryShrink(v string) *ListDatasetsShrinkRequest
	GetDatasetQueryShrink() *string
	SetOpTenantId(v int64) *ListDatasetsShrinkRequest
	GetOpTenantId() *int64
}

type ListDatasetsShrinkRequest struct {
	// The request body.
	DatasetQueryShrink *string `json:"DatasetQuery,omitempty" xml:"DatasetQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDatasetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsShrinkRequest) GetDatasetQueryShrink() *string {
	return s.DatasetQueryShrink
}

func (s *ListDatasetsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDatasetsShrinkRequest) SetDatasetQueryShrink(v string) *ListDatasetsShrinkRequest {
	s.DatasetQueryShrink = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetOpTenantId(v int64) *ListDatasetsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDatasetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
