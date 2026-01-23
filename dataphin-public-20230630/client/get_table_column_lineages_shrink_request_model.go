// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterQueryShrink(v string) *GetTableColumnLineagesShrinkRequest
	GetFilterQueryShrink() *string
	SetOpTenantId(v int64) *GetTableColumnLineagesShrinkRequest
	GetOpTenantId() *int64
	SetTableGuid(v string) *GetTableColumnLineagesShrinkRequest
	GetTableGuid() *string
}

type GetTableColumnLineagesShrinkRequest struct {
	FilterQueryShrink *string `json:"FilterQuery,omitempty" xml:"FilterQuery,omitempty"`
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
	// 1121
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetTableColumnLineagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineagesShrinkRequest) GetFilterQueryShrink() *string {
	return s.FilterQueryShrink
}

func (s *GetTableColumnLineagesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableColumnLineagesShrinkRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableColumnLineagesShrinkRequest) SetFilterQueryShrink(v string) *GetTableColumnLineagesShrinkRequest {
	s.FilterQueryShrink = &v
	return s
}

func (s *GetTableColumnLineagesShrinkRequest) SetOpTenantId(v int64) *GetTableColumnLineagesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableColumnLineagesShrinkRequest) SetTableGuid(v string) *GetTableColumnLineagesShrinkRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableColumnLineagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
