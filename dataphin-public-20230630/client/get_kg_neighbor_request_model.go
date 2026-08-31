// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgNeighborRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityDataId(v string) *GetKgNeighborRequest
	GetEntityDataId() *string
	SetEntityType(v string) *GetKgNeighborRequest
	GetEntityType() *string
	SetNeighborsQuery(v *GetKgNeighborRequestNeighborsQuery) *GetKgNeighborRequest
	GetNeighborsQuery() *GetKgNeighborRequestNeighborsQuery
	SetOpTenantId(v int64) *GetKgNeighborRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetKgNeighborRequest
	GetOpUserId() *string
	SetWorkspaceId(v string) *GetKgNeighborRequest
	GetWorkspaceId() *string
}

type GetKgNeighborRequest struct {
	// The entity record data ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EntityDataId *string `json:"EntityDataId,omitempty" xml:"EntityDataId,omitempty"`
	// The entity type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Student
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity record neighbor node query instruction.
	NeighborsQuery *GetKgNeighborRequestNeighborsQuery `json:"NeighborsQuery,omitempty" xml:"NeighborsQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKgNeighborRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborRequest) GoString() string {
	return s.String()
}

func (s *GetKgNeighborRequest) GetEntityDataId() *string {
	return s.EntityDataId
}

func (s *GetKgNeighborRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKgNeighborRequest) GetNeighborsQuery() *GetKgNeighborRequestNeighborsQuery {
	return s.NeighborsQuery
}

func (s *GetKgNeighborRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetKgNeighborRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetKgNeighborRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKgNeighborRequest) SetEntityDataId(v string) *GetKgNeighborRequest {
	s.EntityDataId = &v
	return s
}

func (s *GetKgNeighborRequest) SetEntityType(v string) *GetKgNeighborRequest {
	s.EntityType = &v
	return s
}

func (s *GetKgNeighborRequest) SetNeighborsQuery(v *GetKgNeighborRequestNeighborsQuery) *GetKgNeighborRequest {
	s.NeighborsQuery = v
	return s
}

func (s *GetKgNeighborRequest) SetOpTenantId(v int64) *GetKgNeighborRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetKgNeighborRequest) SetOpUserId(v string) *GetKgNeighborRequest {
	s.OpUserId = &v
	return s
}

func (s *GetKgNeighborRequest) SetWorkspaceId(v string) *GetKgNeighborRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetKgNeighborRequest) Validate() error {
	if s.NeighborsQuery != nil {
		if err := s.NeighborsQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKgNeighborRequestNeighborsQuery struct {
	// The maximum depth of neighbor nodes. Default value: 1.
	//
	// example:
	//
	// 2
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The direction type. Valid values:
	//
	// - in: the current entity is the target node.
	//
	// - out: the current entity is the source node.
	//
	// - both: the current entity is both the source node and the target node.
	//
	// Default value: both.
	//
	// example:
	//
	// both
	DirectionType *string `json:"DirectionType,omitempty" xml:"DirectionType,omitempty"`
	// The list of relation types.
	RelationTypes []*string `json:"RelationTypes,omitempty" xml:"RelationTypes,omitempty" type:"Repeated"`
}

func (s GetKgNeighborRequestNeighborsQuery) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborRequestNeighborsQuery) GoString() string {
	return s.String()
}

func (s *GetKgNeighborRequestNeighborsQuery) GetDepth() *int32 {
	return s.Depth
}

func (s *GetKgNeighborRequestNeighborsQuery) GetDirectionType() *string {
	return s.DirectionType
}

func (s *GetKgNeighborRequestNeighborsQuery) GetRelationTypes() []*string {
	return s.RelationTypes
}

func (s *GetKgNeighborRequestNeighborsQuery) SetDepth(v int32) *GetKgNeighborRequestNeighborsQuery {
	s.Depth = &v
	return s
}

func (s *GetKgNeighborRequestNeighborsQuery) SetDirectionType(v string) *GetKgNeighborRequestNeighborsQuery {
	s.DirectionType = &v
	return s
}

func (s *GetKgNeighborRequestNeighborsQuery) SetRelationTypes(v []*string) *GetKgNeighborRequestNeighborsQuery {
	s.RelationTypes = v
	return s
}

func (s *GetKgNeighborRequestNeighborsQuery) Validate() error {
	return dara.Validate(s)
}
