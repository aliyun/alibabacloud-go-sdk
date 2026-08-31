// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityArchiveTablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListQualityArchiveTablesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListQualityArchiveTablesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListQualityArchiveTablesShrinkRequest
	GetOpUserId() *string
}

type ListQualityArchiveTablesShrinkRequest struct {
	// The input parameters for querying the anomaly archived table list.
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

func (s ListQualityArchiveTablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityArchiveTablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQualityArchiveTablesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListQualityArchiveTablesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityArchiveTablesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListQualityArchiveTablesShrinkRequest) SetListQueryShrink(v string) *ListQualityArchiveTablesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListQualityArchiveTablesShrinkRequest) SetOpTenantId(v int64) *ListQualityArchiveTablesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityArchiveTablesShrinkRequest) SetOpUserId(v string) *ListQualityArchiveTablesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListQualityArchiveTablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
