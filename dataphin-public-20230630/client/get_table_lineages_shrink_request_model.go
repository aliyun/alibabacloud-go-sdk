// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterQueryShrink(v string) *GetTableLineagesShrinkRequest
	GetFilterQueryShrink() *string
	SetOpTenantId(v int64) *GetTableLineagesShrinkRequest
	GetOpTenantId() *int64
	SetTableGuid(v string) *GetTableLineagesShrinkRequest
	GetTableGuid() *string
}

type GetTableLineagesShrinkRequest struct {
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

func (s GetTableLineagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTableLineagesShrinkRequest) GetFilterQueryShrink() *string {
	return s.FilterQueryShrink
}

func (s *GetTableLineagesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableLineagesShrinkRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableLineagesShrinkRequest) SetFilterQueryShrink(v string) *GetTableLineagesShrinkRequest {
	s.FilterQueryShrink = &v
	return s
}

func (s *GetTableLineagesShrinkRequest) SetOpTenantId(v int64) *GetTableLineagesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableLineagesShrinkRequest) SetTableGuid(v string) *GetTableLineagesShrinkRequest {
	s.TableGuid = &v
	return s
}

func (s *GetTableLineagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
