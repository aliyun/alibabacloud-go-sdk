// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallStatisticsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServiceApiCallStatisticsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServiceApiCallStatisticsShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceApiCallStatisticsShrinkRequest
	GetProjectId() *int32
}

type ListDataServiceApiCallStatisticsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceApiCallStatisticsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallStatisticsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServiceApiCallStatisticsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceApiCallStatisticsShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiCallStatisticsShrinkRequest) SetListQueryShrink(v string) *ListDataServiceApiCallStatisticsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsShrinkRequest) SetOpTenantId(v int64) *ListDataServiceApiCallStatisticsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsShrinkRequest) SetProjectId(v int32) *ListDataServiceApiCallStatisticsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
