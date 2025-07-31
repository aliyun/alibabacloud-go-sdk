// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineageByTaskIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetTableColumnLineageByTaskIdShrinkRequest
	GetOpTenantId() *int64
	SetTableColumnLineageByTaskIdQueryShrink(v string) *GetTableColumnLineageByTaskIdShrinkRequest
	GetTableColumnLineageByTaskIdQueryShrink() *string
}

type GetTableColumnLineageByTaskIdShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	TableColumnLineageByTaskIdQueryShrink *string `json:"TableColumnLineageByTaskIdQuery,omitempty" xml:"TableColumnLineageByTaskIdQuery,omitempty"`
}

func (s GetTableColumnLineageByTaskIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineageByTaskIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineageByTaskIdShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableColumnLineageByTaskIdShrinkRequest) GetTableColumnLineageByTaskIdQueryShrink() *string {
	return s.TableColumnLineageByTaskIdQueryShrink
}

func (s *GetTableColumnLineageByTaskIdShrinkRequest) SetOpTenantId(v int64) *GetTableColumnLineageByTaskIdShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdShrinkRequest) SetTableColumnLineageByTaskIdQueryShrink(v string) *GetTableColumnLineageByTaskIdShrinkRequest {
	s.TableColumnLineageByTaskIdQueryShrink = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
