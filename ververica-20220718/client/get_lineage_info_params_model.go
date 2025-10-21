// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageInfoParams interface {
	dara.Model
	String() string
	GoString() string
	SetDepth(v int64) *GetLineageInfoParams
	GetDepth() *int64
	SetDirection(v string) *GetLineageInfoParams
	GetDirection() *string
	SetId(v string) *GetLineageInfoParams
	GetId() *string
	SetIdType(v string) *GetLineageInfoParams
	GetIdType() *string
	SetIsColumnLevel(v bool) *GetLineageInfoParams
	GetIsColumnLevel() *bool
	SetIsTemporary(v bool) *GetLineageInfoParams
	GetIsTemporary() *bool
	SetNamespace(v string) *GetLineageInfoParams
	GetNamespace() *string
	SetWorkspace(v string) *GetLineageInfoParams
	GetWorkspace() *string
}

type GetLineageInfoParams struct {
	Depth         *int64  `json:"depth,omitempty" xml:"depth,omitempty"`
	Direction     *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	IdType        *string `json:"idType,omitempty" xml:"idType,omitempty"`
	IsColumnLevel *bool   `json:"isColumnLevel,omitempty" xml:"isColumnLevel,omitempty"`
	IsTemporary   *bool   `json:"isTemporary,omitempty" xml:"isTemporary,omitempty"`
	Namespace     *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Workspace     *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetLineageInfoParams) String() string {
	return dara.Prettify(s)
}

func (s GetLineageInfoParams) GoString() string {
	return s.String()
}

func (s *GetLineageInfoParams) GetDepth() *int64 {
	return s.Depth
}

func (s *GetLineageInfoParams) GetDirection() *string {
	return s.Direction
}

func (s *GetLineageInfoParams) GetId() *string {
	return s.Id
}

func (s *GetLineageInfoParams) GetIdType() *string {
	return s.IdType
}

func (s *GetLineageInfoParams) GetIsColumnLevel() *bool {
	return s.IsColumnLevel
}

func (s *GetLineageInfoParams) GetIsTemporary() *bool {
	return s.IsTemporary
}

func (s *GetLineageInfoParams) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLineageInfoParams) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetLineageInfoParams) SetDepth(v int64) *GetLineageInfoParams {
	s.Depth = &v
	return s
}

func (s *GetLineageInfoParams) SetDirection(v string) *GetLineageInfoParams {
	s.Direction = &v
	return s
}

func (s *GetLineageInfoParams) SetId(v string) *GetLineageInfoParams {
	s.Id = &v
	return s
}

func (s *GetLineageInfoParams) SetIdType(v string) *GetLineageInfoParams {
	s.IdType = &v
	return s
}

func (s *GetLineageInfoParams) SetIsColumnLevel(v bool) *GetLineageInfoParams {
	s.IsColumnLevel = &v
	return s
}

func (s *GetLineageInfoParams) SetIsTemporary(v bool) *GetLineageInfoParams {
	s.IsTemporary = &v
	return s
}

func (s *GetLineageInfoParams) SetNamespace(v string) *GetLineageInfoParams {
	s.Namespace = &v
	return s
}

func (s *GetLineageInfoParams) SetWorkspace(v string) *GetLineageInfoParams {
	s.Workspace = &v
	return s
}

func (s *GetLineageInfoParams) Validate() error {
	return dara.Validate(s)
}
