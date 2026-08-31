// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineageByTaskIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetTableLineageByTaskIdShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetTableLineageByTaskIdShrinkRequest
	GetOpUserId() *string
	SetTableLineageByTaskIdQueryShrink(v string) *GetTableLineageByTaskIdShrinkRequest
	GetTableLineageByTaskIdQueryShrink() *string
}

type GetTableLineageByTaskIdShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The data structure for querying table lineage.
	//
	// This parameter is required.
	TableLineageByTaskIdQueryShrink *string `json:"TableLineageByTaskIdQuery,omitempty" xml:"TableLineageByTaskIdQuery,omitempty"`
}

func (s GetTableLineageByTaskIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineageByTaskIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTableLineageByTaskIdShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableLineageByTaskIdShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetTableLineageByTaskIdShrinkRequest) GetTableLineageByTaskIdQueryShrink() *string {
	return s.TableLineageByTaskIdQueryShrink
}

func (s *GetTableLineageByTaskIdShrinkRequest) SetOpTenantId(v int64) *GetTableLineageByTaskIdShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableLineageByTaskIdShrinkRequest) SetOpUserId(v string) *GetTableLineageByTaskIdShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetTableLineageByTaskIdShrinkRequest) SetTableLineageByTaskIdQueryShrink(v string) *GetTableLineageByTaskIdShrinkRequest {
	s.TableLineageByTaskIdQueryShrink = &v
	return s
}

func (s *GetTableLineageByTaskIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
